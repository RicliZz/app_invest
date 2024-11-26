package Utils

import "github.com/gin-gonic/gin"

func GetUserFromContext(c *gin.Context) int64 {
	userID, err := c.Get("userID")
	if !err {
		return 0
	}
	return userID.(int64)
}
