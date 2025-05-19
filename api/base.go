package api

import (
	"github.com/gin-gonic/gin"
	"im-chat/utils"
)

type BaseApi struct{}

func (b *BaseApi) GetUserID(c *gin.Context) int {
	return c.GetInt(utils.ContextUserIDKey)
}

func (b *BaseApi) GetUserName(c *gin.Context) string {
	return c.GetString(utils.ContextUserNameKey)
}

func (b *BaseApi) GetPhone(c *gin.Context) string {
	return c.GetString(utils.ContextPhoneKey)
}
