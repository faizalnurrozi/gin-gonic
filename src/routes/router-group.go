package routes

import (
	"gin-gonic/src/routes/v1"
	v2 "gin-gonic/src/routes/v2"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	Router *gin.Engine
}

func (r RouterGroup) RegisterRoute() {
	routesV1 := r.Router.Group("/v1")
	{
		v1.WelcomeRoutes{RouterGroup: routesV1}.RegisterRouterGroup()
		v1.WorldRoutes{RouterGroup: routesV1}.RegisterRouterGroup()
	}

	routesV2 := r.Router.Group("/v2")
	{
		v2.WelcomeRoutes{RouterGroup: routesV2}.RegisterRouterGroup()
		v2.WorldRoutes{RouterGroup: routesV2}.RegisterRouterGroup()
	}
}
