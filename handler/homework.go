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
		AllowLate   bool   `json:"allow_late"`
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
	name, b := c.Get("user")
	if !b {
		SendResponse(c, nil, pkg.TokenErr)
		return
	}
	err = service.HomeworkService.LaunchHomework(req.Title, req.Description, name.(string), req.AllowLate, deadline)
	if err != nil {
		SendResponse(c, nil, err)
		return
	}
	homework, err := service.HomeworkService.GetHomework(req.Title)
	if err != nil {
		SendResponse(c, nil, err)
		return
	}
	SendResponse(c, map[string]interface{}{
		"id":               homework.ID,
		"title":            homework.Title,
		"department":       model.DeptNameMap[homework.Department],
		"department_label": model.DeptLabelMap[homework.Department],
		"deadline":         deadline.Format("2006-01-02 15:04:05"),
		"allowLate":        req.AllowLate,
	}, nil, "发布成功")
}
func (h *homework) DeleteHomework(c *web.Context) {
	var err error

	param, err := c.Param("id")
	if err != nil {
		SendResponse(c, nil, pkg.ParamError)
	}
	id, err := strconv.ParseUint(param, 10, 64)
	dept, b := c.Get("department")
	if !b {
		SendResponse(c, nil, pkg.ServerError)
	}
	err = service.HomeworkService.DeleteHomework(uint(id), dept.(model.Department))
	if err != nil {
		SendResponse(c, nil, err)
		return
	}
	SendResponse(c, nil, nil, "删除成功")
}

func (h *homework) UpdateHomework(c *web.Context) {
	var req struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
		Deadline    *string `json:"deadline"`
		AllowLate   *bool   `json:"allow_late"`
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
	userDept, flag := c.Get("department")
	if !flag {
		SendResponse(c, nil, pkg.ServerError)
	}

	if req.Title != nil {
		updates["title"] = *req.Title
	}
	if req.Description != nil {
		updates["description"] = *req.Description
	}
	if req.AllowLate != nil {
		updates["allow_late"] = *req.AllowLate
	}

	if req.Deadline != nil {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", *req.Deadline, time.Local)
		if err != nil {
			SendResponse(c, nil, pkg.ParamError)
			return
		}
		updates["deadline"] = t
	}

	err = service.HomeworkService.UpdateHomeworkSecure(uint(id), userDept.(model.Department), updates)

	if err != nil {
		SendResponse(c, nil, err)
		return
	}
	SendResponse(c, map[string]interface{}{
		"id":          id,
		"title":       *req.Title,
		"description": *req.Description,
	}, nil, "修改成功")
}
func (h *homework) GetHomework(c *web.Context) {

	param, err2 := c.Param("id")
	parseUint, err2 := strconv.ParseUint(param, 10, 64)
	if err2 != nil {
		SendResponse(c, nil, pkg.ParamError)
	}
	if err2 != nil {
		SendResponse(c, nil, pkg.ParamError)
	}
	homework, err := service.HomeworkService.GetHomeworkId(uint(parseUint))
	if err != nil {
		SendResponse(c, nil, err)
		return
	}
	resp := homework.ToResponse()
	id, b := c.Get("userID")
	if !b {
		SendResponse(c, nil, pkg.ServerError)
	}
	sub, err := service.HomeworkService.DetectSub(homework, id.(uint))
	if err == nil && sub.ID != 0 {
		resp.MySubmission = &model.SubmissionInfo{
			ID:          sub.ID,
			Score:       *sub.Score,
			IsExcellent: sub.IsExcellent,
		}
	} else {
		resp.MySubmission = nil
	}
	SendResponse(c, resp, nil)
}
func (h *homework) GetHomeworkList(c *web.Context) {
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")
	depStr := c.Query("department")
	val, ok := model.Depart[depStr]
	if !ok {
		SendResponse(c, nil, pkg.ErrDepartmentWorkNotFound)
		return
	}

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	resp, err := service.HomeworkService.GetDepartmentWork(val, page, pageSize)
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
