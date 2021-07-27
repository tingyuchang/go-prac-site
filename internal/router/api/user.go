package api

import (
	"github.com/julienschmidt/httprouter"
	"go-prac-site/e"
	"go-prac-site/internal/models"
	"go-prac-site/internal/router/response"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		response.ReturnError(w, r, models.ResponseError{
			Err: err,
			Code: e.ERROR_USER,
			Desc: e.GetErrorMsg(e.ERROR_USER),

		})
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
		response.ReturnError(w, r, models.ResponseError{
			Err: err,
			Code: e.ERROR_USER_CREATE,
			Desc: e.GetErrorMsg(e.ERROR_USER_CREATE),

		})
		return
	}
	response.DefaultResponse(w, user.ToMapData())
}

func GetUserInfomation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")

	var user models.User
	var err error
	// user call it self
	if uid == "me" {
		username := ps.ByName("username")
		user, err = models.GetUserByUserName(username)
	} else {
		user, err = models.GetUserById(uid)
	}

	if err != nil {
		response.ReturnError(w, r, models.ResponseError{
			Err: err,
			Code: e.ERROR_USER_NOT_FOUND,
			Desc: e.GetErrorMsg(e.ERROR_USER_NOT_FOUND),

		})
		return
	}

	response.DefaultResponse(w, user.ToMapData())
}

