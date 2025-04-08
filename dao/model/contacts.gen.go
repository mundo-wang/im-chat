// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameContacts = "contacts"

// Contacts mapped from table <contacts>
type Contacts struct {
	ID        int            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	OwnerID   int            `gorm:"column:owner_id;not null" json:"owner_id"`
	TargetID  int            `gorm:"column:target_id;not null;comment:对应的人/群ID" json:"target_id"` // 对应的人/群ID
	Type      int            `gorm:"column:type;not null;comment:1为好友，2为群组" json:"type"`          // 1为好友，2为群组
	CreatedAt time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Contacts's table name
func (*Contacts) TableName() string {
	return TableNameContacts
}
