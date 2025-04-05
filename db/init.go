package db

import (
	"context"
	"fmt"
	"search-basic/db/common"
	"search-basic/db/migrations"
	"search-basic/pkg"
)

func Init() {
	ctx := context.Background()
	client := pkg.ConnectDB()

	tool := common.NewTool(client)
	tool.LoadVersions(ctx)
	fmt.Println("aliases loaded")

	items := []common.NewMigration{
		migrations.NewCreateLanguageV1Index,
	}

	skipped := 0
	for _, create := range items {
		migration := create(tool)
		name := migration.Name()

		exists := migration.Exists()
		if exists {
			skipped++

			continue
		}

		fmt.Print(name)

		err := migration.Make(ctx)
		if err != nil {
			fmt.Printf(" -> error \n%v\n", err.Error())

			return
		}

		fmt.Println(" -> done")
	}

	fmt.Println("skipped", skipped)
}
