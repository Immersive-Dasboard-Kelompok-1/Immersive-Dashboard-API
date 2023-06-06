package routers

import (
	"alta/immersive-dashboard-api/app/middlewares"
	userData "alta/immersive-dashboard-api/features/users/data"
	userHandler "alta/immersive-dashboard-api/features/users/handler"
	userService "alta/immersive-dashboard-api/features/users/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouters(db *gorm.DB, e *echo.Echo) {
	UserData := userData.New(db)
	UserService := userService.New(UserData)
	UserHandler := userHandler.New(UserService)

	e.POST("/users", UserHandler.PostUserHandler, middlewares.JWTMiddleware())
	e.PUT("/users/:id", UserHandler.PutUserHandler, middlewares.JWTMiddleware())
	e.GET("/users", UserHandler.GetAllUsersHandler, middlewares.JWTMiddleware())
	e.DELETE("/users/:id", UserHandler.DeleteUserHandler, middlewares.JWTMiddleware())
	e.POST("/login", UserHandler.PostLoginUserHandler)

}