package main

import (
	"github.com/sandronister/mongo-go/configs"
	"github.com/sandronister/mongo-go/internal/infra/database/connection"
	"github.com/sandronister/mongo-go/internal/infra/di"
	"github.com/sandronister/mongo-go/internal/infra/web/server"
	"github.com/sandronister/mongo-go/pkg/errors"
)

func main() {
	config, err := configs.LoadConfig("../../")
	errors.Catch(err)
	database, err := connection.GetDatabase(config)
	errors.Catch(err)

	orderUseCase := di.OrderUseCaseDI(database)

	webServer := server.NewWebServer(config.WebPort)
	webServer.AddHandlerOrder(orderUseCase)
	webServer.Start()

}
