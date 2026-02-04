package handler

import (
	"homework_submit/pkg"
	"homework_submit/service"
	"strconv"

	"github.com/jack-wang-176/Maple/web"
)

type submission struct{}

var Sub submission

func (s *submission) CreateSub(c *web.Context) {
	var req struct {
		Content string `json:"content"`
	}
	if err := c.BindJson(&req); err != nil {
		SendResponse(c, nil, pkg.ParamError)
	}

	user, flag := c.Get("user")
	if flag {
		err := service.SubService.CreateSub(user.(string), req.Content)
		if err != nil {
			SendResponse(c, nil, pkg.ServerError)
		}
		SendResponse(c, nil, pkg.Success)
	}
}

func (s *submission) MySub(c *web.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	name, b := c.Get("user")
	if !b {
		SendResponse(c, nil, pkg.ServerError)
	}

	_, err := service.SubService.MySub(name.(string), page, pageSize)
	if err != nil {
		SendResponse(c, nil, pkg.ServerError)
	}
	SendResponse(c, nil, pkg.Success)
}

func (s *submission) ChangeSub(c *web.Context) {
	var req struct {
		Title        string `json:"title"`
		SubmissionID uint   `json:"submission_id"`
		Score        int    `json:"score"`
		Comment      string `json:"comment"`
		IsExcellent  int    `json:"is_excellent"`
	}

	if err := c.BindJson(&req); err != nil {
		SendResponse(c, nil, pkg.ServerError)
	}

	user, b := c.Get("user")
	if !b {
		SendResponse(c, nil, pkg.ServerError)
	}
	err := service.SubService.ChangeSub(req.Title, user.(string), req.Comment, req.Score, req.IsExcellent)
	if err != nil {
		SendResponse(c, nil, pkg.ServerError)
	}
	SendResponse(c, nil, pkg.Success)
}

func (s *submission) GetExcellentList(c *web.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	_, err := service.SubService.GetExcellentList(page, pageSize)
	if err != nil {
		SendResponse(c, nil, pkg.ServerError)
	}
	SendResponse(c, nil, pkg.Success)
}
