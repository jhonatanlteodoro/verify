package handlers

import (
	"fmt"
	"strconv"

	"github.com/jhonatanlteodoro/fasthttp_test/app/helpers"
	"github.com/jhonatanlteodoro/fasthttp_test/app/models"
	routing "github.com/qiangxue/fasthttp-routing"
)

func (repo *Repository) GetAllUsers(ctx *routing.Context) error {
	// Add pagination

	users, err := repo.DB.GetAllUsers()
	if err != nil {
		fmt.Printf("error: %v", err)
		helpers.RespondServerError(ctx, "server error")
		return nil
	}

	var data []models.APIResponseUser
	for _, user := range users {
		data = append(data, user.ToApiUser())
	}

	helpers.RespondOK(ctx, data)
	return nil
}

func (repo *Repository) GetUserById(ctx *routing.Context) error {
	id, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		fmt.Println(err)
		helpers.RespondBadRequest(ctx, "invalid user id")
		return nil
	}

	user, err := repo.DB.GetUser(id)
	if err != nil {
		fmt.Println(err)
		helpers.RespondNotFound(ctx, nil)
		return nil
	}

	helpers.RespondOK(ctx, user.ToApiUser())
	return nil
}
