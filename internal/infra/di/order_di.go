package di

import (
	"github.com/sandronister/mongo-go/internal/infra/database"
	"github.com/sandronister/mongo-go/internal/infra/enum"
	"github.com/sandronister/mongo-go/internal/infra/web/handler"
)

func ConfigOrderHandlerDI(hc *handler.HandlerConfig) *handler.OrderHandler {
	collection := hc.Db.Collection(string(enum.Order))
	repository := database.NewOrderRepository(collection)
	return handler.NewOrderHandler(hc.Ctx, repository)
}
