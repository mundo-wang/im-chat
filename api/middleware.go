package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mundo-wang/wtool/wlog"
	"github.com/mundo-wang/wtool/wresp"
	"im-chat/utils"
)

func CheckAuthorization(c *gin.Context) error {
	jwtToken := c.GetHeader(utils.Authorization)
	userClaims, err := utils.ParseJwtToken(jwtToken)
	if err != nil {
		if !wresp.IsErrorCode(err) {
			wlog.Error("call utils.ParseJwtToken failed").Err(err).Field("jwtToken", jwtToken).Log()
		}
		return err
	}
	c.Set("userId", userClaims.UserID)
	c.Set("userName", userClaims.UserName)
	c.Set("phone", userClaims.Phone)
	return nil
}
