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

func newCommunities(db *gorm.DB, opts ...gen.DOOption) communities {
	_communities := communities{}

	_communities.communitiesDo.UseDB(db, opts...)
	_communities.communitiesDo.UseModel(&model.Communities{})

	tableName := _communities.communitiesDo.TableName()
	_communities.ALL = field.NewAsterisk(tableName)
	_communities.ID = field.NewInt(tableName, "id")
	_communities.CommunityCode = field.NewString(tableName, "community_code")
	_communities.Name = field.NewString(tableName, "name")
	_communities.OwnerID = field.NewInt(tableName, "owner_id")
	_communities.Avatar = field.NewString(tableName, "avatar")
	_communities.Type = field.NewInt(tableName, "type")
	_communities.Description = field.NewString(tableName, "description")
	_communities.CreatedAt = field.NewTime(tableName, "created_at")
	_communities.UpdatedAt = field.NewTime(tableName, "updated_at")
	_communities.DeletedAt = field.NewField(tableName, "deleted_at")

	_communities.fillFieldMap()

	return _communities
}

type communities struct {
	communitiesDo

	ALL           field.Asterisk
	ID            field.Int
	CommunityCode field.String
	Name          field.String
	OwnerID       field.Int
	Avatar        field.String
	Type          field.Int // 0.默认 1.兴趣爱好 2.行业交流 3.生活休闲
	Description   field.String
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Field

	fieldMap map[string]field.Expr
}

func (c communities) Table(newTableName string) *communities {
	c.communitiesDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c communities) As(alias string) *communities {
	c.communitiesDo.DO = *(c.communitiesDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *communities) updateTableName(table string) *communities {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt(table, "id")
	c.CommunityCode = field.NewString(table, "community_code")
	c.Name = field.NewString(table, "name")
	c.OwnerID = field.NewInt(table, "owner_id")
	c.Avatar = field.NewString(table, "avatar")
	c.Type = field.NewInt(table, "type")
	c.Description = field.NewString(table, "description")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")

	c.fillFieldMap()

	return c
}

func (c *communities) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *communities) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 10)
	c.fieldMap["id"] = c.ID
	c.fieldMap["community_code"] = c.CommunityCode
	c.fieldMap["name"] = c.Name
	c.fieldMap["owner_id"] = c.OwnerID
	c.fieldMap["avatar"] = c.Avatar
	c.fieldMap["type"] = c.Type
	c.fieldMap["description"] = c.Description
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
}

func (c communities) clone(db *gorm.DB) communities {
	c.communitiesDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c communities) replaceDB(db *gorm.DB) communities {
	c.communitiesDo.ReplaceDB(db)
	return c
}

type communitiesDo struct{ gen.DO }

type ICommunitiesDo interface {
	gen.SubQuery
	Debug() ICommunitiesDo
	WithContext(ctx context.Context) ICommunitiesDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICommunitiesDo
	WriteDB() ICommunitiesDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICommunitiesDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICommunitiesDo
	Not(conds ...gen.Condition) ICommunitiesDo
	Or(conds ...gen.Condition) ICommunitiesDo
	Select(conds ...field.Expr) ICommunitiesDo
	Where(conds ...gen.Condition) ICommunitiesDo
	Order(conds ...field.Expr) ICommunitiesDo
	Distinct(cols ...field.Expr) ICommunitiesDo
	Omit(cols ...field.Expr) ICommunitiesDo
	Join(table schema.Tabler, on ...field.Expr) ICommunitiesDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICommunitiesDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICommunitiesDo
	Group(cols ...field.Expr) ICommunitiesDo
	Having(conds ...gen.Condition) ICommunitiesDo
	Limit(limit int) ICommunitiesDo
	Offset(offset int) ICommunitiesDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICommunitiesDo
	Unscoped() ICommunitiesDo
	Create(values ...*model.Communities) error
	CreateInBatches(values []*model.Communities, batchSize int) error
	Save(values ...*model.Communities) error
	First() (*model.Communities, error)
	Take() (*model.Communities, error)
	Last() (*model.Communities, error)
	Find() ([]*model.Communities, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Communities, err error)
	FindInBatches(result *[]*model.Communities, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Communities) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICommunitiesDo
	Assign(attrs ...field.AssignExpr) ICommunitiesDo
	Joins(fields ...field.RelationField) ICommunitiesDo
	Preload(fields ...field.RelationField) ICommunitiesDo
	FirstOrInit() (*model.Communities, error)
	FirstOrCreate() (*model.Communities, error)
	FindByPage(offset int, limit int) (result []*model.Communities, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICommunitiesDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c communitiesDo) Debug() ICommunitiesDo {
	return c.withDO(c.DO.Debug())
}

func (c communitiesDo) WithContext(ctx context.Context) ICommunitiesDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c communitiesDo) ReadDB() ICommunitiesDo {
	return c.Clauses(dbresolver.Read)
}

