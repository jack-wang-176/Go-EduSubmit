package handler

import (
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
		return
	}

	user, flage := c.Get("user")
	if flage {
		err := service.SubService.CreateSub(string(user), req.Content)
		if err != nil {

			return
		}

	}

}

func (s *submission) Mysub(c *web.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	name, b := c.Get("user")
	if !b {

		return
	}

	resp, err := service.SubService.MySub(name, page, pageSize)
	if err != nil {

		return
	}

}

func (s *submission) ChangeSub(c *web.Context) {

	var req struct {
		Title        string `json:"title"`
		Name         string `json:"name"`
		SubmissionID uint   `json:"submission_id"`
		Score        int    `json:"score"`
		Comment      string `json:"comment"`
		IsExcellent  int    `json:"is_excellent"`
	}

	if err := c.BindJson(&req); err != nil {

		return
	}

	user, b := c.Get("user")
	if !b {
		return
	}
	err := service.SubService.ChangeSub(req.Title, req.Name, user, req.Comment, req.Score, req.IsExcellent)
	if err != nil {

		return
	}

}

func (s *submission) GetExcellentList(c *web.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	resp, err := service.SubService.GetExcellentList(page, pageSize)
	if err != nil {

		return
	}

}
