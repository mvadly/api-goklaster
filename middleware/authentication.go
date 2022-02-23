package middleware

import (
	"api-goklaster/util"

	"github.com/gin-gonic/gin"
)

func TokenValidator() gin.HandlerFunc {
	return func(c *gin.Context) {

		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(401, gin.H{"status": false, "message": "Authorization is invalid"})
			c.AbortWithStatus(401)
			return
		}

		lenBearer := len(BEARER_SCHEMA)
		tokenString := authHeader[lenBearer:]
		_, err := util.CheckToken(tokenString)

		if err != nil {
			c.JSON(401, gin.H{"status": false, "message": err.Error()})
			c.AbortWithStatus(401)
			return
		}

		c.Next()
	}
}
