package language

import (
	"fmt"
	"time"
)

func FormatItem(item Item) string {
	return fmt.Sprintf(
		"\nLanguage: %v\nName: %v\nPopularity: %v\nIs typed: %v\nCreated At: %v \nDescription: \n%v\nCreation purpose: \n%v\nFamous projects: \n%v",
		item.ID,
		item.Name,
		item.Popularity,
		item.IsTyped,
		item.CreatedAt.Format(time.DateTime),
		item.Description,
		item.CreationPurpose,
		item.FamousProjects,
	)
}

func FormatManyResult(result ManyResult) string {
	text := fmt.Sprintf("\nTotal: %v", result.Total)
	for _, item := range result.Items {
		text += fmt.Sprintf("\n%v", FormatItem(item))
	}

	return text
}
