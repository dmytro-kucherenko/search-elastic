package pkg

import (
	"context"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/joho/godotenv"
)

func ConnectDB() *elasticsearch.TypedClient {
	godotenv.Load()
	elasticEndpoint := os.Getenv("DB_ENDPOINT")
	elasticUsername := os.Getenv("DB_USERNAME")
	elasticPassword := os.Getenv("DB_PASSWORD")

	client, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{elasticEndpoint},
		Username:  elasticUsername,
		Password:  elasticPassword,
	})

	if err != nil {
		panic(err.Error())
	}

	_, err = client.Info().Do(context.TODO())
	if err != nil {
		panic(err.Error())
	}

	return client
}
