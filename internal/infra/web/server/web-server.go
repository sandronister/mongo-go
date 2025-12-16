package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sandronister/mongo-go/internal/infra/web/handler"
	"github.com/sandronister/mongo-go/internal/usecase"
)

type WebServer struct {
	router        *fiber.App
	webServerPort string
}

func NewWebServer(port string) *WebServer {
	webserver := &WebServer{
		router:        fiber.New(),
		webServerPort: port,
	}

	return webserver
}

func (w *WebServer) AddHandlerOrder(order *usecase.OrderUseCase) {
	orderHandler := handler.NewOrderHandler(order)

	orderRoute := w.router.Group("orders")
	{
		orderRoute.Post("", orderHandler.Create)
		orderRoute.Get("", orderHandler.List)
	}
}

func (w *WebServer) Start() {
	w.router.Listen(fmt.Sprintf(":%s", w.webServerPort))
}
