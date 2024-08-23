package engines

import (
	"fmt"
	"math/rand"
	"os/exec"
	"regexp"
	"strings"
	"sync"

	"github.com/gocolly/colly/v2"
)

type GoogleSearchEngine struct {
	emailRegex       *regexp.Regexp
	websiteCollector *colly.Collector
	emailCollector   *colly.Collector
}

func getRandomUserAgent() string {
	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		// Add more User-Agents as needed
	}
	return userAgents[rand.Intn(len(userAgents))]
}

func NewGoogleSearchEngine() *GoogleSearchEngine {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	websiteCollector := colly.NewCollector(
		colly.AllowedDomains("www.google.com"),
		colly.UserAgent(getRandomUserAgent()), // Use a random User-Agent
		colly.Async(true),                     // Enable asynchronous requests
	)

	emailCollector := colly.NewCollector(colly.Async(true))

	return &GoogleSearchEngine{
		emailRegex:       re,
		websiteCollector: websiteCollector,
		emailCollector:   emailCollector,
	}
}

// GetWebsitesByQuery returns a list of websites based on the query by scraping Google search results.
func (g GoogleSearchEngine) GetWebsitesByQuery(query string) []string {
	query = strings.ReplaceAll(query, " ", "+")

	var websites []string

	// On every anchor element which has an href attribute call the callback
	g.websiteCollector.OnHTML("a", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		if strings.HasPrefix(href, "/url?q=") {
			// Extract the URL from the href attribute
			website := strings.Split(strings.TrimPrefix(href, "/url?q="), "&")[0]
			websites = append(websites, website)
		}
	})

	// Visit multiple pages to get more results
	for start := 0; start < 200; start += 10 {
		if len(websites) >= 200 {
			break
		}
		searchURL := fmt.Sprintf("https://www.google.com/search?q=%s&start=%d", query, start)
		g.websiteCollector.Visit(searchURL)
	}
	return websites
}

// GetEmailsAddressByQuery returns a map of email addresses found on websites based on the query.
func (g GoogleSearchEngine) GetEmailsAddressByQuery(query string) map[string]map[string]struct{} {
	query = strings.ReplaceAll(query, " ", "+")

	websites := g.GetWebsitesByQuery(query)
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

func (g GoogleSearchEngine) Open(url string) {
	cmd := exec.Command("open", url)
	cmd.Run()
}

func (g GoogleSearchEngine) isEmail(word string) bool {
	// Return whether the word matches the email pattern
	return g.emailRegex.MatchString(word)
}
