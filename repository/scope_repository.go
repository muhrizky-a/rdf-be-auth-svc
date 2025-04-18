package repository

import (
	"time"

	"github.com/ryakadev/rdf-be-auth-svc/domain"
	"gorm.io/gorm"
)

type ScopeRepository struct {
	db *gorm.DB
}

func NewScopeRepository(db *gorm.DB) domain.ScopeRepository {
	return &ScopeRepository{db}
}

func (r *ScopeRepository) Create(scope *domain.Scope) (*domain.Scope, error) {
	now := time.Now()
	newScope := &domain.Scope{
		Name:        scope.Name,
		Description: scope.Description,
		CreatedAt:   &now,
	}
	err := r.db.Create(&newScope).Error
	if err != nil {
		return nil, err
	}
	return newScope, nil
}

func (r *ScopeRepository) FindAll() ([]*domain.Scope, error) {
	var scopes []*domain.Scope
	err := r.db.Find(&scopes).Where("deleted_at IS NULL").Error
	if err != nil {
		return nil, err
	}
	return scopes, nil
}

func (r *ScopeRepository) Update(scope *domain.Scope) (*domain.Scope, error) {
	err := r.db.Save(scope).Error
	if err != nil {
		return nil, err
	}

	return scope, nil
}

func (r *ScopeRepository) Delete(scope *domain.Scope) error {
	now := time.Now()
	scope.DeletedAt = &now

	err := r.db.Save(scope).Error
	if err != nil {
		return err
	}
	return nil
}
