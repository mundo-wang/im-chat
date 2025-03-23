// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"im-chat/dao/model"
)

func newMessages(db *gorm.DB, opts ...gen.DOOption) messages {
	_messages := messages{}

	_messages.messagesDo.UseDB(db, opts...)
	_messages.messagesDo.UseModel(&model.Messages{})

	tableName := _messages.messagesDo.TableName()
	_messages.ALL = field.NewAsterisk(tableName)
	_messages.ID = field.NewInt32(tableName, "id")
	_messages.FormID = field.NewInt32(tableName, "form_id")
	_messages.TargetID = field.NewInt32(tableName, "target_id")
	_messages.Type = field.NewInt32(tableName, "type")
	_messages.Content = field.NewString(tableName, "content")
	_messages.Pic = field.NewString(tableName, "pic")
	_messages.URL = field.NewString(tableName, "url")
	_messages.Description = field.NewString(tableName, "description")
	_messages.CreatedAt = field.NewTime(tableName, "created_at")
	_messages.UpdatedAt = field.NewTime(tableName, "updated_at")
	_messages.DeletedAt = field.NewField(tableName, "deleted_at")

	_messages.fillFieldMap()

	return _messages
}

type messages struct {
	messagesDo

	ALL         field.Asterisk
	ID          field.Int32
	FormID      field.Int32
	TargetID    field.Int32
	Type        field.Int32
	Content     field.String
	Pic         field.String
	URL         field.String
	Description field.String
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field

	fieldMap map[string]field.Expr
}

func (m messages) Table(newTableName string) *messages {
	m.messagesDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m messages) As(alias string) *messages {
	m.messagesDo.DO = *(m.messagesDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *messages) updateTableName(table string) *messages {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt32(table, "id")
	m.FormID = field.NewInt32(table, "form_id")
	m.TargetID = field.NewInt32(table, "target_id")
	m.Type = field.NewInt32(table, "type")
	m.Content = field.NewString(table, "content")
	m.Pic = field.NewString(table, "pic")
	m.URL = field.NewString(table, "url")
	m.Description = field.NewString(table, "description")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")
	m.DeletedAt = field.NewField(table, "deleted_at")

	m.fillFieldMap()

	return m
}

func (m *messages) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *messages) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 11)
	m.fieldMap["id"] = m.ID
	m.fieldMap["form_id"] = m.FormID
	m.fieldMap["target_id"] = m.TargetID
	m.fieldMap["type"] = m.Type
	m.fieldMap["content"] = m.Content
	m.fieldMap["pic"] = m.Pic
	m.fieldMap["url"] = m.URL
	m.fieldMap["description"] = m.Description
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["deleted_at"] = m.DeletedAt
}

func (m messages) clone(db *gorm.DB) messages {
	m.messagesDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m messages) replaceDB(db *gorm.DB) messages {
	m.messagesDo.ReplaceDB(db)
	return m
}

type messagesDo struct{ gen.DO }

type IMessagesDo interface {
	gen.SubQuery
	Debug() IMessagesDo
	WithContext(ctx context.Context) IMessagesDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMessagesDo
	WriteDB() IMessagesDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMessagesDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMessagesDo
	Not(conds ...gen.Condition) IMessagesDo
	Or(conds ...gen.Condition) IMessagesDo
	Select(conds ...field.Expr) IMessagesDo
	Where(conds ...gen.Condition) IMessagesDo
	Order(conds ...field.Expr) IMessagesDo
	Distinct(cols ...field.Expr) IMessagesDo
	Omit(cols ...field.Expr) IMessagesDo
	Join(table schema.Tabler, on ...field.Expr) IMessagesDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMessagesDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMessagesDo
	Group(cols ...field.Expr) IMessagesDo
	Having(conds ...gen.Condition) IMessagesDo
	Limit(limit int) IMessagesDo
	Offset(offset int) IMessagesDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMessagesDo
	Unscoped() IMessagesDo
	Create(values ...*model.Messages) error
	CreateInBatches(values []*model.Messages, batchSize int) error
	Save(values ...*model.Messages) error
	First() (*model.Messages, error)
	Take() (*model.Messages, error)
	Last() (*model.Messages, error)
	Find() ([]*model.Messages, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Messages, err error)
	FindInBatches(result *[]*model.Messages, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Messages) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMessagesDo
	Assign(attrs ...field.AssignExpr) IMessagesDo
	Joins(fields ...field.RelationField) IMessagesDo
	Preload(fields ...field.RelationField) IMessagesDo
	FirstOrInit() (*model.Messages, error)
	FirstOrCreate() (*model.Messages, error)
	FindByPage(offset int, limit int) (result []*model.Messages, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMessagesDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m messagesDo) Debug() IMessagesDo {
	return m.withDO(m.DO.Debug())
}

func (m messagesDo) WithContext(ctx context.Context) IMessagesDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m messagesDo) ReadDB() IMessagesDo {
	return m.Clauses(dbresolver.Read)
}

