package service

import (
	"homework_submit/dao"
	"homework_submit/model"
	"time"
)

type submission struct{}

var Sub submission

func (Sub *submission) CreateSub(creator, title string) error {
	student, err := dao.UserDao.GetUserByName(creator)
	if err != nil {
		//TODO
	}
	homework, err := dao.HomeworkDao.GetHomeworkByTitle(title)
	if err != nil {
		//TODO
	}
	if !homework.AllowLate && homework.Deadline.Before(time.Now()) {
		//TODO
	}
	sub := model.Submission{
		HomeworkID:  homework.ID,
		StudentID:   student.ID,
		SubmittedAt: time.Now(),
		IsLate:      homework.Deadline.Before(time.Now()),
	}
	err = dao.SubDao.CreateSub(&sub)
	if err != nil {
		//TODO
	}
	return nil
}
func (Sub *submission) MySub(name string, page, pageSize int) (*model.PageResponse, error) {
	subs, total, err := dao.SubDao.MySubs(name, page, pageSize)
	if err != nil {
		//TODO
	}
	return &model.PageResponse{
		Total:    total,
		Page:     page,
		ListSub:  subs,
		PageSize: pageSize,
	}, err
}
func (Sub *submission) DepartmentSub(department model.Department) (*[]model.Submission, error) {
	subs, err := dao.SubDao.DepartmentSubs(department)
	if err != nil {
		//TODO
	}
	return subs, err
}
func (Sub *submission) ChangeSub(title, name, reviewer, comment string, score, excellent int) error {
	sub, err := dao.SubDao.GetSub(name, reviewer)
	if err != nil {
		//TODO
	}
	err = dao.SubDao.ChangeSub(sub, reviewer, score, comment, excellent)
	if err != nil {
		//TODO
	}
	return nil
}
func (Sub *submission) GetExcellentList(page, pageSize int) (*model.PageResponse, error) {
	subs, total, err := dao.SubDao.GetExcellentList(page, pageSize)
	if err != nil {
		//TODO
	}
	return &model.PageResponse{
		ListHomework: &subs,
		Total:        total,
		Page:         page,
		PageSize:     pageSize,
	}, nil
}
