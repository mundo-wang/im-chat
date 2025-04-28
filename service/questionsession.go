package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mundo-wang/wtool/wlog"
	"im-chat/dao/model"
	"im-chat/utils"
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
	c.SSEvent("question", resp)
	c.Writer.Flush()
	return nil
}
