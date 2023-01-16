package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/kryz81/go_api_and_react/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserService struct {
	Collection *mongo.Collection
	Ctx        context.Context
}

func (service UserService) FindUsers() ([]models.User, error) {
	users := make([]models.User, 0)
	cursor, err := service.Collection.Find(service.Ctx, bson.M{})
	defer cursor.Close(service.Ctx)
	if err != nil {
		return []models.User{}, err
	}

	var user models.User
	for cursor.Next(service.Ctx) {
		cursor.Decode(&user)
		users = append(users, user)
	}

	return users, nil
}

func (service UserService) FindUserById(id string) (models.User, error) {
	result := service.Collection.FindOne(service.Ctx, bson.M{"_id": id})
	if result.Err() == mongo.ErrNoDocuments {
		return models.User{}, result.Err()
	}
	var user models.User
	result.Decode(&user)

	return user, nil
}

func (service UserService) AddUser(userDto models.AddUserDto) (models.User, error) {
	user := models.User{
		Id:        uuid.NewString(),
		Name:      userDto.Name,
		Age:       userDto.Age,
		CreatedAt: time.Now(),
	}

	_, err := service.Collection.InsertOne(service.Ctx, user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (service UserService) DeleteUser(id string) (bool, error) {
	_, err := service.FindUserById(id)
	if err != nil {
		return false, err
	}

	result, err := service.Collection.DeleteOne(service.Ctx, bson.M{"_id": id})
	if err != nil {
		return false, err
	}

	return result.DeletedCount == 1, nil
}