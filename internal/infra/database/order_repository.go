package database

import (
	"context"
	"fmt"

	"github.com/sandronister/mongo-go/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	Collection *mongo.Collection
}

func NewOrderRepository(collection *mongo.Collection) *OrderRepository {
	return &OrderRepository{Collection: collection}
}

func (r *OrderRepository) Save(ctx context.Context, order *entity.Order) error {
	req, err := r.Collection.InsertOne(ctx, order)

	if err != nil {
		return err
	}

	order.ID = fmt.Sprintf("%v", req.InsertedID)
	return nil
}

func (r *OrderRepository) List(ctx context.Context) error {
	cur, err := r.Collection.Find(ctx, bson.D{})

	if err != nil {
		return err
	}

	var orders []bson.M

	defer cur.Close(ctx)

	for cur.Next(ctx) {

		var order bson.M

		if err := cur.Decode(&order); err != nil {
			return err
		}

		orders = append(orders, order)
	}

	return nil
}
