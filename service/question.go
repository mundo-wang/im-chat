package service

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/mundo-wang/wtool/wlog"
	"im-chat/utils"
	"math"
)

type QuestionService struct {
}

func (q *QuestionService) GetQuestionsPage(req *GetQuestionsPageReq) (*utils.PageResult[GetQuestionPageResp], error) {
	questionsCond := questionsQ.WithContext(context.Background())
	questionsCond = questionsCond.Where(questionsQ.Status.Eq(utils.QuestionStatusPublished))
	if req.PositionId != 0 {
		questionsCond = questionsCond.Where(questionsQ.PositionID.Eq(req.PositionId))
	}
	if req.Type != 0 {
		questionsCond = questionsCond.Where(questionsQ.Type.Eq(req.Type))
	}
	if !req.OperationTimeStart.IsZero() && !req.OperationTimeEnd.IsZero() {
		questionsCond = questionsCond.Where(questionsQ.UpdatedAt.Between(req.OperationTimeStart, req.OperationTimeEnd))
	}
	offset := (req.Page - 1) * req.Size
	questionPage, _, err := questionsCond.FindByPage(offset, req.Size)
	if err != nil {
		wlog.Error("call questionsCond.FindByPage failed").Err(err).Field("req", req).Log()
		return nil, err
	}
	respList := make([]GetQuestionPageResp, 0)
	for _, question := range questionPage {
		resp := GetQuestionPageResp{}
		err = copier.Copy(&resp, question)
		if err != nil {
			wlog.Error("call copier.Copy failed").Err(err).Field("question", question).Log()
			return nil, err
		}
		optionList := make([]Options, 0)
		options, err := questionOptionsQ.Where(questionOptionsQ.QuestionID.Eq(question.ID)).Find()
		if err != nil {
			wlog.Error("call questionOptionsQ.Find failed").Err(err).Field("questionId", question.ID).Log()
			return nil, err
		}
		err = copier.Copy(&optionList, options)
		if err != nil {
			wlog.Error("call copier.Copy failed").Err(err).Field("options", options).Log()
			return nil, err
		}
		resp.Options = optionList
		respList = append(respList, resp)
	}
	count, err := questionsCond.Count()
	if err != nil {
		wlog.Error("call questionsQB.Count failed").Log()
		return nil, err
	}
	pageCount := int(math.Ceil(float64(count) / float64(req.Size)))
	page := &utils.PageResult[GetQuestionPageResp]{
		Page:      req.Page,
		Size:      req.Size,
		Total:     count,
		PageCount: pageCount,
		Records:   respList,
	}
	return page, nil
}
