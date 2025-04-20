package grpc

import (
	"context"

	"github.com/ryakadev/rdf-be-auth-svc/domain"
	"github.com/ryakadev/rdf-be-auth-svc/helper"
	"github.com/ryakadev/rdf-be-auth-svc/proto"
)

type ScopeGRPC struct {
	proto.UnimplementedScopeServiceServer
	ScopeUsecase domain.ScopeUsecase
}

func NewScopeGRPC(usecase domain.ScopeUsecase) *ScopeGRPC {
	return &ScopeGRPC{ScopeUsecase: usecase}
}

func (g *ScopeGRPC) CreateScope(ctx context.Context, req *proto.CreateScopeRequest) (*proto.ScopeResponse, error) {
	scope := &domain.Scope{
		Name: req.Name,
	}

	scope, err := g.ScopeUsecase.CreateScope(scope)
	if err != nil {
		return nil, err
	}

	res := &proto.ScopeResponse{
		Scope: &proto.Scope{
			Id:        scope.Id,
			Name:      scope.Name,
			CreatedAt: helper.SafeTimeString(scope.CreatedAt),
			UpdatedAt: helper.SafeTimeString(scope.UpdatedAt),
			DeletedAt: helper.SafeTimeString(scope.DeletedAt),
		},
	}

	return res, nil
}

func (g *ScopeGRPC) GetScope(ctx context.Context, req *proto.GetScopeRequest) (*proto.ScopeResponse, error) {
	scope := &domain.Scope{
		Id: req.Id,
	}

	scope, err := g.ScopeUsecase.GetScope(scope)
	if err != nil {
		return nil, err
	}

	res := &proto.ScopeResponse{
		Scope: &proto.Scope{
			Id:        scope.Id,
			Name:      scope.Name,
			CreatedAt: helper.SafeTimeString(scope.CreatedAt),
			UpdatedAt: helper.SafeTimeString(scope.UpdatedAt),
			DeletedAt: helper.SafeTimeString(scope.DeletedAt),
		},
	}

	return res, nil
}

func (g *ScopeGRPC) ShowScopes(ctx context.Context, req *proto.ListScopeRequest) (*proto.ListScopeResponse, error) {

	scopes, err := g.ScopeUsecase.ShowScopes()
	if err != nil {
		return nil, err
	}

	res := &proto.ListScopeResponse{Scopes: make([]*proto.Scope, 0, len(scopes))}
	for _, v := range scopes {

		scope := &proto.Scope{
			Id:        v.Id,
			Name:      v.Name,
			CreatedAt: helper.SafeTimeString(v.CreatedAt),
			UpdatedAt: helper.SafeTimeString(v.UpdatedAt),
			DeletedAt: helper.SafeTimeString(v.DeletedAt),
		}
		res.Scopes = append(res.Scopes, scope)
	}

	return res, nil
}

func (g *ScopeGRPC) UpdateScope(ctx context.Context, req *proto.UpdateScopeRequest) (*proto.ScopeResponse, error) {
	scope := &domain.Scope{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
	}

	scope, err := g.ScopeUsecase.UpdateScope(scope)
	if err != nil {
		return nil, err
	}

	res := &proto.ScopeResponse{
		Scope: &proto.Scope{
			Id:        scope.Id,
			Name:      scope.Name,
			CreatedAt: helper.SafeTimeString(scope.CreatedAt),
			UpdatedAt: helper.SafeTimeString(scope.UpdatedAt),
			DeletedAt: helper.SafeTimeString(scope.DeletedAt),
		},
	}

	return res, nil
}

func (g *ScopeGRPC) DeleteScope(ctx context.Context, req *proto.DeleteScopeRequest) (*proto.DeleteResponse, error) {
	scope := &domain.Scope{
		Id: req.Id,
	}

	err := g.ScopeUsecase.DeleteScope(scope)
	if err != nil {
		return nil, err
	}

	return &proto.DeleteResponse{}, nil
}
