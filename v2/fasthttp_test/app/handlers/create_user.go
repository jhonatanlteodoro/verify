package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/jhonatanlteodoro/fasthttp_test/app/helpers"
	"github.com/jhonatanlteodoro/fasthttp_test/app/models"
	routing "github.com/qiangxue/fasthttp-routing"
)

func (repo *Repository) CreateUser(ctx *routing.Context) error {

	var user models.User

	err := json.Unmarshal(ctx.PostBody(), &user)
	if err != nil {
		fmt.Println(err)
		helpers.RespondBadRequest(ctx, "invalid body data")
		return nil
	}

	hashedPass, err := helpers.GenerateHashPassword(user.Password)
	if err != nil {
		helpers.RespondServerError(ctx, "server error")
	}

	user.Password = hashedPass

	err = repo.DB.CreateUser(user)
	if err != nil {
		fmt.Println(err)
		helpers.RespondServerError(ctx, "fail creating user")
		return nil
	}

	helpers.RespondOK(ctx, nil)
	return nil
}
