package utils

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ResponseModel struct {
	Code   string      `json:"code"`
	Result interface{} `json:"result"`
	Msg    string      `json:"msg"`
}

func ResponseJSON(ctx echo.Context, code int, msg string, object interface{}) error {
	if code != 200 {
		fmt.Println(msg)
	}
	return ctx.JSON(code, ResponseModel{Code: strconv.Itoa(code), Result: object, Msg: msg})
}
