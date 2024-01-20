package handler

import (
	"github.com/sandronister/mongo-go/configs"
	"go.mongodb.org/mongo-driver/mongo"
)

type HandlerConfig struct {
	Db   *mongo.Database
	Port string
}

func NewHandlerConfig(config *configs.Conf, db *mongo.Database) *HandlerConfig {
	return &HandlerConfig{
		Port: config.WebPort,
		Db:   db,
	}
}
