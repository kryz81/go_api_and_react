package models

import "time"

type AddUserDto struct {
	Name string `json:"name" binding:"required,min=3"`
	Age  int    `json:"age" binding:"required,gt=0"`
}

type User struct {
	Id        string    `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	Age       int       `json:"age" bson:"age"`
	CreatedAt time.Time `json:"createdAt" bson:"created_at"`
}
