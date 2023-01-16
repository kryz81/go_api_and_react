package types

import (
	"context"
	"github.com/kryz81/go_api_and_react/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceContainer struct {
	UserService services.UserService
}

type AppContext struct {
	Db       *mongo.Database
	Ctx      context.Context
	Services ServiceContainer
}

type BodyError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
