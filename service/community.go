package service

import (
	"github.com/jinzhu/copier"
	"github.com/mundo-wang/wtool/wlog"
	"im-chat/utils"
)

type CommunityService struct {
}

func (c *CommunityService) LoadByUserId(userId int) ([]*LoadByUserIdResp, error) {
	communityIds := make([]int, 0)
	err := contactsQ.Select(contactsQ.TargetID).Where(contactsQ.OwnerID.Eq(userId),
		contactsQ.Type.Eq(utils.ContactTypeGroup)).Scan(&communityIds)
	if err != nil {
		wlog.Error("call contactsQ.Scan failed").Err(err).Field("userId", userId).Log()
		return nil, err
	}
	communities, err := communitiesQ.Where(communitiesQ.ID.In(communityIds...)).Find()
	if err != nil {
		wlog.Error("call communitiesQ.Find failed").Err(err).Field("communityIds", communityIds).Log()
		return nil, err
	}
	resp := make([]*LoadByUserIdResp, 0)
	err = copier.Copy(&resp, communities)
	if err != nil {
		wlog.Error("call copier.Copy failed").Err(err).Field("communities", communities).Log()
		return nil, err
	}
	return resp, nil
}
