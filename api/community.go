package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mundo-wang/wtool/wlog"
	"im-chat/service"
	"strconv"
)

type CommunityApi struct {
	service.CommunityService
}

func GetCommunityApi() *CommunityApi {
	return &CommunityApi{}
}

func (api *CommunityApi) LoadByUserId(c *gin.Context) (interface{}, error) {
	userIdStr := c.Query("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		wlog.Error("call strconv.Atoi failed").Err(err).Field("userId", userIdStr).Log()
		return nil, err
	}
	communities, err := api.CommunityService.LoadByUserId(userId)
	if err != nil {
		wlog.Error("call api.CommunityService.LoadByUserId failed").Err(err).Field("userId", userId).Log()
		return nil, err
	}
	return communities, nil
}
