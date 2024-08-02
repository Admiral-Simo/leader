package main

import (
	"flag"
	"os"
	"server/engines"
	"strings"
)

func main() {
	query := flag.String("query", "", "query to search for.")
	onlyEmails := flag.Bool("onlyEmails", false, "only emails flag.")
	flag.Parse()
	if *query == "" {
		flag.Usage()
		os.Exit(1)
	}

	*query = strings.ReplaceAll(*query, " ", "+")

	var searchEngine engines.SearchEngine

	searchEngine = engines.NewGoogleSearchEngine()
	result := searchEngine.GetEmailsAddressByQuery(*query)

	if *onlyEmails {
		showEmails(result)
	} else {
		prettierFormat(result)
	}
}

func prettierFormat(result map[string]map[string]struct{}) {
	for site, emails := range result {
		os.Stdout.WriteString(site + "\n")
		for email := range emails {
			os.Stdout.WriteString("  " + email + "\n")
		}
	}
}

func showEmails(result map[string]map[string]struct{}) {
	for _, emails := range result {
		for email := range emails {
			os.Stdout.WriteString(email + "\n")
		}
	}
}
