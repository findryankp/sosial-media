package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title    string `json:"title" gorm:"not null" valid:"required~Please insert title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" gorm:"not null" valid:"required~Please insert photo URL"`
	UserId   int    `json:"user_id"`
	User     *User  `gorm:"foreignKey:UserId"`
}

type PhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (photo *Photo) PhotoToPhotoResponse() PhotoResponse {
	return PhotoResponse{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserId:    photo.UserId,
		CreatedAt: photo.CreatedAt,
		UpdatedAt: photo.UpdatedAt,
	}
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (photo *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(photo)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
