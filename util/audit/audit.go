package audit

import (
	"gorm.io/gorm"
	"time"
)

type FullAudit struct {
	ID          int64          `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null;default:now()" json:"created_at,omitempty"`
	CreatedUser int64          `gorm:"column:created_user;" json:"created_user,omitempty"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;default:now()" json:"updated_at,omitempty"`
	UpdatedUser int64          `gorm:"column:updated_user;" json:"updated_user,omitempty"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at,omitempty"`
	DeletedUser int64          `gorm:"column:deleted_user;" json:"deleted_user,omitempty"`
}

type CreateAudit struct {
	ID          int       `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at,omitempty"`
	CreatedUser int64     `gorm:"column:created_user;" json:"created_user,omitempty"`
}

type CreateUpdateAudit struct {
	ID          int       `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at,omitempty"`
	CreatedUser int64     `gorm:"column:created_user;" json:"created_user,omitempty"`
	UpdatedAt   time.Time `gorm:"column:updated_at;" json:"updated_at,omitempty"`
	UpdatedUser int64     `gorm:"column:updated_user;" json:"updated_user,omitempty"`
}

type CreateDeleteAudit struct {
	ID          int            `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null;default:now()" json:"created_at,omitempty"`
	CreatedUser int64          `gorm:"column:created_user;" json:"created_user,omitempty"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at,omitempty"`
	DeletedUser int64          `gorm:"column:deleted_user;" json:"deleted_user,omitempty"`
}

type LookupAudit struct {
	ID int `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
}
