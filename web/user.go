package web

import (
	"basket-keeper/model"
	"basket-keeper/util"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type createUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type loginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func CreateUser(c echo.Context) error {
	var db, derr = util.ConnectToSQLite()
	if derr != nil {
		return c.String(http.StatusInternalServerError, "Error connecting to the database")
	}
	var req createUserRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}
	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := model.CreateUser(db, req.Username, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, user)
}

func LoginUser(c echo.Context) error {
	var db, derr = util.ConnectToSQLite()
	if derr != nil {
		return c.String(http.StatusInternalServerError, "Error connecting to the database")
	}
	var req loginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}
	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	token, err := model.LoginUser(db, req.Username, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func DeleteUser(c echo.Context) error {
	var db, derr = util.ConnectToSQLite()
	if derr != nil {
		return c.String(http.StatusInternalServerError, "Error connecting to the database")
	}
	var userIDStr = c.Param("id")
	var userID, err = strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}
	err = model.DeleteUser(db, userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func GetUserIDFromContext(c echo.Context) (int64, error) {
	var userID = c.Get("user_id")
	if userID == nil {
		return 0, errors.New("user ID not found in context")
	}
	var userIDInt64, ok = userID.(int64)
	if !ok {
		return 0, errors.New("unexpected user ID type in context")
	}
	return userIDInt64, nil
}
