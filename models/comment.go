package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId  int    `json:"user_id"`
	User    *User  `gorm:"foreignKey:UserId"`
	PhotoId int    `json:"photo_id"`
	Photo   *Photo `gorm:"foreignKey:PhotoId"`
	Message string `json:"message" gorm:"not null" valid:"required~Please insert message"`
}

type CommentResponse struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoId   int       `json:"photo_id"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (comment *Comment) CommentToCommentResponse() CommentResponse {
	return CommentResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoId:   comment.PhotoId,
		UserId:    comment.UserId,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
}

func (comment *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(comment)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (comment *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(comment)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
