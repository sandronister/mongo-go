package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sandronister/mongo-go/internal/infra/web/handler"
	"github.com/sandronister/mongo-go/internal/usecase"
)

type WebServer struct {
	router        *gin.Engine
	webServerPort string
}

func NewWebServer(port string) *WebServer {
	webserver := &WebServer{
		router:        gin.Default(),
		webServerPort: port,
	}

	return webserver
}

func (w *WebServer) AddHandlerOrder(order *usecase.OrderUseCase) {
	orderHandler := handler.NewOrderHandler(order)

	orderRoute := w.router.Group("orders")
	{
		orderRoute.POST("", orderHandler.Create)
		orderRoute.GET("", orderHandler.List)
	}
}

func (w *WebServer) Start() {
	w.router.Run(w.webServerPort)
}
