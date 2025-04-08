package service

import (
	"github.com/jinzhu/copier"
	"github.com/mundo-wang/wtool/wlog"
	"im-chat/code"
	"im-chat/dao/model"
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

func (c *CommunityService) Create(req *CreateCommunityReq) error {
	count, err := communitiesQ.Where(communitiesQ.Name.Eq(req.Name)).Count()
	if err != nil {
		wlog.Error("call communitiesQ.Count failed").Err(err).Field("req", req).Log()
		return err
	}
	if count != 0 {
		return code.GroupNameAlreadyExist
	}
	community := &model.Communities{}
	err = copier.Copy(community, req)
	if err != nil {
		wlog.Error("call copier.Copy failed").Err(err).Field("req", req).Log()
		return err
	}
	var communityCode string
	for {
		communityCode, err = utils.GenerateRandomDigits(10)
		if err != nil {
			wlog.Error("call utils.GenerateRandomDigits failed").Err(err).Log()
			return err
		}
		count, err = communitiesQ.Where(communitiesQ.CommunityCode.Eq(communityCode)).Count()
		if err != nil {
			wlog.Error("call usersQ.Count failed").Err(err).Field("communityCode", communityCode).Log()
			return err
		}
		if count == 0 {
			break
		}
	}
	community.CommunityCode = communityCode
	err = communitiesQ.Create(community)
	if err != nil {
		wlog.Error("call communitiesQ.Create failed").Err(err).Field("community", community).Log()
		return err
	}
	contact := &model.Contacts{
		OwnerID:  req.OwnerID,
		TargetID: community.ID,
		Type:     utils.ContactTypeGroup,
	}
	err = contactsQ.Create(contact)
	if err != nil {
		wlog.Error("call contactsQ.Create failed").Err(err).Field("contact", contact).Log()
		return err
	}
	return nil
}
