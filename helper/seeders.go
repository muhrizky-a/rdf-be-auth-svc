package helper

import (
	"fmt"
	"log"
	"strings"

	"gorm.io/gorm"
)

func getScopeSeed() []*map[string]string {
	return []*map[string]string{
		{
			"name":        "Scope:Create",
			"description": "Create a scope",
		},
		{
			"name":        "Scope:ShowAll",
			"description": "Show all scopes",
		},
		{
			"name":        "Scope:Update",
			"description": "Update a scope",
		},
		{
			"name":        "Scope:Delete",
			"description": "Delete a scope",
		},
		{
			"name":        "Role:Create",
			"description": "Create a role",
		},
		{
			"name":        "Role:ShowAll",
			"description": "Show all roles",
		},
		{
			"name":        "Role:Update",
			"description": "Update a role",
		},
		{
			"name":        "Role:Delete",
			"description": "Delete a role",
		},
		{
			"name":        "Account:Create",
			"description": "Create an account",
		},
		{
			"name":        "Account:ShowAll",
			"description": "Show all accounts",
		},
		{
			"name":        "Account:Update",
			"description": "Update an account",
		},
		{
			"name":        "Account:Delete",
			"description": "Delete an account",
		},
	}
}

func runScopeSeed(db *gorm.DB) error {
	// Create sample scopes
	scopes := getScopeSeed()

	// Building query
	/// Final query with three data insertion should be:
	/// INSERT INTO scopes (name, description) VALUES (?,?), (?,?), (?,?) ON CONFLICT (name) DO NOTHING;
	query := "INSERT INTO scopes (name, description) VALUES "
	for _, scope := range scopes {
		query = query + fmt.Sprintf("('%s', '%s'),", (*scope)["name"], (*scope)["description"])
	}
	/// Slice the last comma
	query = strings.TrimSuffix(query, ",")
	query = query + " ON CONFLICT (name) DO NOTHING;"

	// Create scopes in the database
	err := db.Exec(
		query,
	).Error
	if err != nil {
		log.Fatalf("Error when create scopes:: ", err)
	} else {
		log.Println("Success create scopes")
	}

	return err
}

func RunSeeds(db *gorm.DB) error {
	return runScopeSeed(db)
}
