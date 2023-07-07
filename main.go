package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func main() {
	//var j middleware.JwtCustomClaims
	err := model.InitModel()
	if err != nil {
		print(err)
	}
	e := echo.New()
	// e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
	// 	LogStatus: true,
	// 	LogURI:    true,
	// 	BeforeNextFunc: func(c echo.Context) {
	// 		c.Set("customValueFromContext", 42)
	// 	},
	// 	LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
	// 		value, _ := c.Get("customValueFromContext").(int)
	// 		fmt.Printf("REQUEST: uri: %v, status: %v, custom-value: %v\n", v.URI, v.Status, value)
	// 		return nil
	// 	},
	// }))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	router.InitRouter(e)
	// e.POST("/user/login", loginHandler)
	// e.GET("/user/getInfo", getInfoHandler)
	e.Logger.Fatal(e.Start("0.0.0.0:5000"))
}
