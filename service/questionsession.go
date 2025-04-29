package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/mundo-wang/wtool/wlog"
	"gorm.io/gorm"
	"im-chat/code"
	"im-chat/dao/model"
	"im-chat/utils"
	"time"
)

type QuestionSessionService struct {
}

func (q *QuestionSessionService) GenerateQuestions(c *gin.Context, req *GenerateQuestionsReq) error {
	sessionId := uuid.NewString()
	questionSession := &model.QuestionSession{
		SessionID:  sessionId,
		Status:     utils.GenerateStatusGenerating,
		PositionID: req.PositionId,
		AgentCode:  req.AgentCode,
		UserID:     utils.GetUserId(c),
		Username:   utils.GetUserName(c),
		CreatedBy:  utils.GetUserName(c),
	}
	err := questionSessionQ.Create(questionSession)
	if err != nil {
		wlog.Error("call questionSessionQ.Create failed").Err(err).Field("questionSession", questionSession).Log()
		return err
	}
	question := &model.Questions{
		Title:        "测试题目",
		Type:         req.Type,
		Answer:       "A",
		Status:       utils.QuestionStatusUnpublished,
		AgentCode:    req.AgentCode,
		PositionID:   req.PositionId,
		SessionRefID: sessionId,
		CreatedBy:    utils.GetUserName(c),
	}
	err = questionsQ.Create(question)
	if err != nil {
		wlog.Error("call questionsQ.Create failed").Err(err).Field("question", question).Log()
		return err
	}
	options := []*model.QuestionOptions{
		{QuestionID: question.ID, OptionKey: "A", OptionText: "测试选项A"},
		{QuestionID: question.ID, OptionKey: "B", OptionText: "测试选项B"},
		{QuestionID: question.ID, OptionKey: "C", OptionText: "测试选项C"},
		{QuestionID: question.ID, OptionKey: "D", OptionText: "测试选项D"},
	}
	err = questionOptionsQ.Create(options...)
	if err != nil {
		wlog.Error("call questionOptionsQ.Create failed").Err(err).Field("options", options).Log()
		return err
	}
	resp := &GenerateQuestionsResp{
		ID:     question.ID,
		Title:  question.Title,
		Answer: question.Answer,
		Options: []Options{
			{OptionKey: "A", OptionText: "测试选项A"},
			{OptionKey: "B", OptionText: "测试选项B"},
			{OptionKey: "C", OptionText: "测试选项C"},
			{OptionKey: "D", OptionText: "测试选项D"},
		},
	}
	c.SSEvent("resp", resp)
	c.Writer.Flush()
	_, err = questionSessionQ.Where(questionSessionQ.ID.Eq(questionSession.ID)).
		UpdateColumn(questionSessionQ.Status, utils.GenerateStatusGenerated)
	if err != nil {
		wlog.Error("call questionSessionQ.UpdateColumn failed").Err(err).Field("sessionId", sessionId).Log()
		return err
	}
	return nil
}

func (q *QuestionSessionService) CheckUnpublishedSession(c *gin.Context) string {
	questionSession, err := questionSessionQ.Where(
		questionSessionQ.Status.In(utils.GenerateStatusGenerating, utils.GenerateStatusGenerated),
		questionSessionQ.CreatedBy.Eq(utils.GetUserName(c))).Order(questionSessionQ.CreatedAt.Desc()).First()
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return ""
	} else {
		return questionSession.SessionID
	}
}

func (q *QuestionSessionService) GetSessionQuestions(sessionId string) ([]GetSessionQuestionsResp, error) {
	questionSession, err := questionSessionQ.Where(questionSessionQ.SessionID.Eq(sessionId)).First()
	if err != nil {
		wlog.Error("call questionSessionQ.First failed").Err(err).Field("sessionId", sessionId).Log()
		return nil, err
	}
	if questionSession.Status == utils.GenerateStatusGenerating {
		createdAt := questionSession.CreatedAt
		if time.Now().Sub(createdAt) > 10*time.Minute {
			questionSession.Status = utils.GenerateStatusFailed
			_, err = questionSessionQ.Updates(questionSession)
			if err != nil {
				wlog.Error("call questionSessionQ.Updates failed").Err(err).Field("questionSession", questionSession).Log()
				return nil, err
			}
		}
	}
	if questionSession.Status == utils.GenerateStatusFailed {
		return nil, code.QuestionGenerating
	}
	respList := make([]GetSessionQuestionsResp, 0)
	questionList, err := questionsQ.Where(questionsQ.SessionRefID.Eq(sessionId), questionsQ.Status.Eq(utils.QuestionStatusUnpublished)).Find()
	if err != nil {
		wlog.Error("call questionsQ.Find failed").Err(err).Field("sessionId", sessionId).Log()
		return nil, err
	}
	for _, question := range questionList {
		resp := GetSessionQuestionsResp{}
		err = copier.Copy(&resp, question)
		if err != nil {
			wlog.Error("call copier.Copy failed").Err(err).Log()
			return nil, err
		}
		optionList, err := questionOptionsQ.Where(questionOptionsQ.QuestionID.Eq(question.ID)).Find()
		if err != nil {
			wlog.Error("call questionOptionsQ.Find failed").Err(err).Field("questionId", question.ID).Log()
			return nil, err
		}
		optionResp := make([]Options, 0)
		err = copier.Copy(&optionResp, optionList)
		if err != nil {
			wlog.Error("call copier.Copy failed").Err(err).Field("optionList", optionList).Log()
			return nil, err
		}
		resp.Options = optionResp
		respList = append(respList, resp)
	}
	return respList, nil
}
