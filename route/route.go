package route

import (
	"mini_project/constants"
	"mini_project/controller"
	"mini_project/middleware"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func StartRoute() *echo.Echo {
	e := echo.New()

	middleware.LogMiddleware(e)

	noAuth := e.Group("")
	noAuth.POST("/register", controller.Register)
	noAuth.POST("/login", controller.Login)

	Auth := e.Group("/auth", echojwt.JWT([]byte(constants.SECRET_JWT)))
	Auth.POST("/task", controller.CreateTaskController)
	Auth.GET("/task/all", controller.GetTasksController)
	Auth.PUT("/task/:id", controller.UpdateTaskController)
	Auth.DELETE("/task/:id", controller.DeleteTaskController)
	Auth.GET("/task", controller.GetTaskController)
	Auth.PUT("/task/status/:id", controller.ChangeStatusController)
	Auth.GET("/task/timeline", controller.ShowTimelineController)

	return e
}
