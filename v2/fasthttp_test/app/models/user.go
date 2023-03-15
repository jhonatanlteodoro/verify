package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Age      int
	Email    string
	Password string
	Address  string
}

type APIResponseUser struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func (u *User) ToApiUser() APIResponseUser {
	return APIResponseUser{
		ID:      int(u.ID),
		Name:    u.Name,
		Age:     u.Age,
		Email:   u.Email,
		Address: u.Address,
	}
}
