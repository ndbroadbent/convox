package algoliasearch

type IndexOperation struct {
	Destination string   `json:"destination"`
	Operation   string   `json:"operation"`
	Scopes      []string `json:"scope,omitempty"`
}
