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
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := infrastructure.ConnectDB()

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
