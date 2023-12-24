package routes

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sandronister/mongo-go/internal/infra/di"
	"go.mongodb.org/mongo-driver/mongo"
)

func orderHandler(ctx context.Context, r *gin.Engine, db *mongo.Database) {
	orderHandler := di.ConfigOrderHandlerDI(ctx, db)

	orderRoute := r.Group("orders")
	{
		orderRoute.POST("", orderHandler.Create)
		orderRoute.GET("", orderHandler.List)
	}
}

func HandlerRequest(ctx context.Context, port string, db *mongo.Database) {
	r := gin.Default()
	orderHandler(ctx, r, db)
	r.Run(port)
}
