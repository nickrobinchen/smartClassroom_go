package router

import (
	"github.com/labstack/echo/v4"
	"github.com/nickrobinchen/smartClassroom_go/controller"
	"github.com/nickrobinchen/smartClassroom_go/middleware"
)

func initUserAPI(group *echo.Group) {
	group.POST("/login", controller.LoginHandler)
	group.GET("/getInfo", controller.GetUserInfoHandler, middleware.GetUserToken)
	// group.GET("/:id", controller.HandlerGetDeviceV1, middleware.JWT)
	// group.PUT("/:id/warning", controller.HandlerTurnOffDeviceWarningV1, middleware.JWT)
	// group.POST("/data", controller.HandlerAddReportDataV1) // for IoT device
	// group.GET("/data", controller.HandlerGetReportDataV1, middleware.JWT)
}
