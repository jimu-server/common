package resp

const (
	Err = 999 // 业务失败
	Ok  = 200 // 业务成功

	WebErr     = 1000 // 系统错误
	WebArgsErr = 1001 // 参数解析错误
	TokenErr   = 1002 // 身份验证失败
)
