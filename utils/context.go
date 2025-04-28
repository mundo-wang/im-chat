package utils

import "github.com/gin-gonic/gin"

func GetUserId(c *gin.Context) int {
	return c.GetInt(ContextUserIDKey) // GetInt方法，如果没有对应key，或者断言int失败，都会返回默认值0
}

func GetUserName(c *gin.Context) string {
	return c.GetString(ContextUserNameKey)
}

func getPhone(c *gin.Context) string {
	return c.GetString(ContextPhoneKey)
}
