package controllers

import (
	"net/http"
	"sosialMedia/configs"
	"sosialMedia/helpers"
	"sosialMedia/middlewares"
	"sosialMedia/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	configs.DB.Find(&users)
	helpers.ResponseJson(c, http.StatusOK, true, "-", users)
}

func GetUser(c *gin.Context) {
	var user models.User

	isExist, err := helpers.ById(c.Param("id"), &user)
	if !isExist {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "Data not found", err)
		return
	}

	helpers.ResponseJson(c, http.StatusOK, true, "-", user)
}

func UpdateUser(c *gin.Context) {
	var user models.User

	isExist, err := helpers.ById(c.Param("id"), &user)
	if !isExist {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "Data not found", err)
		return
	}

	if !middlewares.CanAccess(int(user.ID), c) {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "Don't have an authorization for this endpoint", err)
		return
	}

	var formUpdateUser models.User
	if err := c.ShouldBindJSON(&formUpdateUser); err != nil {
		helpers.ResponseJson(c, http.StatusBadRequest, false, err.Error(), user)
		return
	}

	// formUpdateUser.HashPassword(formUpdateUser.Password)

	configs.DB.Model(&user).Updates(formUpdateUser)
	response := user.UserToUserResponse()
	helpers.ResponseJson(c, http.StatusOK, true, "Data Updated Successfully", response)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	isExist, err := helpers.ById(c.Param("id"), &user)
	if !isExist {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "Data not found", err)
		return
	}

	if !middlewares.CanAccess(int(user.ID), c) {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "Don't have an authorization for this endpoint", err)
		return
	}

	configs.DB.Delete(&user)
	helpers.ResponseJson(c, http.StatusOK, true, "Data Deleted Successfully", nil)
}
