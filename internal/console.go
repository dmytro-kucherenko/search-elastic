package internal

import (
	"context"
	"fmt"
	"os"
	"search-basic/internal/common/utils"
	"search-basic/internal/modules/language"
)

type Console struct {
	languageService *language.Service
}

func NewConsole(languageService *language.Service) *Console {
	return &Console{languageService}
}

func (console *Console) Init(ctx context.Context) {
	handle := utils.CreateHandler(ctx)

	for {
		fmt.Println("\nChoose an option:",
			"\n1 - Get Many Simple",
			"\n2 - Get Many Full Text",
			"\n3 - Create",
			"\n4 - Delete",
			"\n5 - Exit",
		)
		fmt.Println()

		var action string
		fmt.Scanln(&action)

		fmt.Println()

		switch action {
		case "1":
			handle(console.getManySimple)
		case "2":
			handle(console.getManyFullText)
		case "3":
			handle(console.create)
		case "4":
			handle(console.delete)
		case "5":
			os.Exit(0)
		}
	}
}

func (console *Console) getManySimple(ctx context.Context) (text string, err error) {
	params, err := utils.ScanFields[language.GetManySimpleParams]([]utils.ScanField{
		{Title: "Name: ", Field: "name", IsOptional: true},
		{Title: "Popularity: ", Field: "popularity", IsOptional: true},
		{Title: "Typed (true/false): ", Field: "isTyped", IsOptional: true},
		{Title: "Created at (2006-01-02 15:04:05): ", Field: "createdAt", IsOptional: true},
	})

	if err != nil {
		return
	}

	result, err := console.languageService.GetManySimple(params, ctx)
	if err != nil {
		return
	}

	text = fmt.Sprintf("%v\n", language.FormatManyResult(result))

	return
}

func (console *Console) getManyFullText(ctx context.Context) (text string, err error) {
	params, err := utils.ScanFields[language.GetManyFullTextParams]([]utils.ScanField{
		{Title: "Description:\n", Field: "description", IsOptional: true},
		{Title: "Creation purpose:\n", Field: "creationPurpose", IsOptional: true},
		{Title: "Famous projects:\n", Field: "famousProjects", IsOptional: true},
	})

	if err != nil {
		return
	}

	result, err := console.languageService.GetManyFullText(params, ctx)
	if err != nil {
		return
	}

	text = fmt.Sprintf("%v\n", language.FormatManyResult(result))

	return
}

func (console *Console) create(ctx context.Context) (text string, err error) {
	params, err := utils.ScanFields[language.CreateParams]([]utils.ScanField{
		{Title: "Name: ", Field: "name"},
		{Title: "Popularity: ", Field: "popularity"},
		{Title: "Typed (true/false): ", Field: "isTyped"},
		{Title: "Created at (2006-01-02 15:04:05): ", Field: "createdAt"},
		{Title: "Description:\n", Field: "description"},
		{Title: "Creation purpose:\n", Field: "creationPurpose"},
		{Title: "Famous projects:\n", Field: "famousProjects"},
	})

	if err != nil {
		return
	}

	result, err := console.languageService.Create(params, ctx)
	if err != nil {
		return
	}

	text = fmt.Sprintf("%v\n", language.FormatItem(result))

	return
}

func (console *Console) delete(ctx context.Context) (text string, err error) {
	params, err := utils.ScanFields[language.DeleteParams]([]utils.ScanField{
		{Title: "ID: ", Field: "id"},
	})

	if err != nil {
		return
	}

	err = console.languageService.Delete(params, ctx)
	if err != nil {
		return
	}

	text = fmt.Sprintf("Deleted: %v\n", params.ID)

	return
}
