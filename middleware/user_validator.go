package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kryz81/go_api_and_react/models"
	"github.com/kryz81/go_api_and_react/utils"
	"net/http"
)

func UserValidator(ctx *gin.Context) {
	var userDto models.AddUserDto
	if err := ctx.ShouldBindJSON(&userDto); err != nil {
		var validationErrors validator.ValidationErrors
		var errorMessage any
		if errors.As(err, &validationErrors) {
			errorMessage = utils.ExtractValidationErrors(validationErrors)
		} else {
			errorMessage = err.Error()
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errorMessage})
		return
	}
	ctx.Next()
}
