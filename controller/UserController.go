package controller

import (
	"golang-pos/model"
	UserService "golang-pos/services"
)

func UserStore() {
	user := model.User{
		Name:     "alex",
		Email:    "alex@gmail.com",
		Password: "alex@123",
	}

	UserService.Create(user)
}
