package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/ryakadev/rdf-be-auth-svc/delivery/grpc"
	"github.com/ryakadev/rdf-be-auth-svc/helper"
	"github.com/ryakadev/rdf-be-auth-svc/infrastructure"
	"github.com/ryakadev/rdf-be-auth-svc/proto"
	"github.com/ryakadev/rdf-be-auth-svc/repository"
	"github.com/ryakadev/rdf-be-auth-svc/usecase"

	gogrpc "google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	db := infrastructure.ConnectDB()

	// Run migrations
	if err := helper.RunMigrations(db); err != nil {
		log.Fatal("Migrations failed: ", err)
	}

	// Run seeds
	if err := helper.RunSeeds(db); err != nil {
		log.Fatal("Seeding failed: ", err)
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	validator := helper.NewValidator()
	scopeRepo := repository.NewScopeRepository(db)
	roleScopeRepo := repository.NewRoleScopeRepository(db)
	scopeUC := usecase.NewScopeUsecase(scopeRepo, roleScopeRepo)
	scopeGRPC := grpc.NewScopeGRPC(scopeUC, validator)

	gRPCServer := gogrpc.NewServer()
	proto.RegisterScopeServiceServer(gRPCServer, scopeGRPC)

	log.Println("gRPC server is running on port " + port)
	err = gRPCServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve gRPC server over port "+port+": %v", err)
	}
}
