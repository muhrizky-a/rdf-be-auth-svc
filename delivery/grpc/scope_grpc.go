package grpc

import (
	"context"

	"github.com/ryakadev/rdf-be-auth-svc/domain"
	"github.com/ryakadev/rdf-be-auth-svc/exceptions"
	"github.com/ryakadev/rdf-be-auth-svc/helper"
	"github.com/ryakadev/rdf-be-auth-svc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ScopeGRPC struct {
	proto.UnimplementedScopeServiceServer
	ScopeUsecase        domain.ScopeUsecase
	GRPCValidator       *helper.GRPCValidator
	GRPCErrorTranslator *exceptions.GRPCErrorTranslator
}

func NewScopeGRPC(usecase domain.ScopeUsecase,
	validator *helper.GRPCValidator,
	gRPCErrorTranslator *exceptions.GRPCErrorTranslator,
) *ScopeGRPC {
	return &ScopeGRPC{
		ScopeUsecase:        usecase,
		GRPCValidator:       validator,
		GRPCErrorTranslator: gRPCErrorTranslator,
	}
}

func (g *ScopeGRPC) CreateScope(ctx context.Context, req *proto.CreateScopeRequest) (*proto.ScopeResponse, error) {
	createScopeRequest := &domain.CreateScopeRequest{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := g.GRPCValidator.Validator.Validate(createScopeRequest); err != nil {
		return nil, g.GRPCValidator.CreateGRPCValidationError(err)
	}

	scope := &domain.Scope{
		Name:        req.Name,
		Description: req.Description,
	}

	scope, err := g.ScopeUsecase.CreateScope(scope)
	if err != nil {
		gRPCError := g.GRPCErrorTranslator.Translate(err)
		errorf := status.Errorf(
			codes.Code(gRPCError.StatusCode),
			"%v", gRPCError.Message,
		)
		return nil, errorf
	}

	res := &proto.ScopeResponse{
		StatusCode: int32(codes.OK),
		Message:    "Scope created succesfully",
		Body: &proto.ScopeResponseBody{
			Scope: &proto.Scope{
				Id:          scope.Id,
				Name:        scope.Name,
				Description: scope.Description,
				CreatedAt:   helper.SafeTimeString(scope.CreatedAt),
				UpdatedAt:   helper.SafeTimeString(scope.UpdatedAt),
				DeletedAt:   helper.SafeTimeString(scope.DeletedAt),
			},
		},
	}

	return res, nil
}

func (g *ScopeGRPC) GetScope(ctx context.Context, req *proto.GetScopeRequest) (*proto.ScopeResponse, error) {
	getScopeRequest := &domain.GetScopeRequest{
		Id: req.Id,
	}

	if err := g.GRPCValidator.Validator.Validate(getScopeRequest); err != nil {
		return nil, g.GRPCValidator.CreateGRPCValidationError(err)
	}

	scope := &domain.Scope{
		Id: req.Id,
	}

	scope, err := g.ScopeUsecase.GetScope(scope)
	if err != nil {
		gRPCError := g.GRPCErrorTranslator.Translate(err)
		errorf := status.Errorf(
			codes.Code(gRPCError.StatusCode),
			"%v", gRPCError.Message,
		)
		return nil, errorf
	}

	res := &proto.ScopeResponse{
		StatusCode: int32(codes.OK),
		Message:    "Scope retrieved succesfully",
		Body: &proto.ScopeResponseBody{
			Scope: &proto.Scope{
				Id:          scope.Id,
				Name:        scope.Name,
				Description: scope.Description,
				CreatedAt:   helper.SafeTimeString(scope.CreatedAt),
				UpdatedAt:   helper.SafeTimeString(scope.UpdatedAt),
				DeletedAt:   helper.SafeTimeString(scope.DeletedAt),
			},
		},
	}

	return res, nil
}

func (g *ScopeGRPC) ListScope(ctx context.Context, req *proto.ListScopeRequest) (*proto.ListScopeResponse, error) {
	scopes, err := g.ScopeUsecase.ShowScopes()
	if err != nil {
		gRPCError := g.GRPCErrorTranslator.Translate(err)
		errorf := status.Errorf(
			codes.Code(gRPCError.StatusCode),
			"%v", gRPCError.Message,
		)
		return nil, errorf
	}

	listScopeResponseBody := &proto.ListScopeResponseBody{Scopes: make([]*proto.Scope, 0, len(scopes))}
	for _, v := range scopes {

		scope := &proto.Scope{
			Id:          v.Id,
			Name:        v.Name,
			Description: v.Description,
			CreatedAt:   helper.SafeTimeString(v.CreatedAt),
			UpdatedAt:   helper.SafeTimeString(v.UpdatedAt),
			DeletedAt:   helper.SafeTimeString(v.DeletedAt),
		}
		listScopeResponseBody.Scopes = append(listScopeResponseBody.Scopes, scope)
	}

	res := &proto.ListScopeResponse{
		StatusCode: int32(codes.OK),
		Message:    "Scope retrieved succesfully",
		Body:       listScopeResponseBody,
	}

	return res, nil
}

func (g *ScopeGRPC) UpdateScope(ctx context.Context, req *proto.UpdateScopeRequest) (*proto.ScopeResponse, error) {
	updateScopeRequest := &domain.UpdateScopeRequest{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
	}

	if err := g.GRPCValidator.Validator.Validate(updateScopeRequest); err != nil {
		return nil, g.GRPCValidator.CreateGRPCValidationError(err)
	}

	scope := &domain.Scope{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
	}

	scope, err := g.ScopeUsecase.UpdateScope(scope)
	if err != nil {
		gRPCError := g.GRPCErrorTranslator.Translate(err)
		errorf := status.Errorf(
			codes.Code(gRPCError.StatusCode),
			"%v", gRPCError.Message,
		)
		return nil, errorf
	}

	res := &proto.ScopeResponse{
		StatusCode: int32(codes.OK),
		Message:    "Scope updated succesfully",
		Body: &proto.ScopeResponseBody{
			Scope: &proto.Scope{
				Id:          scope.Id,
				Name:        scope.Name,
				Description: scope.Description,
				CreatedAt:   helper.SafeTimeString(scope.CreatedAt),
				UpdatedAt:   helper.SafeTimeString(scope.UpdatedAt),
				DeletedAt:   helper.SafeTimeString(scope.DeletedAt),
			},
		},
	}
	return res, nil
}

func (g *ScopeGRPC) DeleteScope(ctx context.Context, req *proto.DeleteScopeRequest) (*proto.DeleteResponse, error) {
	deleteScopeRequest := &domain.DeleteScopeRequest{
		Id: req.Id,
	}

	if err := g.GRPCValidator.Validator.Validate(deleteScopeRequest); err != nil {
		return nil, g.GRPCValidator.CreateGRPCValidationError(err)
	}

	scope := &domain.Scope{
		Id: req.Id,
	}

	err := g.ScopeUsecase.DeleteScope(scope)
	if err != nil {
		gRPCError := g.GRPCErrorTranslator.Translate(err)
		errorf := status.Errorf(
			codes.Code(gRPCError.StatusCode),
			"%v", gRPCError.Message,
		)
		return nil, errorf
	}

	return &proto.DeleteResponse{
		StatusCode: int32(codes.OK),
		Message:    "Scope deleted succesfully",
	}, nil
}
