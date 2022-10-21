package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
)

type SosialMedia struct {
	gorm.Model
	Name           string `json:"name" gorm:"not null" valid:"required~Please insert name"`
	SosialMediaUrl string `json:"sosial_media_url" gorm:"not null" valid:"required~Please insert sosial media url"`
	UserId         int    `json:"user_id"`
	User           *User  `gorm:"foreignKey:UserId"`
}

type SosialMediaResponse struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SosialMediaUrl string    `json:"sosial_media_url"`
	UserId         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (sosialMedia *SosialMedia) SosialMediaToSosialMediaResponse() SosialMediaResponse {
	return SosialMediaResponse{
		ID:             sosialMedia.ID,
		Name:           sosialMedia.Name,
		SosialMediaUrl: sosialMedia.SosialMediaUrl,
		UserId:         sosialMedia.UserId,
		CreatedAt:      sosialMedia.CreatedAt,
		UpdatedAt:      sosialMedia.UpdatedAt,
	}
}

func (sosialMedia *SosialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(sosialMedia)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (sosialMedia *SosialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(sosialMedia)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
