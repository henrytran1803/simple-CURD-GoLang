package repository

import (
	"context"
	"simple-CURD-GoLang/service/user/model"
	_ "simple-CURD-GoLang/service/user/model"
)

func (repo *mysqlRepo) GetUser(ctx context.Context, id int) (*model.User, error) {
	var user model.User
	if err := repo.db.Table(model.User{}.TableName()).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
