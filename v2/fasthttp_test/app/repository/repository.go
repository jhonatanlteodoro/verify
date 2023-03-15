package repository

import (
	"github.com/jhonatanlteodoro/fasthttp_test/app/models"
)

type DatabaseRepo interface {
	GetAllUsers() ([]models.User, error)
	GetUser(userID int) (models.User, error)
	CreateUser(user models.User) error
	UpdateUser(user models.User) error
	DeleteUser(userID int) error
}
