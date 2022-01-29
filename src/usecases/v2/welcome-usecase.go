package v2

import (
	"gin-gonic/src/usecases"
	"github.com/gin-gonic/gin"
)

type WelcomeParam struct {
	Greet string `uri:"greet" binding:"required"`
	Day   string `uri:"day" binding:"required"`
}

type WelcomeUseCase struct {
	*usecases.Contract
}

func NewWelcomeUseCase() *WelcomeUseCase {
	return &WelcomeUseCase{}
}

func (uc WelcomeUseCase) WelcomeGet(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "[GET] Welcome v2",
		"error":   false,
	})
}

func (uc WelcomeUseCase) WelcomePost(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "[POST] Welcome v2",
		"error":   false,
	})
}

func (uc WelcomeUseCase) WelcomePostValidation(context *gin.Context) {
	var param WelcomeParam
	if err := context.ShouldBindJSON(&param); err != nil {
		context.JSON(400, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		context.Abort()
	}
}

func (uc WelcomeUseCase) WelcomePut(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "[PUT] Welcome v2",
		"error":   false,
	})
}

func (uc WelcomeUseCase) WelcomePatch(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "[Patch] Welcome v2",
		"error":   false,
	})
}

func (uc WelcomeUseCase) WelcomeDelete(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "[DELETE] Welcome v2",
		"error":   false,
	})
}
