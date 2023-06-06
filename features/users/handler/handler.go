package handler

import (
	"alta/immersive-dashboard-api/app/helper"
	"alta/immersive-dashboard-api/app/middlewares"
	"alta/immersive-dashboard-api/features/users"
	"alta/immersive-dashboard-api/features/users/data"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService users.UserServiceInterface
}

func New(service users.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) PostUserHandler(c echo.Context) error {
	payload := data.Users{}
	if err := c.Bind(&payload); err != nil {
		if err == echo.ErrBadRequest {
			return helper.StatusBadRequestResponse(c, "error bind payload " + err.Error())
		} else {
			return helper.StatusInternalServerError(c)
		}
	}

	userId := middlewares.ExtracTokenUserId(c)
	userLoggedIn, err := handler.userService.GetUser(uint(userId))
	if err != nil {
		return err
	}
	
	if userLoggedIn.Role == "admin" {
		payloadMap := users.Core{
			FullName: payload.FullName,
			Email: payload.Email,
			Password: payload.Password,
			Team: payload.Team,
			Role: payload.Role,
			Status: payload.Status,
		}
		if err := handler.userService.AddUser(payloadMap); err != nil {
			if strings.Contains(err.Error(), "validation") {
				return helper.StatusBadRequestResponse(c, "error validate payload: " + err.Error())
			} else {
				return helper.StatusInternalServerError(c)
			}
		} 
		return helper.StatusCreated(c, "User berhasil ditambahkan", userId)
	} else {
		return helper.StatusForbiddenResponse(c, "Anda haruslah seorang admin agar bisa menambahkan user")
	}
}

func (handler *UserHandler) PutUserHandler(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))
	newData := data.Users{}
	if errBind := c.Bind(&newData); errBind != nil {
		if errBind == echo.ErrBadRequest {
			return helper.StatusBadRequestResponse(c, "error bind payload " + errBind.Error())
		}
	}

	user := middlewares.ExtracTokenUserId(c)
	userLoggedIn, err := handler.userService.GetUser(uint(user))
	if err != nil {
		return err
	}
	
	if userLoggedIn.Role == "admin" {
		newDataMap := users.Core{
			FullName:newData.FullName,
			Email: newData.Email,
			Password:newData.Password,
			Team:newData.Team,
			Role:newData.Role,
			Status:newData.Status,
		}
		if err := handler.userService.EditUser(uint(userId), newDataMap); err != nil {
			if strings.Contains(err.Error(), "validation") {
				return helper.StatusBadRequestResponse(c, "error validate payload: " + err.Error())
			} else {
				return helper.StatusInternalServerError(c)
			}
		}
		return helper.StatusOK(c, "Berhasil memperbarui data pengguna")
	} else {
		return helper.StatusForbiddenResponse(c, "Anda harus admin jika ingin mengedit resource ini")
	}
}

func (handler *UserHandler) GetAllUsersHandler(c echo.Context) error {
	allUsers, errGetAll := handler.userService.GetAllUser()
	if errGetAll != nil {
		return helper.StatusInternalServerError(c)
	}
	return helper.StatusOKWithData(c, "Berhasil mendapatkan semua data pengguna terdaftar", allUsers)
}

func (handler *UserHandler) DeleteUserHandler(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))
	user := middlewares.ExtracTokenUserId(c)
	userLoggedIn, err := handler.userService.GetUser(uint(user))
	if err != nil {
		return err
	}

	if userLoggedIn.Role == "admin" {
		if errDelete := handler.userService.DeleteUser(uint(userId)); errDelete != nil {
			if strings.Contains(err.Error(), "validation") {
				return helper.StatusBadRequestResponse(c, "error validate payload: " + err.Error())
			} else {
				return helper.StatusInternalServerError(c)
			}
		}
		return helper.StatusOK(c, "Berhasil menghapus data pengguna")
	} else {
		return helper.StatusForbiddenResponse(c, "Anda harus admin jika ingin menghapus resource ini")
	}
}

func (handler *UserHandler) PostLoginUserHandler(c echo.Context) error {
	var payload users.LoginUser
	if errBind := c.Bind(&payload); errBind != nil {
		if errBind == echo.ErrBadRequest {
			return helper.StatusBadRequestResponse(c, "error bind payload " + errBind.Error())
		}
	}

	userId, err := handler.userService.LoginUser(payload.Email, payload.Password)
	if err != nil {
		if strings.Contains(err.Error(), "kredensial tidak cocok") {
			return helper.StatusBadRequestResponse(c, "Kredensial yang anda berikan tidak valid") 
		}
	}

	accessToken, err := middlewares.CreateToken(userId)
	if err != nil {
		return err
	}

	return helper.StatusCreated(c, "Login Berhasil", map[string]any{
		"accessToken": accessToken, 
	})
}
