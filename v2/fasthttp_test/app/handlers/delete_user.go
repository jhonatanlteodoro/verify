package handlers

import (
	"fmt"
	"strconv"

	"github.com/jhonatanlteodoro/fasthttp_test/app/helpers"
	routing "github.com/qiangxue/fasthttp-routing"
)

func (repo *Repository) DeleteUserById(ctx *routing.Context) error {
	id, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		helpers.RespondBadRequest(ctx, "invalid user id")
		return nil
	}

	err = repo.DB.DeleteUser(id)
	if err.Error() == "resource not found" {
		helpers.RespondNotFound(ctx, "user not found")
		return nil

	} else if err != nil {
		fmt.Println(err)
		helpers.RespondServerError(ctx, "fail deleting user")
		return nil
	}

	helpers.RespondOK(ctx, nil)
	return nil
}
