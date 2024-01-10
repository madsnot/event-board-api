package opensearch

import (
	"github.com/google/uuid"
	"time"
)

type IndexDocument struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
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
