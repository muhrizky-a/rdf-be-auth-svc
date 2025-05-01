package grpc

import (
	"context"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/ryakadev/rdf-be-auth-svc/domain"
	"github.com/ryakadev/rdf-be-auth-svc/exceptions"
	"github.com/ryakadev/rdf-be-auth-svc/helper"
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
	// Load .env in project root directory
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Println("Error loading .env file in ScopeGRPCTestSuite.SetupTest(). Using default env...")
	}
}

func (suite *ScopeGRPCTestSuite) TearDownSuite() {
	db := infrastructure.ConnectDB()
	db.Exec("DELETE FROM role_scopes WHERE 1=1")
	db.Exec("DELETE FROM scopes WHERE 1=1")
}

func (suite *ScopeGRPCTestSuite) TestScopeHandler() {
	httpErrorTranslator := exceptions.NewHTTPErrorTranslator()
	validator := helper.NewValidator()
	grpcValidator := helper.NewGRPCValidator(validator, httpErrorTranslator)

	db := infrastructure.ConnectDB()
	scopeRepo := repository.NewScopeRepository(db)
	roleScopesRepo := repository.NewRoleScopeRepository(db)
	scopeUC := usecase.NewScopeUsecase(scopeRepo, roleScopesRepo)
	scopeGRPC := NewScopeGRPC(scopeUC, grpcValidator, httpErrorTranslator)

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
		res, err := scopeGRPC.ListScope(context.Background(), &proto.ListScopeRequest{})
		assert.Nil(suite.T(), err)
		assert.NotNil(suite.T(), res)
		scope.Id = res.Body.Scopes[0].Id
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
