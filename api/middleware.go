package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mundo-wang/wtool/wlog"
	"github.com/mundo-wang/wtool/wresp"
	"im-chat/code"
	"im-chat/utils"
)

func CheckAuthorization(c *gin.Context) error {
	jwtToken := c.GetHeader(utils.Authorization)
	if jwtToken == "" {
		return code.JwtTokenNotExist
	}
	userClaims, err := utils.ParseJwtToken(jwtToken)
	if err != nil {
		if !wresp.IsErrorCode(err) {
			wlog.Error("call utils.ParseJwtToken failed").Err(err).Field("jwtToken", jwtToken).Log()
		}
		return err
	}
	c.Set(utils.ContextUserIDKey, userClaims.UserID)
	c.Set(utils.ContextUserNameKey, userClaims.UserName)
	c.Set(utils.ContextPhoneKey, userClaims.Phone)
	return nil
}
