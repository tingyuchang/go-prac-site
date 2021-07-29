package models

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"go-prac-site/internal/types"
	"gorm.io/gorm"
	"time"
)

type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Password string `json:"password"`

}

type User struct {
	gorm.Model
	Id         string         `json:"uid" gorm:"column:id;autoIncrement"`
	Name       string         `json:"username" gorm:"column:name" validate:"required"`
	Password   string         `json:"-" gorm:"column:password" `
	Email      string         `json:"email" gorm:"column:email" validate:"required,email"`
	Phone      string         `json:"phone" gorm:"column:phone"`
	Gender     string         `json:"gender" gorm:"column:gender"`
	Birthday   types.NullTime `json:"birthday" gorm:"column:birthday"`
	RegisterAt time.Time      `json:"register_at" gorm:"column:register_at"`
	LastLogin  types.NullTime `json:"last_login_at" gorm:"column:last_login_at"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}

func (u User) ToMapData() map[string]interface{} {
	var returnValue map[string]interface{}
	marshalValue, _ := json.Marshal(u)
	json.Unmarshal(marshalValue, &returnValue)
	return returnValue
}

func Login(username, password string) (User, error) {
	if err := CompareUserAndPassword(username, password); err != nil {
		return User{}, err
	}
	user, err := GetUserByUserName(username)
	UpdateUserLastLoginTime(user)
	return user, err
}

func (u *User) validate() error {
	validate := validator.New()
	return validate.Struct(u)

}
// checkExistUser check name & email is exist or not, if exist return error
func (u *User) checkExistUser() error {
	if u.Id == "" && u.Name == "" {
		return fmt.Errorf("invalid user")
	}

	if u.Id != "" {
		_,err := GetUserById(u.Id)
		if err == nil {
			return fmt.Errorf("user exist")
		}
	}

	if u.Name != "" {
		_,err := GetUserByUserName(u.Name)
		if err == nil {
			return fmt.Errorf("user exist")
		}
	}

	return nil
}

func (u *User) TableName() string {
	return "user"
}