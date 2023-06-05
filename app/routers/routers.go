package routers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"alta/immersive-dashboard-api/app/middlewares"
	_classData "alta/immersive-dashboard-api/features/classes/data"
	_classHandler "alta/immersive-dashboard-api/features/classes/handler"
	_classService "alta/immersive-dashboard-api/features/classes/service"
)

func InitRouters(db *gorm.DB, e *echo.Echo) {
	classData := _classData.New(db)
	classService := _classService.New(classData)
	classHandlerAPI := _classHandler.New(classService)
	e.POST("/classes",classHandlerAPI.CreateClass,middlewares.JWTMiddleware())
	e.PUT("/classes/:id",classHandlerAPI.UpdateClass,middlewares.JWTMiddleware())
	e.DELETE("/classes/:id",classHandlerAPI.DeleteClass,middlewares.JWTMiddleware())
}
