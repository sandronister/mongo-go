package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func (o *OrderHandler) Create(c *gin.Context) {
	var orderDTO usecase.OrderInputDTO

	if err := c.ShouldBindJSON(&orderDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	output, err := o.orderUseCase.Save(orderDTO)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, output)
}

func (o *OrderHandler) List(c *gin.Context) {
	output, err := o.orderUseCase.List()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}
