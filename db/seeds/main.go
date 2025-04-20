package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ryakadev/rdf-be-auth-svc/config"
	"github.com/ryakadev/rdf-be-auth-svc/db/seeds/scopes"
	"github.com/ryakadev/rdf-be-auth-svc/infrastructure"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	env := os.Getenv("APP_ENV")
	if env != "PROD" {
		config := config.NewDatabaseConfig()
		os.Setenv("DB_HOST", config.Host)
		os.Setenv("DB_PORT", config.Port)
		os.Setenv("DB_USER", config.User)
		os.Setenv("DB_PASS", config.Pass)
		os.Setenv("DB_NAME", config.Name)
	}

	db := infrastructure.ConnectDB()
	return db
}

func main() {
	db := initDB()

	tableFlag := flag.String("table", "all_table", "specify the table")

	flag.Parse()

	// Retrieve the value of the flag
	table := *tableFlag

	// Your program logic goes here...
	switch {
	case table == "scopes":
		scopes.ScopeSeeds(db)
	default:
		fmt.Println("your argument not found")
	}
}
