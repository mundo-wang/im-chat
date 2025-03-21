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

func newUserBasic(db *gorm.DB, opts ...gen.DOOption) userBasic {
	_userBasic := userBasic{}

	_userBasic.userBasicDo.UseDB(db, opts...)
	_userBasic.userBasicDo.UseModel(&model.UserBasic{})

	tableName := _userBasic.userBasicDo.TableName()
	_userBasic.ALL = field.NewAsterisk(tableName)
	_userBasic.ID = field.NewInt32(tableName, "id")
	_userBasic.Name = field.NewString(tableName, "name")
	_userBasic.PassWord = field.NewString(tableName, "pass_word")
	_userBasic.Phone = field.NewString(tableName, "phone")
	_userBasic.Email = field.NewString(tableName, "email")
	_userBasic.Identity = field.NewString(tableName, "identity")
	_userBasic.ClientIP = field.NewString(tableName, "client_ip")
	_userBasic.ClientPort = field.NewString(tableName, "client_port")
	_userBasic.LoginTime = field.NewTime(tableName, "login_time")
	_userBasic.HeartbeatTime = field.NewTime(tableName, "heartbeat_time")
	_userBasic.LoginOutTime = field.NewTime(tableName, "login_out_time")
	_userBasic.IsLogout = field.NewBool(tableName, "is_logout")
	_userBasic.DeviceInfo = field.NewString(tableName, "device_info")
	_userBasic.Salt = field.NewString(tableName, "salt")
	_userBasic.Avatar = field.NewString(tableName, "avatar")
	_userBasic.CreatedAt = field.NewTime(tableName, "created_at")
	_userBasic.UpdatedAt = field.NewTime(tableName, "updated_at")
	_userBasic.DeletedAt = field.NewField(tableName, "deleted_at")

	_userBasic.fillFieldMap()

	return _userBasic
}

type userBasic struct {
	userBasicDo

	ALL           field.Asterisk
	ID            field.Int32
	Name          field.String
	PassWord      field.String
	Phone         field.String
	Email         field.String
	Identity      field.String
	ClientIP      field.String
	ClientPort    field.String
	LoginTime     field.Time
	HeartbeatTime field.Time
	LoginOutTime  field.Time
	IsLogout      field.Bool
	DeviceInfo    field.String
	Salt          field.String
	Avatar        field.String
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Field

	fieldMap map[string]field.Expr
}

func (u userBasic) Table(newTableName string) *userBasic {
	u.userBasicDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userBasic) As(alias string) *userBasic {
	u.userBasicDo.DO = *(u.userBasicDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userBasic) updateTableName(table string) *userBasic {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt32(table, "id")
	u.Name = field.NewString(table, "name")
	u.PassWord = field.NewString(table, "pass_word")
	u.Phone = field.NewString(table, "phone")
	u.Email = field.NewString(table, "email")
	u.Identity = field.NewString(table, "identity")
	u.ClientIP = field.NewString(table, "client_ip")
	u.ClientPort = field.NewString(table, "client_port")
	u.LoginTime = field.NewTime(table, "login_time")
	u.HeartbeatTime = field.NewTime(table, "heartbeat_time")
	u.LoginOutTime = field.NewTime(table, "login_out_time")
	u.IsLogout = field.NewBool(table, "is_logout")
	u.DeviceInfo = field.NewString(table, "device_info")
	u.Salt = field.NewString(table, "salt")
	u.Avatar = field.NewString(table, "avatar")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")

	u.fillFieldMap()

	return u
}

func (u *userBasic) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userBasic) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 18)
	u.fieldMap["id"] = u.ID
	u.fieldMap["name"] = u.Name
	u.fieldMap["pass_word"] = u.PassWord
	u.fieldMap["phone"] = u.Phone
	u.fieldMap["email"] = u.Email
	u.fieldMap["identity"] = u.Identity
	u.fieldMap["client_ip"] = u.ClientIP
	u.fieldMap["client_port"] = u.ClientPort
	u.fieldMap["login_time"] = u.LoginTime
	u.fieldMap["heartbeat_time"] = u.HeartbeatTime
	u.fieldMap["login_out_time"] = u.LoginOutTime
	u.fieldMap["is_logout"] = u.IsLogout
	u.fieldMap["device_info"] = u.DeviceInfo
	u.fieldMap["salt"] = u.Salt
	u.fieldMap["avatar"] = u.Avatar
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
}

func (u userBasic) clone(db *gorm.DB) userBasic {
	u.userBasicDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userBasic) replaceDB(db *gorm.DB) userBasic {
	u.userBasicDo.ReplaceDB(db)
	return u
}

type userBasicDo struct{ gen.DO }

