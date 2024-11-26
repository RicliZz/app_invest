package Utils

import "github.com/gin-gonic/gin"

func GetUserIDFromContext(c *gin.Context) int64 {
	userID, isExist := c.Get("userID")
	if !isExist {
		return 0
	}
	return userID.(int64)
}
