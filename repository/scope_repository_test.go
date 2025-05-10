package repository

import (
	"errors"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/ryakadev/rdf-be-auth-svc/domain"
	"github.com/ryakadev/rdf-be-auth-svc/infrastructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ScopeTestSuite struct {
	suite.Suite
}

func (suite *ScopeTestSuite) SetupTest() {
	// Load .env in project root directory
	err := godotenv.Load("../.env")

	if err != nil {
		log.Println("Error loading .env file in ScopeTestSuite.SetupTest(). Using default env...")
	}
}

func (suite *ScopeTestSuite) TearDownTest() {
	db := infrastructure.ConnectDB()
	db.Exec("DELETE FROM scopes WHERE 1=1")
}
func (suite *ScopeTestSuite) TestCreateScope() {
	// Arrange
	scope := &domain.Scope{
		Name:        "Scope:Create",
		Description: "Create a scope",
	}
	db := infrastructure.ConnectDB()
	scopeRepo := NewScopeRepository(db)

	// Action
	scope, err := scopeRepo.Create(scope)

	// Assert
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), scope)
	assert.Equal(
		suite.T(),
		scope,
		&domain.Scope{
			Id:          scope.Id,
			Name:        "Scope:Create",
			Description: "Create a scope",
			CreatedAt:   scope.CreatedAt,
			UpdatedAt:   scope.UpdatedAt,
		},
	)

	/// Make sure there are one scope persists in database
	scopes, err := scopeRepo.FindAll()
	assert.Equal(suite.T(), len(scopes), 1)
}

func (suite *ScopeTestSuite) TestCreateScopeWithExistingName() {
	// Arrange
	scope := &domain.Scope{
		Name:        "Scope:Create",
		Description: "Create a scope",
	}
	db := infrastructure.ConnectDB()
	scopeRepo := NewScopeRepository(db)

	/// Creating new scope
	_, err := scopeRepo.Create(scope)
	assert.Nil(suite.T(), err)

	// Action
	scope, err = scopeRepo.Create(scope)

	// Assert
	assert.Nil(suite.T(), scope)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), err, errors.New("SCOPE_REPOSITORY.DUPLICATE_NAME"))
}

func (suite *ScopeTestSuite) TestShowScopes() {
	// Arrange
	db := infrastructure.ConnectDB()
	scopeRepo := NewScopeRepository(db)

	/// Creating two scopes
	_, err := scopeRepo.Create(
		&domain.Scope{
			Name:        "Scope:Create",
			Description: "Create a scope",
		},
	)
	assert.Nil(suite.T(), err)
	_, err = scopeRepo.Create(
		&domain.Scope{
			Name:        "Scope:ShowAll",
			Description: "Show all scopes",
		},
	)
	assert.Nil(suite.T(), err)

	// Action
	scopes, err := scopeRepo.FindAll()
	firstScope, secondScope := scopes[0], scopes[1]

	// Assert
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), scopes)
	assert.Equal(suite.T(), len(scopes), 2)
	assert.Contains(
		suite.T(),
		scopes,
		&domain.Scope{
			Id:          firstScope.Id,
			Name:        "Scope:Create",
			Description: "Create a scope",
			CreatedAt:   firstScope.CreatedAt,
			UpdatedAt:   firstScope.UpdatedAt,
		},
	)
	assert.Equal(
		suite.T(),
		firstScope,
		&domain.Scope{
			Id:          firstScope.Id,
			Name:        "Scope:Create",
			Description: "Create a scope",
			CreatedAt:   firstScope.CreatedAt,
			UpdatedAt:   firstScope.UpdatedAt,
		},
	)
	assert.Contains(
		suite.T(),
		scopes,
		&domain.Scope{
			Id:          secondScope.Id,
			Name:        "Scope:ShowAll",
			Description: "Show all scopes",
			CreatedAt:   secondScope.CreatedAt,
			UpdatedAt:   secondScope.UpdatedAt,
		},
	)
	assert.Equal(
		suite.T(),
		secondScope,
		&domain.Scope{
			Id:          secondScope.Id,
			Name:        "Scope:ShowAll",
			Description: "Show all scopes",
			CreatedAt:   secondScope.CreatedAt,
			UpdatedAt:   secondScope.UpdatedAt,
		},
	)
}

