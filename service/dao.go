package service

import "im-chat/dao/query"

var (
	usersQ           = query.Users
	contactsQ        = query.Contacts
	communitiesQ     = query.Communities
	messagesQ        = query.Messages
	questionsQ       = query.Questions
	questionOptionsQ = query.QuestionOptions
	questionSessionQ = query.QuestionSession
	examRecordsQ     = query.ExamRecords
	answerRecordsQ   = query.AnswerRecords
)

func InitDao() {
	usersQ = query.Users
	contactsQ = query.Contacts
	communitiesQ = query.Communities
	messagesQ = query.Messages
	questionsQ = query.Questions
	questionOptionsQ = query.QuestionOptions
	questionSessionQ = query.QuestionSession
	examRecordsQ = query.ExamRecords
	answerRecordsQ = query.AnswerRecords
}
