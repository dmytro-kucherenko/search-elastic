package common

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8"
)

type Tool struct {
	*elasticsearch.TypedClient
	versions Versions
}

func NewTool(client *elasticsearch.TypedClient) *Tool {
	return &Tool{client, make(Versions)}
}

func (tool *Tool) LoadVersions(ctx context.Context) error {
	aliases, err := tool.Cat.Aliases().Do(ctx)
	if err != nil {
		return err
	}

	for _, record := range aliases {
		version, err := getVersion(*record.Index)
		if err == nil {
			tool.versions[*record.Alias] = version
		}
	}

	return nil
}

func (tool *Tool) AddVersion(alias string, version uint) {
	tool.versions[alias] = version
}

func (tool *Tool) ExistsVersion(alias string, version uint) bool {
	current, ok := tool.versions[alias]
	if !ok {
		return false
	}

	return current >= version
}

func getVersion(index string) (uint, error) {
	re := regexp.MustCompile(`v(\d+)$`)
	matches := re.FindStringSubmatch(index)
	if matches == nil {
		return 0, fmt.Errorf("invalid index format")
	}

	version, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, fmt.Errorf("invalid index versoion")
	}

	return uint(version), nil
}
