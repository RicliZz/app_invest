package Utils

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AdminOnlyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, role, err := ExtractUserFromJWT(c)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}
		if role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{})
			return
		}
		c.Next()
	}
}
