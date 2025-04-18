package usecase

import (
	"os"
	"testing"

	"github.com/ryakadev/rdf-be-auth-svc/config"
	"github.com/ryakadev/rdf-be-auth-svc/domain"
	"github.com/ryakadev/rdf-be-auth-svc/infrastructure"
	"github.com/ryakadev/rdf-be-auth-svc/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type RepoRoleScopeMock struct {
	mock.Mock
}

type ScopeUseCaseTestSuite struct {
	suite.Suite
}

func (r *RepoRoleScopeMock) FindRoleScopesByScopeId(id int64) ([]*domain.RoleScope, error) {
	args := r.Called(id)
	return args.Get(0).([]*domain.RoleScope), args.Error(1)
}

func (suite *ScopeUseCaseTestSuite) SetupTest() {
	config := config.NewDatabaseConfig()
	os.Setenv("DB_HOST", config.Host)
	os.Setenv("DB_PORT", config.Port)
	os.Setenv("DB_USER", config.User)
	os.Setenv("DB_PASS", config.Pass)
	os.Setenv("DB_NAME", config.Name)
}

func (suite *ScopeUseCaseTestSuite) TearDownTest() {
	db := infrastructure.ConnectDB()
	db.Exec("DELETE FROM role_scopes WHERE 1=1")
	db.Exec("DELETE FROM scopes WHERE 1=1")
}

func (suite *ScopeUseCaseTestSuite) TestCreateScope() {

	db := infrastructure.ConnectDB()
	scopeRepo := repository.NewScopeRepository(db)
	roleScopeRepo := repository.NewRoleScopeRepository(db)
	scopeUC := NewScopeUsecase(scopeRepo, roleScopeRepo)
	scope := &domain.Scope{
		Name:        "Account:Create",
		Description: "Create an account",
	}
	// Create a new Scope with mock
	scope, err := scopeUC.CreateScope(scope)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), scope)
}

func (suite *ScopeUseCaseTestSuite) TestShowScope() {

	db := infrastructure.ConnectDB()
	scopeRepo := repository.NewScopeRepository(db)
	roleScopeRepo := repository.NewRoleScopeRepository(db)
	ucScope := NewScopeUsecase(scopeRepo, roleScopeRepo)

	//Show a Scope
	scopes, err := ucScope.ShowScopes()
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), scopes)
}

func (suite *ScopeUseCaseTestSuite) TestUpdateScope() {
	db := infrastructure.ConnectDB()
	scopeRepo := repository.NewScopeRepository(db)
	roleScopeRepo := repository.NewRoleScopeRepository(db)
	scopeUC := NewScopeUsecase(scopeRepo, roleScopeRepo)
	scope := &domain.Scope{
		Name:        "WrongAccount:Update",
		Description: "Update an account",
	}

	//Update a Scope
	scope, err := scopeUC.CreateScope(scope)
	assert.Nil(suite.T(), err)

	scope.Name = "Account:Update"
	_, err = scopeUC.UpdateScope(scope)
	assert.Nil(suite.T(), err)
}

func (suite *ScopeUseCaseTestSuite) TestDeleteScope() {
	db := infrastructure.ConnectDB()
	scopeRepo := repository.NewScopeRepository(db)
	roleScopeRepo := repository.NewRoleScopeRepository(db)
	scopeUC := NewScopeUsecase(scopeRepo, roleScopeRepo)
	scope := &domain.Scope{
		Name:        "Account:Delete",
		Description: "Delete an account",
	}

	scope, err := scopeUC.CreateScope(scope)
	assert.Nil(suite.T(), err)

	//Delete a Scope
	err = scopeUC.DeleteScope(scope)
	assert.Nil(suite.T(), err)
}

func (suite *ScopeUseCaseTestSuite) TestDeleteScopeWithRolesAttached() {
	db := infrastructure.ConnectDB()
	scopeRepo := repository.NewScopeRepository(db)
	roleScopeRepo := repository.NewRoleScopeRepository(db)
	scopeUC := NewScopeUsecase(scopeRepo, roleScopeRepo)
	scope := &domain.Scope{
		Name:        "Account:Delete",
		Description: "Delete an account",
	}

	scope, err := scopeUC.CreateScope(scope)
	assert.Nil(suite.T(), err)

	//Delete a Scope
	roleScope := &domain.RoleScope{
		RoleId:  1,
		ScopeId: scope.Id,
	}

	//Delete a Scope
	_, err = roleScopeRepo.Create(roleScope)
	assert.Nil(suite.T(), err)

	err = scopeUC.DeleteScope(scope)
	assert.NotNil(suite.T(), err)
}

func TestScopeUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(ScopeUseCaseTestSuite))
}
