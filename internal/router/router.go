package router

import (
	"20210703/internal/models"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

var router *httprouter.Router

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	addressAndTime := fmt.Sprintf("%s_%v", r.RemoteAddr, time.Now())
	id := fmt.Sprintf("%x", md5.Sum([]byte(addressAndTime)))
	response := models.Response{
		Id: id,
		Data: map[string]interface{}{
			"msg": "success",
		},
		Datetime: time.Now().Unix(),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	email := r.Form.Get("email")
	phone := r.Form.Get("phone")

	user := &models.User{
		Name: username,
		Password: password,
		Email: email,
		Phone: phone,
	}
	w.Header().Set("Content-Type", "application/json")
	addressAndTime := fmt.Sprintf("%s_%v", r.RemoteAddr, time.Now())
	id := fmt.Sprintf("%x", md5.Sum([]byte(addressAndTime)))
	response := models.Response{
		Id: id,
		Data: user.ToMapData(),
		Datetime: time.Now().Unix(),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	user, err := models.Login(username, password)

	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	addressAndTime := fmt.Sprintf("%s_%v", r.RemoteAddr, time.Now())
	id := fmt.Sprintf("%x", md5.Sum([]byte(addressAndTime)))
	response := models.Response{
		Id: id,
		Data: user.ToMapData(),
		Datetime: time.Now().Unix(),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}


func NewRouter() *httprouter.Router {
	router = httprouter.New()

	router.GET("/", Index)
	router.POST("/user", CreateUser)
	router.POST("/login", Login)

	return router
}
