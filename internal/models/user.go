package models

import (
	"20210703/internal/types"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Id string 					`json:"uid" db:"id"`
	Name string 				`json:"username" db:"name" validate:"required"`
	Password string 			`json:"-" db:"password" `
	Email string 				`json:"email" db:"email" validate:"required,email"`
	Phone string 				`json:"phone" db:"phone"`
	Gender string 				`json:"gender" db:"gender"`
	Birthday types.NullTime  	`json:"birthday" db:"birthday"`
	Regist_at types.NullTime 	`json:"register_at" db:"regist_at"`
	LastLogin types.NullTime 	`json:"last_login_at" db:"last_login_at"`
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