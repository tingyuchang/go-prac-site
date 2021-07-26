package router

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
	"go-prac-site/e"
	"go-prac-site/internal/middleware"
	"go-prac-site/internal/models"
	"go-prac-site/internal/services"
	"net/http"
	"time"
)

var router *httprouter.Router

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.BasicAuth()
	defaultResponse(w, map[string]interface{}{
		"msg": "success",
	})
}

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		ResponseError(w, r, err)
		return
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	email := r.Form.Get("email")
	phone := r.Form.Get("phone")
	gender := r.Form.Get("gender")

	user := models.User{
		Name: username,
		Password: password,
		Email: email,
		Phone: phone,
		Gender: gender,
	}

	user, err = models.CreateUser(user)

	if err != nil {
		ResponseError(w, r, err)
		return
	}
	defaultResponse(w, user.ToMapData())
}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	err := r.ParseForm()
	if err != nil {
		// ERROR_AUTH_PARSEFORM
		// save error to log system
		// transfer to ERROR_CODE
		ResponseError(w, r, fmt.Errorf(e.GetErrorMsg(e.ERROR_AUTH_PARSEFORM)))
		return
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	user, err := models.Login(username, password)

	if err != nil {
		ResponseError(w, r, err)
		return
	}

	tokenStr, err := services.CreateTokenString(username, password)
	if err != nil {
		ResponseError(w, r, fmt.Errorf(e.GetErrorMsg(e.ERROR_AUTH)))
		return
	}
	responseData := user.ToMapData()
	responseData["token"] = tokenStr

	defaultResponse(w, responseData)
}

func NewRouter() *httprouter.Router {
	router = httprouter.New()
	router.GET("/", middleware.NotifyForHttpRouter()(Index))
	router.POST("/user", CreateUser)
	router.POST("/login", Login)
	router.GET("/user/:id", middleware.JWT(GetUserInfomation))
	return router
}

func NewNegroni() *negroni.Negroni {
	router = httprouter.New()
	router.GET("/", Index)
	router.POST("/user", CreateUser)
	router.POST("/login", Login)
	router.GET("/user/:username", middleware.JWT(GetUserInfomation))
	n := negroni.Classic()
	n.UseHandler(router)

	return n
}

func GetUserInfomation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Get user info")
}

func ResponseError(w http.ResponseWriter, r *http.Request, err error) {
	data := make(map[string]interface{})
	data["fail"] = err.Error()
	defaultResponse(w, data)
}

func defaultResponse(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	md5Data,_ := json.Marshal(data)
	id := fmt.Sprintf("%x", md5.Sum(md5Data))
	response := models.Response{
		Id: id,
		Data: data,
		Datetime: time.Now().Unix(),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}


