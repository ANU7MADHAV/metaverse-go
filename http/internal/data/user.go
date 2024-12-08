package data

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserType string

const (
	admin UserType = "admin"
	user  UserType = "user"
)

type RegisterInput struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Type     UserType `json:"type"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserModel struct {
	db *gorm.DB
}

type User struct {
	gorm.Model
	Username string   `gorm:"unique;not null"`
	Password string   `gorm:"not null"`
	Type     UserType `gorm:"not null"`
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{db: db}
}

func (u *UserModel) Create(input *RegisterInput) error {

	if input.Username == "" || input.Password == "" {
		return errors.New("username and password are required")
	}

	var existingUser User

	result := u.db.Where("username = ?", input.Username).First(&existingUser)

	if result.Error == nil {
		return errors.New("username already exist")
	}

	if result.Error != gorm.ErrRecordNotFound {
		return result.Error
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user := User{
		Username: input.Username,
		Password: string(hashPassword),
		Type:     input.Type,
	}

	return u.db.Create(&user).Error
}
