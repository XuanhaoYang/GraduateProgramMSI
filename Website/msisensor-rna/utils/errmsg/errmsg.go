package errmsg

const (
	SUCCESS = 200
	ERROR   = 500
	// code = 1000... 用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_NOT_EXIST  = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008
	// code = 2000... 样本模块的错误
	ERROR_SAMPNAME_USED  = 4001
	ERROR_SAMP_NOT_EXIST = 4002
	// code = 3000... 上传的错误
	ERROR_UPLOAD_WRONG = 3001
)

var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在!",
	ERROR_PASSWORD_WRONG:   "密码错误!",
	ERROR_USER_NOT_EXIST:   "用户不存在!",
	ERROR_TOKEN_NOT_EXIST:  "token不存在!",
	ERROR_TOKEN_RUNTIME:    "token已过期!",
	ERROR_TOKEN_WRONG:      "token不正确",
	ERROR_TOKEN_TYPE_WRONG: "token格式错误",
	ERROR_USER_NO_RIGHT:    "用户无权限",
	ERROR_SAMPNAME_USED:    "样本ID已存在",
	ERROR_SAMP_NOT_EXIST:   "样本不存在",
	ERROR_UPLOAD_WRONG:     "样本上传错误",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
