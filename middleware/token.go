package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/nickrobinchen/smartClassroom_go/utils"
)

type UserInfoClaim struct {
	UserID int    `json:"user_id`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func gen_token(user_id int) {

}

func parseToken(tokenStr string) (*jwt.Token, error) {
	var claims UserInfoClaim
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}
	return token, nil
}

func GetUserToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenStr := c.Request().Header.Get("Authorization")
		fmt.Printf("tokenStr: %v\n", tokenStr)
		if !strings.HasPrefix(tokenStr, "Bearer ") {
			return utils.ResponseJSON(c, http.StatusUnauthorized, "bad token without bearer", nil)
		}
		tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)
		token, err := parseToken(tokenStr)
		if err != nil {
			return utils.ResponseJSON(c, 401, "bad token(parsing error)", nil)
		}
		fmt.Printf("token: %v\n", token.Claims)
		claims := token.Claims.(*UserInfoClaim)
		fmt.Printf("claims: %v\n", claims)
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)
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

func ValidateJWTToken(tokenStr string) {
	panic("unimplemented")
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
