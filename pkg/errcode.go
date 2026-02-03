package pkg

import "net/http"

// 通用错误
var (
	Success     = New(0, "success", http.StatusOK)
	ServerError = New(10000, "服务器内部错误", http.StatusInternalServerError)
	ParamError  = New(10001, "参数错误", http.StatusBadRequest)
)

// 用户模块错误 (200xx)
var (
	ErrUserNotFound      = New(20001, "用户不存在", http.StatusOK) // 业务上算失败，但HTTP可以是200也可以是404，看你规范
	ErrPasswordIncorrect = New(20002, "密码错误", http.StatusOK)
	ErrUserExists        = New(20003, "用户已存在", http.StatusOK)
)

// 认证模块错误 (300xx)
var (
	ErrTokenInvalid = New(30001, "Token无效", http.StatusUnauthorized)
)
