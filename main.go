package main

import (
	"alta/immersive-dashboard-api/app/config"
	"alta/immersive-dashboard-api/app/database"
	"alta/immersive-dashboard-api/app/routers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := config.InitConfig()
	database := database.InitDB(config)

	// 
	echo := echo.New()
	echo.Pre(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	routers.InitRouters(database, echo)

	echo.Logger.Fatal(echo.Start(":8080"))
}