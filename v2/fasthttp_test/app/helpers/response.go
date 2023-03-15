package helpers

import (
	"encoding/json"
	"fmt"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func responseDefault(ctx *routing.Context, data []byte, statusCode int) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	ctx.SetBody(data)
}

func RespondOK(ctx *routing.Context, data interface{}) {
	resp, err := json.Marshal(map[string]interface{}{
		"data": data,
	})

	if err != nil {
		fmt.Println("Fail marsheling")
		return
	}

	responseDefault(ctx, resp, fasthttp.StatusOK)
}

func RespondNotFound(ctx *routing.Context, data interface{}) {

	resp, err := json.Marshal(map[string]interface{}{
		"data": data,
	})

	if err != nil {
		fmt.Println("Fail marsheling")
		return
	}
	responseDefault(ctx, resp, fasthttp.StatusNotFound)
}

func RespondBadRequest(ctx *routing.Context, data interface{}) {

	resp, err := json.Marshal(map[string]interface{}{
		"msg": data,
	})

	if err != nil {
		fmt.Println("Fail marsheling")
		return
	}

	responseDefault(ctx, resp, fasthttp.StatusBadRequest)
}

func RespondServerError(ctx *routing.Context, data interface{}) {

	resp, err := json.Marshal(map[string]interface{}{
		"error": data,
	})

	if err != nil {
		fmt.Println("Fail marsheling")
		return
	}

	responseDefault(ctx, resp, fasthttp.StatusInternalServerError)
}
