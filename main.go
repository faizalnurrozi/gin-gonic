package main

import (
	"gin-gonic/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.RouterGroup{Router: router}.RegisterRoute()

	_ = router.Run()
}
