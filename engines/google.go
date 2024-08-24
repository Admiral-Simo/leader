package engines

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/gocolly/colly/v2"
)

type GoogleSearchEngine struct {
	cx               string
	apiKey           string
	emailRegex       *regexp.Regexp
	websiteCollector *colly.Collector
	emailCollector   *colly.Collector
}

type GoogleSearchResponse struct {
	Items []struct {
		Link string `json:"link"`
	} `json:"items"`
}

func NewGoogleSearchEngine() *GoogleSearchEngine {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	apiKey := os.Getenv("API_KEY")
	cx := os.Getenv("CX")

	websiteCollector := colly.NewCollector(
		colly.AllowedDomains("www.google.com"),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 8_2_2) AppleWebKit/536.30 (KHTML, like Gecko) Chrome/53.0.3698.216 Safari/602"),
		colly.Async(true), // Enable asynchronous requests
	)

	emailCollector := colly.NewCollector(colly.Async(true))

	return &GoogleSearchEngine{
		cx:               cx,
		apiKey:           apiKey,
		emailRegex:       re,
		websiteCollector: websiteCollector,
		emailCollector:   emailCollector,
	}
}

func (g GoogleSearchEngine) GetWebsitesByQuery(query string) []string {
    query = strings.ReplaceAll(query, " ", "+")

    var websites []string
    baseURL := "https://www.googleapis.com/customsearch/v1"

    for start := 1; start <= 100; start += 10 { // Adjust the range to get more results
        reqURL, err := url.Parse(baseURL)
        if err != nil {
            fmt.Println("Error parsing URL:", err)
            return websites
        }

        params := url.Values{}
        params.Add("key", g.apiKey)
        params.Add("cx", g.cx)
        params.Add("q", query)
        params.Add("start", fmt.Sprintf("%d", start)) // Start parameter to paginate results
        reqURL.RawQuery = params.Encode()

        resp, err := http.Get(reqURL.String())
        if err != nil {
            fmt.Println("Error making HTTP request:", err)
            return websites
        }
        defer resp.Body.Close()

        var searchResponse GoogleSearchResponse
        if err := json.NewDecoder(resp.Body).Decode(&searchResponse); err != nil {
            fmt.Println("Error decoding JSON response:", err)
            return websites
        }

        for _, item := range searchResponse.Items {
            websites = append(websites, item.Link)
        }

        // Break early if less than 10 results were returned (meaning no more results)
        if len(searchResponse.Items) < 10 {
            break
        }
    }

    return websites
}

// GetEmailsAddressByQuery returns a map of email addresses found on websites based on the query.
func (g GoogleSearchEngine) GetEmailsAddressByQuery(query string) map[string]map[string]struct{} {
	websites := g.GetWebsitesByQuery(query)

	fmt.Println(websites)
	emailData := make(map[string]map[string]struct{})
	var mu sync.Mutex
	var wg sync.WaitGroup

	g.emailCollector.OnHTML("body", func(e *colly.HTMLElement) {
		text := e.Text
		words := strings.Fields(text)
		emails := make(map[string]struct{})
		for _, word := range words {
			if g.isEmail(word) {
				emails[word] = struct{}{}
			}
		}

		if len(emails) > 0 {
			mu.Lock()
			emailData[e.Request.URL.String()] = emails
			mu.Unlock()
		}
	})

	// Visit each website concurrently
	for _, website := range websites {
		wg.Add(1)
		go func(site string) {
			defer wg.Done()
			g.emailCollector.Visit(site)
		}(website)
	}

	wg.Wait()
	g.emailCollector.Wait()
	return emailData
}

func (g GoogleSearchEngine) isEmail(word string) bool {
	// Return whether the word matches the email pattern
	return g.emailRegex.MatchString(word)
}
