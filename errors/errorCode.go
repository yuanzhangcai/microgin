// Package common 通用方法
// creator zacyuan
// date 2019-12-30
// 定义所有返回错误码

package errors

// ErrorCode 错误码类型
type ErrorCode int64

const (
	// ErrorSystem 系统错误
	ErrorSystem ErrorCode = 9999 - iota
	// ErrorNoLogin 没有登录
	ErrorNoLogin
	// ErrorCheckLoginCreateRequestFailed 登录验证，创建请求失败
	ErrorCheckLoginCreateRequestFailed
	// ErrorCheckLoginSendRequestFailed 登录验证，发送请求失败
	ErrorCheckLoginSendRequestFailed
	// ErrorCheckLoginReadResponseFailed 登录验证，读取登录验证返回数据失败。
	ErrorCheckLoginReadResponseFailed
	// ErrorCheckLoginUnmarshalJSONailed 登录验证，解析登录验证返回数据失败。
	ErrorCheckLoginUnmarshalJSONailed

	// ErrorOK 系统正常
	ErrorOK ErrorCode = 0
)
