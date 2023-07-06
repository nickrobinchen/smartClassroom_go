package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type User struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
type Manage struct {
	gorm.Model
	Account  string `json:"account"`
	Password string `json:"password"`
}

func main() {
	e := echo.New()
	// e.POST("/user/login", func(c echo.Context) error {
	// 	var u User
	// 	//fmt.Printf("c.QueryParam(\"username\"): %v\n", c.QueryParam("username"))
	// 	// 调用 echo.Context 的 Bind 函数将请求参数和 User 对象进行绑定。
	// 	err := c.Bind(&u)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	fmt.Printf("u: %v\n", u)
	// 	return c.JSON(http.StatusOK, "ss") //c.String(http.StatusOK, "Hello World! I am here!")
	// })

	// e.Logger.Fatal(e.Start("0.0.0.0:5000"))
}
