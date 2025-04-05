package migrations

import (
	"context"
	"search-basic/db/common"

	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	types "github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	name    = "1-create-language-v1-index"
	index   = "language_v1"
	alias   = "language"
	version = 1
)

type migration struct {
	tool *common.Tool
}

func NewCreateLanguageV1Index(tool *common.Tool) common.Migration {
	return &migration{tool}
}

func (migration migration) Name() string {
	return name
}

func (migration migration) Exists() bool {
	return migration.tool.ExistsVersion(alias, version)
}

func (migration migration) Make(ctx context.Context) error {
	htmlAnalyzer := types.NewCustomAnalyzer()
	htmlAnalyzer.Type = "custom"
	htmlAnalyzer.Tokenizer = "standard"
	htmlAnalyzer.Filter = []string{"lowercase", "porter_stem"}
	htmlAnalyzer.CharFilter = []string{"html_strip"}
	htmlAnalyzerName := "html"

	request := &create.Request{
		Settings: &types.IndexSettings{
			Analysis: &types.IndexSettingsAnalysis{
				Analyzer: map[string]types.Analyzer{htmlAnalyzerName: *htmlAnalyzer},
			},
		},
		Mappings: &types.TypeMapping{
			Properties: map[string]types.Property{
				"name":             types.NewKeywordProperty(),
				"popularity":       types.NewIntegerNumberProperty(),
				"is_typed":         types.NewBooleanProperty(),
				"description":      types.TextProperty{Analyzer: &[]string{"standard"}[0]},
				"creation_purpose": types.TextProperty{Analyzer: &[]string{"english"}[0]},
				"famous_projects":  types.TextProperty{Analyzer: &htmlAnalyzerName},
				"created_at":       types.NewDateProperty(),
			},
		},
	}

	_, err := migration.tool.Indices.Create(index).Request(request).Do(ctx)
	if err != nil {
		return err
	}

	_, err = migration.tool.Indices.PutAlias(index, alias).Do(ctx)
	if err != nil {
		return err
	}

	migration.tool.AddVersion(alias, version)

	return nil
}
