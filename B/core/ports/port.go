package ports

import (
	"B/infr/adapters/model"
	"context"
)

type RepositoryUser interface {
	FindAll(context.Context) ([]*model.User, error)
	FindByEmail(context.Context, string) (*model.User, error)
	Create(context.Context, *model.User) (*model.User, error)
	UpdateByEmail(context.Context, string, *model.UserUpdate) (bool, error)
	DeleteByEmail(context.Context, string) (bool, error)
}
