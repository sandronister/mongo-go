package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandronister/mongo-go/internal/entity"
	"github.com/sandronister/mongo-go/internal/usecase"
)

type OrderHandler struct {
	repository entity.OrderRepositoryInterface
	ctx        context.Context
}

func NewOrderHandler(ctx context.Context, repository entity.OrderRepositoryInterface) *OrderHandler {
	return &OrderHandler{
		ctx:        ctx,
		repository: repository,
	}
}

func (o *OrderHandler) Create(c *gin.Context) {
	var orderDTO usecase.OrderInputDTO

	if err := c.ShouldBindJSON(&orderDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	usecase := usecase.NewOrderUseCase(o.ctx, o.repository)
	output, err := usecase.Save(orderDTO)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, output)
}

func (o *OrderHandler) List(c *gin.Context) {
	usecase := usecase.NewOrderUseCase(o.ctx, o.repository)
	output, err := usecase.List()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}
