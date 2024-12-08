package main

import (
	"context"
	"encoding/json"

	"github.com/arystanbek2002/swe/api"
	"github.com/arystanbek2002/swe/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://team1:swe@db:27017"))
	if err != nil {
		panic(err)
	}

	store, err := storage.NewMongoStore(ctx, client)
	if err != nil {
		panic(err)
	}

	server := api.NewAPIServer(":8080", store)
	server.Run()
}
