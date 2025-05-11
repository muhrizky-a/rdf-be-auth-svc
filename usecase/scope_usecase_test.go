package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/ryakadev/rdf-be-auth-svc/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type RoleScopeRepoMock struct {
	mock.Mock
}

type ScopeRepoMock struct {
	mock.Mock
}

func newRoleScopeRepoMock() *RoleScopeRepoMock {
	return &RoleScopeRepoMock{}
}

func newScopeRepoMock() *ScopeRepoMock {
	return &ScopeRepoMock{}
}

type ScopeUseCaseTestSuite struct {
	suite.Suite
}

func (r *RoleScopeRepoMock) Create(roleScope *domain.RoleScope) (*domain.RoleScope, error) {
	args := r.Called(roleScope)
	return args.Get(0).(*domain.RoleScope), args.Error(1)
}

func (r *RoleScopeRepoMock) FindByScopeId(id int64) ([]*domain.RoleScope, error) {
	args := r.Called(id)
	return args.Get(0).([]*domain.RoleScope), args.Error(1)
}

func (r *ScopeRepoMock) Create(scope *domain.Scope) (*domain.Scope, error) {
	args := r.Called(scope)
	return args.Get(0).(*domain.Scope), args.Error(1)
}

func (r *ScopeRepoMock) FindAll() ([]*domain.Scope, error) {
	args := r.Called()
	return args.Get(0).([]*domain.Scope), args.Error(1)
}

func (r *ScopeRepoMock) FindById(id int64) (*domain.Scope, error) {
	args := r.Called(id)
	return args.Get(0).(*domain.Scope), args.Error(1)
}

func (r *ScopeRepoMock) Update(scope *domain.Scope) (*domain.Scope, error) {
	args := r.Called(scope)
	return args.Get(0).(*domain.Scope), args.Error(1)
}

func (r *ScopeRepoMock) Delete(scope *domain.Scope) error {
	args := r.Called(scope)
	return args.Error(0)
}

func (suite *ScopeUseCaseTestSuite) TestCreateScopeMock() {
	// Arrange
	scope := &domain.Scope{
		Name:        "ScopeTest:Create",
		Description: "Create a scope",
	}
	now := time.Now()

	/// Mocking ScopeRepository.Create()
	scopeRepoMock := newScopeRepoMock()
	scopeRepoMock.On("Create", scope).Return(
		&domain.Scope{
			Id:          1,
			Name:        "ScopeTest:Create",
			Description: "Create a scope",
			CreatedAt:   &now,
			UpdatedAt:   &now,
		},
		nil,
	)

	scopeUC := NewScopeUsecase(scopeRepoMock, newRoleScopeRepoMock())

	// Action
	newScope, err := scopeUC.CreateScope(scope)

	// Assert
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), newScope)
	assert.Equal(
		suite.T(),
		newScope,
		&domain.Scope{
			Id:          1,
			Name:        "ScopeTest:Create",
			Description: "Create a scope",
			CreatedAt:   &now,
			UpdatedAt:   &now,
		},
	)
	scopeRepoMock.AssertCalled(suite.T(), "Create", scope)
	scopeRepoMock.AssertNumberOfCalls(suite.T(), "Create", 1)
}

func (suite *ScopeUseCaseTestSuite) TestGetScope() {
	// Arrange
	now := time.Now()
	scope := &domain.Scope{
		Id: 1,
	}

	scopeRepoMock := newScopeRepoMock()
	scopeRepoMock.On("FindById", scope.Id).Return(
		&domain.Scope{
			Id:          1,
			Name:        "ScopeTest:Create",
			Description: "Create a scope",
			CreatedAt:   &now,
			UpdatedAt:   &now,
		},
		nil,
	)

	scopeUC := NewScopeUsecase(scopeRepoMock, newRoleScopeRepoMock())
	// Action
	foundScope, err := scopeUC.GetScope(scope)

	// Assert
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), foundScope)
	assert.Equal(
		suite.T(),
		foundScope,
		&domain.Scope{
			Id:          1,
			Name:        "ScopeTest:Create",
			Description: "Create a scope",
			CreatedAt:   &now,
			UpdatedAt:   &now,
		},
	)
	scopeRepoMock.AssertCalled(suite.T(), "FindById", scope.Id)
	scopeRepoMock.AssertNumberOfCalls(suite.T(), "FindById", 1)
}

