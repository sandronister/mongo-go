package repository

import (
	"context"

	"github.com/sandronister/mongo-go/internal/entity"
	"github.com/sandronister/mongo-go/internal/infra/database/mongo/schemas"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	Collection *mongo.Collection
}

func NewOrderRepository(collection *mongo.Collection) *OrderRepository {
	return &OrderRepository{Collection: collection}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	ctx := context.Background()

	orderSchmea := schemas.Order{
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}

	req, err := r.Collection.InsertOne(ctx, orderSchmea)

	if err != nil {
		return err
	}

	order.ID = req.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (r *OrderRepository) List() ([]entity.Order, error) {
	ctx := context.Background()
	cur, err := r.Collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	var orders []entity.Order

	defer cur.Close(ctx)

	if err = cur.All(ctx, &orders); err != nil {
		return orders, err
	}

	return orders, nil
}
