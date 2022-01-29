package v1

import "github.com/gin-gonic/gin"

type WelcomeRoutes struct {
	RouterGroup *gin.RouterGroup
}

func (rg WelcomeRoutes) RegisterRouterGroup() *gin.RouterGroup {
	rg.RouterGroup.GET("/welcome", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "[GET] Welcome v1",
			"error":   false,
		})
	})
	rg.RouterGroup.POST("/welcome", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "[POST] Welcome v1",
			"error":   false,
		})
	})
	rg.RouterGroup.PUT("/welcome", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "[PUT] Welcome v1",
			"error":   false,
		})
	})
	rg.RouterGroup.PATCH("/welcome", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "[PATCH] Welcome v1",
			"error":   false,
		})
	})
	rg.RouterGroup.DELETE("/welcome", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "[DELETE] Welcome v1",
			"error":   false,
		})
	})

	return rg.RouterGroup
}