type IUserBasicDo interface {
	gen.SubQuery
	Debug() IUserBasicDo
	WithContext(ctx context.Context) IUserBasicDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserBasicDo
	WriteDB() IUserBasicDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserBasicDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserBasicDo
	Not(conds ...gen.Condition) IUserBasicDo
	Or(conds ...gen.Condition) IUserBasicDo
	Select(conds ...field.Expr) IUserBasicDo
	Where(conds ...gen.Condition) IUserBasicDo
	Order(conds ...field.Expr) IUserBasicDo
	Distinct(cols ...field.Expr) IUserBasicDo
	Omit(cols ...field.Expr) IUserBasicDo
	Join(table schema.Tabler, on ...field.Expr) IUserBasicDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserBasicDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserBasicDo
	Group(cols ...field.Expr) IUserBasicDo
	Having(conds ...gen.Condition) IUserBasicDo
	Limit(limit int) IUserBasicDo
	Offset(offset int) IUserBasicDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserBasicDo
	Unscoped() IUserBasicDo
	Create(values ...*model.UserBasic) error
	CreateInBatches(values []*model.UserBasic, batchSize int) error
	Save(values ...*model.UserBasic) error
	First() (*model.UserBasic, error)
	Take() (*model.UserBasic, error)
	Last() (*model.UserBasic, error)
	Find() ([]*model.UserBasic, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserBasic, err error)
	FindInBatches(result *[]*model.UserBasic, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserBasic) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserBasicDo
	Assign(attrs ...field.AssignExpr) IUserBasicDo
	Joins(fields ...field.RelationField) IUserBasicDo
	Preload(fields ...field.RelationField) IUserBasicDo
	FirstOrInit() (*model.UserBasic, error)
	FirstOrCreate() (*model.UserBasic, error)
	FindByPage(offset int, limit int) (result []*model.UserBasic, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserBasicDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userBasicDo) Debug() IUserBasicDo {
	return u.withDO(u.DO.Debug())
}

func (u userBasicDo) WithContext(ctx context.Context) IUserBasicDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userBasicDo) ReadDB() IUserBasicDo {
	return u.Clauses(dbresolver.Read)
}

func (u userBasicDo) WriteDB() IUserBasicDo {
	return u.Clauses(dbresolver.Write)
}

func (u userBasicDo) Session(config *gorm.Session) IUserBasicDo {
	return u.withDO(u.DO.Session(config))
}

func (u userBasicDo) Clauses(conds ...clause.Expression) IUserBasicDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userBasicDo) Returning(value interface{}, columns ...string) IUserBasicDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userBasicDo) Not(conds ...gen.Condition) IUserBasicDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userBasicDo) Or(conds ...gen.Condition) IUserBasicDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userBasicDo) Select(conds ...field.Expr) IUserBasicDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userBasicDo) Where(conds ...gen.Condition) IUserBasicDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userBasicDo) Order(conds ...field.Expr) IUserBasicDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userBasicDo) Distinct(cols ...field.Expr) IUserBasicDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userBasicDo) Omit(cols ...field.Expr) IUserBasicDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userBasicDo) Join(table schema.Tabler, on ...field.Expr) IUserBasicDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userBasicDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserBasicDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userBasicDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserBasicDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userBasicDo) Group(cols ...field.Expr) IUserBasicDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userBasicDo) Having(conds ...gen.Condition) IUserBasicDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userBasicDo) Limit(limit int) IUserBasicDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userBasicDo) Offset(offset int) IUserBasicDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userBasicDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserBasicDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userBasicDo) Unscoped() IUserBasicDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userBasicDo) Create(values ...*model.UserBasic) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userBasicDo) CreateInBatches(values []*model.UserBasic, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userBasicDo) Save(values ...*model.UserBasic) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userBasicDo) First() (*model.UserBasic, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserBasic), nil
	}
}

func (u userBasicDo) Take() (*model.UserBasic, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserBasic), nil
	}
}

func (u userBasicDo) Last() (*model.UserBasic, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserBasic), nil
	}
}

func (u userBasicDo) Find() ([]*model.UserBasic, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserBasic), err
}

func (u userBasicDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserBasic, err error) {
	buf := make([]*model.UserBasic, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userBasicDo) FindInBatches(result *[]*model.UserBasic, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userBasicDo) Attrs(attrs ...field.AssignExpr) IUserBasicDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userBasicDo) Assign(attrs ...field.AssignExpr) IUserBasicDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userBasicDo) Joins(fields ...field.RelationField) IUserBasicDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userBasicDo) Preload(fields ...field.RelationField) IUserBasicDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userBasicDo) FirstOrInit() (*model.UserBasic, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserBasic), nil
	}
}

func (u userBasicDo) FirstOrCreate() (*model.UserBasic, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserBasic), nil
	}
}

func (u userBasicDo) FindByPage(offset int, limit int) (result []*model.UserBasic, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userBasicDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userBasicDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userBasicDo) Delete(models ...*model.UserBasic) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userBasicDo) withDO(do gen.Dao) *userBasicDo {
	u.DO = *do.(*gen.DO)
	return u
}
