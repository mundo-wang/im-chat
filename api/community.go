package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mundo-wang/wtool/wlog"
	"github.com/mundo-wang/wtool/wresp"
	"im-chat/service"
	"im-chat/utils"
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
	userId := utils.GetUserId(c)
	req := &service.CreateCommunityReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		wlog.Error("call c.ShouldBindJSON failed").Err(err).Log()
		return nil, err
	}
	req.OwnerID = userId
	err = api.CommunityService.Create(req)
	if err != nil {
		if !wresp.IsErrorCode(err) {
			wlog.Error("call api.CommunityService.Create failed").Err(err).Field("req", req).Log()
		}
		return nil, err
	}
	return nil, nil
}

func (api *CommunityApi) JoinGroup(c *gin.Context) (interface{}, error) {
	userId := utils.GetUserId(c)
	groupCode := c.Query("groupCode")
	err := api.CommunityService.JoinGroup(groupCode, userId)
	if err != nil {
		if !wresp.IsErrorCode(err) {
			wlog.Error("call api.CommunityService.JoinGroup failed").Err(err).Field("groupCode", groupCode).Log()
		}
		return nil, err
	}
	return nil, nil
}
