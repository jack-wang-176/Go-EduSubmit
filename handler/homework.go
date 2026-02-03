package handler

import (
	"homework_submit/model"
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
		//todo
	}
	deadline, err := time.ParseInLocation("2006-01-02 15:04:05", req.Deadline, time.Local)
	if err != nil {
		//todo
	}
	name, b := c.Get("Username")
	if !b {
		//todo
	}
	err = service.HomeworkService.LaunchHomework(req.Title, req.Description, name, req.AllowLate, deadline)
	if err != nil {
		//todo
	}
}
func (h *homework) DeleteHomework(c *web.Context) {
	var err error
	var req struct {
		Title string `json:"title"`
	}
	err = c.BindJson(req)
	if err != nil {
		//todo
	}
	name, b := c.Get("Username")
	if !b {
		//todo
	}
	detectUser, err := service.UserService.DetectUser(name.(string))
	if err != nil {

	}
	if !detectUser {

	}
	err = service.HomeworkService.DeleteHomework(req.Title)
	if err != nil {
		//todo
	}
	c.JsonResp()
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

	// 2. 绑定参数
	if err := c.BindJson(&req); err != nil {

		return
	}

	newDeadline, err := time.ParseInLocation("2006-01-02 15:04:05", req.Deadline, time.Local)
	if err != nil {
		return
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
		return
	}

}
func (h *homework) GetHomework(c *web.Context) {
	query := c.Query("title")
	getHomework, err := service.HomeworkService.GetHomework(query)
	if err != nil {

	}
}
func (h *homework) GetHomeworkList(c *web.Context) {
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")
	depStr := c.Query("department")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	department, _ := strconv.Atoi(depStr)
	departmentWork, err := service.HomeworkService.GetDepartmentWork(department, page, pageSize)
	if err != nil {

	}
}
