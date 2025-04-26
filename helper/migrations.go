package helper

import (
	"log"

	"github.com/ryakadev/rdf-be-auth-svc/domain"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	if err := db.AutoMigrate(&domain.Scope{}); err != nil {
		log.Fatalf("Migration Scope failed: %v", err)
		return err
	}

	if err := db.AutoMigrate(&domain.RoleScope{}); err != nil {
		log.Fatalf("Migration RoleScope failed: %v", err)
		return err
	}

	return nil
}
