package controllers

import (
	"net/http"
	"sosialMedia/configs"
	"sosialMedia/helpers"
	"sosialMedia/models"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "-", err.Error())
		return
	}

	record := configs.DB.Create(&user)
	if record.Error != nil {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "-", record.Error.Error())
		return
	}

	response := user.UserToUserResponse()

	helpers.ResponseJson(c, http.StatusCreated, true, "-", response)
}

func LoginUser(c *gin.Context) {
	var user, request models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "-", err.Error())
		return
	}

	record := configs.DB.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "-", record.Error.Error())
		return
	}

	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		helpers.ResponseJson(c, http.StatusUnauthorized, false, "Invalid Credential", nil)
		return
	}

	tokenString, err := helpers.GenerateJWT(int(user.ID), user.Role)
	if err != nil {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "-", err.Error())
		return
	}

	response := map[string]string{
		"token": tokenString,
	}

	helpers.ResponseJson(c, http.StatusOK, true, "-", response)
}

func ClaimToken(c *gin.Context) {
	claim := helpers.ClaimToken(c)
	helpers.ResponseJson(c, http.StatusOK, true, "-", claim)
}

func ChangePasswordUser(c *gin.Context) {
	var input models.ChangePassword
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.ResponseJson(c, http.StatusBadRequest, false, err.Error(), nil)
		c.Abort()
		return
	}

	var user models.User

	credentialError := user.CheckPassword(input.PasswordOld)
	if credentialError != nil {
		helpers.ResponseJson(c, http.StatusBadRequest, false, credentialError.Error(), nil)
		c.Abort()
		return
	}

	data := models.User{
		Password: models.HashPassword(input.PasswordNew),
	}

	configs.DB.Model(&user).Updates(data)

	helpers.ResponseJson(c, http.StatusBadRequest, true, "Password changed successfull", nil)
}
