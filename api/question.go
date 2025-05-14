package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mundo-wang/wtool/wlog"
	"im-chat/service"
	"strconv"
)

type QuestionApi struct {
	service.QuestionService
}

func GetQuestionApi() *QuestionApi {
	return &QuestionApi{}
}

func (api *QuestionApi) GetQuestionsPage(c *gin.Context) (interface{}, error) {
	req := &service.GetQuestionsPageReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		wlog.Error("call c.ShouldBindJSON failed").Err(err).Log()
		return nil, err
	}
	page, err := api.QuestionService.GetQuestionsPage(req)
	if err != nil {
		wlog.Error("call api.QuestionService.GetQuestionsPage failed").Err(err).Field("req", req).Log()
		return nil, err
	}
	return page, nil
}

func (api *QuestionApi) GetQuestionInfo(c *gin.Context) (interface{}, error) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		wlog.Error("call strconv.Atoi failed").Err(err).Field("idStr", idStr).Log()
		return nil, err
	}
	resp, err := api.QuestionService.GetQuestionInfo(id)
	if err != nil {
		wlog.Error("call api.QuestionService.GetQuestionInfo failed").Err(err).Field("id", id).Log()
		return nil, err
	}
	return resp, nil
}

func (api *QuestionApi) UpdateQuestion(c *gin.Context) (interface{}, error) {
	req := &service.UpdateQuestionReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		wlog.Error("call c.ShouldBindJSON failed").Err(err).Log()
		return nil, err
	}
	err = api.QuestionService.UpdateQuestion(req)
	if err != nil {
		wlog.Error("call api.QuestionService.UpdateQuestion failed").Err(err).Field("req", req).Log()
		return nil, err
	}
	return nil, nil
}

func (api *QuestionApi) DeleteQuestion(c *gin.Context) (interface{}, error) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		wlog.Error("call strconv.Atoi failed").Err(err).Field("idStr", idStr).Log()
		return nil, err
	}
	err = api.QuestionService.DeleteQuestion(id)
	if err != nil {
		wlog.Error("call api.QuestionService.DeleteQuestion failed").Err(err).Field("id", id).Log()
		return nil, err
	}
	return nil, nil
}

func (api *QuestionApi) FetchRandomQuestions(c *gin.Context) (interface{}, error) {
	req := &service.FetchRandomQuestionsReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		wlog.Error("call c.ShouldBindQuery failed").Err(err).Log()
		return nil, err
	}
	resp, err := api.QuestionService.FetchRandomQuestions(req.PositionId, req.Count)
	if err != nil {
		wlog.Error("call api.QuestionService.FetchRandomQuestions failed").Err(err).Field("req", req).Log()
		return nil, err
	}
	return resp, nil
}

func (api *QuestionApi) CalculateScore(c *gin.Context) (interface{}, error) {
	req := &service.CalculateScoreReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		wlog.Error("call c.ShouldBindJSON failed").Err(err).Log()
		return nil, err
	}
	resp, err := api.QuestionService.CalculateScore(c, req)
	if err != nil {
		wlog.Error("call api.QuestionService.CalculateScore failed").Err(err).Field("req", req).Log()
		return nil, err
	}
	return resp, nil
}