func (suite *ScopeUseCaseTestSuite) TestGetScopeByNonexistentId() {
	// Arrange
	nonexistentScope := &domain.Scope{
		Id: -1337,
	}

	scopeRepoMock := newScopeRepoMock()
	/// Must return empty Scope instead of nil due to Mock constrain
	scopeRepoMock.On("FindById", nonexistentScope.Id).Return(
		&domain.Scope{},
		errors.New("SCOPE_REPOSITORY.SCOPE_NOT_FOUND"),
	)

	scopeUC := NewScopeUsecase(scopeRepoMock, newRoleScopeRepoMock())

	// Action
	foundScope, err := scopeUC.GetScope(nonexistentScope)

	// Assert
	assert.Nil(suite.T(), foundScope)
	assert.NotNil(suite.T(), err)
	assert.Equal(
		suite.T(),
		err,
		errors.New("SCOPE_REPOSITORY.SCOPE_NOT_FOUND"),
	)
	scopeRepoMock.AssertCalled(suite.T(), "FindById", nonexistentScope.Id)
	scopeRepoMock.AssertNumberOfCalls(suite.T(), "FindById", 1)
}

func (suite *ScopeUseCaseTestSuite) TestShowScopes() {
	// Arrange
	scopeRepoMock := newScopeRepoMock()

	now := time.Now()
	scopeRepoMock.On("FindAll").Return(
		[]*domain.Scope{
			&domain.Scope{
				Id:          1,
				Name:        "ScopeTest:Create",
				Description: "Create a scope",
				CreatedAt:   &now,
				UpdatedAt:   &now,
			},
			&domain.Scope{
				Id:          2,
				Name:        "ScopeTest:ShowAll",
				Description: "Show all scopes",
				CreatedAt:   &now,
				UpdatedAt:   &now,
			},
		},
		nil,
	)
	scopeUC := NewScopeUsecase(scopeRepoMock, newRoleScopeRepoMock())

	// Action
	scopes, err := scopeUC.ShowScopes()
	firstScope, secondScope := scopes[0], scopes[1]

	// Assert
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), scopes)
	assert.Contains(
		suite.T(),
		scopes,
		&domain.Scope{
			Id:          1,
			Name:        "ScopeTest:Create",
			Description: "Create a scope",
			CreatedAt:   &now,
			UpdatedAt:   &now,
		},
	)
	assert.Equal(
		suite.T(),
		firstScope,
		&domain.Scope{
			Id:          1,
			Name:        "ScopeTest:Create",
			Description: "Create a scope",
			CreatedAt:   &now,
			UpdatedAt:   &now,
		},
	)
	assert.Contains(
		suite.T(),
		scopes,
		&domain.Scope{
			Id:          2,
			Name:        "ScopeTest:ShowAll",
			Description: "Show all scopes",
			CreatedAt:   &now,
			UpdatedAt:   &now,
		},
	)
	assert.Equal(
		suite.T(),
		secondScope,
		&domain.Scope{
			Id:          2,
			Name:        "ScopeTest:ShowAll",
			Description: "Show all scopes",
			CreatedAt:   &now,
			UpdatedAt:   &now,
		},
	)
	assert.Equal(suite.T(), len(scopes), 2)
	scopeRepoMock.AssertCalled(suite.T(), "FindAll")
	scopeRepoMock.AssertNumberOfCalls(suite.T(), "FindAll", 1)
}

