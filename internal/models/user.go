package models

import "encoding/json"

type User struct {
	Name string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	lastLogin string
}

func (u User) ToMapData() map[string]interface{} {
	var returnValue map[string]interface{}
	marshalValue, _ := json.Marshal(u)
	json.Unmarshal(marshalValue, &returnValue)
	return returnValue
}




func Login(username, password string) (User, error) {

	user := User{
		Name: username,
		Password: password,
	}

	return user, nil
}