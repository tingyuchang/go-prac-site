package models

type Response struct {
	Id string 							`json:"id"`
	Data map[string]interface{} 		`json:"data"`
	Datetime int64 						`json:"datetime"`
}
type ResponseError struct {
	Err error
	Code int
	Desc string
}

func (err ResponseError) ToMapData() (data map[string]interface{}) {
	data = make(map[string]interface{})
	data["code"] = err.Code
	data["desc"] = err.Desc
	data["err_msg"] = err.Err.Error()
	return
}
