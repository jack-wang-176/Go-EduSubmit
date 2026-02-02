package handler

import (
	"homework_submit/model"
	"homework_submit/service"
	"net/http"

	"github.com/jack-wang-176/Maple/web"
)

type submission struct{}

var Sub submission

func (s *submission) CreateSub(c *web.Context) error {
	var err error
	var sub model.Submission
	err = c.BindJson(&sub)
	service.SubService.CreateSub()

}
func (s *submission) Mysub(c *web.Context) error {
	var err error
	var sub model.Submission
	err = c.BindJson(&sub)
	if err != nil {
		//todo
	}
	service.SubService.MySub()
}

func (s *submission) ChangeSub(c *web.Context) error {
	var err error
	var sub model.Submission
	err = c.BindJson(&sub)
	if err != nil {
		//todo
	}
	service.SubService.ChangeSub()
}
func (s *submission) GetExcellentList(c *web.Context) error {
	var err error
	var sub model.Submission
	err = c.BindJson(&sub)
	if err != nil {
		//todo
	}
	service.SubService.GetExcellentList()
}
func (s *submission) DepartmentSub(c *web.Context) error {
	var err error
	var sub model.Submission
	err = c.BindJson(&sub)
	if err != nil {
		//todo
	}
	service.SubService.DepartmentSub()
}
