package controllers

import (
	"net/http"
	"sosialMedia/configs"
	"sosialMedia/helpers"
	"sosialMedia/middlewares"
	"sosialMedia/models"

	"github.com/gin-gonic/gin"
)

func byIdSosialMedia(id, model interface{}) (bool, interface{}) {
	query := configs.DB.Joins("User").First(&model, id)
	if query.Error != nil {
		return false, query.Error
	}
	return true, nil
}

func GetSosialMedias(c *gin.Context) {
	var sosialMedias []models.SosialMedia
	configs.DB.Joins("User").Find(&sosialMedias)
	helpers.ResponseJson(c, http.StatusOK, true, "-", sosialMedias)
}

func GetSosialMedia(c *gin.Context) {
	var sosialMedia models.SosialMedia
	isExist, err := byIdSosialMedia(c.Param("id"), &sosialMedia)
	if !isExist {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "Data not found", err)
		return
	}

	helpers.ResponseJson(c, http.StatusOK, true, "-", sosialMedia)
}

func CreateSosialMedia(c *gin.Context) {
	var sosialMedia models.SosialMedia
	if err := c.ShouldBindJSON(&sosialMedia); err != nil {
		helpers.ResponseJson(c, http.StatusBadRequest, false, err.Error(), sosialMedia)
		return
	}

	sosialMedia.UserId = helpers.ClaimToken(c).ID

	record := configs.DB.Create(&sosialMedia)
	if record.Error != nil {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "-", record.Error.Error())
		return
	}

	response := sosialMedia.SosialMediaToSosialMediaResponse()

	helpers.ResponseJson(c, http.StatusCreated, true, "Data Created Successfully", response)
}

func UpdateSosialMedia(c *gin.Context) {
	var sosialMedia models.SosialMedia
	isExist, err := byIdSosialMedia(c.Param("id"), &sosialMedia)
	if !isExist {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "Data not found", err)
		return
	}

	if !middlewares.CanAccess(sosialMedia.UserId, c) {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "Don't have an authorization for this endpoint", err)
		return
	}

	var formUpdateSosialMedia models.SosialMedia
	if err := c.ShouldBindJSON(&formUpdateSosialMedia); err != nil {
		helpers.ResponseJson(c, http.StatusBadRequest, false, err.Error(), formUpdateSosialMedia)
		return
	}

	record := configs.DB.Model(&sosialMedia).Updates(formUpdateSosialMedia)
	if record.Error != nil {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "-", record.Error.Error())
		return
	}

	response := sosialMedia.SosialMediaToSosialMediaResponse()
	helpers.ResponseJson(c, http.StatusOK, true, "Data Updated Successfully", response)
}

func DeleteSosialMedia(c *gin.Context) {
	var sosialMedia models.SosialMedia
	isExist, err := byIdSosialMedia(c.Param("id"), &sosialMedia)
	if !isExist {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "Data not found", err)
		return
	}

	if !middlewares.CanAccess(sosialMedia.UserId, c) {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "Don't have an authorization for this endpoint", err)
		return
	}

	configs.DB.Delete(&sosialMedia)
	helpers.ResponseJson(c, http.StatusOK, true, "Data Deleted Successfully", nil)
}
