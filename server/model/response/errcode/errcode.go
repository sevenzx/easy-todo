package errcode

type ErrCode int

func (ec ErrCode) Error() string {
	return ec.String()
}

// 1. 安装stringer工具: go install golang.org/x/tools/cmd/stringer@latest
// 2. 定义好ErrCode以及Message之后，运行以下命令自动生成新的错误码和错误信息
//go:generate stringer -type ErrCode -linecomment

const Success ErrCode = 0 // success

// 1开头: 服务级错误码
const (
	// ServerError 内部错误
	ServerError     ErrCode = iota + 10001 // 服务内部错误
	ParamsError                            // 参数信息有误
	TokenAuthFail                          // Token鉴权失败
	TokenIsNotExist                        // Token不存在
)

// 2开头: 用户级错误码
const (
	UserNotExist     ErrCode = iota + 20001 // 用户不存在
	UserAlreadyExist                        // 用户已存在
	PasswordError                           // 密码错误
)
