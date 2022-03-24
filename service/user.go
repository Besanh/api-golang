package service

import (
	"api-golang/repository"
	"context"
	"net/http"
)

type (
	IUser interface {
		GetUser(g context.Context, limit, offset int) (int, interface{})
	}
	User struct {
	}
)

// struct User dang implement IUser
func NewUser() IUser {
	return &User{}
}

func (u *User) GetUser(g context.Context, limit, offset int) (int, interface{}) {
	users, err := repository.UserRepo.SelectUsers(g, limit, offset)
	if err != nil {
		return http.StatusServiceUnavailable, map[string]interface{}{
			"error": err,
		}
	}
	return http.StatusOK, map[string]interface{}{
		"success": users,
	}
}
