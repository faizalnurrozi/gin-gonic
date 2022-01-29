package v1

import "github.com/gin-gonic/gin"

type WorldRoutes struct {
	RouterGroup *gin.RouterGroup
}

func (rg WorldRoutes) RegisterRouterGroup() *gin.RouterGroup {
	rg.RouterGroup.GET("/world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "[GET] World v1",
			"error":   false,
		})
	})
	rg.RouterGroup.POST("/world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "[POST] World v1",
			"error":   false,
		})
	})
	rg.RouterGroup.PUT("/world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "[PUT] World v1",
			"error":   false,
		})
	})
	rg.RouterGroup.PATCH("/world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "[PATCH] World v1",
			"error":   false,
		})
	})
	rg.RouterGroup.DELETE("/world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "[DELETE] World v1",
			"error":   false,
		})
	})

	return rg.RouterGroup
}
