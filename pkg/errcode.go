package pkg

import "net/http"

// 通用错误
var (
	Success     = New(0, "success", http.StatusOK)
	ServerError = New(10000, "服务器内部错误", http.StatusInternalServerError)
	ParamError  = New(10001, "参数错误", http.StatusBadRequest)
)

// 用户模块 (20000 - 20099)
var (
	ErrUserNotFound      = New(20001, "用户不存在", http.StatusOK)
	ErrPasswordIncorrect = New(20002, "密码错误", http.StatusOK)
	ErrUserExists        = New(20003, "用户已存在", http.StatusOK)
)

// 作业模块 (30000 - 30099)
var (
	ErrAlreadyLate            = New(30001, "作业提交已截止", http.StatusOK)
	ErrHomeworkNotFound       = New(30002, "找不到该作业", http.StatusOK)
	ErrDepartmentWorkNotFound = New(30003, "没有找到该部门的作业", http.StatusOK)
)

var (
	ErrDepartmentSubNotFound = New(40001, "没有当前部门的提交", http.StatusOK)
	ErrNoSuchSub             = New(40002, "未找到相应的提交", http.StatusOK)
)

var (
	NoInput = New(50001, "没有输入", http.StatusOK)
)
