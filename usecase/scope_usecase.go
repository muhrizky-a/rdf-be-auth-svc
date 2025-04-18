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

func (u *ScopeUsecase) ShowScopes() ([]*domain.Scope, error) {
	return u.ScopeRepository.FindAll()
}

func (u *ScopeUsecase) UpdateScope(scope *domain.Scope) (*domain.Scope, error) {
	oldScope, err := u.ScopeRepository.FindById(scope.Id)
	if err != nil || oldScope == nil {
		return nil, err
	}
	now := time.Now()
	oldScope.Name = scope.Name
	oldScope.UpdatedAt = &now

	return u.ScopeRepository.Update(oldScope)
}

func (u *ScopeUsecase) DeleteScope(scope *domain.Scope) error {
	_, err := u.ScopeRepository.FindById(scope.Id)
	if err != nil {
		return err
	}

	roleScopes, err := u.RoleScopeRepository.FindByScopeId(scope.Id)
	if err != nil {
		return err
	}

	if len(roleScopes) > 0 {
		return errors.New("Scope tidak bisa dihapus jika masih terhubung dengan role")
	}

	now := time.Now()
	scope.DeletedAt = &now
	return u.ScopeRepository.Delete(scope)
}
