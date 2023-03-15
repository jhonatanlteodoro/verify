package db_repos

import (
	"github.com/jhonatanlteodoro/fasthttp_test/app/models"
)

func (m *mysqlDBRepo) GetAllUsers() ([]models.User, error) {
	var users []models.User
	return users, nil
}

func (m *mysqlDBRepo) GetUser(userID int) (models.User, error) {
	var user models.User
	return user, nil
}

func (m *mysqlDBRepo) CreateUser(user models.User) error {
	return nil
}

func (m *mysqlDBRepo) UpdateUser(user models.User) error {
	return nil
}

func (m *mysqlDBRepo) DeleteUser(userID int) error {
	return nil
}
