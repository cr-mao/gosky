package errcode

/**
  项目组代号:10
  服务代号:01
  模块代号:0~99
  错误码：0~99
  | 错误标识                | 错误码   | HTTP状态码 | 描述                          |
  | ----------------------- | -------- | ---------- | ----------------------------- |
  | ErrNo                   | 10010000 | 200        |  OK                            |
  | ErrInternalServer       | 10010001 | 500        |  Internal server error （服务器内部错误）      |
  | ErrParams               | 10010002 | 400        |  Illegal params  (请求参数不合法)                |
  | ErrAuthenticationHeader | 10010003 | 401        |  Authentication header Illegal  (要登录的接口，头的token认证失败,失败跳登录页面)|
  | ErrAuthentication       | 10010004 | 401        |  Authentication failed  (登录失败，输入账户、密码失败)|
  | ErrNotFound             | 10010005 | 404        |  Route not found     (请求路由找不到）             |
  | ErrPermission           | 10010006 | 403        |  Permission denied (没有权限,一些接口可能没请求权限， 这个估计暂时用不到)            |
  | ErrTooFast              | 10010007 | 429        |  Too Many Requests （用户在给定的时间内发送了太多请求）            |
  | ErrTimeout              | 10010008 | 504        |  Server response timeout   （go服务这边不会返回，一般是nginx、网关超时 才返回504）|
  | ErrMysqlServer          | 10010101 | 500        |  Mysql server error      （mysql 服务错误)       |
  | ErrMysqlSQL             | 10010102 | 500        |  Illegal SQL               (sql 代码错误）       |
  | ErrRedisServer          | 10010201 | 500        |  Redis server error        （redis 服务错误）    |

客户端要关心的是http code 是     ErrParams  400 ，ErrAuthenticationHeader 401，ErrAuthentication 401，ErrPermission 403, ErrTooFast 429
当返回http code= 401 时，   ErrAuthenticationHeader,ErrAuthentication 根据具体Code 去区分处理
*/

type ErrCode struct {
	Code     int
	HTTPCode int
	Desc     string
}

type errCodes struct {
	ErrNo                   ErrCode
	ErrInternalServer       ErrCode
	ErrParams               ErrCode
	ErrAuthenticationHeader ErrCode
	ErrAuthentication       ErrCode
	ErrNotFound             ErrCode
	ErrPermission           ErrCode
	ErrTooFast              ErrCode
	ErrTimeout              ErrCode
	ErrMysqlServer          ErrCode
	ErrMysqlSQL             ErrCode
	ErrRedisServer          ErrCode
}

var ErrCodes = errCodes{
	ErrNo: ErrCode{
		Code:     0,
		HTTPCode: 200,
		Desc:     "",
	},
	ErrInternalServer: ErrCode{
		Code:     10010001,
		HTTPCode: 500,
		Desc:     "Internal server error",
	},
	ErrParams: ErrCode{
		Code:     10010002,
		HTTPCode: 400,
		Desc:     "Illegal params",
	},
	ErrAuthenticationHeader: ErrCode{
		Code:     10010003,
		HTTPCode: 401,
		Desc:     "Authentication header Illegal",
	},
	ErrAuthentication: ErrCode{
		Code:     10010004,
		HTTPCode: 401,
		Desc:     "Authentication failed",
	},
	ErrNotFound: ErrCode{
		Code:     10010005,
		HTTPCode: 404,
		Desc:     "Route not found",
	},
	ErrPermission: ErrCode{
		Code:     10010006,
		HTTPCode: 403,
		Desc:     "Permission denied",
	},
	ErrTooFast: ErrCode{
		Code:     10010007,
		HTTPCode: 429,
		Desc:     "Too Many Requests",
	},
	ErrTimeout: ErrCode{
		Code:     10010008,
		HTTPCode: 504,
		Desc:     "Server response timeout",
	},
	ErrMysqlServer: ErrCode{
		Code:     10010101,
		HTTPCode: 500,
		Desc:     "Mysql server error",
	},
	ErrMysqlSQL: ErrCode{
		Code:     10010102,
		HTTPCode: 500,
		Desc:     "Illegal SQL",
	},
	ErrRedisServer: ErrCode{
		Code:     10010201,
		HTTPCode: 500,
		Desc:     "Redis server error",
	},
}
