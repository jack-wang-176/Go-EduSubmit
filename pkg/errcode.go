package pkg

import "net/http"

var (
	ServerError = New(10000, "服务器内部错误", http.StatusInternalServerError)
	ParamError  = New(10001, "参数错误", http.StatusBadRequest)
	TokenErr    = New(10002, "token验证失败", http.StatusBadRequest)
	TokenFailed = New(10003, "token失效或过期", http.StatusBadRequest)
)

var (
	ErrUserNotFound      = New(20001, "用户不存在", http.StatusOK)
	ErrPasswordIncorrect = New(20002, "密码错误", http.StatusOK)
	ErrUserExists        = New(20003, "用户已存在", http.StatusOK)
	ErrUserNotAdmin      = New(20004, "用户没有管理员权限", http.StatusOK)
	ErrDeleteUserFailed  = New(20005, "删除用户失败", http.StatusOK)
	ErrWrongDepartment   = New(20006, "你没有这个部门的权限", http.StatusOK)
)

var (
	ErrAlreadyLate            = New(30001, "作业提交已截止", http.StatusOK)
	ErrHomeworkNotFound       = New(30002, "找不到该作业", http.StatusOK)
	ErrDepartmentWorkNotFound = New(30003, "没有找到该部门的作业", http.StatusOK)
)

var (
	ErrDepartmentSubNotFound = New(40001, "没有当前部门的提交", http.StatusOK)
	ErrNoSuchSub             = New(40002, "未找到相应的提交", http.StatusOK)
	ErrWrongHomeID           = New(40003, "对应的作业没有提交", http.StatusOK)
	ErrSubBeChanged          = New(40004, "对应的提交已经被其他人修改", http.StatusOK)
)

var (
	NoInput = New(50001, "没有输入", http.StatusOK)
)
