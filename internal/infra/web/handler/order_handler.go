package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sandronister/mongo-go/internal/usecase"
)

type OrderHandler struct {
	orderUseCase *usecase.OrderUseCase
}

func NewOrderHandler(orderUseCase *usecase.OrderUseCase) *OrderHandler {
	return &OrderHandler{
		orderUseCase: orderUseCase,
	}
}

func (o *OrderHandler) Create(c *fiber.Ctx) error {
	var orderDTO usecase.OrderInputDTO

	if err := c.BodyParser(&orderDTO); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error()})
	}

	output, err := o.orderUseCase.Save(orderDTO)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(output)
}

func (o *OrderHandler) List(c *fiber.Ctx) error {
	output, err := o.orderUseCase.List()

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(output)
}
