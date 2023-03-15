package handlers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/jhonatanlteodoro/fasthttp_test/app/helpers"
	"github.com/jhonatanlteodoro/fasthttp_test/app/models"
	routing "github.com/qiangxue/fasthttp-routing"
)

func (repo *Repository) UpdateUser(ctx *routing.Context) error {

	id, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		helpers.RespondBadRequest(ctx, "invalid user id")
		return nil
	}

	var user models.User
	err = json.Unmarshal(ctx.PostBody(), &user)
	if err != nil {
		fmt.Println("fail unmarsheling body")
		helpers.RespondBadRequest(ctx, "invalid body data")
		return nil
	}

	user.ID = uint(id)
	err = repo.DB.UpdateUser(user)
	if err.Error() == "resource not found" {
		helpers.RespondNotFound(ctx, "user not found")
		return nil

	} else if err != nil {
		fmt.Println(err)
		helpers.RespondServerError(ctx, "fail updating user")
		return nil
	}

	helpers.RespondOK(ctx, nil)
	return nil
}
