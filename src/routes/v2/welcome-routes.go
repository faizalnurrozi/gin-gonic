package v2

import (
	v2 "gin-gonic/src/usecases/v2"
	"github.com/gin-gonic/gin"
)

type WelcomeRoutes struct {
	RouterGroup *gin.RouterGroup
}

func (rg WelcomeRoutes) RegisterRouterGroup() *gin.RouterGroup {
	ucWelcome := v2.NewWelcomeUseCase()

	rg.RouterGroup.GET("/welcome", ucWelcome.WelcomeGet)
	rg.RouterGroup.POST("/welcome", ucWelcome.WelcomePostValidation, ucWelcome.WelcomePost)
	rg.RouterGroup.PUT("/welcome", ucWelcome.WelcomePut)
	rg.RouterGroup.PATCH("/welcome", ucWelcome.WelcomePatch)
	rg.RouterGroup.DELETE("/welcome", ucWelcome.WelcomeDelete)

	return rg.RouterGroup
}
