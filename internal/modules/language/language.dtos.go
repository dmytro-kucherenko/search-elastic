package language

import (
	"time"
)

type GetManySimpleParams struct {
	Name       *string    `json:"name" mapstructure:"name"`
	Popularity *uint8     `json:"popularity" mapstructure:"popularity" validate:"omitempty,min=1,max=10"`
	IsTyped    *bool      `json:"is_typed" mapstructure:"isTyped"`
	CreatedAt  *time.Time `json:"created_at" mapstructure:"createdAt"`
}

type GetManyFullTextParams struct {
	Description     *string `json:"description" mapstructure:"description"`
	CreationPurpose *string `json:"creation_purpose" mapstructure:"creationPurpose"`
	FamousProjects  *string `json:"famous_projects" mapstructure:"famousProjects"`
}

type GetManyParams struct {
	GetManySimpleParams
	GetManyFullTextParams
}

type CreateParams struct {
	Name            string    `json:"name" mapstructure:"name" validate:"required"`
	Popularity      uint8     `json:"popularity" mapstructure:"popularity" validate:"required,min=1,max=10"`
	IsTyped         bool      `json:"is_typed" mapstructure:"isTyped"`
	Description     string    `json:"description" mapstructure:"description" validate:"required,min=8"`
	CreationPurpose string    `json:"creation_purpose" mapstructure:"creationPurpose" validate:"required,min=8"`
	FamousProjects  string    `json:"famous_projects" mapstructure:"famousProjects" validate:"required,min=8"`
	CreatedAt       time.Time `json:"created_at" mapstructure:"createdAt" validate:"required"`
}

type DeleteParams struct {
	ID string `json:"id" mapstructure:"id"`
}

type Item struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Popularity      uint8     `json:"popularity"`
	IsTyped         bool      `json:"is_typed"`
	Description     string    `json:"description"`
	CreationPurpose string    `json:"creation_purpose"`
	FamousProjects  string    `json:"famous_projects"`
	CreatedAt       time.Time `json:"created_at"`
}

type ManyResult struct {
	Items []Item `json:"items"`
	Total uint32 `json:"total"`
}
