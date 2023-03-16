package handler_grpc

import (
	"B/core/usecase"
	"B/infr/adapters/mapper"
	"B/infr/adapters/model"
	"B/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverGrpc struct {
	*pb.UnimplementedUserServiceServer
	handleGrpc *usecase.RepositoryUppercase
}

func NewHandlerGrpc(grpc *usecase.RepositoryUppercase) *serverGrpc {
	return &serverGrpc{
		handleGrpc: grpc,
	}
}
func (p *serverGrpc) List(ctx context.Context, req *pb.RequestList) (*pb.ReposeList, error) {
	user, err := p.handleGrpc.FindAllUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
	}
	return &pb.ReposeList{User: mapper.ConvertListUserInModelToListUserPb(user)}, nil
}
func (p *serverGrpc) FindByEmail(ctx context.Context, req *pb.RequestFindByEmail) (*pb.ReposeFindByEmail, error) {
	user, err := p.handleGrpc.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "method FindByEmail not implemented")
	}
	return &pb.ReposeFindByEmail{User: mapper.ConvertUserInModelToUserPb(user)}, nil
}
func (p *serverGrpc) Create(ctx context.Context, req *pb.RequestCreate) (*pb.User, error) {
	user, err := p.handleGrpc.Create(ctx, mapper.ConvertUserInModelToRequestCreate(req))
	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
	}
	return &pb.User{
		Id:      user.Id,
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}, nil
}
func (p *serverGrpc) UpdateByEmail(ctx context.Context, req *pb.RequestUpdateByEmail) (*pb.ResponseUpdateByEmail, error) {
	user := model.UserUpdate{
		Name:    req.Fields.Name,
		Address: req.Fields.Address,
		Phone:   req.Fields.Phone,
	}
	message, err := p.handleGrpc.UpdateByEmail(ctx, req.Email, &user)

	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "method UpdateByEmail not implemented")
	}
	return &pb.ResponseUpdateByEmail{Message: message}, nil
}
func (p *serverGrpc) DeleteByEmail(ctx context.Context, req *pb.RequestDeleteByEmail) (*pb.ReposeDeleteByEmail, error) {
	mess, err := p.handleGrpc.DeleteByEmail(ctx, req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "method DeleteByEmail not implemented")
	}
	return &pb.ReposeDeleteByEmail{Message: mess}, nil
}
