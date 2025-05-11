package helper

import (
	"log"
	"os"
	"path/filepath"

	"gorm.io/gorm"
)

func runMigrationFromDatabaseFiles(db *gorm.DB, databasePaths []string) error {
	for _, databasePath := range databasePaths {
		// Read the SQL file content
		path, err := filepath.Abs(databasePath)
		if err != nil {
			log.Fatalf("Error reading SQL file path: %v", err)
			return err
		}

		sqlFile, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("Error reading SQL file: %v", err)
			return err
		}

		// Execute the SQL command
		sqlCommand := string(sqlFile)
		if err := db.Exec(sqlCommand).Error; err != nil {
			log.Printf("Error executing Scope migration command: %v", err)
			return err
		}
	}

	return nil
}

func RunMigrations(db *gorm.DB) error {
	return runMigrationFromDatabaseFiles(
		db,
		[]string{
			"./db/migrations/20250418024648_create-table-scopes.up.sql",
			"./db/migrations/20250418101548_create-table-role-scopes.up.sql",
		},
	)
}
