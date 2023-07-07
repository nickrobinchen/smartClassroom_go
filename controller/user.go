package controller

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/nickrobinchen/smartClassroom_go/model"
	"github.com/nickrobinchen/smartClassroom_go/utils"
	"gorm.io/gorm"
)

func LoginHandler(c echo.Context) error {
	//var params LoginParams
	var manager model.Manager
	var msg string
	var params struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}
	err := c.Bind(&params)
	if err != nil {
		return utils.ResponseJSON(c, 205, "Login error(binding)", nil)
	}
	fmt.Printf("params: %v\n", params)
	result := model.DB.Where("account = ? AND password = ?", params.Account, params.Password).First(&manager)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		msg = "No such account or incorrect password!"
		return utils.ResponseJSON(c, 205, msg, nil)
	} else {
		msg = "Login succeed"
		claims := &struct {
			UserID int    `json:"user_id"`
			Role   string `json:"role"`
			//Expire int    `json:expire_time`
			jwt.RegisteredClaims
		}{
			int(manager.ID),
			"manager",
			//Expire: int(jwt.NewNumericDate(time.Now().Add(time.Hour * 72)).Unix()),
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		data := echo.Map{
			"token":       t,
			"roles":       []string{"manager"},
			"expire_time": claims.ExpiresAt.Unix()}
		return utils.ResponseJSON(c, 200, "Login Success", data)
	}
	//return utils.ResponseJSON(c, 205, "Unknown server error.", nil)
}

func GetUserInfoHandler(c echo.Context) error {
	user_id := c.Get("user_id")
	role := c.Get("role").(string)
	name := ""
	if role == "manager" {
		var manager model.Manager
		model.DB.Where("id = ?", user_id).First(&manager)
		name = manager.Name

	}
	//dict(roles=[dict(roleName=role, value=role)], userId=user_id, username=user.name, realName=user.name,
	//                avatar='', homePath=homePath)
	data := echo.Map{
		"roles":    [1]echo.Map{{"roleName": role, "value": role}},
		"userId":   user_id,
		"username": name}
	return utils.ResponseJSON(c, 200, "Get user info success", data)
}
