package language

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/optype"
)

const (
	alias = "language"
)

type Repository struct {
	client *elasticsearch.TypedClient
}

func NewRepository(client *elasticsearch.TypedClient) *Repository {
	return &Repository{client}
}

func (repository *Repository) GetMany(params GetManyParams, ctx context.Context) (result ManyResult, err error) {
	queries := make([]types.Query, 0, 7)
	if params.Name != nil {
		schema := fmt.Sprintf(`%v.*`, regexp.QuoteMeta(*params.Name))
		insensitive := true

		queries = append(
			queries,
			types.Query{Regexp: map[string]types.RegexpQuery{"name": {
				CaseInsensitive: &insensitive,
				Value:           schema,
			}}},
		)
	}

	if params.Popularity != nil {
		popularity := types.Float64(*params.Popularity)
		queries = append(
			queries,
			types.Query{Range: map[string]types.RangeQuery{"popularity": types.NumberRangeQuery{
				Gt: &popularity,
			}}},
		)
	}

	if params.IsTyped != nil {
		queries = append(
			queries,
			types.Query{Term: map[string]types.TermQuery{"is_typed": {
				Value: params.IsTyped,
			}}},
		)
	}

	if params.CreatedAt != nil {
		createdAt := params.CreatedAt.UTC().Format(time.RFC3339)
		queries = append(
			queries,
			types.Query{Range: map[string]types.RangeQuery{"created_at": &types.DateRangeQuery{
				Gt: &createdAt,
			}}},
		)
	}

	if params.Description != nil {
		analyzer := "standard"
		queries = append(
			queries,
			types.Query{Match: map[string]types.MatchQuery{"description": {
				Analyzer: &analyzer,
				Query:    *params.Description,
			}}},
		)
	}

	if params.CreationPurpose != nil {
		analyzer := "english"
		queries = append(
			queries,
			types.Query{Match: map[string]types.MatchQuery{"creation_purpose": {
				Analyzer: &analyzer,
				Query:    *params.CreationPurpose,
			}}},
		)
	}

	if params.FamousProjects != nil {
		analyzer := "html"
		queries = append(
			queries,
			types.Query{Match: map[string]types.MatchQuery{"famous_projects": {
				Analyzer: &analyzer,
				Query:    *params.FamousProjects,
			}}},
		)
	}

	request := &search.Request{
		Query: &types.Query{
			Bool: &types.BoolQuery{Must: queries},
		},
		TrackTotalHits: true,
	}

	response, err := repository.client.Search().Index(alias).Request(request).Do(ctx)
	if err != nil {
		return
	}

	result.Total = uint32(response.Hits.Total.Value)
	for _, hit := range response.Hits.Hits {
		item, err := parse(*hit.Id_, hit.Source_)
		if err != nil {
			return result, err
		}

		result.Items = append(result.Items, item)
	}

	return
}

func (repository *Repository) Create(params CreateParams, ctx context.Context) (ID string, err error) {
	response, err := repository.client.Index(alias).OpType(optype.Create).Request(params).Do(ctx)
	if err != nil {
		return
	}

	return response.Id_, nil
}

func (repository *Repository) Delete(params DeleteParams, ctx context.Context) (bool, error) {
	response, err := repository.client.Delete(alias, params.ID).Do(ctx)

	return response.Result.Name == "deleted", err
}

func parse(ID string, source json.RawMessage) (item Item, err error) {
	item.ID = ID
	err = json.Unmarshal(source, &item)
	if err != nil {
		return
	}

	return
}
