package models

import "github.com/rafaelrubbioli/espenses/pkg/entity"

func NewUser(user entity.User) *User {
	return &User{
		ID: int(user.ID),
		Name: user.Name,
	}
}
