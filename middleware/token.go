package middleware

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// JwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type JwtCustomClaims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.Claims
}

func (j JwtCustomClaims) Valid() error {
	//TODO implement me
	panic("implement me")
}

func gen_token(user_id int) {

}

func GetUserToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token) // by default token is stored under `user` key
		if !true {
			fmt.Println("JWT token missing or invalid")
			return errors.New("JWT token missing or invalid")
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.ErrBadRequest
		}
		c.Set("user_id", claims["user_id"])
		c.Set("role", claims["role"])
		return next(c)
		// tokenStr := ctx.Request().Header.Get("Authorization")
		// if !strings.HasPrefix(tokenStr, "Bearer ") {
		// 	return util.FailedResp(ctx, http.StatusUnauthorized, "bad token", "token not set")
		// }
		// tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)
		// idHex, err := util.ValidateJWTToken(tokenStr)
		// if err != nil {
		// 	return util.FailedResp(ctx, http.StatusUnauthorized, "bad token", err.Error())
		// }
		// id, err := primitive.ObjectIDFromHex(idHex)
		// if err != nil {
		// 	return util.FailedResp(ctx, http.StatusUnauthorized, "bad token", err.Error())
		// }
		// ctx.Set("id", id)
		// return next(ctx)
	}
}

// func login(c echo.Context) error {
// 	username := c.FormValue("username")
// 	password := c.FormValue("password")

// 	// Throws unauthorized error
// 	if username != "jon" || password != "shhh!" {
// 		return echo.ErrUnauthorized
// 	}

// 	// Set custom claims

// 	// Generate encoded token and send it as response.
// 	t, err := token.SignedString([]byte("secret"))
// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, echo.Map{
// 		"token": t,
// 	})
// }

// func accessible(c echo.Context) error {
// 	return c.String(http.StatusOK, "Accessible")
// }

// func restricted(c echo.Context) error {
// 	user := c.Get("user_id").(*jwt.Token)
// 	claims := user.Claims.(*JwtCustomClaims)
// 	name := claims.UserID

// 	return c.String(http.StatusOK, "Welcome "+name+"!")
// }

// func main() {
// 	e := echo.New()

// 	// Middleware
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	// Login route
// 	e.POST("/login", login)

// 	// Unauthenticated route
// 	e.GET("/", accessible)

// 	// Restricted group
// 	r := e.Group("/restricted")

// 	// Configure middleware with the custom claims type
// 	config := echojwt.Config{
// 		NewClaimsFunc: func(c echo.Context) jwt.Claims {
// 			return new(JwtCustomClaims)
// 		},
// 		SigningKey: []byte("secret"),
// 	}
// 	r.Use(echojwt.WithConfig(config))
// 	r.GET("", restricted)

// 	e.Logger.Fatal(e.Start(":1323"))
// }
