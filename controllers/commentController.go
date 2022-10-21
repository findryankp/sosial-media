package controllers

import (
	"net/http"
	"sosialMedia/configs"
	"sosialMedia/helpers"
	"sosialMedia/middlewares"
	"sosialMedia/models"

	"github.com/gin-gonic/gin"
)

func byIdComment(id string, model *models.Comment) (bool, interface{}) {
	query := configs.DB.Preload("User").Preload("Photo").First(&model, id)
	if query.Error != nil {
		return false, query.Error
	}
	return true, nil
}

func GetComments(c *gin.Context) {
	var comments []models.Comment
	configs.DB.Preload("User").Preload("Photo").Find(&comments)
	helpers.ResponseJson(c, http.StatusOK, true, "-", comments)
}

func GetComment(c *gin.Context) {
	var comment models.Comment
	isExist, err := byIdComment(c.Param("id"), &comment)
	if !isExist {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "Data not found", err)
		return
	}

	helpers.ResponseJson(c, http.StatusOK, true, "-", comment)
}

func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		helpers.ResponseJson(c, http.StatusBadRequest, false, err.Error(), comment)
		return
	}
	comment.UserId = helpers.ClaimToken(c).ID

	record := configs.DB.Create(&comment)
	if record.Error != nil {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "-", record.Error.Error())
		return
	}

	response := comment.CommentToCommentResponse()

	helpers.ResponseJson(c, http.StatusCreated, true, "Data Created Successfully", response)
}

func UpdateComment(c *gin.Context) {
	var comment models.Comment
	isExist, err := byIdComment(c.Param("id"), &comment)
	if !isExist {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "Data not found", err)
		return
	}

	if !middlewares.CanAccess(comment.UserId, c) {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "Don't have an authorization for this endpoint", err)
		return
	}

	var formUpdateComment models.Comment
	if err := c.ShouldBindJSON(&formUpdateComment); err != nil {
		helpers.ResponseJson(c, http.StatusBadRequest, false, err.Error(), formUpdateComment)
		return
	}

	record := configs.DB.Model(&comment).Updates(formUpdateComment)
	if record.Error != nil {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "-", record.Error.Error())
		return
	}

	response := comment.CommentToCommentResponse()
	helpers.ResponseJson(c, http.StatusOK, true, "Data Updated Successfully", response)
}

func DeleteComment(c *gin.Context) {
	var comment models.Comment
	isExist, err := byIdComment(c.Param("id"), &comment)
	if !isExist {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "Data not found", err)
		return
	}

	if !middlewares.CanAccess(comment.UserId, c) {
		helpers.ResponseJson(c, http.StatusBadRequest, false, "Don't have an authorization for this endpoint", err)
		return
	}

	configs.DB.Delete(&comment)
	helpers.ResponseJson(c, http.StatusOK, true, "Data Deleted Successfully", nil)
}
