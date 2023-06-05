package handler

import (
	"alta/immersive-dashboard-api/app/helper"
	"alta/immersive-dashboard-api/app/middlewares"
	"alta/immersive-dashboard-api/features/classes"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ClassHandler struct {
	classService classes.ClassServiceInterface
}

func New(service classes.ClassServiceInterface) *ClassHandler{
	return &ClassHandler{
		classService: service,
	}
}

func (handler *ClassHandler) CreateClass(c echo.Context) error{
	userId := middlewares.ExtracTokenUserId(c)
	classInput := ClassRequest{}
	errBind := c.Bind(&classInput)
	if errBind != nil{
		return c.JSON(http.StatusBadRequest,helper.FailedResponse("error bind data"))
	}
	classCore := RequestToCore(classInput)

	err := handler.classService.Create(classCore,userId )
	if err != nil{
		if strings.Contains(err.Error(),"validation"){
			return c.JSON(http.StatusBadRequest,helper.FailedResponse(err.Error()))
		}else{
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error insert data"+err.Error()))
		}
	}
	return c.JSON(http.StatusOK,helper.SuccessResponse("insert successfuly"))

}

func (handler *ClassHandler) UpdateClass(c echo.Context) error{
	userId := middlewares.ExtracTokenUserId(c)
	id := c.Param("id")
	classInput := ClassRequest{}
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil{
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("id error"))
	}
	errBind := c.Bind(&classInput)
	if errBind != nil{
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("bind error, update failed"))
	}
	classCore :=RequestToCore(classInput)
	errUpdate := handler.classService.Edit(idConv,userId,classCore)
	if errUpdate != nil{
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error update data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("update successfuly"))
}

func (handler *ClassHandler) DeleteClass(c echo.Context) error{
	userId := middlewares.ExtracTokenUserId(c)
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil{
		return c.JSON(http.StatusBadRequest,helper.FailedResponse("Delete error"))
	}

	err := handler.classService.Deleted(idConv,userId)
	if err != nil{
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error delete class"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("delete successfuly"))
}