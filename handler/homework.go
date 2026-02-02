package handler

import (
	"homework_submit/model"
	"homework_submit/service"
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
	var err error
	var homework model.Homework
	err = c.BindJson(&homework)
	if err != nil {
		//todo
	}
	err = service.HomeworkService.UpdateHomework(homework.Title)
}
func (h *homework) GetHomework(c *web.Context) {
	var err error
	var homework model.Homework
	err = c.BindJson(&homework)
	if err != nil {
		//todo
	}
	getHomework, err := service.HomeworkService.GetHomework(homework.Title)
	if err != nil {
		//todo
	}

}
func (h *homework) GetHomeworkList(c *web.Context) {
	var err error
	var homework model.Homework
	err = c.BindJson(&homework)
	if err != nil {
		//todo
	}
	service.HomeworkService.GetDepartmentWork()
}
