package v2

import (
	"gin-gonic/src/usecases"
	"github.com/gin-gonic/gin"
)

type WorldUseCase struct {
	*usecases.Contract
}

func NewWorldUseCase() *WorldUseCase {
	return &WorldUseCase{}
}

func (uc WorldUseCase) WorldGet(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "[GET] World v2",
		"error":   false,
	})
}

func (uc WorldUseCase) WorldPost(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "[POST] World v2",
		"error":   false,
	})
}

func (uc WorldUseCase) WorldPut(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "[PUT] World v2",
		"error":   false,
	})
}

func (uc WorldUseCase) WorldPatch(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "[Patch] World v2",
		"error":   false,
	})
}

func (uc WorldUseCase) WorldDelete(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "[DELETE] World v2",
		"error":   false,
	})
}
