package internal

import (
	"context"
	"search-basic/internal/modules/language"
	"search-basic/pkg"
)

func Init() {
	client := pkg.ConnectDB()
	languageRepository := language.NewRepository(client)
	languageService := language.NewService(languageRepository)

	console := NewConsole(languageService)
	console.Init(context.Background())
}
