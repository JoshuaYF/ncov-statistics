package router

const (
	SUCCESS             = 0
	FAILURE             = 10000
	SYSTEM_ERROR        = 99999
	PARAM_NOT_COMPLETED = 10001
	TOO_MANY_REQUEST    = 10002
)

var MsgFlags = map[int]string{
	SUCCESS:             "ok",
	FAILURE:             "fail",
	SYSTEM_ERROR:        "服务端错误",
	PARAM_NOT_COMPLETED: "参数不完整",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[FAILURE]
}
