package api

import (
	"github.com/auth-service/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartServer() {
	r := gin.Default()
	// User Group
	userGroup := r.Group("/user")
	{
		userGroup.POST("/registration", func(c *gin.Context) {
			var user types.UserRegister
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			err := RegisterNewUser(user)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"status": "ok",
					"error":  err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": "user registration success",
			})
			return
		})

		userGroup.POST("/login", func(c *gin.Context) {
			var user types.UserRegister
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			access, err := UserLogin(user)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"status": "failed",
					"error":  err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message":      "login success",
				"access_token": access,
			})
		})
	}

	// Admin Group
	adminGroup := r.Group("/admin")
	{
		adminGroup.GET("/find-user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Admin API Home",
			})
		})
	}

	r.Run(":8080")
}
