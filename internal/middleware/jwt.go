package middleware

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-prac-site/internal/models"
	"go-prac-site/internal/services"
	"net/http"
	"strings"
	"time"
)

func JWT(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		username := ps.ByName("username")
		authorizations := strings.Split(r.Header.Get("Authorization"), " ")

		if len(authorizations) != 2 {
			authFailedResponse(w, fmt.Errorf("authorization not found"))
			return
		}

		tokenStr := strings.Split(r.Header.Get("Authorization"), " ")[1]

		err := services.Check(username, tokenStr)

		if err != nil {
			authFailedResponse(w, err)
			return
		}

		next(w,r,ps)
	}
}

func authFailedResponse(w http.ResponseWriter, err error) {
	data := make(map[string]interface{})
	data["result"] = err.Error()
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