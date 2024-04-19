package schemas

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Price      float64            `bson:"price,omitempty"`
	Tax        float64            `bson:"tax,omitempty"`
	FinalPrice float64            `bson:"final,omitempty"`
}
