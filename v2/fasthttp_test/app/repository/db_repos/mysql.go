package db_repos

import (
	"errors"
	"fmt"

	"github.com/jhonatanlteodoro/fasthttp_test/app/models"
)

func (m *testDBRepo) GetAllUsers() ([]models.User, error) {
	var users []models.User

	result := m.DB.Find(&users)
	if result.Error != nil {
		fmt.Println("error loading users")
		fmt.Println(result.Error)
	}
	return users, nil
}

func (m *testDBRepo) GetUser(userID int) (models.User, error) {
	var user models.User
	result := m.DB.First(&user, userID)
	if result.Error != nil {
		fmt.Println("error loading users")
		fmt.Println(result.Error)
		return user, result.Error
	}
	return user, nil
}

func (m *testDBRepo) CreateUser(user models.User) error {
	result := m.DB.Create(&user)
	if result.Error != nil {
		fmt.Println("error creating user")
		fmt.Println(result.Error)
		return result.Error
	}
	fmt.Println("user created")
	return nil
}

func (m *testDBRepo) UpdateUser(user models.User) error {

	result := m.DB.Model(&user).Updates(&user)
	if result.Error != nil {
		fmt.Println("error updating user")
		fmt.Println(result.Error)
		return result.Error
	}

	if result.RowsAffected != 1 {
		fmt.Println("resource not found")
		return errors.New("resource not found")
	}

	fmt.Println("user updated")
	return nil
}

func (m *testDBRepo) DeleteUser(userID int) error {

	result := m.DB.Delete(&models.User{}, userID)
	if result.Error != nil {
		fmt.Println("error deleting user")
		fmt.Println(result.Error)
		return result.Error
	}

	fmt.Println("user deleted")
	return nil
}
