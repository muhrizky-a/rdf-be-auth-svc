package usecase

import (
	"errors"
	"time"

	"github.com/ryakadev/rdf-be-auth-svc/domain"
)

type ScopeUsecase struct {
	ScopeRepository     domain.ScopeRepository
	RoleScopeRepository domain.RoleScopeRepository
}

func NewScopeUsecase(repo domain.ScopeRepository, roleScopeRepo domain.RoleScopeRepository) domain.ScopeUsecase {
	return &ScopeUsecase{ScopeRepository: repo, RoleScopeRepository: roleScopeRepo}
}

func (u *ScopeUsecase) CreateScope(scope *domain.Scope) (*domain.Scope, error) {
	return u.ScopeRepository.Create(scope)
}

func (u *ScopeUsecase) GetScope(scope *domain.Scope) (*domain.Scope, error) {
	findScope, err := u.ScopeRepository.FindById(scope.Id)
	if err != nil || findScope == nil {
		return nil, err
	}

	return findScope, nil
}

func (u *ScopeUsecase) ShowScopes() ([]*domain.Scope, error) {
	return u.ScopeRepository.FindAll()
}

func (u *ScopeUsecase) UpdateScope(scope *domain.Scope) (*domain.Scope, error) {
	findScope, err := u.ScopeRepository.FindById(scope.Id)
	if err != nil || findScope == nil {
		return nil, err
	}
	now := time.Now()
	findScope.UpdatedAt = &now

	return u.ScopeRepository.Update(findScope)
}

func (u *ScopeUsecase) DeleteScope(scope *domain.Scope) error {
	oldScope, err := u.ScopeRepository.FindById(scope.Id)
	if err != nil || oldScope == nil {
		return err
	}

	roleScopes, err := u.RoleScopeRepository.FindByScopeId(scope.Id)
	if err != nil {
		return err
	}

	if len(roleScopes) > 0 {
		return errors.New("SCOPE_USE_CASE.SCOPE_TIED_TO_ROLES")
	}

	return u.ScopeRepository.Delete(oldScope)
}
