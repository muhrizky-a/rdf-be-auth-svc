package grpc

import (
	"context"
	"os"
	"testing"

	"github.com/ryakadev/rdf-be-auth-svc/config"
	"github.com/ryakadev/rdf-be-auth-svc/domain"
	"github.com/ryakadev/rdf-be-auth-svc/infrastructure"
	"github.com/ryakadev/rdf-be-auth-svc/proto"
	"github.com/ryakadev/rdf-be-auth-svc/repository"
	"github.com/ryakadev/rdf-be-auth-svc/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ScopeGRPCTestSuite struct {
	suite.Suite
}

func (suite *ScopeGRPCTestSuite) SetupTest() {
	config := config.NewDatabaseConfig()
	os.Setenv("DB_HOST", config.Host)
	os.Setenv("DB_PORT", config.Port)
	os.Setenv("DB_USER", config.User)
	os.Setenv("DB_PASS", config.Pass)
	os.Setenv("DB_NAME", config.Name)
}

func (suite *ScopeGRPCTestSuite) TearDownSuite() {
	db := infrastructure.ConnectDB()
	db.Exec("DELETE FROM role_scopes WHERE 1=1")
	db.Exec("DELETE FROM scopes WHERE 1=1")
}

func (suite *ScopeGRPCTestSuite) TestScopeHandler() {
	db := infrastructure.ConnectDB()
	scopeRepo := repository.NewScopeRepository(db)
	roleScopesRepo := repository.NewRoleScopeRepository(db)
	scopeUC := usecase.NewScopeUsecase(scopeRepo, roleScopesRepo)
	scopeGRPC := NewScopeGRPC(scopeUC)

	scope := domain.Scope{
		Name:        "Scope:Create",
		Description: "Create a scope",
	}

	suite.T().Run("Create a new Scope with mock", func(t *testing.T) {
		res, err := scopeGRPC.CreateScope(
			context.Background(),
			&proto.CreateScopeRequest{Name: scope.Name, Description: scope.Description},
		)
		assert.Nil(suite.T(), err)
		assert.NotNil(suite.T(), res)
	})

	suite.T().Run("Create additional Scope with the same mock", func(t *testing.T) {
		res, err := scopeGRPC.CreateScope(
			context.Background(),

			&proto.CreateScopeRequest{Name: scope.Name, Description: scope.Description},
		)
		assert.Nil(suite.T(), res)
		assert.NotNil(suite.T(), err)
	})

	suite.T().Run("Show Scopes", func(t *testing.T) {
		res, err := scopeGRPC.ShowScopes(context.Background(), &proto.ListScopeRequest{})
		assert.Nil(suite.T(), err)
		assert.NotNil(suite.T(), res)
		scope.Id = res.Scopes[0].Id
	})

	suite.T().Run("Get a Scope", func(t *testing.T) {
		res, err := scopeGRPC.GetScope(
			context.Background(),
			&proto.GetScopeRequest{Id: scope.Id},
		)
		assert.Nil(suite.T(), err)
		assert.NotNil(suite.T(), res)
	})

	suite.T().Run("Update a Scope", func(t *testing.T) {
		res, err := scopeGRPC.UpdateScope(
			context.Background(),
			&proto.UpdateScopeRequest{
				Id:          scope.Id,
				Name:        "Scope:Update",
				Description: "Update a scope",
			},
		)
		assert.Nil(suite.T(), err)
		assert.NotNil(suite.T(), res)
	})

	suite.T().Run("Delete a Scope", func(t *testing.T) {
		res, err := scopeGRPC.DeleteScope(
			context.Background(),
			&proto.DeleteScopeRequest{Id: scope.Id},
		)
		assert.Nil(suite.T(), err)
		assert.NotNil(suite.T(), res)
	})
}

func TestScopeGRPCTestSuite(t *testing.T) {
	suite.Run(t, new(ScopeGRPCTestSuite))
}
