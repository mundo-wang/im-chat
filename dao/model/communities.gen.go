// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCommunities = "communities"

// Communities mapped from table <communities>
type Communities struct {
	ID          int            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name        string         `gorm:"column:name;not null" json:"name"`
	OwnerID     int            `gorm:"column:owner_id;not null" json:"owner_id"`
	Img         string         `gorm:"column:img" json:"img"`
	Description string         `gorm:"column:description" json:"description"`
	CreatedAt   time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Communities's table name
func (*Communities) TableName() string {
	return TableNameCommunities
}
