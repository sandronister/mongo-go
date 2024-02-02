package di

import (
	"github.com/sandronister/mongo-go/internal/infra/database"
	"github.com/sandronister/mongo-go/internal/infra/enum"
	"github.com/sandronister/mongo-go/internal/usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

func OrderUseCaseDI(db *mongo.Database) *usecase.OrderUseCase {
	collection := db.Collection(string(enum.Order))
	repository := database.NewOrderRepository(collection)
	return usecase.NewOrderUseCase(repository)
}
