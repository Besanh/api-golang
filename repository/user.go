package repository

import (
	"api-golang/common/model"
	"context"
)

type (
	IUser interface {
		SelectUsers(c context.Context, limit, offset int) (*[]model.User, error)
	}
	User struct{}
)

var UserRepo IUser

func NewUserRepo() IUser {
	return &User{}
}

func (u *User) SelectUsers(c context.Context, limit, offset int) (*[]model.User, error) {
	// Select struct can dung new()
	users := new([]model.User)
	query := Db.NewSelect().Model(users).Limit(limit).Offset(offset)
	// Scan query result
	err := query.Scan(c)
	if err != nil {
		return users, err
	}
	return users, nil
}
