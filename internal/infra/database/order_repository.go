package database

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

func (r *OrderRepository) Save(ctx context.Context, order *entity.Order) error {
	req, err := r.Collection.InsertOne(ctx, order)

	if err != nil {
		return err
	}

	order.ID = req.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (r *OrderRepository) List(ctx context.Context) ([]entity.Order, error) {
	cur, err := r.Collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	var orders []entity.Order

	defer cur.Close(ctx)

	for cur.Next(ctx) {

		var order entity.Order

		if err := cur.Decode(&order); err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}
