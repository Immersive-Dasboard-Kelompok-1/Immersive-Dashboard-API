package handler

import (
	"alta/immersive-dashboard-api/app/helper"
	"alta/immersive-dashboard-api/app/middlewares"
	"alta/immersive-dashboard-api/features/classes"
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
		return helper.StatusBadRequestResponse(c, "error bind data")
	}
	classCore := RequestToCore(classInput)

	err := handler.classService.Create(classCore,userId )
	if err != nil{
		if strings.Contains(err.Error(),"validation"){
			return helper.StatusBadRequestResponse(c, err.Error())
		} else {
			return helper.StatusInternalServerError(c)
		}
	}
	return helper.StatusOK(c, "insert successfuly")

}

func (handler *ClassHandler) UpdateClass(c echo.Context) error{
	userId := middlewares.ExtracTokenUserId(c)
	id := c.Param("id")
	classInput := ClassRequest{}
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil{
		return helper.StatusBadRequestResponse(c, "id error")
	}
	errBind := c.Bind(&classInput)
	if errBind != nil{
		return helper.StatusBadRequestResponse(c, "bind error, update failed")
	}
	classCore :=RequestToCore(classInput)
	errUpdate := handler.classService.Edit(idConv,userId,classCore)
	if errUpdate != nil{
		return helper.StatusBadRequestResponse(c, "error update data")
	}
	return helper.StatusOK(c, "update successfuly")
}

func (handler *ClassHandler) DeleteClass(c echo.Context) error{
	userId := middlewares.ExtracTokenUserId(c)
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil{
		return helper.StatusBadRequestResponse(c, "Delete error")
	}

	err := handler.classService.Deleted(idConv,userId)
	if err != nil{
		return helper.StatusBadRequestResponse(c, "error delete class")
	}
	return helper.StatusOK(c, "delete successfuly")
}

func (handler *ClassHandler) GetAll(c echo.Context) error{

	dataClass, err := handler.classService.GetAll()
	if err != nil{
		return helper.StatusBadRequestResponse(c, "error read class")
	}
	var ClassResponAll []Response
	for _,value := range dataClass{
		dataResponse :=CoreToResponse(value)
		ClassResponAll = append(ClassResponAll, dataResponse)
	}
	return helper.StatusOKWithData(c, "Success read data class", ClassResponAll)
}