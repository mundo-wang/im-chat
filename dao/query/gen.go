// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q               = new(Query)
	AnswerRecords   *answerRecords
	Communities     *communities
	Contacts        *contacts
	ExamRecords     *examRecords
	Messages        *messages
	QuestionOptions *questionOptions
	QuestionSession *questionSession
	Questions       *questions
	Users           *users
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	AnswerRecords = &Q.AnswerRecords
	Communities = &Q.Communities
	Contacts = &Q.Contacts
	ExamRecords = &Q.ExamRecords
	Messages = &Q.Messages
	QuestionOptions = &Q.QuestionOptions
	QuestionSession = &Q.QuestionSession
	Questions = &Q.Questions
	Users = &Q.Users
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:              db,
		AnswerRecords:   newAnswerRecords(db, opts...),
		Communities:     newCommunities(db, opts...),
		Contacts:        newContacts(db, opts...),
		ExamRecords:     newExamRecords(db, opts...),
		Messages:        newMessages(db, opts...),
		QuestionOptions: newQuestionOptions(db, opts...),
		QuestionSession: newQuestionSession(db, opts...),
		Questions:       newQuestions(db, opts...),
		Users:           newUsers(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	AnswerRecords   answerRecords
	Communities     communities
	Contacts        contacts
	ExamRecords     examRecords
	Messages        messages
	QuestionOptions questionOptions
	QuestionSession questionSession
	Questions       questions
	Users           users
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:              db,
		AnswerRecords:   q.AnswerRecords.clone(db),
		Communities:     q.Communities.clone(db),
		Contacts:        q.Contacts.clone(db),
		ExamRecords:     q.ExamRecords.clone(db),
		Messages:        q.Messages.clone(db),
		QuestionOptions: q.QuestionOptions.clone(db),
		QuestionSession: q.QuestionSession.clone(db),
		Questions:       q.Questions.clone(db),
		Users:           q.Users.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:              db,
		AnswerRecords:   q.AnswerRecords.replaceDB(db),
		Communities:     q.Communities.replaceDB(db),
		Contacts:        q.Contacts.replaceDB(db),
		ExamRecords:     q.ExamRecords.replaceDB(db),
		Messages:        q.Messages.replaceDB(db),
		QuestionOptions: q.QuestionOptions.replaceDB(db),
		QuestionSession: q.QuestionSession.replaceDB(db),
		Questions:       q.Questions.replaceDB(db),
		Users:           q.Users.replaceDB(db),
	}
}

type queryCtx struct {
	AnswerRecords   IAnswerRecordsDo
	Communities     ICommunitiesDo
	Contacts        IContactsDo
	ExamRecords     IExamRecordsDo
	Messages        IMessagesDo
	QuestionOptions IQuestionOptionsDo
	QuestionSession IQuestionSessionDo
	Questions       IQuestionsDo
	Users           IUsersDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		AnswerRecords:   q.AnswerRecords.WithContext(ctx),
		Communities:     q.Communities.WithContext(ctx),
		Contacts:        q.Contacts.WithContext(ctx),
		ExamRecords:     q.ExamRecords.WithContext(ctx),
		Messages:        q.Messages.WithContext(ctx),
		QuestionOptions: q.QuestionOptions.WithContext(ctx),
		QuestionSession: q.QuestionSession.WithContext(ctx),
		Questions:       q.Questions.WithContext(ctx),
		Users:           q.Users.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
