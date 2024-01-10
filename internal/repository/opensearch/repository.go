package opensearch

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/madsnot/event-board-api/internal/config"
	"github.com/madsnot/event-board-api/internal/domain/models"
	"github.com/madsnot/event-board-api/internal/repository"
	"github.com/opensearch-project/opensearch-go/v2"
	"github.com/opensearch-project/opensearch-go/v2/opensearchapi"
	"io"
	"strings"
	"time"
)

type Client struct {
	cfg    config.OpensearchConfig
	client *opensearch.Client
}

func NewOpensearchClient(cfg config.OpensearchConfig, client *opensearch.Client) Client {
	return Client{
		cfg:    cfg,
		client: client,
	}
}

func (c Client) Index(ctx context.Context, event models.Event) error {
	var end time.Time

	if event.EndDate != nil {
		end = *event.EndDate
	}

	document := IndexDocument{
		Title:       event.Title,
		Description: event.Description,
		StartDate:   event.StartDate,
		EndDate:     end,
	}

	body, err := json.Marshal(document)
	if err != nil {
		return repository.ErrInternal.Wrap(err)
	}

	req := opensearchapi.IndexRequest{
		Index:      c.cfg.Index,
		DocumentID: event.ID.String(),
		Body:       strings.NewReader(string(body)),
	}

	_, err = req.Do(ctx, c.client)
	if err != nil {
		return repository.ErrInternal.Wrap(err)
	}

	return nil
}

func (c Client) Search(ctx context.Context, query string) ([]uuid.UUID, error) {
	var (
		ids []uuid.UUID
		res SearchResponse
	)

	content := strings.NewReader(fmt.Sprintf(`{
    "query": {
        "multi_match": {
        	"query": "%s",
        	"fields": ["title", "description"]
        }
    }
}`, query))

	search := opensearchapi.SearchRequest{
		Index: []string{c.cfg.Index},
		Body:  content,
	}

	searchResponse, err := search.Do(ctx, c.client)
	if err != nil {
		return nil, repository.ErrInternal.Wrap(err)
	}

	bytes, err := io.ReadAll(searchResponse.Body)
	if err != nil {
		return nil, repository.ErrInternal.Wrap(err)
	}

	if err = json.Unmarshal(bytes, &res); err != nil {
		return nil, repository.ErrInternal.Wrap(err)
	}

	for _, doc := range res.hits.docs {
		ids = append(ids, doc.id)
	}

	return ids, nil
}
