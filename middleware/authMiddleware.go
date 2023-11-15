package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	helpers "github.com/1shubham7/jwt/helpers"
)

func Authenticate() gin.HandlerFunc{
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error":fmt.Sprintf("No Authorization header provided")})
			c.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)

	
	}
}