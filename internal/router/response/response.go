package response

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"go-prac-site/internal/models"
	"net/http"
	"time"
)

func ReturnError(w http.ResponseWriter, r *http.Request, err models.ResponseError) {
	DefaultResponse(w, err.ToMapData())
}

func DefaultResponse(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	md5Data,_ := json.Marshal(data)
	id := fmt.Sprintf("%x", md5.Sum(md5Data))
	resp := models.Response{}
	resp.Id = id
	resp.Data = data
	resp.Datetime = time.Now().Unix()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}