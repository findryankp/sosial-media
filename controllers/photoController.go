package controllers

import (
	"net/http"
	"sosialMedia/configs"
	"sosialMedia/helpers"
	"sosialMedia/middlewares"
	"sosialMedia/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func byIdPhoto(id string, model *models.Photo) (bool, interface{}) {
	query := configs.DB.Joins("User").First(&model, id)
	if query.Error != nil {
		return false, query.Error
	}
	return true, nil
}

func GetPhotos(c *gin.Context) {
	var photos []models.Photo
	configs.DB.Joins("User").Find(&photos)
	helpers.ResponseJson(c, http.StatusOK, true, "-", photos)
}

func GetPhoto(c *gin.Context) {
	var photos models.Photo
	isExist, err := byIdPhoto(c.Param("id"), &photos)
	if !isExist {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "Data not found", err)
		return
	}

	helpers.ResponseJson(c, http.StatusOK, true, "-", photos)
}

func GetPhotoByIdUser(c *gin.Context) {
	var photos []models.Photo

	userIdParams := c.Param("id")
	userId, _ := strconv.Atoi(userIdParams)

	if !middlewares.CanAccess(userId, c) {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "Don't have an authorization for this endpoint", nil)
		return
	}

	configs.DB.Joins("User").Where("user_id = ?", userId).Find(&photos)
	helpers.ResponseJson(c, http.StatusOK, true, "-", photos)
}

func CreatePhoto(c *gin.Context) {
	var photo models.Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		helpers.ResponseJson(c, http.StatusBadRequest, false, err.Error(), photo)
		return
	}

	photo.UserId = helpers.ClaimToken(c).ID

	record := configs.DB.Create(&photo)
	if record.Error != nil {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "-", record.Error.Error())
		return
	}

	response := photo.PhotoToPhotoResponse()
	helpers.ResponseJson(c, http.StatusCreated, true, "Data Created Successfully", response)
}

func UpdatePhoto(c *gin.Context) {
	var photo models.Photo
	isExist, err := byIdPhoto(c.Param("id"), &photo)
	if !isExist {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "Data not found", err)
		return
	}

	if !middlewares.CanAccess(photo.UserId, c) {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "Don't have an authorization for this endpoint", err)
		return
	}

	var formUpdatePhoto models.Photo
	if err := c.ShouldBindJSON(&formUpdatePhoto); err != nil {
		helpers.ResponseJson(c, http.StatusBadRequest, false, err.Error(), formUpdatePhoto)
		return
	}

	record := configs.DB.Model(&photo).Updates(formUpdatePhoto)
	if record.Error != nil {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "-", record.Error.Error())
		return
	}

	response := photo.PhotoToPhotoResponse()
	helpers.ResponseJson(c, http.StatusOK, true, "Data Updated Successfully", response)
}

func DeletePhoto(c *gin.Context) {
	var photo models.Photo
	isExist, err := byIdPhoto(c.Param("id"), &photo)
	if !isExist {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "Data not found", err)
		return
	}

	if !middlewares.CanAccess(photo.UserId, c) {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "Don't have an authorization for this endpoint", err)
		return
	}

	configs.DB.Delete(&photo)
	helpers.ResponseJson(c, http.StatusOK, true, "Data Deleted Successfully", nil)
}
