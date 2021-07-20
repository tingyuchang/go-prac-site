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


func NewRouter() *httprouter.Router {
	router = httprouter.New()

	router.GET("/", Index)

	return router
}