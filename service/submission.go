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
func (Sub *submission) MySub(name string) (*[]model.Submission, error) {
	subs, err := dao.SubDao.MySubs(name)
	if err != nil {
		//TODO
	}
	return subs, err
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
func (Sub *submission) ExcellentSub() (*[]model.Homework, error) {
	excellents, err := dao.SubDao.DetectExcellent()
	if err != nil {
		//TODO
	}
	var homeworks []model.Homework
	for _, excellent := range *excellents {
		homework, err := dao.HomeworkDao.GetHomeworkByID(excellent)
		if err != nil {
			//TODO
		}
		homeworks = append(homeworks, *homework)
	}
	return &homeworks, nil
}
