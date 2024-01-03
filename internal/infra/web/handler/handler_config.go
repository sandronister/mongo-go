package handler

import (
	"context"

	"github.com/sandronister/mongo-go/configs"
	"go.mongodb.org/mongo-driver/mongo"
)

type HandlerConfig struct {
	Ctx  *context.Context
	Db   *mongo.Database
	Port string
}

func NewHandlerConfig(ctx *context.Context, config *configs.Conf, db *mongo.Database) *HandlerConfig {
	return &HandlerConfig{
		Ctx:  ctx,
		Port: config.WebPort,
		Db:   db,
	}
}
