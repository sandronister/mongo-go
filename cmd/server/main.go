package main

import (
	"context"

	"github.com/sandronister/mongo-go/configs"
	"github.com/sandronister/mongo-go/internal/infra/web/routes"
	"github.com/sandronister/mongo-go/pkg/catch"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	config, err := configs.LoadConfig(".")
	catch.Execute(err)
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(config.DBHost))
	catch.Execute(err)
	routes.HandlerRequest(config.WebPort, mongoClient.Database(config.DBName))
}
