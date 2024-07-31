package engines

type SearchEngine interface {
	GetWebsitesByQuery(query string) []string
	GetEmailsAddressByQuery(query string) map[string]map[string]struct{}
	Open(url string)
}