func (m messagesDo) WriteDB() IMessagesDo {
	return m.Clauses(dbresolver.Write)
}

func (m messagesDo) Session(config *gorm.Session) IMessagesDo {
	return m.withDO(m.DO.Session(config))
}

func (m messagesDo) Clauses(conds ...clause.Expression) IMessagesDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m messagesDo) Returning(value interface{}, columns ...string) IMessagesDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m messagesDo) Not(conds ...gen.Condition) IMessagesDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m messagesDo) Or(conds ...gen.Condition) IMessagesDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m messagesDo) Select(conds ...field.Expr) IMessagesDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m messagesDo) Where(conds ...gen.Condition) IMessagesDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m messagesDo) Order(conds ...field.Expr) IMessagesDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m messagesDo) Distinct(cols ...field.Expr) IMessagesDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m messagesDo) Omit(cols ...field.Expr) IMessagesDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m messagesDo) Join(table schema.Tabler, on ...field.Expr) IMessagesDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m messagesDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMessagesDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m messagesDo) RightJoin(table schema.Tabler, on ...field.Expr) IMessagesDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m messagesDo) Group(cols ...field.Expr) IMessagesDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m messagesDo) Having(conds ...gen.Condition) IMessagesDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m messagesDo) Limit(limit int) IMessagesDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m messagesDo) Offset(offset int) IMessagesDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m messagesDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMessagesDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m messagesDo) Unscoped() IMessagesDo {
	return m.withDO(m.DO.Unscoped())
}

func (m messagesDo) Create(values ...*model.Messages) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m messagesDo) CreateInBatches(values []*model.Messages, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m messagesDo) Save(values ...*model.Messages) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m messagesDo) First() (*model.Messages, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Messages), nil
	}
}

func (m messagesDo) Take() (*model.Messages, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Messages), nil
	}
}

func (m messagesDo) Last() (*model.Messages, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Messages), nil
	}
}

func (m messagesDo) Find() ([]*model.Messages, error) {
	result, err := m.DO.Find()
	return result.([]*model.Messages), err
}

func (m messagesDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Messages, err error) {
	buf := make([]*model.Messages, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m messagesDo) FindInBatches(result *[]*model.Messages, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m messagesDo) Attrs(attrs ...field.AssignExpr) IMessagesDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m messagesDo) Assign(attrs ...field.AssignExpr) IMessagesDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m messagesDo) Joins(fields ...field.RelationField) IMessagesDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m messagesDo) Preload(fields ...field.RelationField) IMessagesDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m messagesDo) FirstOrInit() (*model.Messages, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Messages), nil
	}
}

func (m messagesDo) FirstOrCreate() (*model.Messages, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Messages), nil
	}
}

func (m messagesDo) FindByPage(offset int, limit int) (result []*model.Messages, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m messagesDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m messagesDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m messagesDo) Delete(models ...*model.Messages) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *messagesDo) withDO(do gen.Dao) *messagesDo {
	m.DO = *do.(*gen.DO)
	return m
}
