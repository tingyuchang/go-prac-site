package models

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	Id string `json:"id"`
	Data map[string]interface{} `json:"data"`
	Datetime int64 `json:"datetime"`
}

func ResponseError(w http.ResponseWriter, r *http.Request, err error) {
	data := make(map[string]interface{})
	data["fail"] = err.Error()
	DefaultResponse(w, data)
}

func DefaultResponse(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	md5Data,_ := json.Marshal(data)
	id := fmt.Sprintf("%x", md5.Sum(md5Data))
	resp := Response{}
	resp.Id = id
	resp.Data = data
	resp.Datetime = time.Now().Unix()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}