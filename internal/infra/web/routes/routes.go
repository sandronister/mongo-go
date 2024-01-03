package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sandronister/mongo-go/internal/infra/di"
	"github.com/sandronister/mongo-go/internal/infra/web/handler"
)

func orderHandler(hc *handler.HandlerConfig, r *gin.Engine) {
	orderHandler := di.ConfigOrderHandlerDI(hc)

	orderRoute := r.Group("orders")
	{
		orderRoute.POST("", orderHandler.Create)
		orderRoute.GET("", orderHandler.List)
	}
}

func HandlerRequest(config *handler.HandlerConfig) {
	r := gin.Default()
	orderHandler(config, r)
	r.Run(config.Port)
}
