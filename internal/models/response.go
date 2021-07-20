package models

type Response struct {
	Id string `json:"id"`
	Data map[string]interface{} `json:"data"`
	Datetime int64 `json:"datetime"`
}