func (suite *ScopeUseCaseTestSuite) TestUpdateScope() {
	// Arrange
	scopeRepoMock := newScopeRepoMock()

	now := time.Now()
	scope := &domain.Scope{
		Id:          1,
		Name:        "WrongScope:Update",
		Description: "Update a scope",
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	scopeRepoMock.On("FindById", scope.Id).Return(
		&domain.Scope{
			Id:          1,
			Name:        "WrongScope:Update",
			Description: "Update a scope",
			CreatedAt:   &now,
			UpdatedAt:   &now,
		},
		nil,
	)

	updatedTime := time.Now().Add(1 * time.Second)
	scopeRepoMock.On(
		"Update",
		&domain.Scope{
			Id:          1,
			Name:        "WrongScope:Update",
			Description: "Update a scope",
			CreatedAt:   &now,
			UpdatedAt:   &now,
		},
	).Return(
		&domain.Scope{
			Id:          1,
			Name:        "ScopeTest:Update",
			Description: "Update a scope",
			CreatedAt:   &now,
			UpdatedAt:   &updatedTime,
		},
		nil,
	)
	scopeUC := NewScopeUsecase(scopeRepoMock, newRoleScopeRepoMock())

	// Action
	updatedScope, err := scopeUC.UpdateScope(scope)

	// Assert
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), updatedScope)
	assert.Equal(suite.T(), updatedScope.Name, "ScopeTest:Update")
	assert.Equal(
		suite.T(),
		updatedScope,
		&domain.Scope{
			Id:          1,
			Name:        "ScopeTest:Update",
			Description: "Update a scope",
			CreatedAt:   &now,
			UpdatedAt:   &updatedTime,
		},
	)
	assert.NotEqual(
		suite.T(),

		scope.UpdatedAt,
		updatedScope.UpdatedAt,
	)
	assert.True(suite.T(), updatedScope.UpdatedAt.After(*scope.UpdatedAt))
	scopeRepoMock.AssertCalled(suite.T(), "FindById", scope.Id)
	scopeRepoMock.AssertNumberOfCalls(suite.T(), "FindById", 1)
	scopeRepoMock.AssertCalled(
		suite.T(),
		"Update",
		&domain.Scope{
			Id:          1,
			Name:        "WrongScope:Update",
			Description: "Update a scope",
			CreatedAt:   &now,
			UpdatedAt:   &now,
		},
	)
	scopeRepoMock.AssertNumberOfCalls(suite.T(), "Update", 1)
}

func (suite *ScopeUseCaseTestSuite) TestUpdateScopeWithNonexistentId() {
	nonexistentScope := &domain.Scope{
		Id: -1337,
	}

	scopeRepoMock := newScopeRepoMock()
	/// Must return empty Scope instead of nil due to Mock constrain
	scopeRepoMock.On("FindById", nonexistentScope.Id).Return(
		&domain.Scope{},
		errors.New("SCOPE_REPOSITORY.SCOPE_NOT_FOUND"),
	)

	scopeUC := NewScopeUsecase(scopeRepoMock, newRoleScopeRepoMock())

	// Action
	foundScope, err := scopeUC.UpdateScope(nonexistentScope)

	// Assert
	assert.Nil(suite.T(), foundScope)
	assert.NotNil(suite.T(), err)
	assert.Equal(
		suite.T(),
		err,
		errors.New("SCOPE_REPOSITORY.SCOPE_NOT_FOUND"),
	)
	scopeRepoMock.AssertCalled(suite.T(), "FindById", nonexistentScope.Id)
	scopeRepoMock.AssertNumberOfCalls(suite.T(), "FindById", 1)
	scopeRepoMock.AssertNumberOfCalls(suite.T(), "Update", 0)
}

