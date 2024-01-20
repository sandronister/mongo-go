package main

import (
	"github.com/sandronister/mongo-go/configs"
	"github.com/sandronister/mongo-go/internal/infra/database/connection"
	"github.com/sandronister/mongo-go/internal/infra/web/handler"
	"github.com/sandronister/mongo-go/internal/infra/web/routes"
	"github.com/sandronister/mongo-go/pkg/errors"
)

func main() {
	config, err := configs.LoadConfig("../../")
	errors.Catch(err)
	database, err := connection.GetDatabase(config)
	errors.Catch(err)
	routes.HandlerRequest(handler.NewHandlerConfig(config, database))
}
