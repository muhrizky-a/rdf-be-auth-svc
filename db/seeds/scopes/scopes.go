package scopes

import (
	"fmt"

	"github.com/ryakadev/rdf-be-auth-svc/domain"
	"gorm.io/gorm"
)

func _GetSeed() []*domain.Scope {
	return []*domain.Scope{
		{
			Name:        "Scope:Create",
			Description: "Create a scope",
		},
		{
			Name:        "Scope:ShowAll",
			Description: "Show all scopes",
		},
		{
			Name:        "Scope:Update",
			Description: "Update a scope",
		},
		{
			Name:        "Scope:Delete",
			Description: "Delete a scope",
		},
	}
}

func ScopeSeeds(db *gorm.DB) {
	// Create sample scopes
	scopes := _GetSeed()

	// Create scopes in the database
	for _, scope := range scopes {
		err := db.Save(&scope).Error
		if err != nil {
			fmt.Printf("Error when create scopes: %s\n", scope.Name)
		} else {
			fmt.Printf("Success create scopes: %s\n", scope.Name)
		}
	}
}
