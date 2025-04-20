package repository

import (
	"os"
	"testing"
	"time"

	"github.com/ryakadev/rdf-be-auth-svc/config"
	"github.com/ryakadev/rdf-be-auth-svc/domain"
	"github.com/ryakadev/rdf-be-auth-svc/infrastructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type RepoScopeMock struct {
	mock.Mock
}

type ScopeTestSuite struct {
	suite.Suite
}

func (suite *ScopeTestSuite) SetupTest() {
	config := config.NewDatabaseConfig()
	os.Setenv("DB_HOST", config.Host)
	os.Setenv("DB_PORT", config.Port)
	os.Setenv("DB_USER", config.User)
	os.Setenv("DB_PASS", config.Pass)
	os.Setenv("DB_NAME", config.Name)
}

func (suite *ScopeTestSuite) TearDownTest() {
	db := infrastructure.ConnectDB()
	db.Exec("DELETE FROM scopes WHERE 1=1")
}

func (r *RepoScopeMock) Create(scope *domain.Scope) (*domain.Scope, error) {
	args := r.Called(scope)
	return args.Get(0).(*domain.Scope), args.Error(1)
}

func (r *RepoScopeMock) FindAll() ([]*domain.Scope, error) {
	args := r.Called()
	return args.Get(0).([]*domain.Scope), args.Error(1)
}

func (r *RepoScopeMock) Update(scope *domain.Scope) (*domain.Scope, error) {
	args := r.Called(scope)
	return args.Get(0).(*domain.Scope), args.Error(1)
}

func (r *RepoScopeMock) Delete(scope *domain.Scope) error {
	args := r.Called(scope)
	return args.Error(0)
}

func (suite *ScopeTestSuite) TestCreateScope() {
	// Create a new Scope with mock
	repoScopeMock := RepoScopeMock{}
	scope := &domain.Scope{
		Name:        "Scope:Create",
		Description: "Create a scope",
	}
	now := time.Now()
	CreateScopeReponse := &domain.Scope{
		Id:          1,
		Name:        "Scope:Create",
		Description: "Create a scope",
		CreatedAt:   &now,
	}
	repoScopeMock.On("Create", scope).Return(CreateScopeReponse, nil)

	scope, err := repoScopeMock.Create(scope)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), scope)

	// Create a new Scope to DB
	db := infrastructure.ConnectDB()
	repoScope := NewScopeRepository(db)

	scope, err = repoScope.Create(scope)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), scope)
}

func (suite *ScopeTestSuite) TestCreateScopeWithExistingName() {
	db := infrastructure.ConnectDB()
	repoScope := NewScopeRepository(db)
	scope := &domain.Scope{
		Name:        "Scope:ShowAll",
		Description: "Show all scopes",
	}

	// Create a new Scope with Existing Name to DB
	scope, err := repoScope.Create(scope)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), scope)

	scope, err = repoScope.Create(scope)
	assert.Nil(suite.T(), scope)
	assert.NotNil(suite.T(), err)

}

func (suite *ScopeTestSuite) TestShowScope() {
	db := infrastructure.ConnectDB()
	scopeRepo := NewScopeRepository(db)

	// Show a Scope
	scopes, err := scopeRepo.FindAll()
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), scopes)
}

func (suite *ScopeTestSuite) TestShowScopeById() {
	db := infrastructure.ConnectDB()
	repoScope := NewScopeRepository(db)
	scope := &domain.Scope{
		Name:        "Scope:Show",
		Description: "Show a scope",
	}

	// Create a new Scope with Existing Name to DB
	scope, err := repoScope.Create(scope)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), scope)

	// Show a Scope
	scope, err = repoScope.FindById(scope.Id)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), scope)
}

func (suite *ScopeTestSuite) TestUpdateScope() {

	repoScopeMock := RepoScopeMock{}
	scope := &domain.Scope{
		Name:        "WrongScope:Update",
		Description: "Update a scope",
	}
	now := time.Now()
	CreateScopeReponse := &domain.Scope{
		Id:          1,
		Name:        "WrongScope:Update",
		Description: "Update a scope",
		CreatedAt:   &now,
	}
	repoScopeMock.On("Create", scope).Return(CreateScopeReponse, nil)
	CreateScopeReponse = &domain.Scope{
		Id:          1,
		Name:        "Scope:Update",
		Description: "Update a scope",
		CreatedAt:   &now,
	}
	now = time.Now()
	UpdateScopeResponse := &domain.Scope{
		Id:          1,
		Name:        "Scope:Update",
		Description: "Update a scope",
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
	repoScopeMock.On("Update", CreateScopeReponse).Return(UpdateScopeResponse, nil)

	// Update a Scope with mock
	scope, err := repoScopeMock.Create(scope)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), scope)

	scope, err = repoScopeMock.Update(CreateScopeReponse)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "Scope:Update", scope.Name)

	db := infrastructure.ConnectDB()
	repoScope := NewScopeRepository(db)

	//Update a Scope to DB
	scope, err = repoScope.Create(scope)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), scope)

	scope.Name = "Scope:Update"
	scope, err = repoScope.Update(scope)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "Scope:Update", scope.Name)

}

func (suite *ScopeTestSuite) TestDeleteScope() {
	db := infrastructure.ConnectDB()
	repoScope := NewScopeRepository(db)
	scope := &domain.Scope{
		Name:        "Scope:Delete",
		Description: "Delete a scope",
	}

	//Delete a Scope
	scope, err := repoScope.Create(scope)
	assert.Nil(suite.T(), err)

	err = repoScope.Delete(scope)
	assert.Nil(suite.T(), err)

}

func TestScopeTestSuite(t *testing.T) {
	suite.Run(t, new(ScopeTestSuite))
}
