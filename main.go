package main

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/nickrobinchen/smartClassroom_go/model"
	"github.com/nickrobinchen/smartClassroom_go/router"
)

type Data struct {
	Roles []string `json:"roles"`
	Token string   `json:"token"`
}
type LoginResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Data   `json:"result"`
}

func initDB() (err error) {
	return nil
	// dsn := "root:1a2b3c@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// return err
}

type UserJWT struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func main() {
	//var j middleware.JwtCustomClaims
	err := model.InitModel()
	if err != nil {
		print(err)
	}
	e := echo.New()
	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"),
	}))
	router.InitRouter(e)
	// e.POST("/user/login", loginHandler)
	// e.GET("/user/getInfo", getInfoHandler)

	e.Logger.Fatal(e.Start("0.0.0.0:5000"))
}
