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
		SendResponse(c, nil, pkg.ParamError)
		return
	}
	deadline, err := time.ParseInLocation("2006-01-02 15:04:05", req.Deadline, time.Local)
	if err != nil {
		SendResponse(c, nil, pkg.ParamError)
		return
	}
	name, b := c.Get("Username")
	if !b {
		SendResponse(c, nil, err)
		return
	}
	err = service.HomeworkService.LaunchHomework(req.Title, req.Description, name.(string), req.AllowLate, deadline)
	if err != nil {
		SendResponse(c, nil, err)
		return
	}
	SendResponse(c, nil, nil)
}
func (h *homework) DeleteHomework(c *web.Context) {
	var err error
	var req struct {
		Title string `json:"title"`
	}
	err = c.BindJson(&req)
	if err != nil {
		SendResponse(c, nil, pkg.ParamError)
		return
	}
	name, b := c.Get("Username")
	if !b {
		SendResponse(c, nil, pkg.ServerError)
		return
	}
	detectUser, err := service.UserService.DetectUser(name.(string))
	if err != nil {
		SendResponse(c, nil, err)
		return
	}
	if !detectUser {
		SendResponse(c, nil, pkg.ErrUserNotAdmin)
		return
	}
	err = service.HomeworkService.DeleteHomework(req.Title)
	if err != nil {
		SendResponse(c, nil, err)
		return
	}
	SendResponse(c, nil, nil)
}

func (h *homework) UpdateHomework(c *web.Context) {
	var req struct {
		Title       *string           `json:"title"`
		Description *string           `json:"description"`
		Deadline    *string           `json:"deadline"`
		AllowLate   *bool             `json:"allow_late"`
		Department  *model.Department `json:"department"`
		Version     int               `json:"version"`
	}
	idStr, err := c.Param("id")
	if err != nil {
		SendResponse(c, nil, pkg.ParamError)
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		SendResponse(c, nil, pkg.ParamError)
		return
	}
	if err := c.BindJson(&req); err != nil {
		SendResponse(c, nil, pkg.ParamError)
		return
	}

	updates := make(map[string]interface{})

	if req.Title != nil {
		updates["title"] = *req.Title
	}
	if req.Description != nil {
		updates["description"] = *req.Description
	}
	if req.AllowLate != nil {
		updates["allow_late"] = *req.AllowLate
	}
	if req.Department != nil {
		updates["department"] = *req.Department
	}
	if req.Deadline != nil {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", *req.Deadline, time.Local)
		if err != nil {
			SendResponse(c, nil, pkg.ParamError)
			return
		}
		updates["deadline"] = t
	}
	err = service.HomeworkService.UpdateHomework(uint(id), updates, req.Version)

	if err != nil {
		SendResponse(c, nil, err)
		return
	}
	SendResponse(c, nil, nil)
}
func (h *homework) GetHomework(c *web.Context) {
	query := c.Query("title")
	homework, err := service.HomeworkService.GetHomework(query)
	var list *model.HomeworkResponse
	if err != nil {
		SendResponse(c, nil, err)
		return
	}
	list = homework.ToResponse()
	SendResponse(c, &list, nil)
}
func (h *homework) GetHomeworkList(c *web.Context) {
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")
	depStr := c.Query("department")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	department, _ := strconv.Atoi(depStr)
	depart := model.Department(department)
	resp, err := service.HomeworkService.GetDepartmentWork(depart, page, pageSize)
	if err != nil || resp == nil {
		SendResponse(c, nil, err)
		return
	}
	var list []model.HomeworkResponse
	if resp.ListHomework != nil {
		for _, hw := range *resp.ListHomework {
			list = append(list, *hw.ToResponse())
		}
	}
	SendResponse(c, map[string]interface{}{
		"list":      list,
		"total":     resp.Total,
		"page":      page,
		"page_size": pageSize,
	}, nil)
}
