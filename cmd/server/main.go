package main

import (
	"context"

	"github.com/sandronister/mongo-go/configs"
	"github.com/sandronister/mongo-go/internal/infra/database/connection"
	"github.com/sandronister/mongo-go/internal/infra/web/handler"
	"github.com/sandronister/mongo-go/internal/infra/web/routes"
	"github.com/sandronister/mongo-go/pkg/errors"
)

func main() {
	ctx := context.Background()
	config, err := configs.LoadConfig("../../")
	errors.Catch(err)
	database, err := connection.GetDatabase(&ctx, config)
	errors.Catch(err)
	routes.HandlerRequest(handler.NewHandlerConfig(&ctx, config, database))
}
