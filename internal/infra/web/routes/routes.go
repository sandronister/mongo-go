package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sandronister/mongo-go/internal/infra/di"
	"go.mongodb.org/mongo-driver/mongo"
)

func orderHandler(r *gin.Engine, db *mongo.Database) {
	orderHandler := di.ConfigOrderHandlerDI(db)

	orderRoute := r.Group("orders")
	{
		orderRoute.POST("", orderHandler.Create)
		orderRoute.GET("", orderHandler.List)
	}
}

func HandlerRequest(port string, db *mongo.Database) {
	r := gin.Default()
	orderHandler(r, db)
	r.Run(port)
}
