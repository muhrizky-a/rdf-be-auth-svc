package repository

import (
	"github.com/ryakadev/rdf-be-auth-svc/domain"
	"gorm.io/gorm"
)

type RoleScopeRepository struct {
	db *gorm.DB
}

func NewRoleScopeRepository(db *gorm.DB) domain.RoleScopeRepository {
	return &RoleScopeRepository{db}
}

func (r *RoleScopeRepository) Create(roleScope *domain.RoleScope) (*domain.RoleScope, error) {
	newRoleScope := &domain.RoleScope{
		RoleId:  roleScope.RoleId,
		ScopeId: roleScope.ScopeId,
	}
	err := r.db.Create(&newRoleScope).Error
	if err != nil {
		return nil, err
	}
	return newRoleScope, nil
}

func (r *RoleScopeRepository) FindByScopeId(id int64) ([]*domain.RoleScope, error) {
	var roleScopes []*domain.RoleScope
	err := r.db.Find(&roleScopes).Where("scope_id = ? AND deleted_at IS NULL", id).Error
	if err != nil {
		return nil, err
	}
	return roleScopes, nil
}
