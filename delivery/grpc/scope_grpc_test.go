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
	gRPCErrorTranslator := exceptions.NewGRPCErrorTranslator()
	validator := helper.NewValidator()
	grpcValidator := helper.NewGRPCValidator(validator, gRPCErrorTranslator)

	db := infrastructure.ConnectDB()
	scopeRepo := repository.NewScopeRepository(db)
	roleScopesRepo := repository.NewRoleScopeRepository(db)
	scopeUC := usecase.NewScopeUsecase(scopeRepo, roleScopesRepo)
	suite.scopeGRPC = NewScopeGRPC(scopeUC, grpcValidator, gRPCErrorTranslator)
}

func (suite *ScopeGRPCTestSuite) TearDownTest() {
	db := infrastructure.ConnectDB()
	db.Exec("DELETE FROM role_scopes WHERE created_at >= ?", suite.startTime)
	db.Exec("DELETE FROM scopes WHERE created_at >= ?", suite.startTime)
}

func (suite *ScopeGRPCTestSuite) TestCreateScopeHandler() {
	// Arrange
	scope := domain.Scope{
		Name:        "ScopeTest:Create",
		Description: "Create a scope",
	}

	suite.T().Run("Create a new Scope", func(t *testing.T) {
		// Action
		res, err := suite.scopeGRPC.CreateScope(
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

	suite.T().Run("Create Scope with empty payload", func(t *testing.T) {
		// Action
		res, err := suite.scopeGRPC.CreateScope(
			context.Background(),
			&proto.CreateScopeRequest{},
		)

		// Assert
		assert.Nil(suite.T(), res)
		assert.NotNil(suite.T(), err)
		assert.Equal(
			suite.T(),
			err.Error(),
			"rpc error: code = InvalidArgument desc = validation error: missing request bodies",
		)
	})

	suite.T().Run("Create duplicate Scope", func(t *testing.T) {
		// Action
		res, err := suite.scopeGRPC.CreateScope(
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
}

func (suite *ScopeGRPCTestSuite) TestShowScopesHandler() {
	// Arrange
	_, _ = suite.scopeGRPC.CreateScope(
		context.Background(),
		&proto.CreateScopeRequest{
			Name:        "ScopeTest:Create",
			Description: "Create a scope",
		},
	)
	_, _ = suite.scopeGRPC.CreateScope(
		context.Background(),
		&proto.CreateScopeRequest{
			Name:        "ScopeTest:ShowAll",
			Description: "Show all scopes",
		},
	)

	suite.T().Run("Show Scopes with two recently-created Scopes", func(t *testing.T) {
		// Action
		res, err := suite.scopeGRPC.ListScope(context.Background(), &proto.ListScopeRequest{})
		firstScope, secondScope := res.Body.Scopes[len(res.Body.Scopes)-2], res.Body.Scopes[len(res.Body.Scopes)-1]

		// Assert
		assert.Nil(suite.T(), err)
		assert.NotNil(suite.T(), res)
		assert.Equal(suite.T(), res.StatusCode, int32(codes.OK))
		assert.Equal(suite.T(), res.Message, "Scope retrieved succesfully")
		assert.NotNil(suite.T(), res.Body.Scopes)
		assert.GreaterOrEqual(suite.T(), len(res.Body.Scopes), 2)
		assert.Contains(
			suite.T(),
			res.Body.Scopes,
			&proto.Scope{
				Id:          firstScope.Id,
				Name:        "ScopeTest:Create",
				Description: "Create a scope",
				CreatedAt:   firstScope.CreatedAt,
				UpdatedAt:   firstScope.UpdatedAt,
			},
		)
		assert.Equal(
			suite.T(),
			firstScope,
			&proto.Scope{
				Id:          firstScope.Id,
				Name:        "ScopeTest:Create",
				Description: "Create a scope",
				CreatedAt:   firstScope.CreatedAt,
				UpdatedAt:   firstScope.UpdatedAt,
			},
		)
		assert.Contains(
			suite.T(),
			res.Body.Scopes,
			&proto.Scope{
				Id:          secondScope.Id,
				Name:        "ScopeTest:ShowAll",
				Description: "Show all scopes",
				CreatedAt:   secondScope.CreatedAt,
				UpdatedAt:   secondScope.UpdatedAt,
			},
		)
		assert.Equal(
			suite.T(),
			secondScope,
			&proto.Scope{
				Id:          secondScope.Id,
				Name:        "ScopeTest:ShowAll",
				Description: "Show all scopes",
				CreatedAt:   secondScope.CreatedAt,
				UpdatedAt:   secondScope.UpdatedAt,
			},
		)
	})
}

func (suite *ScopeGRPCTestSuite) TestGetScopeHandler() {
	// Arrange
	res, _ := suite.scopeGRPC.CreateScope(
		context.Background(),
		&proto.CreateScopeRequest{
			Name:        "ScopeTest:Show",
			Description: "Show a scope",
		},
	)
	scopeId := res.Body.Scope.Id

	suite.T().Run("Get a Scope", func(t *testing.T) {
		// Action
		res, err := suite.scopeGRPC.GetScope(
			context.Background(),
			&proto.GetScopeRequest{Id: scopeId},
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
				Name:        "ScopeTest:Show",
				Description: "Show a scope",
				CreatedAt:   res.Body.Scope.CreatedAt,
				UpdatedAt:   res.Body.Scope.UpdatedAt,
			},
		)
	})

	suite.T().Run("Get a Scope with nonexistent Id", func(t *testing.T) {
		// Action
		res, err := suite.scopeGRPC.GetScope(
			context.Background(),
			&proto.GetScopeRequest{Id: -1337},
		)

		// // Assert
		assert.Nil(suite.T(), res)
		assert.NotNil(suite.T(), err)
		assert.Equal(
			suite.T(),
			err.Error(),
			"rpc error: code = NotFound desc = Scope not found",
		)
	})
}

func (suite *ScopeGRPCTestSuite) TestUpdateScopeHandler() {
	// Arrange
	res, _ := suite.scopeGRPC.CreateScope(
		context.Background(),
		&proto.CreateScopeRequest{
			Name:        "WrongScope:Update",
			Description: "Update a wrong scope",
		},
	)
	scopeId := res.Body.Scope.Id

	suite.T().Run("Update a Scope", func(t *testing.T) {
		// Action
		res, err := suite.scopeGRPC.UpdateScope(
			context.Background(),
			&proto.UpdateScopeRequest{
				Id:          scopeId,
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

	suite.T().Run("Update a Scope with empty payload", func(t *testing.T) {
		// Action
		res, err := suite.scopeGRPC.UpdateScope(
			context.Background(),
			&proto.UpdateScopeRequest{},
		)

		// Assert
		assert.Nil(suite.T(), res)
		assert.NotNil(suite.T(), err)
		assert.Equal(
			suite.T(),
			err.Error(),
			"rpc error: code = InvalidArgument desc = validation error: missing request bodies",
		)
	})

	suite.T().Run("Update a Scope with nonexistent Id", func(t *testing.T) {
		// Action
		res, err := suite.scopeGRPC.UpdateScope(
			context.Background(),
			&proto.UpdateScopeRequest{
				Id:          -1337,
				Name:        "ScopeTest:Update",
				Description: "Update a scope",
			},
		)

		// // Assert
		assert.Nil(suite.T(), res)
		assert.NotNil(suite.T(), err)
		assert.Equal(
			suite.T(),
			err.Error(),
			"rpc error: code = NotFound desc = Scope not found",
		)
	})
}

func (suite *ScopeGRPCTestSuite) TestDeleteScopeHandler() {
	// Arrange
	res, _ := suite.scopeGRPC.CreateScope(
		context.Background(),
		&proto.CreateScopeRequest{
			Name:        "ScopeTest:Delete",
			Description: "Delete a scope",
		},
	)
	scopeId := res.Body.Scope.Id

	suite.T().Run("Delete a Scope", func(t *testing.T) {
		// Action
		res, err := suite.scopeGRPC.DeleteScope(
			context.Background(),
			&proto.DeleteScopeRequest{Id: scopeId},
		)

		// Assert
		assert.Nil(suite.T(), err)
		assert.NotNil(suite.T(), res)
		assert.Equal(suite.T(), res.StatusCode, int32(codes.OK))
		assert.Equal(suite.T(), res.Message, "Scope deleted succesfully")
	})

	suite.T().Run("Delete a Scope with empty payload", func(t *testing.T) {
		// Action
		res, err := suite.scopeGRPC.DeleteScope(
			context.Background(),
			&proto.DeleteScopeRequest{},
		)

		// Assert
		assert.Nil(suite.T(), res)
		assert.NotNil(suite.T(), err)
		assert.Equal(
			suite.T(),
			err.Error(),
			"rpc error: code = InvalidArgument desc = validation error: missing request bodies",
		)
	})

	suite.T().Run("Delete a Scope with nonexistent Id", func(t *testing.T) {
		// Action
		res, err := suite.scopeGRPC.DeleteScope(
			context.Background(),
			&proto.DeleteScopeRequest{Id: -1337},
		)

		// // Assert
		assert.Nil(suite.T(), res)
		assert.NotNil(suite.T(), err)
		assert.Equal(
			suite.T(),
			err.Error(),
			"rpc error: code = NotFound desc = Scope not found",
		)
	})
}

func TestScopeGRPCTestSuite(t *testing.T) {
	suite.Run(t, new(ScopeGRPCTestSuite))
}
