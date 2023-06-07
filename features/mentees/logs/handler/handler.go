package handler

import (
	"alta/immersive-dashboard-api/app/helper"
	"alta/immersive-dashboard-api/app/middlewares"
	"alta/immersive-dashboard-api/features/mentees/logs"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type LogsHandler struct {
	logsService logs.LogsServiceInterface
}

func New(service logs.LogsServiceInterface) *LogsHandler {
	return &LogsHandler{
		logsService: service,
	}
}

func (handler *LogsHandler) CreateLogs(c echo.Context) error{
	userId := middlewares.ExtracTokenUserId(c)
	logsInput := LogsRequest{}
	errBind := c.Bind(&logsInput)
	if errBind != nil{
		return helper.StatusBadRequestResponse(c, "error bind data")
	}
	logsCore := RequestToCoreLogs(logsInput)

	err := handler.logsService.Add(logsCore,uint(userId) )
	if err != nil{
		if strings.Contains(err.Error(),"validation"){
			return helper.StatusBadRequestResponse(c, err.Error())
		} else {
			return helper.StatusInternalServerError(c, err.Error())
		}
	}
	return helper.StatusOK(c, "insert successfuly")

}

func (handler *LogsHandler) EditLogs(c echo.Context) error{

	userId := middlewares.ExtracTokenUserId(c)
	logsInput := LogsRequest{}

	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil{
		return helper.StatusBadRequestResponse(c, "id error")
	}

	errBind := c.Bind(&logsInput)
	if errBind != nil{
		return helper.StatusBadRequestResponse(c, "bind error, update failed")
	}

	logsCore :=RequestToCoreLogs(logsInput)
	logsCore.UserID = uint(userId)
	errUpdate := handler.logsService.Edit(logsCore,uint(idConv))
	if errUpdate != nil{
		return helper.StatusBadRequestResponse(c, "error update data")
	}

	return helper.StatusOK(c, "update successfuly")
	}