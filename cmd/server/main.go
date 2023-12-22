package main

import (
	"context"
	"fmt"

	"github.com/sandronister/mongo-go/internal/infra/web/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		fmt.Println(err.Error())
	}

	routes.HandlerRequest("8080", mongoClient.Database("go_mongo"))
}
