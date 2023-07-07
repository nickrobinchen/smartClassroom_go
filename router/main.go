package router

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func InitRouter(echo *echo.Echo) {
	// groupAPI := echo.Group("/api")

	// groupAPIV1 := groupAPI.Group("/v1")

	groupUser := echo.Group("/user")
	initUserAPI(groupUser)
	// groupTeacher := echo.Group("/teacher")
	// initDeviceGroupV1(groupDeviceV1)
	// initAPIGroupV1(groupAPIV1)
	fmt.Printf("[Router] Init done")
}

func initAPIGroupV1(group *echo.Group) {
}
