package code

const (
	DBInsertErrorCode         = 101
	DBUpdateErrorCode         = 102
	DBQueryErrorCode          = 103
	DBDeleteErrorCode         = 104
	PasswordErrorCode         = 201
	RedisQueryErrorCode       = 301
	AppointmentQueryErrorCode = 401
)

const (
	DBInsertErrorMsg         = "数据插入异常"
	DBUpdateErrorMsg         = "数据更新异常"
	DBQueryErrorMsg          = "数据查询异常"
	DBDeleteErrorMsg         = "数据删除异常"
	PasswordErrorMsg         = "密码错误"
	RedisQueryErrorMsg       = "缓存查询异常"
	AppointmentQueryErrorMsg = "预约查询异常"
)

var ErrorMap map[int]string

func init() {
	ErrorMap = make(map[int]string)
	ErrorMap[DBInsertErrorCode] = DBInsertErrorMsg
	ErrorMap[DBUpdateErrorCode] = DBUpdateErrorMsg
	ErrorMap[DBQueryErrorCode] = DBQueryErrorMsg
	ErrorMap[DBDeleteErrorCode] = DBDeleteErrorMsg
	ErrorMap[PasswordErrorCode] = PasswordErrorMsg
	ErrorMap[RedisQueryErrorCode] = RedisQueryErrorMsg
	ErrorMap[AppointmentQueryErrorCode] = AppointmentQueryErrorMsg
}
