package service

import (
	"homework_submit/dao"
	"homework_submit/model"
	"homework_submit/pkg"
	"time"
)

type homeworkService struct {
}

var HomeworkService = homeworkService{}

func (s *homeworkService) LaunchHomework(title, des, creator string, late bool, deadline time.Time) error {
	create, err := dao.UserDao.GetUserByName(creator)
	if err != nil {
		return pkg.ErrUserNotFound
	}
	if create == nil {
		return pkg.ErrUserNotFound
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
		return pkg.ErrorPkg.WithCause(err)
	}
	return nil
}
func (s *homeworkService) DeleteHomework(title string) error {
	homework, err := dao.HomeworkDao.GetHomeworkByTitle(title)
	if err != nil {
		return pkg.ErrHomeworkNotFound
	}
	err = dao.HomeworkDao.DeleteHomework(&homework)
	if err != nil {
		return pkg.ErrorPkg.WithCause(err)
	}
	return nil
}

func (s *homeworkService) UpdateHomework(id uint, title, des string, department model.Department, deadline time.Time, allow bool, version int) error {

	h, err := dao.HomeworkDao.GetHomeworkByID(id)
	if err != nil {
		return pkg.ErrHomeworkNotFound
	}

	h.Version = &version

	err = dao.HomeworkDao.UpdateHomework(h, title, des, department, deadline, allow)
	if err != nil {
		return pkg.ErrorPkg.WithCause(err)
	}
	return nil
}
func (s *homeworkService) GetHomework(title string) (*model.Homework, error) {
	h, err := dao.HomeworkDao.GetHomeworkByTitle(title)
	if err != nil {
		return nil, pkg.ErrHomeworkNotFound
	}
	return &h, nil
}
func (s *homeworkService) GetDepartmentWork(department model.Department, page, pageSize int) (*model.PageResponse, error) {
	homeworks, total, err := dao.HomeworkDao.GetHomeworkByDepartment(department, page, pageSize)
	if err != nil {
		return nil, pkg.ErrDepartmentWorkNotFound
	}
	return &model.PageResponse{
		ListHomework: &homeworks,
		Total:        total,
		PageSize:     pageSize,
		Page:         page,
	}, nil
}
