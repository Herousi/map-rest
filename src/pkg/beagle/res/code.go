package res

// xxx 前三位采用谷歌http错误码  xx 模块 xx 错误类型
// 400 客户端错误
// 401 权限错误
// 500 服务器错误
// 501 数据库错误

var (
	// 成功
	OK = BgRes{Code: 200, Msg: "执行成功"}
	// 参数错误
	ParamsParserError = BgRes{Code: 4000001, Msg: "参数类型不匹配，解析失败"}
	ParamsEmptyError  = BgRes{Code: 4000002, Msg: "参数不能为空"}

	// 权限错误
	LoginTimeOutError = BgRes{Code: 4010001, Msg: "登录超时"}
	UnauthorizedError = BgRes{Code: 4010003, Msg: "无权操作"}

	// 服务器错误
	InternalServerError   = BgRes{Code: 5000001, Msg: "系统异常"}
	DataFailError         = BgRes{Code: 5000002, Msg: "数据获取失败"}
	LogoutError           = BgRes{Code: 5000003, Msg: "退出失败"}
	AccountNotEnableError = BgRes{Code: 5000004, Msg: "该账户未启用,不能登录"}

	// 数据库错误
	DbConnectError = BgRes{Code: 5010001, Msg: "数据库连接失败"}
)
