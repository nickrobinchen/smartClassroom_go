package utils

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type ResponseModel struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
	Msg    string      `json:"msg"`
}

func ResponseJSON(ctx echo.Context, code int, msg string, object interface{}) error {
	if code != 200 {
		fmt.Println(msg)
	}
	return ctx.JSON(code, ResponseModel{Code: code, Result: object, Msg: msg})
}
