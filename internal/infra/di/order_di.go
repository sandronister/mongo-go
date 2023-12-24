package di

import (
	"context"

	"github.com/sandronister/mongo-go/internal/infra/database"
	"github.com/sandronister/mongo-go/internal/infra/web/handler"
	"github.com/sandronister/mongo-go/internal/pkg/enum"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConfigOrderHandlerDI(ctx context.Context, mongodb *mongo.Database) *handler.OrderHandler {
	collection := mongodb.Collection(string(enum.Order))
	repository := database.NewOrderRepository(collection)
	return handler.NewOrderHandler(ctx, repository)
}
