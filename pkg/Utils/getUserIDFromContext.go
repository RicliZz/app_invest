package Utils

import "github.com/gin-gonic/gin"

func GetUserFromContext(c *gin.Context) (int64, any) {
	userID, isExist := c.Get("userID")
	if !isExist {
		return 0, ""
	}
	role, isExist := c.Get("role")
	if !isExist {
		return 0, ""
	}
	return userID.(int64), role
}