func (suite *ScopeUseCaseTestSuite) TestDeleteScope() {
	// Arrange
	scopeRepoMock := newScopeRepoMock()
	roleScopeRepoMock := newRoleScopeRepoMock()

	now := time.Now()
	scope := &domain.Scope{
		Id:          1,
		Name:        "Scope:Delete",
		Description: "Delete a scope",
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
	scopeRepoMock.On("FindById", scope.Id).Return(
		&domain.Scope{
			Id:          1,
			Name:        "Scope:Delete",
			Description: "Delete a scope",
			CreatedAt:   &now,
			UpdatedAt:   &now,
		},
		nil,
	)

	roleScopeRepoMock.On("FindByScopeId", scope.Id).Return([]*domain.RoleScope{}, nil)
	scopeRepoMock.On("Delete", scope).Return(nil)
	scopeUC := NewScopeUsecase(scopeRepoMock, roleScopeRepoMock)

	// Action
	err := scopeUC.DeleteScope(scope)

	// Assert
	assert.Nil(suite.T(), err)
	scopeRepoMock.AssertCalled(suite.T(), "FindById", scope.Id)
	scopeRepoMock.AssertNumberOfCalls(suite.T(), "FindById", 1)
	scopeRepoMock.AssertCalled(suite.T(), "Delete", scope)
	scopeRepoMock.AssertNumberOfCalls(suite.T(), "Delete", 1)
}

func (suite *ScopeUseCaseTestSuite) TestDeleteScopeWithNonexistentId() {
	nonexistentScope := &domain.Scope{
		Id: -1337,
	}

	scopeRepoMock := newScopeRepoMock()
	/// Must return empty Scope instead of nil due to Mock constrain
	scopeRepoMock.On("FindById", nonexistentScope.Id).Return(
		&domain.Scope{},
		errors.New("SCOPE_REPOSITORY.SCOPE_NOT_FOUND"),
	)

	scopeUC := NewScopeUsecase(scopeRepoMock, newRoleScopeRepoMock())

	// Action
	err := scopeUC.DeleteScope(nonexistentScope)

	// Assert
	assert.NotNil(suite.T(), err)
	assert.Equal(
		suite.T(),
		err,
		errors.New("SCOPE_REPOSITORY.SCOPE_NOT_FOUND"),
	)
	scopeRepoMock.AssertCalled(suite.T(), "FindById", nonexistentScope.Id)
	scopeRepoMock.AssertNumberOfCalls(suite.T(), "FindById", 1)
	scopeRepoMock.AssertNumberOfCalls(suite.T(), "Delete", 0)
}

func (suite *ScopeUseCaseTestSuite) TestDeleteScopeWithRolesAttached() {
	// Arrange
	scopeRepoMock := newScopeRepoMock()
	roleScopeRepoMock := newRoleScopeRepoMock()

	now := time.Now()
	scope := &domain.Scope{
		Id:          1,
		Name:        "Scope:Delete",
		Description: "Delete a scope",
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
	scopeRepoMock.On("FindById", scope.Id).Return(
		&domain.Scope{
			Id:          1,
			Name:        "Scope:Delete",
			Description: "Delete a scope",
			CreatedAt:   &now,
			UpdatedAt:   &now,
		},
		nil,
	)

	roleScopeRepoMock.On("FindByScopeId", scope.Id).Return(
		[]*domain.RoleScope{
			&domain.RoleScope{
				RoleId:  1,
				ScopeId: scope.Id,
			},
		},
		nil,
	)
	scopeUC := NewScopeUsecase(scopeRepoMock, roleScopeRepoMock)

	// Action
	err := scopeUC.DeleteScope(scope)

	// Assert
	assert.NotNil(suite.T(), err)
	assert.Equal(
		suite.T(),
		err,
		errors.New("SCOPE_USE_CASE.SCOPE_TIED_TO_ROLES"),
	)
	scopeRepoMock.AssertCalled(suite.T(), "FindById", scope.Id)
	scopeRepoMock.AssertNumberOfCalls(suite.T(), "FindById", 1)
	scopeRepoMock.AssertNumberOfCalls(suite.T(), "Delete", 0)
}

func TestScopeUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(ScopeUseCaseTestSuite))
}
