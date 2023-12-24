package usecase

import (
	"context"

	"github.com/sandronister/mongo-go/internal/entity"
)

type OrderInputDTO struct {
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type OrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type OrderUseCase struct {
	ctx        context.Context
	repository entity.OrderRepositoryInterface
}

func NewOrderUseCase(ctx context.Context, repository entity.OrderRepositoryInterface) *OrderUseCase {
	return &OrderUseCase{
		ctx:        ctx,
		repository: repository,
	}
}

func (o *OrderUseCase) Save(input OrderInputDTO) (OrderOutputDTO, error) {
	order, err := entity.NewOrder(input.Price, input.Tax)

	if err != nil {
		return OrderOutputDTO{}, nil
	}

	err = order.CalculateFinalPrice()

	if err != nil {
		return OrderOutputDTO{}, nil
	}

	if err := o.repository.Save(o.ctx, order); err != nil {
		return OrderOutputDTO{}, nil
	}

	return OrderOutputDTO{
		ID:         order.ID.String(),
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}

func (o *OrderUseCase) List() ([]OrderOutputDTO, error) {
	orders, err := o.repository.List(o.ctx)

	if err != nil {
		return []OrderOutputDTO{}, err
	}

	orderOutput := []OrderOutputDTO{}

	for _, order := range orders {
		orderOutput = append(orderOutput, OrderOutputDTO{
			ID:         order.ID.String(),
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return orderOutput, nil
}
