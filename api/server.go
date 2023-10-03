package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartServer() {
	r := gin.Default()
	// User Group
	userGroup := r.Group("/user")
	{
		userGroup.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "User API Home",
			})
		})

		userGroup.GET("/profile", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "User Profile API",
			})
		})
	}

	// Admin Group
	adminGroup := r.Group("/admin")
	{
		adminGroup.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Admin API Home",
			})
		})

		adminGroup.GET("/dashboard", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Admin Dashboard API",
			})
		})
	}

	r.Run(":8080")
}
