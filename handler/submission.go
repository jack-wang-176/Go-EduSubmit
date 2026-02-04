package handler

import (
	"homework_submit/model"
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
		return
	}

	user, flag := c.Get("user")
	if flag {
		err := service.SubService.CreateSub(user.(string), req.Content)
		if err != nil {
			SendResponse(c, nil, err)
			return
		}
		SendResponse(c, nil, nil)
	}
}

func (s *submission) MySub(c *web.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	name, b := c.Get("user")
	if !b {
		SendResponse(c, nil, pkg.ServerError)
		return
	}

	subs, err := service.SubService.MySub(name.(string), page, pageSize)
	if err != nil {
		SendResponse(c, nil, err)
		return
	}
	var resSubs []model.SubmissionResponse
	for _, sub := range *subs.ListSub {
		resSubs = append(resSubs, *sub.ToResponse())
	}

	SendResponse(c, subs, nil)
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
		return
	}

	user, b := c.Get("user")
	if !b {
		SendResponse(c, nil, pkg.ServerError)
		return
	}
	err := service.SubService.ChangeSub(req.Title, user.(string), req.Comment, req.Score, req.IsExcellent)
	if err != nil {
		SendResponse(c, nil, err)
		return
	}
	SendResponse(c, nil, nil)
}

func (s *submission) GetExcellentList(c *web.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	list, err := service.SubService.GetExcellentList(page, pageSize)
	if err != nil {
		SendResponse(c, nil, err)
		return
	}
	var resSubs []model.SubmissionResponse
	for _, sub := range *list.ListSub {
		resSubs = append(resSubs, *sub.ToResponse())
	}
	SendResponse(c, resSubs, nil)
}
