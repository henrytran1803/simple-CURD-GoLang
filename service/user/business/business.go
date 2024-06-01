package business

import (
	"context"
	"simple-CURD-GoLang/service/user/model"
)

type UserRepository interface {
	CreateNewUser(ctx context.Context, user *model.CreateUserReq) error
	DeleteUser(ctx context.Context, id int) error
	GetUser(ctx context.Context, id int) (*model.User, error)
}

type Business struct {
	repository UserRepository
}

func NewBusiness(repository UserRepository) *Business {
	return &Business{repository: repository}
}

func (biz *Business) CreateUser(ctx context.Context, user *model.CreateUserReq) error {
	err := biz.repository.CreateNewUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (biz *Business) DeleteUser(ctx context.Context, id int) error {
	err := biz.repository.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (biz *Business) GetUser(ctx context.Context, id int) (*model.User, error) {
	user, err := biz.repository.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
