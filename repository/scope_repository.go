package repository

import (
	"errors"
	"strings"
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
		isDuplicateKeyFirstCheckTrue := errors.Is(err, gorm.ErrDuplicatedKey)
		isDuplicateKeySecondCheckTrue := strings.Contains(err.Error(), "duplicate")
		if isDuplicateKeyFirstCheckTrue || isDuplicateKeySecondCheckTrue {
			return nil, errors.New("SCOPE_REPOSITORY.DUPLICATE_NAME")
		}

		return nil, err
	}
	return newScope, nil
}

func (r *ScopeRepository) FindAll() ([]*domain.Scope, error) {
	var scopes []*domain.Scope
	err := r.db.Where("deleted_at IS NULL").Find(&scopes).Error
	if err != nil {
		return nil, err
	}
	return scopes, nil
}

func (r *ScopeRepository) FindById(id int64) (*domain.Scope, error) {
	var scope domain.Scope
	err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&scope).Error
	if err != nil {
		isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
		if isNotFound {
			return nil, errors.New("SCOPE_REPOSITORY.SCOPE_NOT_FOUND")
		}
		return nil, err
	}
	return &scope, nil
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
