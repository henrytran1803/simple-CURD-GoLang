package repository

import (
	"context"
	"simple-CURD-GoLang/service/user/model"
)

func (repo *mysqlRepo) DeleteUser(ctx context.Context, id int) error {
	if err := repo.db.Table(model.CreateUserReq{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{"status": 1}).Error; err != nil {
		return err
	}
	return nil
}
