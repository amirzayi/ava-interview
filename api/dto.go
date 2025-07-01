package api

import (
	"database/sql"

	"github.com/amirzayi/ava-interview/database/model"
)

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (u User) CreateParam() model.CreateUserParams {
	isValid := len(u.Phone) > 0
	param := model.CreateUserParams{
		Name:  u.Name,
		Phone: sql.NullString{String: u.Phone, Valid: isValid},
	}
	return param
}

func ModelToDTO(user model.User) User {
	return User{
		ID:    user.ID,
		Name:  user.Name,
		Phone: user.Phone.String,
	}
}

func ModelsToDTOs(users []model.User) []User {
	var dtos []User
	for _, user := range users {
		dtos = append(dtos, ModelToDTO(user))
	}
	return dtos
}
