package usecase

import (
	"B/core/ports"
	"B/infr/adapters/model"
	"context"
)

type RepositoryUppercase struct {
	userCase ports.RepositoryUser
}

func NewUserCase(userCase ports.RepositoryUser) *RepositoryUppercase {
	return &RepositoryUppercase{
		userCase: userCase,
	}
}
func (u *RepositoryUppercase) FindAllUser(ctx context.Context) ([]*model.User, error) {
	user, err := u.userCase.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (u *RepositoryUppercase) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := u.userCase.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *RepositoryUppercase) Create(ctx context.Context, user *model.User) (*model.User, error) {
	human, err := u.userCase.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return human, nil
}

func (u *RepositoryUppercase) UpdateByEmail(ctx context.Context, email string, req *model.UserUpdate) (bool, error) {
	mess, err := u.userCase.UpdateByEmail(ctx, email, req)
	if err != nil {
		return false, err
	}
	return mess, nil
}

func (u *RepositoryUppercase) DeleteByEmail(ctx context.Context, email string) (bool, error) {
	mess, err := u.userCase.DeleteByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	return mess, nil
}