func (c communitiesDo) WriteDB() ICommunitiesDo {
	return c.Clauses(dbresolver.Write)
}

func (c communitiesDo) Session(config *gorm.Session) ICommunitiesDo {
	return c.withDO(c.DO.Session(config))
}

func (c communitiesDo) Clauses(conds ...clause.Expression) ICommunitiesDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c communitiesDo) Returning(value interface{}, columns ...string) ICommunitiesDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c communitiesDo) Not(conds ...gen.Condition) ICommunitiesDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c communitiesDo) Or(conds ...gen.Condition) ICommunitiesDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c communitiesDo) Select(conds ...field.Expr) ICommunitiesDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c communitiesDo) Where(conds ...gen.Condition) ICommunitiesDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c communitiesDo) Order(conds ...field.Expr) ICommunitiesDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c communitiesDo) Distinct(cols ...field.Expr) ICommunitiesDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c communitiesDo) Omit(cols ...field.Expr) ICommunitiesDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c communitiesDo) Join(table schema.Tabler, on ...field.Expr) ICommunitiesDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c communitiesDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICommunitiesDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c communitiesDo) RightJoin(table schema.Tabler, on ...field.Expr) ICommunitiesDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c communitiesDo) Group(cols ...field.Expr) ICommunitiesDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c communitiesDo) Having(conds ...gen.Condition) ICommunitiesDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c communitiesDo) Limit(limit int) ICommunitiesDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c communitiesDo) Offset(offset int) ICommunitiesDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c communitiesDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICommunitiesDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c communitiesDo) Unscoped() ICommunitiesDo {
	return c.withDO(c.DO.Unscoped())
}

func (c communitiesDo) Create(values ...*model.Communities) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c communitiesDo) CreateInBatches(values []*model.Communities, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c communitiesDo) Save(values ...*model.Communities) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c communitiesDo) First() (*model.Communities, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Communities), nil
	}
}

func (c communitiesDo) Take() (*model.Communities, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Communities), nil
	}
}

func (c communitiesDo) Last() (*model.Communities, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Communities), nil
	}
}

func (c communitiesDo) Find() ([]*model.Communities, error) {
	result, err := c.DO.Find()
	return result.([]*model.Communities), err
}

func (c communitiesDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Communities, err error) {
	buf := make([]*model.Communities, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c communitiesDo) FindInBatches(result *[]*model.Communities, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c communitiesDo) Attrs(attrs ...field.AssignExpr) ICommunitiesDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c communitiesDo) Assign(attrs ...field.AssignExpr) ICommunitiesDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c communitiesDo) Joins(fields ...field.RelationField) ICommunitiesDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c communitiesDo) Preload(fields ...field.RelationField) ICommunitiesDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c communitiesDo) FirstOrInit() (*model.Communities, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Communities), nil
	}
}

func (c communitiesDo) FirstOrCreate() (*model.Communities, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Communities), nil
	}
}

func (c communitiesDo) FindByPage(offset int, limit int) (result []*model.Communities, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c communitiesDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c communitiesDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c communitiesDo) Delete(models ...*model.Communities) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *communitiesDo) withDO(do gen.Dao) *communitiesDo {
	c.DO = *do.(*gen.DO)
	return c
}
