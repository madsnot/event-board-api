package dto

type MultiMatchQuery struct {
	MultiMatch MultiMatch `json:"multi_match"`
}

type MultiMatch struct {
	Query               string   `json:"query"`
	Fields              []string `json:"fields"`
	Analyzer            string   `json:"analyzer"`
	FuzzyTranspositions bool     `json:"fuzzy_transpositions"`
}

type SearchQuery struct {
	From  int             `json:"from"`
	Size  int             `json:"size"`
	Query MultiMatchQuery `json:"query"`
}
