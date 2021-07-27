package e

var msgFlags = map[int]string {
	ERROR: 							"未知的錯誤",
	ERROR_AUTH:						"授權認證失敗",
	ERROR_AUTH_PARSEFORM:			"發送參數有誤",
	ERROR_AUTH_INVALID_PASSWORD:	"帳戶與密碼不符合",
	ERROR_AUTH_INVALID_TOKEN:		"授權憑證過期",
	ERROR_USER:						"使用者錯誤",
	ERROR_USER_CREATE:              "建立用戶錯誤",
	ERROR_USER_NOT_FOUND:			"找不到用戶",
}
func GetErrorMsg(code int) string {
	msg, ok := msgFlags[code]
	if ok {
		return msg
	}
	return msgFlags[ERROR]
}