package routers

import (
	"alta/immersive-dashboard-api/app/middlewares"
	userData "alta/immersive-dashboard-api/features/users/data"
	userHandler "alta/immersive-dashboard-api/features/users/handler"
	userService "alta/immersive-dashboard-api/features/users/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_classData "alta/immersive-dashboard-api/features/classes/data"
	_classHandler "alta/immersive-dashboard-api/features/classes/handler"
	_classService "alta/immersive-dashboard-api/features/classes/service"

	_logsData "alta/immersive-dashboard-api/features/mentees/logs/data"
	_logsHandler "alta/immersive-dashboard-api/features/mentees/logs/handler"
	_logsService "alta/immersive-dashboard-api/features/mentees/logs/service"
)

func InitRouters(db *gorm.DB, e *echo.Echo) {
	UserData := userData.New(db)
	UserService := userService.New(UserData)
	UserHandler := userHandler.New(UserService)
  
  	classData := _classData.New(db)
	classService := _classService.New(classData)
	classHandlerAPI := _classHandler.New(classService)

	logsData := _logsData.New(db)
	logsService := _logsService.New(logsData)
	logsHandlerAPI := _logsHandler.New(logsService)

	e.POST("/users", UserHandler.PostUserHandler, middlewares.JWTMiddleware())
	e.PUT("/users/:id", UserHandler.PutUserHandler, middlewares.JWTMiddleware())
	e.GET("/users", UserHandler.GetAllUsersHandler, middlewares.JWTMiddleware())
	e.DELETE("/users/:id", UserHandler.DeleteUserHandler, middlewares.JWTMiddleware())
	e.POST("/login", UserHandler.PostLoginUserHandler)

	e.POST("/classes",classHandlerAPI.CreateClass,middlewares.JWTMiddleware())
	e.PUT("/classes/:id",classHandlerAPI.UpdateClass,middlewares.JWTMiddleware())
	e.DELETE("/classes/:id",classHandlerAPI.DeleteClass,middlewares.JWTMiddleware())
	e.GET("/classes",classHandlerAPI.GetAll)

	e.POST("/logs",logsHandlerAPI.CreateLogs,middlewares.JWTMiddleware())
}
