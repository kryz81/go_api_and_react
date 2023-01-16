package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kryz81/go_api_and_react/models"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func (handlerCtx HandlerContext) UsersListHandler(ctx *gin.Context) {
	users, err := handlerCtx.Services.UserService.FindUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (handlerCtx HandlerContext) UserDetailsHandler(ctx *gin.Context) {
	user, err := handlerCtx.Services.UserService.FindUserById(ctx.Param("id"))
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("user with ID %s not found", ctx.Param("id"))})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (handlerCtx HandlerContext) UserAddHandler(ctx *gin.Context) {
	var userDto models.AddUserDto
	ctx.ShouldBindJSON(&userDto)

	result, err := handlerCtx.Services.UserService.AddUser(userDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (handlerCtx HandlerContext) UserUpdateHandler(ctx *gin.Context) {

	var userDto models.AddUserDto
	ctx.ShouldBindJSON(&userDto)

	result, err := handlerCtx.Services.UserService.AddUser(userDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (handlerCtx HandlerContext) UserDeleteHandler(ctx *gin.Context) {
	_, err := handlerCtx.Services.UserService.DeleteUser(ctx.Param("id"))
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("user with ID %s not found", ctx.Param("id"))})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
