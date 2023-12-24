package di

import (
	"context"

	"github.com/sandronister/mongo-go/internal/infra/database"
	"github.com/sandronister/mongo-go/internal/infra/web/handler"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConfigOrderHandlerDI(ctx context.Context, mongodb *mongo.Database) *handler.OrderHandler {
	collection := mongodb.Collection("orders")
	repository := database.NewOrderRepository(collection)
	return handler.NewOrderHandler(ctx, repository)
}
