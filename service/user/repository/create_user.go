package repository

import (
	"context"
	"github.com/pkg/errors"
	"simple-CURD-GoLang/service/user/model"
)

func (repo *mysqlRepo) CreateNewUser(ctx context.Context, user *model.CreateUserReq) error {
	if err := repo.db.Table(user.TableName()).Create(user).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
