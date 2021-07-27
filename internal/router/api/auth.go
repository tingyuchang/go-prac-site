package api

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-prac-site/e"
	"go-prac-site/internal/models"
	"go-prac-site/internal/router/response"
	"go-prac-site/internal/services"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	err := r.ParseForm()
	if err != nil {
		// ERROR_AUTH_PARSEFORM
		// save error to log system
		// transfer to ERROR_CODE
		response.ReturnError(w, r, fmt.Errorf(e.GetErrorMsg(e.ERROR_AUTH_PARSEFORM)))
		return
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	user, err := models.Login(username, password)

	if err != nil {
		response.ReturnError(w, r, err)
		return
	}

	tokenStr, err := services.CreateTokenString(username, password)
	if err != nil {
		response.ReturnError(w, r, fmt.Errorf(e.GetErrorMsg(e.ERROR_AUTH)))
		return
	}
	responseData := user.ToMapData()
	responseData["token"] = tokenStr

	response.DefaultResponse(w, responseData)
}
