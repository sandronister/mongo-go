package di

import (
	"github.com/sandronister/mongo-go/internal/infra/database/mongo/repository"
	"github.com/sandronister/mongo-go/internal/infra/enum"
	"github.com/sandronister/mongo-go/internal/usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewOrderUseCase(db *mongo.Database) *usecase.OrderUseCase {
	collection := db.Collection(string(enum.Order))
	repository := repository.NewOrderRepository(collection)
	return usecase.NewOrderUseCase(repository)
}
