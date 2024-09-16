package middlewares

import (
	"log"
	"main/src/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token = c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized to perform request. Please get a valid API key"})
			return
		}

		// Extract Bearer token
		const bearerPrefix = "Bearer "
		splitToken := strings.Split(token, bearerPrefix)
		var reqToken = splitToken[1]

		if reqToken == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token format. Bearer token required"})
			return
		}

		claims, err := models.DecodeToken(reqToken)
		if err != nil {
			log.Printf("Error decoding token: %v\n", err)
			return
		}

		log.Printf("Claims: %v\n", claims)

		c.Set("userId", 1)

		c.Next()
	}
}

// Register middleware on the base router
func RegisterMiddlewares(router *gin.Engine) {
	router.Use(AuthMiddleware())
}
