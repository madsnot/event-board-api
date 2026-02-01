package dto

import "github.com/google/uuid"

type SearchDocument struct {
	Hits struct {
		Hits []struct {
			ID uuid.UUID `json:"_id"`
		} `json:"hits"`
	} `json:"hits"`
}
