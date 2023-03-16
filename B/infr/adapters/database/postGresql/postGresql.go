package postGresql

import (
	"B/infr/adapters/mapper"
	"B/infr/adapters/model"
	"context"
	"errors"
	"gorm.io/gorm"
)

type postGreSql struct {
	collection *gorm.DB
}

func NewPostGreSql(collection *gorm.DB) *postGreSql {
	return &postGreSql{
		collection: collection,
	}
}
func (k *postGreSql) FindAll(ctx context.Context) ([]*model.User, error) {
	user := make([]*model.User, 0)
	result := k.collection.Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
func (k *postGreSql) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user *model.User
	result := k.collection.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
func (k *postGreSql) Create(ctx context.Context, user *model.User) (*model.User, error) {
	result := k.collection.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
func (k *postGreSql) UpdateByEmail(ctx context.Context, email string, req *model.UserUpdate) (bool, error) {
	result := k.collection.Model(&model.User{}).Where("email = ?", email).Updates(mapper.MapUserUpdateToUser(req))
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, errors.New("no user found with the provided email")
	}
	return true, nil
}
func (k *postGreSql) DeleteByEmail(ctx context.Context, email string) (bool, error) {
	result := k.collection.Where("email = ?", email).Delete(&model.User{})
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, errors.New("no user found with the provided email")
	}
	return true, nil
}
