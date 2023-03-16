package mapper

import (
	"B/infr/adapters/model"
	"B/pb"
)

func MapUserUpdateToUser(userUpdate *model.UserUpdate) *model.User {
	return &model.User{
		Name:    userUpdate.Name,
		Address: userUpdate.Address,
		Phone:   userUpdate.Phone,
	}
}
func ConvertUserInModelToUserPb(req *model.User) *pb.User {
	return &pb.User{
		Id:      req.Id,
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
		Phone:   req.Phone,
	}
}
func ConvertListUserInModelToListUserPb(users []*model.User) []*pb.User {
	var pbUsers []*pb.User
	for _, user := range users {
		pbUser := ConvertUserInModelToUserPb(user)
		pbUsers = append(pbUsers, pbUser)
	}
	return pbUsers
}
func ConvertUserInModelToRequestCreate(user *pb.RequestCreate) *model.User {
	req := user.User
	return &model.User{
		Id:      req.Id,
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
		Phone:   req.Phone,
	}
}
func ConvertUserRequestTo() {

}
