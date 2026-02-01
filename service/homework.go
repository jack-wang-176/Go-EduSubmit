package service

import (
	"homework_submit/dao"
	"homework_submit/model"
	"time"
)

type homeworkService struct {
}

var HomeworkService = homeworkService{}

func (s *homeworkService) LaunchHomework(title, des, creator string, late bool, deadline time.Time) error {
	create, err := dao.UserDao.GetUserByName(creator)
	if err != nil {
		//TODO 错误检查
	}
	if create == nil {
		//TODO 错误检查
	}
	err = dao.HomeworkDao.LaunchHomework(&model.Homework{
		Title:       title,
		Description: des,
		AllowLate:   late,
		CreatorID:   create.ID,
		Department:  create.Department,
		Deadline:    deadline,
	})
	if err != nil {
		//TODO 错误检查
	}
	return nil
}
func (s *homeworkService) DeleteHomework(title string) error {
	homework, err := dao.HomeworkDao.GetHomeworkByTitle(title)
	if err != nil {
		//TODO 错误去处理
	}
	err = dao.HomeworkDao.DeleteHomework(&homework)
	if err != nil {
		//TODO
	}
	return nil
}
func (s *homeworkService) UpdateHomework(title, newTitle, des string, department model.Department, deadline time.Time, allow bool) error {
	h, err := dao.HomeworkDao.GetHomeworkByTitle(title)
	if err != nil {
		//TODO
	}
	err = dao.HomeworkDao.UpdateHomework(&h, newTitle, des, department, deadline, allow)
	if err != nil {
		//TODO
	}
	return nil
}
func (s *homeworkService) GetHomework(title string) (*model.Homework, error) {
	h, err := dao.HomeworkDao.GetHomeworkByTitle(title)
	if err != nil {
		//TODO
	}
	return &h, nil
}
func (s *homeworkService) GetDepartmentWork(department model.Department) (*[]model.Homework, error) {
	homeworks, err := dao.HomeworkDao.GetHomeworkByDepartment(department)
	if err != nil {
		//TODO
	}
	return &homeworks, nil
}
