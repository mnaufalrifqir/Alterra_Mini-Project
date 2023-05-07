package controller

import (
	"net/http"
	"mini_project/database"
	"mini_project/lib"
	"mini_project/model"
	"mini_project/util"

	"github.com/labstack/echo/v4"
)

// Register
func Register(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	errs := util.UserValidate(user)
	if errs != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Failed register input invalid",
			"Errors":  errs,
		})
	}

	hashPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashPassword

	if err := database.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Register",
	})
}

// Login
func Login(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	Token, err := lib.LoginUsers(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Login",
		"token":   Token,
	})
}
