package repository

import (
	"context"

	"github.com/sandronister/mongo-go/internal/entity"
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
	req, err := r.Collection.InsertOne(ctx, order)

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
