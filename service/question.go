package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/mundo-wang/wtool/wlog"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"im-chat/dao/model"
	"im-chat/dao/query"
	"im-chat/utils"
	"math"
)

type QuestionService struct {
}

func (q *QuestionService) GetQuestionsPage(req *GetQuestionsPageReq) (*utils.PageResult[GetQuestionPageResp], error) {
	conds := []gen.Condition{}
	conds = append(conds, questionsQ.Status.Eq(utils.QuestionStatusPublished))
	if req.PositionId != 0 {
		conds = append(conds, questionsQ.PositionID.Eq(req.PositionId))
	}
	if req.Type != nil {
		conds = append(conds, questionsQ.Type.Eq(*req.Type))
	}
	if !req.OperationTimeStart.IsZero() && !req.OperationTimeEnd.IsZero() {
		conds = append(conds, questionsQ.UpdatedAt.Between(req.OperationTimeStart, req.OperationTimeEnd))
	}
	offset := (req.Page - 1) * req.Size
	questionPage, _, err := questionsQ.Where(conds...).FindByPage(offset, req.Size)
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
	count, err := questionsQ.Where(conds...).Count()
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

func (q *QuestionService) GetQuestionInfo(id int) (*GetQuestionInfoResp, error) {
	question, err := questionsQ.Where(questionsQ.ID.Eq(id)).First()
	if err != nil {
		wlog.Error("call questionsQ.First failed").Err(err).Field("id", id).Log()
		return nil, err
	}
	optionList, err := questionOptionsQ.Where(questionOptionsQ.QuestionID.Eq(id)).Find()
	if err != nil {
		wlog.Error("call questionOptionsQ.Find failed").Err(err).Field("id", id).Log()
		return nil, err
	}
	resp := &GetQuestionInfoResp{}
	err = copier.Copy(resp, question)
	if err != nil {
		wlog.Error("call copier.Copy failed").Err(err).Field("question", question).Log()
		return nil, err
	}
	if question.Type == utils.QuestionTypeChoice {
		options := make([]Options, 0)
		err = copier.Copy(&options, optionList)
		if err != nil {
			wlog.Error("call copier.Copy failed").Err(err).Field("optionList", optionList).Log()
			return nil, err
		}
		resp.Options = options
	}
	return resp, nil
}

func (q *QuestionService) UpdateQuestion(req *UpdateQuestionReq) error {
	tx := query.Q.Begin()
	questionsTX := tx.Questions
	questionOptionsTX := tx.QuestionOptions
	question, err := questionsTX.Where(questionsQ.ID.Eq(req.ID)).First()
	if err != nil {
		wlog.Error("call questionsQ.First failed").Err(err).Field("req", req).Log()
		return err
	}
	err = copier.Copy(question, req)
	if err != nil {
		wlog.Error("call copier.Copy failed").Err(err).Field("req", req).Log()
		return err
	}
	_, err = questionsTX.Updates(question)
	if err != nil {
		wlog.Error("call questionsQ.Updates failed").Err(err).Field("req", req).Log()
		tx.Rollback()
		return err
	}
	_, err = questionOptionsTX.Where(questionOptionsTX.QuestionID.Eq(req.ID)).Delete()
	if err != nil {
		wlog.Error("call questionOptionsTX.Delete failed").Err(err).Field("req", req).Log()
		tx.Rollback()
		return err
	}
	optionList := make([]*model.QuestionOptions, 0)
	err = copier.Copy(&optionList, req.Options)
	if err != nil {
		wlog.Error("call copier.Copy failed").Err(err).Field("req", req).Log()
		tx.Rollback()
		return err
	}
	for _, opt := range optionList {
		opt.QuestionID = req.ID
	}
	err = questionOptionsTX.Create(optionList...)
	if err != nil {
		wlog.Error("call questionOptionsTX.Create failed").Err(err).Field("req", req).Log()
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (q *QuestionService) DeleteQuestion(id int) error {
	_, err := questionsQ.Where(questionsQ.ID.Eq(id)).Delete()
	if err != nil {
		wlog.Error("call questionsQ.Delete failed").Err(err).Field("id", id).Log()
		return err
	}
	return nil
}

func (q *QuestionService) FetchRandomQuestions(positionId, count int) ([]FetchRandomQuestionsResp, error) {
	questionList, err := questionsQ.Where(questionsQ.PositionID.Eq(positionId)).Order(field.Func.Rand()).Limit(count).Find()
	if err != nil {
		wlog.Error("call questionsQ.Find failed").Err(err).Field("positionId", positionId).Log()
		return nil, err
	}
	respList := make([]FetchRandomQuestionsResp, 0)

	for _, question := range questionList {
		resp := FetchRandomQuestionsResp{}
		err = copier.Copy(&resp, question)
		if err != nil {
			wlog.Error("call copier.Copy failed").Err(err).Field("positionId", positionId).Log()
			return nil, err
		}
		options := make([]Options, 0)
		optionList, err := questionOptionsQ.Where(questionOptionsQ.QuestionID.Eq(question.ID)).Find()
		if err != nil {
			wlog.Error("call questionOptionsQ.Find failed").Err(err).Field("positionId", positionId).Log()
			return nil, err
		}
		err = copier.Copy(&options, optionList)
		if err != nil {
			wlog.Error("call copier.Copy failed").Err(err).Field("optionList", optionList).Log()
			return nil, err
		}
		resp.Options = options
		respList = append(respList, resp)
	}
	return respList, nil
}

func (q *QuestionService) CalculateScore(c *gin.Context, req *CalculateScoreReq) (*CalculateScoreResp, error) {
	resp := &CalculateScoreResp{}
	err := query.Q.Transaction(func(tx *query.Query) error {
		examRecordsTX := tx.ExamRecords
		answerRecordsTX := tx.AnswerRecords
		examRecord := &model.ExamRecords{
			UserID:     utils.GetUserId(c),
			PositionID: req.PositionID,
		}
		err := examRecordsTX.Create(examRecord)
		if err != nil {
			wlog.Error("call examRecordsTX.Create failed").Err(err).Field("req", req).Log()
			return err
		}
		score := 0
		answerList := make([]Answer, 0)
		for _, questionAnswer := range req.QuestionAnswerList {
			questionNo := questionAnswer.QuestionNo
			questionID := questionAnswer.QuestionID
			userAnswer := questionAnswer.UserAnswer
			question, err := questionsQ.Where(questionsQ.ID.Eq(questionID)).First()
			if err != nil {
				wlog.Error("call questionsQ.First failed").Err(err).Field("req", req).Log()
				return err
			}
			answer := Answer{
				QuestionNo:    questionNo,
				QuestionID:    questionID,
				Title:         question.Title,
				UserAnswer:    userAnswer,
				CorrectAnswer: question.Answer,
				Type:          question.Type,
			}
			answerRecord := &model.AnswerRecords{}
			if userAnswer == question.Answer {
				score += 10
				answer.IsCorrect = true
				answerRecord.IsCorrect = 1
			}
			err = copier.Copy(answerRecord, answer)
			if err != nil {
				wlog.Error("call copier.Copy failed").Err(err).Field("questionAnswer", questionAnswer).Log()
				return err
			}
			return errors.New("模拟一个错误")
			answerRecord.ExamID = examRecord.ID
			err = answerRecordsTX.Create(answerRecord)
			if err != nil {
				wlog.Error("call answerRecordsTX.Create failed").Err(err).Field("answerRecord", answerRecord).Log()
				return err
			}
			optionList, err := questionOptionsQ.Where(questionOptionsQ.QuestionID.Eq(questionID)).Find()
			if err != nil {
				wlog.Error("call questionOptionsQ.Find failed").Err(err).Field("questionAnswer", questionAnswer).Log()
				return err
			}
			options := make([]Options, 0)
			err = copier.Copy(&options, optionList)
			if err != nil {
				wlog.Error("call copier.Copy failed").Err(err).Field("questionAnswer", questionAnswer).Log()
				return err
			}
			answer.Options = options
			answerList = append(answerList, answer)
		}
		resp = &CalculateScoreResp{
			Score:   score,
			Answers: answerList,
			ExamID:  examRecord.ID,
		}
		examRecord.Score = score
		examRecord.Remark = "继续努力"
		_, err = examRecordsTX.Updates(examRecord)
		if err != nil {
			wlog.Error("call examRecordsTX.Updates failed").Err(err).Field("req", req).Log()
			return err
		}
		return nil
	})
	if err != nil {
		wlog.Error("call query.Q.Transaction failed").Err(err).Field("req", req).Log()
		return nil, err
	}
	return resp, nil
}
