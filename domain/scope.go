package domain

import (
	"time"
)

type Scope struct {
	Id          int64  `gorm:"primaryKey"`
	Name        string `gorm:"unique"`
	Description string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}

type ScopeRepository interface {
	Create(scope *Scope) (*Scope, error)
	FindAll() ([]*Scope, error)
	FindById(int64) (*Scope, error)
	Update(*Scope) (*Scope, error)
	Delete(*Scope) error
}

type ScopeUsecase interface {
	CreateScope(*Scope) (*Scope, error)
	GetScope(*Scope) (*Scope, error)
	ShowScopes() ([]*Scope, error)
	UpdateScope(*Scope) (*Scope, error)
	DeleteScope(*Scope) error
}
