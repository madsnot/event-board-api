package opensearch

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/madsnot/event-board-api/internal/config"
	"github.com/madsnot/event-board-api/internal/domain/models"
	"github.com/madsnot/event-board-api/internal/repository"
	"github.com/madsnot/event-board-api/internal/repository/opensearch/dto"
	"github.com/madsnot/event-board-api/internal/repository/postgres"
	"github.com/opensearch-project/opensearch-go/v2"
	"github.com/opensearch-project/opensearch-go/v2/opensearchapi"
	"io"
	"log"
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

func (c Client) InitIndex(ctx context.Context, er postgres.EventRepository) error {
	if err := c.createIndex(ctx); err != nil {
		return err
	}

	if err := c.migrateDataToIndex(ctx, er); err != nil {

	}

	return nil
}

func (c Client) createIndex(ctx context.Context) error {
	s := dto.InitSettings()

	settings, err := json.Marshal(s)
	if err != nil {
		return repository.ErrInternal.Wrap(err)
	}

	req := opensearchapi.IndicesCreateRequest{
		Index: c.cfg.Index,
		Body:  bytes.NewReader(settings),
	}

	resp, err := req.Do(ctx, c.client)
	if err != nil {
		return repository.ErrInternal.Wrap(err)
	}

	log.Println(resp)

	return nil
}

func (c Client) migrateDataToIndex(ctx context.Context, er postgres.EventRepository) error {
	filters := models.EventFilters{
		Genders: models.EventGender{
			Man:   true,
			Woman: true,
		},
	}

	list, err := er.GetList(ctx, filters)
	if err != nil {
		return err
	}

	if list == nil {
		return nil
	}

	for _, event := range list {
		if err := c.Index(ctx, event); err != nil {
			return err
		}
	}

	return nil
}

func (c Client) Index(ctx context.Context, event models.Event) error {
	var end time.Time

	if event.EndDate != nil {
		end = *event.EndDate
	}

	document := dto.IndexDocument{
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

func (c Client) Search(ctx context.Context, query string, from, size int) ([]uuid.UUID, error) {
	var (
		ids []uuid.UUID
		res dto.SearchDocument
	)

	q := dto.SearchQuery{
		From: from,
		Size: size,
		Query: dto.MultiMatchQuery{
			MultiMatch: dto.MultiMatch{
				Query:               query,
				Fields:              []string{"title^6", "description^4", "title.spell^2", "description.spell"},
				Analyzer:            "spell_analyzer",
				FuzzyTranspositions: true,
			},
		},
	}

	body, err := json.Marshal(q)

	search := opensearchapi.SearchRequest{
		Index: []string{c.cfg.Index},
		Body:  bytes.NewReader(body),
	}

	searchResponse, err := search.Do(ctx, c.client)
	if err != nil {
		return nil, repository.ErrInternal.Wrap(err)
	}

	b, err := io.ReadAll(searchResponse.Body)
	if err != nil {
		return nil, repository.ErrInternal.Wrap(err)
	}

	if err = json.Unmarshal(b, &res); err != nil {
		return nil, repository.ErrInternal.Wrap(err)
	}

	for _, doc := range res.Hits.Hits {
		ids = append(ids, doc.ID)
	}

	return ids, nil
}
