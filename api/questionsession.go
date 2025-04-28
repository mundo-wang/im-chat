package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mundo-wang/wtool/wlog"
	"github.com/mundo-wang/wtool/wresp"
	"im-chat/service"
)

type QuestionSessionApi struct {
	service.QuestionSessionService
}

func GetQuestionSessionApi() *QuestionSessionApi {
	return &QuestionSessionApi{}
}

func (api *QuestionSessionApi) GenerateQuestions(c *gin.Context) error {
	req := &service.GenerateQuestionsReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		wlog.Error("call c.ShouldBindQuery failed").Err(err).Log()
		return err
	}
	err = api.QuestionSessionService.GenerateQuestions(c, req)
	if err != nil {
		if !wresp.IsErrorCode(err) {
			wlog.Error("call api.QuestionSessionService.GenerateQuestions failed").Err(err).Field("req", req).Log()
		}
		return err
	}
	return nil
}
