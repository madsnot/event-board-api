package opensearch

import "github.com/google/uuid"

type IndexDocument struct {
	title       string `json:"title"`
	description string `json:"description"`
}

type SearchDocument struct {
	id uuid.UUID `json:"_id"`
}

type SearchDocuments struct {
	docs []SearchDocument `json:"hits"`
}

type SearchResponse struct {
	hits SearchDocuments `json:"hits"`
}
