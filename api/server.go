package api

import (
	"github.com/auth-service/db"
	"github.com/auth-service/types"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

func StartServer() {
	r := gin.Default()
	// User Group
	userGroup := r.Group("/user")
	{
		userGroup.POST("/registration", func(c *gin.Context) {
			var user types.UserRegister
			if err := c.ShouldBindJSON(&user); err != nil {
				log.Err(err).Msgf("reqbody json binding failed")
				c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
				return
			}
			err := RegisterNewUser(user)
			if err != nil {
				log.Err(err).Msgf("register user failed")
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
				return
			}
			log.Info().Msgf("user registration succeeded")
			c.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": "user registration success",
			})
			return
		})

		userGroup.POST("/login", func(c *gin.Context) {
			var user types.UserRegister
			if err := c.ShouldBindJSON(&user); err != nil {
				log.Err(err).Msgf("reqbody json binding failed")
				c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
				return
			}

			access, err := UserLogin(user)
			if err != nil {
				log.Err(err).Msgf("login user failed")
				c.JSON(http.StatusOK, gin.H{
					"status": "failed",
					"error":  err.Error(),
				})
				return
			}
			log.Info().Msgf("user login succeeded")
			c.JSON(http.StatusOK, gin.H{
				"message":      "login success",
				"access_token": access,
			})
		})
	}

	// Admin Group
	adminGroup := r.Group("/admin")
	{

		adminGroup.POST("/registration", func(c *gin.Context) {
			var admin types.AdminCredential
			if err := c.ShouldBindJSON(&admin); err != nil {
				log.Err(err).Msgf("reqbody json binding failed")
				c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
				return
			}
			err := RegisterNewAdmin(admin)
			if err != nil {
				log.Err(err).Msgf("admin registration failed")
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": "please wait until super admin accept your request",
			})
			return
		})

		adminGroup.POST("/login", func(c *gin.Context) {
			var admin types.AdminCredential
			if err := c.ShouldBindJSON(&admin); err != nil {
				log.Err(err).Msgf("reqbody json binding failed")
				c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
				return
			}

			access, err := AdminLogin(admin)
			if err != nil {
				log.Err(err).Msgf("admin login failed")
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
				return
			}

			log.Info().Msgf("admin login success")
			c.JSON(http.StatusOK, gin.H{
				"message":      "login success",
				"access_token": access,
			})
		})

		adminGroup.GET("/approve-admin", func(c *gin.Context) {
			superAdminaccessToken := c.GetHeader("access-token")
			superAdminUser := c.GetHeader("admin-user")

			if IsAuthorisedAdminRequest(superAdminUser, superAdminaccessToken, true) {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "unauthorized request",
				})
				return
			}

			adminUser := c.Query("admin_user")
			isApproved := c.DefaultQuery("is_approved", "false")
			if isApproved == "false" {
				c.JSON(http.StatusOK, gin.H{
					"message": "admin not approved",
				})
				return
			}

			err := ApproveAdmin(adminUser)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"status":  "failed",
					"message": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "admin approved",
			})

		})

		adminGroup.GET("/find-user", func(c *gin.Context) {
			// authorization admin
			accessToken := c.GetHeader("access-token")
			adminUser := c.GetHeader("admin-user")

			if IsAuthorisedAdminRequest(adminUser, accessToken, false) {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "unauthorized request",
				})
				return
			}

			// if req contain username then it will return the user details
			userName := c.Query("user_name")
			if userName != "" {

				user, err := db.GetUserCredentialsByUserName(userName)
				if err != nil {
					log.Err(err).Msgf("fetching user  failed")
					c.JSON(http.StatusBadRequest, gin.H{
						"message": "unable to proceed",
					})
					return
				}
				log.Err(err).Msgf("fetching users with pagination failed")
				c.JSON(http.StatusOK, gin.H{
					"message": "success",
					"Data":    user,
				})

				// else it will return all users page by page
			} else {
				limitStr := c.DefaultQuery("limit", "10")
				pageStr := c.DefaultQuery("page", "1")

				limit := 10
				l, err := strconv.Atoi(limitStr)
				if err == nil && l != 0 {
					limit = l
				}

				page := 1
				p, err := strconv.Atoi(pageStr)
				if err == nil && p != 0 {
					page = p
				}

				users, err := db.GetUsersWithPagination(limit, page)
				if err != nil {
					log.Err(err).Msgf("fetching users with pagination failed")
					c.JSON(http.StatusBadRequest, gin.H{
						"message": "unable to proceed",
					})
					return
				}
				log.Err(err).Msgf("fetching users with pagination failed")

				c.JSON(http.StatusOK, gin.H{
					"message": "success",
					"Data":    users,
				})

			}

		})
	}

	r.Run(":8080")
}
