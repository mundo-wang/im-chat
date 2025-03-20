// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameContact = "contact"

// Contact mapped from table <contact>
type Contact struct {
	ID          int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	OwnerID     int32          `gorm:"column:owner_id" json:"owner_id"`
	TargetID    int32          `gorm:"column:target_id" json:"target_id"`
	Type        int32          `gorm:"column:type" json:"type"`
	Description string         `gorm:"column:description" json:"description"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Contact's table name
func (*Contact) TableName() string {
	return TableNameContact
}
