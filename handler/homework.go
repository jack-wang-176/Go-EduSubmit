package handler

import (
	"homework_submit/model"
	"homework_submit/pkg"
	"homework_submit/service"
	"strconv"
	"time"

	"github.com/jack-wang-176/Maple/web"
)

type homework struct{}

var HomeworkHandler = &homework{}

func (h *homework) LaunchHomework(c *web.Context) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Deadline    string `json:"deadline"`
		AllowLate   bool   `json:"allowLate"`
	}
	var err error
	err = c.BindJson(&req)
	if err != nil {
		SendResponse(c, nil, pkg.ServerError)
	}
	deadline, err := time.ParseInLocation("2006-01-02 15:04:05", req.Deadline, time.Local)
	if err != nil {
		SendResponse(c, nil, pkg.ParamError)
	}
	name, b := c.Get("Username")
	if !b {
		SendResponse(c, nil, pkg.ServerError)
	}
	err = service.HomeworkService.LaunchHomework(req.Title, req.Description, name.(string), req.AllowLate, deadline)
	if err != nil {
		SendResponse(c, nil, pkg.ParamError)
	}
	SendResponse(c, nil, pkg.Success)
}
func (h *homework) DeleteHomework(c *web.Context) {
	var err error
	var req struct {
		Title string `json:"title"`
	}
	err = c.BindJson(req)
	if err != nil {
		SendResponse(c, nil, pkg.ParamError)
	}
	name, b := c.Get("Username")
	if !b {
		SendResponse(c, nil, pkg.ServerError)
	}
	detectUser, err := service.UserService.DetectUser(name.(string))
	if err != nil {
		SendResponse(c, nil, pkg.ServerError)
	}
	if !detectUser {
		SendResponse(c, nil, pkg.ParamError)
	}
	err = service.HomeworkService.DeleteHomework(req.Title)
	if err != nil {
		SendResponse(c, nil, pkg.ParamError)
	}
	SendResponse(c, nil, pkg.Success)
}
func (h *homework) UpdateHomework(c *web.Context) {

	var req struct {
		ID          uint   `json:"id"`          // 必须传 ID 才知道改哪一个
		Title       string `json:"title"`       // 新标题
		Description string `json:"description"` // 新描述
		Deadline    string `json:"deadline"`    // 时间接收字符串 "2026-02-20 12:00:00"
		AllowLate   bool   `json:"allow_late"`
		Department  int    `json:"department"` // 部门枚举值
		Version     int    `json:"version"`    // 乐观锁必须传旧版本号
	}

	if err := c.BindJson(&req); err != nil {
		SendResponse(c, nil, pkg.ParamError)
	}

	newDeadline, err := time.ParseInLocation("2006-01-02 15:04:05", req.Deadline, time.Local)
	if err != nil {
		SendResponse(c, nil, pkg.ParamError)
	}

	err = service.HomeworkService.UpdateHomework(
		req.ID,
		req.Title,
		req.Description,
		model.Department(req.Department),
		newDeadline,
		req.AllowLate,
		req.Version,
	)

	if err != nil {
		SendResponse(c, nil, pkg.ParamError)
	}
	SendResponse(c, nil, pkg.Success)
}
func (h *homework) GetHomework(c *web.Context) {
	query := c.Query("title")
	_, err := service.HomeworkService.GetHomework(query)
	if err != nil {
		SendResponse(c, nil, pkg.ServerError)
	}
	SendResponse(c, nil, pkg.Success)
}
func (h *homework) GetHomeworkList(c *web.Context) {
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")
	depStr := c.Query("department")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	department, _ := strconv.Atoi(depStr)
	depart := model.Department(department)
	_, err := service.HomeworkService.GetDepartmentWork(depart, page, pageSize)
	if err != nil {
		SendResponse(c, nil, pkg.ServerError)
	}
	SendResponse(c, nil, pkg.Success)
}
