package dto

import (
	"time"
)

type FullAudit struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	CreatedUser string    `json:"created_user,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	UpdatedUser string    `json:"updated_user,omitempty"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
	DeletedUser string    `json:"deleted_user,omitempty"`
}

type CreateAudit struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	CreatedUser string    `json:"created_user,omitempty"`
}

type CreateUpdateAudit struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	CreatedUser string    `json:"created_user,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	UpdatedUser string    `json:"updated_user,omitempty"`
}

type CreateDeleteAudit struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	CreatedUser string    `json:"created_user,omitempty"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
	DeletedUser string    `json:"deleted_user,omitempty"`
}

type LookupAudit struct {
	ID int `json:"id"`
}
