package v2

import (
	v2 "gin-gonic/src/usecases/v2"
	"github.com/gin-gonic/gin"
)

type WorldRoutes struct {
	RouterGroup *gin.RouterGroup
}

func (rg WorldRoutes) RegisterRouterGroup() *gin.RouterGroup {
	ucWorld := v2.NewWorldUseCase()

	rg.RouterGroup.GET("/world", ucWorld.WorldGet)
	rg.RouterGroup.POST("/world", ucWorld.WorldPost)
	rg.RouterGroup.PUT("/world", ucWorld.WorldPut)
	rg.RouterGroup.PATCH("/world", ucWorld.WorldPatch)
	rg.RouterGroup.DELETE("/world", ucWorld.WorldDelete)

	return rg.RouterGroup
}
