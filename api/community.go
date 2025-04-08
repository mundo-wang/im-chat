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

func (api *CommunityApi) Create(c *gin.Context) (interface{}, error) {
	userId := c.GetInt("userID") // GetInt方法，如果没有对应key，或者断言int失败，都会返回默认值0
	req := &service.CreateCommunityReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		wlog.Error("call c.ShouldBindJSON failed").Err(err).Log()
		return nil, err
	}
	req.OwnerID = userId
	err = api.CommunityService.Create(req)
	if err != nil {
		wlog.Error("call api.CommunityService.Create failed").Err(err).Field("req", req).Log()
		return nil, err
	}
	return nil, nil
}
