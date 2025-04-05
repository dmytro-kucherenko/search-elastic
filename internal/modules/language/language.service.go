package language

import (
	"context"
	"errors"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository}
}

func (service *Service) GetManySimple(params GetManySimpleParams, ctx context.Context) (ManyResult, error) {
	return service.repository.GetMany(GetManyParams{GetManySimpleParams: params}, ctx)
}

func (service *Service) GetManyFullText(params GetManyFullTextParams, ctx context.Context) (ManyResult, error) {
	return service.repository.GetMany(GetManyParams{GetManyFullTextParams: params}, ctx)
}

func (service *Service) Create(params CreateParams, ctx context.Context) (item Item, err error) {
	ID, err := service.repository.Create(params, ctx)
	if err != nil {
		return
	}

	item = Item{
		ID:              ID,
		Name:            params.Name,
		Popularity:      params.Popularity,
		IsTyped:         params.IsTyped,
		CreatedAt:       params.CreatedAt,
		Description:     params.Description,
		CreationPurpose: params.CreationPurpose,
		FamousProjects:  params.FamousProjects,
	}

	return
}

func (service *Service) Delete(params DeleteParams, ctx context.Context) error {
	ok, err := service.repository.Delete(params, ctx)
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("language not found")
	}

	return nil
}
