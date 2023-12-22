package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandronister/mongo-go/internal/entity"
	"github.com/sandronister/mongo-go/internal/usecase"
)

type OrderHandler struct {
	repository entity.OrderRepositoryInterface
}

func NewOrderHandler(repository entity.OrderRepositoryInterface) *OrderHandler {
	return &OrderHandler{
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

	usecase := usecase.NewOrderUseCase(c, o.repository)
	output, err := usecase.Save(orderDTO)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, output)
}

func (o *OrderHandler) List(c *gin.Context) {
	usecase := usecase.NewOrderUseCase(c, o.repository)
	output, err := usecase.List()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}
