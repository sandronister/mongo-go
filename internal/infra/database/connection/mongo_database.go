package connection

import (
	"context"

	"github.com/sandronister/mongo-go/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDatabase(config *configs.Conf) (*mongo.Database, error) {
	ctx := context.Background()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(config.DBHost))

	if err != nil {
		return nil, err
	}

	return mongoClient.Database(config.DBName), nil
}
