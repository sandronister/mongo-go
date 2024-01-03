package entity

import (
	"context"
)

type OrderRepositoryInterface interface {
	Save(ctx *context.Context, order *Order) error
	List(ctx *context.Context) ([]Order, error)
}
