package model

import "simple-CURD-GoLang/comon"

type User struct {
	comon.SQLModel
	RoleId  int    `json:"role_id" gorm:"column:role_id"`
	Name    string `json:"name" gorm:"column:name"`
	Email   string `json:"email" gorm:"column:email"`
	Address string `json:"addr" gorm:"column:addr"`
}
type CreateUserReq struct {
	comon.SQLModel
	Name  string `json:"name" gorm:"column:name"`
	Email string `json:"email" gorm:"column:email"`
}

func (CreateUserReq) TableName() string { return "users" }
func (User) TableName() string          { return "users" }
