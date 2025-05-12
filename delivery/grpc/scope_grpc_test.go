package grpc

import (
	"context"
	"log"
	"testing"
	"time"

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
	"google.golang.org/grpc/codes"
)

type ScopeGRPCTestSuite struct {
	suite.Suite
	startTime time.Time
	scopeGRPC *ScopeGRPC
}

func (suite *ScopeGRPCTestSuite) SetupTest() {
	// Load .env in project root directory
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Println("Error loading .env file in ScopeGRPCTestSuite.SetupTest(). Using default env...")
	}
	suite.startTime = time.Now()
}

func (suite *ScopeGRPCTestSuite) TearDownSuite() {
	db := infrastructure.ConnectDB()
	db.Exec("DELETE FROM role_scopes WHERE created_at >= ?", suite.startTime)
	db.Exec("DELETE FROM scopes WHERE created_at >= ?", suite.startTime)
}

func (suite *ScopeGRPCTestSuite) TestScopeHandler() {
	gRPCErrorTranslator := exceptions.NewGRPCErrorTranslator()
	validator := helper.NewValidator()
	grpcValidator := helper.NewGRPCValidator(validator, gRPCErrorTranslator)

	db := infrastructure.ConnectDB()
	scopeRepo := repository.NewScopeRepository(db)
	roleScopesRepo := repository.NewRoleScopeRepository(db)
	scopeUC := usecase.NewScopeUsecase(scopeRepo, roleScopesRepo)
	scopeGRPC := NewScopeGRPC(scopeUC, grpcValidator, gRPCErrorTranslator)

	// Arrange
	scope := domain.Scope{
		Name:        "ScopeTest:Create",
		Description: "Create a scope",
	}

	suite.T().Run("Create a new Scope", func(t *testing.T) {
		// Action
		res, err := scopeGRPC.CreateScope(
			context.Background(),
			&proto.CreateScopeRequest{Name: scope.Name, Description: scope.Description},
		)
		scope.Id = res.Body.Scope.Id

		// Assert
		assert.Nil(suite.T(), err)
		assert.NotNil(suite.T(), res)
		assert.Equal(suite.T(), res.StatusCode, int32(codes.OK))
		assert.Equal(suite.T(), res.Message, "Scope created succesfully")
		assert.NotNil(suite.T(), res.Body.Scope)
		assert.Equal(
			suite.T(),
			res.Body.Scope,
			&proto.Scope{
				Id:          res.Body.Scope.Id,
				Name:        "ScopeTest:Create",
				Description: "Create a scope",
				CreatedAt:   res.Body.Scope.CreatedAt,
				UpdatedAt:   res.Body.Scope.UpdatedAt,
			},
		)
	})

	suite.T().Run("Create duplicate Scope", func(t *testing.T) {
		// Action
		res, err := scopeGRPC.CreateScope(
			context.Background(),
			&proto.CreateScopeRequest{Name: scope.Name, Description: scope.Description},
		)

		// Assert
		assert.Nil(suite.T(), res)
		assert.NotNil(suite.T(), err)
		assert.Equal(
			suite.T(),
			err.Error(),
			"rpc error: code = InvalidArgument desc = Scope with such name already exists",
		)
	})

	suite.T().Run("Show Scopes", func(t *testing.T) {
		// Action
		res, err := scopeGRPC.ListScope(context.Background(), &proto.ListScopeRequest{})
		latestScope := res.Body.Scopes[len(res.Body.Scopes)-1]

		// Assert
		assert.Nil(suite.T(), err)
		assert.NotNil(suite.T(), res)
		assert.Equal(suite.T(), res.StatusCode, int32(codes.OK))
		assert.Equal(suite.T(), res.Message, "Scope retrieved succesfully")
		assert.NotNil(suite.T(), res.Body.Scopes)
		assert.GreaterOrEqual(suite.T(), len(res.Body.Scopes), 1)
		assert.Equal(
			suite.T(),
			latestScope,
			&proto.Scope{
				Id:          latestScope.Id,
				Name:        "ScopeTest:Create",
				Description: "Create a scope",
				CreatedAt:   latestScope.CreatedAt,
				UpdatedAt:   latestScope.UpdatedAt,
			},
		)
	})

	suite.T().Run("Get a Scope", func(t *testing.T) {
		// Action
		res, err := scopeGRPC.GetScope(
			context.Background(),
			&proto.GetScopeRequest{Id: scope.Id},
		)

		// Assert
		assert.Nil(suite.T(), err)
		assert.NotNil(suite.T(), res)
		assert.Equal(suite.T(), res.StatusCode, int32(codes.OK))
		assert.Equal(suite.T(), res.Message, "Scope retrieved succesfully")
		assert.NotNil(suite.T(), res.Body.Scope)
		assert.Equal(
			suite.T(),
			res.Body.Scope,
			&proto.Scope{
				Id:          res.Body.Scope.Id,
				Name:        "ScopeTest:Create",
				Description: "Create a scope",
				CreatedAt:   res.Body.Scope.CreatedAt,
				UpdatedAt:   res.Body.Scope.UpdatedAt,
			},
		)
	})

	suite.T().Run("Update a Scope", func(t *testing.T) {
		// Action
		res, err := scopeGRPC.UpdateScope(
			context.Background(),
			&proto.UpdateScopeRequest{
				Id:          scope.Id,
				Name:        "ScopeTest:Update",
				Description: "Update a scope",
			},
		)

		// Assert
		assert.Nil(suite.T(), err)
		assert.NotNil(suite.T(), res)
		assert.Equal(suite.T(), res.StatusCode, int32(codes.OK))
		assert.Equal(suite.T(), res.Message, "Scope updated succesfully")
		assert.NotNil(suite.T(), res.Body.Scope)
		assert.Equal(
			suite.T(),
			res.Body.Scope,
			&proto.Scope{
				Id:          res.Body.Scope.Id,
				Name:        "ScopeTest:Update",
				Description: "Update a scope",
				CreatedAt:   res.Body.Scope.CreatedAt,
				UpdatedAt:   res.Body.Scope.UpdatedAt,
			},
		)
	})

	suite.T().Run("Delete a Scope", func(t *testing.T) {
		// Action
		res, err := scopeGRPC.DeleteScope(
			context.Background(),
			&proto.DeleteScopeRequest{Id: scope.Id},
		)

		// Assert
		assert.Nil(suite.T(), err)
		assert.NotNil(suite.T(), res)
		assert.Equal(suite.T(), res.StatusCode, int32(codes.OK))
		assert.Equal(suite.T(), res.Message, "Scope deleted succesfully")
	})
}

func TestScopeGRPCTestSuite(t *testing.T) {
	suite.Run(t, new(ScopeGRPCTestSuite))
}