func (suite *ScopeTestSuite) TestShowScopeById() {
	// Arrange
	db := infrastructure.ConnectDB()
	scopeRepo := NewScopeRepository(db)
	scope, err := scopeRepo.Create(
		&domain.Scope{
			Name:        "Scope:Show",
			Description: "Show a scope",
		},
	)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), scope)

	// Action
	scope, err = scopeRepo.FindById(scope.Id)

	// Assert
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), scope)
	assert.Equal(suite.T(),
		scope,
		&domain.Scope{
			Id:          scope.Id,
			Name:        "Scope:Show",
			Description: "Show a scope",
			CreatedAt:   scope.CreatedAt,
			UpdatedAt:   scope.UpdatedAt,
		},
	)
}

func (suite *ScopeTestSuite) TestShowScopeByNonexistentId() {
	// Arrange
	db := infrastructure.ConnectDB()
	scopeRepo := NewScopeRepository(db)
	const nonexistentScopeId = -1337

	// Action
	scope, err := scopeRepo.FindById(nonexistentScopeId)

	// Assert
	assert.Nil(suite.T(), scope)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), err, errors.New("SCOPE_REPOSITORY.SCOPE_NOT_FOUND"))
}

func (suite *ScopeTestSuite) TestUpdateScope() {
	// Arrange
	db := infrastructure.ConnectDB()
	scopeRepo := NewScopeRepository(db)

	/// Creating new scope
	scope, err := scopeRepo.Create(
		&domain.Scope{
			Name:        "WrongScope:Update",
			Description: "Update a scope",
		},
	)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), scope)

	// Action
	scope.Name = "Scope:Update"
	scope, err = scopeRepo.Update(scope)

	// Assert
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), scope)
	assert.Equal(suite.T(), "Scope:Update", scope.Name)

	/// Make sure the updated scope persists in database
	scope, _ = scopeRepo.FindById(scope.Id)
	assert.NotNil(suite.T(), scope)
	assert.Equal(suite.T(),
		scope,
		&domain.Scope{
			Id:          scope.Id,
			Name:        "Scope:Update",
			Description: "Update a scope",
			CreatedAt:   scope.CreatedAt,
			UpdatedAt:   scope.UpdatedAt,
		},
	)
}

func (suite *ScopeTestSuite) TestDeleteScope() {
	// Arrange
	db := infrastructure.ConnectDB()
	scopeRepo := NewScopeRepository(db)

	/// Creating two scopes
	_, err := scopeRepo.Create(
		&domain.Scope{
			Name:        "Scope:Create",
			Description: "Create a scope",
		},
	)
	assert.Nil(suite.T(), err)
	scope, err := scopeRepo.Create(
		&domain.Scope{
			Name:        "Scope:Delete",
			Description: "Delete a scope",
		},
	)
	assert.Nil(suite.T(), err)

	/// Make sure there are two scope persist in database
	scopes, err := scopeRepo.FindAll()
	assert.Equal(suite.T(), len(scopes), 2)

	// Action
	err = scopeRepo.Delete(scope)

	// Assert
	assert.Nil(suite.T(), err)

	/// Make sure there are only one scope persist in database
	scopes, err = scopeRepo.FindAll()
	remainingScope := scopes[0]
	assert.Equal(suite.T(), len(scopes), 1)

	assert.Contains(
		suite.T(),
		scopes,
		&domain.Scope{
			Id:          remainingScope.Id,
			Name:        "Scope:Create",
			Description: "Create a scope",
			CreatedAt:   remainingScope.CreatedAt,
			UpdatedAt:   remainingScope.UpdatedAt,
		},
	)
	assert.Equal(
		suite.T(),
		remainingScope,
		&domain.Scope{
			Id:          remainingScope.Id,
			Name:        "Scope:Create",
			Description: "Create a scope",
			CreatedAt:   remainingScope.CreatedAt,
			UpdatedAt:   remainingScope.UpdatedAt,
		},
	)
}

func TestScopeTestSuite(t *testing.T) {
	suite.Run(t, new(ScopeTestSuite))
}
