package engines

import (
	"fmt"
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

func NewGoogleSearchEngine() *GoogleSearchEngine {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	websiteCollector := colly.NewCollector(
		colly.AllowedDomains("www.google.com"),
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
		colly.Async(true), // Enable asynchronous requests
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

	g.websiteCollector.Wait()
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
