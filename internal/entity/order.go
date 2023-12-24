package entity

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Price      float64            `bson:"price,omitempty"`
	Tax        float64            `bson:"tax,omitempty"`
	FinalPrice float64            `bson:"final,omitempty"`
}

func NewOrder(price, tax float64) (*Order, error) {
	order := &Order{
		Price: price,
		Tax:   tax,
	}

	err := order.IsValid()
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *Order) IsValid() error {
	if o.Price <= 0 {
		return errors.New("invalid Price")
	}

	if o.Tax <= 0 {
		return errors.New("invalid tax")
	}

	return nil
}

func (o *Order) CalculateFinalPrice() error {
	taxPrice := o.Price * (o.Tax / 100)
	o.FinalPrice = o.Price + taxPrice
	err := o.IsValid()

	if err != nil {
		return err
	}

	return nil
}
