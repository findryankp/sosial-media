package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null;size:256;uniqueIndex"  valid:"required~Please insert username"`
	Email    string `json:"email" gorm:"uniqueIndex;not null;size:256" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string `json:"password" gorm:"not null" form:"password" valid:"required~Your password is required,stringlength(6|100)~Minimal character of password is 6"`
	Age      int    `json:"age" gorm:"not null" valid:"range(8|150)"`
	Role     string `json:"role"`
}

type ChangePassword struct {
	PasswordOld string `json:"password_old" binding:"required"`
	PasswordNew string `json:"password_new" binding:"required"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) UserToUserResponse() UserResponse {
	return UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Age:       user.Age,
		UpdatedAt: user.UpdatedAt,
	}
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(user)

	if errCreate != nil {
		err = errCreate
		return
	}
	user.Password = HashPassword(user.Password)

	err = nil
	return
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "err"
	}

	password = string(bytes)
	return password
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
