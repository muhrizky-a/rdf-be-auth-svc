package repository

import (
	"testing"
	"time"

	"github.com/ryakadev/rdf-be-auth-svc/domain"
	"github.com/ryakadev/rdf-be-auth-svc/infrastructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repoScopeMock struct {
	mock.Mock
}

func NewRepoScopeMock() domain.ScopeRepository {
	return &repoScopeMock{}
}

func (r *repoScopeMock) Create(scope *domain.Scope) (*domain.Scope, error) {
	args := r.Called(scope)
	return args.Get(0).(*domain.Scope), args.Error(1)
}

func (r *repoScopeMock) FindAll() ([]*domain.Scope, error) {
	args := r.Called()
	return args.Get(0).([]*domain.Scope), args.Error(1)
}

func (r *repoScopeMock) FindById(id int64) (*domain.Scope, error) {
	args := r.Called(id)
	return args.Get(0).(*domain.Scope), args.Error(1)
}

func (r *repoScopeMock) Update(scope *domain.Scope) (*domain.Scope, error) {
	args := r.Called(scope)
	return args.Get(0).(*domain.Scope), args.Error(1)
}

func (r *repoScopeMock) Delete(scope *domain.Scope) error {
	args := r.Called(scope)
	return args.Error(0)
}

func TestCreateScope(t *testing.T) {
	repoScopeMock := repoScopeMock{}
	scope := &domain.Scope{
		Name:        "Account:Create",
		Description: "Create an account",
	}
	now := time.Now()
	CreateScopeReponse := &domain.Scope{
		Id:          1,
		Name:        "Account:Create",
		Description: "Create an account",
		CreatedAt:   &now,
	}
	repoScopeMock.On("Create", scope).Return(CreateScopeReponse, nil)

	t.Run("Create a new Scope with mock", func(t *testing.T) {
		assert.Equal(t, 123, 123, "they should be equal")
		scope, err := repoScopeMock.Create(scope)
		assert.Nil(t, err)
		assert.NotNil(t, scope)
	})

	db := infrastructure.ConnectDB()
	repoScope := NewScopeRepository(db)

	t.Run("Create a new Scope to DB", func(t *testing.T) {
		scope, err := repoScope.Create(scope)
		assert.NotNil(t, scope.Id)
		assert.Nil(t, err)
	})
}

func TestCreateScopeWithExistingName(t *testing.T) {
	db := infrastructure.ConnectDB()
	repoScope := NewScopeRepository(db)
	scope := &domain.Scope{
		Name:        "Account:Show",
		Description: "Show an account",
	}

	t.Run("Create a new Scope with Existing Name to DB", func(t *testing.T) {
		scope, err := repoScope.Create(scope)
		assert.Nil(t, err)
		assert.NotNil(t, scope)

		scope, err = repoScope.Create(scope)
		assert.Nil(t, scope)
		assert.NotNil(t, err)
	})
}

func TestShowScope(t *testing.T) {

	db := infrastructure.ConnectDB()
	scopeRepo := NewScopeRepository(db)
	t.Run("Show a Scope", func(t *testing.T) {
		scopes, err := scopeRepo.FindAll()
		assert.Nil(t, err)
		assert.NotNil(t, scopes)
	})
}

func TestUpdateScope(t *testing.T) {

	repoScopeMock := repoScopeMock{}
	scope := &domain.Scope{
		Name:        "WrongAccount:Update",
		Description: "Update an account",
	}
	now := time.Now()
	CreateScopeReponse := &domain.Scope{
		Id:          1,
		Name:        "WrongAccount:Update",
		Description: "Update an account",
		CreatedAt:   &now,
	}
	repoScopeMock.On("Create", scope).Return(CreateScopeReponse, nil)
	CreateScopeReponse = &domain.Scope{
		Id:          1,
		Name:        "Account:Update",
		Description: "Update an account",
		CreatedAt:   &now,
	}
	now = time.Now()
	UpdateScopeResponse := &domain.Scope{
		Id:          1,
		Name:        "Account:Update",
		Description: "Update an account",
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
	repoScopeMock.On("Update", CreateScopeReponse).Return(UpdateScopeResponse, nil)
	t.Run("Update a Scope with mock", func(t *testing.T) {

		scope, err := repoScopeMock.Create(scope)
		assert.Nil(t, err)
		assert.NotNil(t, scope)

		scope, err = repoScopeMock.Update(CreateScopeReponse)
		assert.Nil(t, err)
		assert.Equal(t, "Account:Update", scope.Name)
	})

	db := infrastructure.ConnectDB()
	repoScope := NewScopeRepository(db)

	t.Run("Update a Scope to DB", func(t *testing.T) {
		scope, err := repoScope.Create(scope)
		assert.Nil(t, err)
		assert.NotNil(t, scope)

		scope.Name = "Account:Update"
		scope, err = repoScope.Update(scope)
		assert.Nil(t, err)
		assert.Equal(t, "Account:Update", scope.Name)
	})
}

func TestDeleteScope(t *testing.T) {

	db := infrastructure.ConnectDB()
	repoScope := NewScopeRepository(db)
	scope := &domain.Scope{
		Name:        "Account:Delete",
		Description: "Delete an account",
	}

	t.Run("Delete a Scope", func(t *testing.T) {
		scope, err := repoScope.Create(scope)
		assert.Nil(t, err)

		err = repoScope.Delete(scope)
		assert.Nil(t, err)
	})
}
