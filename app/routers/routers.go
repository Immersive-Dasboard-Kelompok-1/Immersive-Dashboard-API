package routers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_classData "alta/immersive-dashboard-api/features/classes/data"
	_classHandler "alta/immersive-dashboard-api/features/classes/handler"
	_classService "alta/immersive-dashboard-api/features/classes/service"
)

func InitRouters(db *gorm.DB, e *echo.Echo) {
	classData := _classData.New(db)
	classService := _classService.New(classData)
	classHandlerAPI := _classHandler.New(classService)
	e.POST("/classes",classHandlerAPI.CreateClass)
	e.PUT("/classes/:id",classHandlerAPI.UpdateClass)
}
