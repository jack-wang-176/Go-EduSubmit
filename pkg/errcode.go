package pkg

import "net/http"

var (
	Success     = New(0, "success", http.StatusOK)
	ServerError = New(10000, "服务器内部错误", http.StatusInternalServerError)
	ParamError  = New(10001, "参数错误", http.StatusBadRequest)
)

var (
	ErrUserNotFound      = New(20001, "用户不存在", http.StatusOK) // 业务上算失败，但HTTP可以是200也可以是404，看你规范
	ErrPasswordIncorrect = New(20002, "密码错误", http.StatusOK)
	ErrUserExists        = New(20003, "用户已存在", http.StatusOK)
	ErrAlreadyLate       = New(20004, "作业提交已经关闭", http.StatusOK)
)

var (
	ErrTokenInvalid = New(30001, "Token无效", http.StatusUnauthorized)
)

var (
	NoInput = New(40001, "没有input", http.StatusBadRequest)
)
