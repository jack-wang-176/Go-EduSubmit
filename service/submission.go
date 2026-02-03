package service

import (
	"homework_submit/dao"
	"homework_submit/model"
	"homework_submit/pkg"
	"time"
)

type submission struct{}

var SubService submission

func (Sub *submission) CreateSub(creator, title string) error {
	student, err := dao.UserDao.GetUserByName(creator)
	if err != nil {
		return pkg.ErrUserNotFound
	}
	homework, err := dao.HomeworkDao.GetHomeworkByTitle(title)
	if err != nil {
		return pkg.ErrorPkg.WithCause(err)
	}
	if !homework.AllowLate && homework.Deadline.Before(time.Now()) {
		return pkg.ErrAlreadyLate
	}
	sub := model.Submission{
		HomeworkID:  homework.ID,
		StudentID:   student.ID,
		SubmittedAt: time.Now(),
		IsLate:      homework.Deadline.Before(time.Now()),
	}
	err = dao.SubDao.CreateSub(&sub)
	if err != nil {
		return pkg.ErrorPkg.WithCause(err)
	}
	return nil
}
func (Sub *submission) MySub(name string, page, pageSize int) (*model.PageResponse, error) {
	subs, total, err := dao.SubDao.MySubs(name, page, pageSize)
	if err != nil {
		return nil, pkg.ErrorPkg.WithCause(err)
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
		return nil, pkg.ErrorPkg.WithCause(err)
	}
	return subs, err
}
func (Sub *submission) ChangeSub(title, name, reviewer, comment string, score, excellent int) error {
	sub, err := dao.SubDao.GetSub(name, reviewer)
	if err != nil {
		return pkg.ErrorPkg.WithCause(err)
	}
	err = dao.SubDao.ChangeSub(sub, reviewer, score, comment, excellent)
	if err != nil {
		return pkg.ErrorPkg.WithCause(err)
	}
	return nil
}
func (Sub *submission) GetExcellentList(page, pageSize int) (*model.PageResponse, error) {
	subs, total, err := dao.SubDao.GetExcellentList(page, pageSize)
	if err != nil {
		return nil, pkg.ErrorPkg.WithCause(err)
	}
	return &model.PageResponse{
		ListHomework: &subs,
		Total:        total,
		Page:         page,
		PageSize:     pageSize,
	}, nil
}
