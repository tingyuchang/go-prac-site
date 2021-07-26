package e

var msgFlags = map[int]string {
	ERROR: 					"未知的錯誤",
	ERROR_AUTH:				"授權認證失敗",

}
func GetErrorMsg(code int) string {
	msg, ok := msgFlags[code]
	if ok {
		return msg
	}
	return msgFlags[ERROR]
}