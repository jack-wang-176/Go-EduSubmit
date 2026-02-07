package service

import (
	"errors"
	"homework_submit/dao"
	"homework_submit/model"
	"homework_submit/pkg"
	"time"

	"gorm.io/gorm"
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
func (s *homeworkService) DeleteHomework(id uint, dept model.Department) error {
	homework, err := dao.HomeworkDao.GetHomeworkByID(id)
	if err != nil {
		return pkg.ErrHomeworkNotFound
	}
	err = dao.HomeworkDao.DeleteHomework(homework)
	if err != nil {
		return pkg.ErrorPkg.WithCause(err)
	}
	return nil
}

func (s *homeworkService) UpdateHomeworkSecure(hwID uint, userDept model.Department, updates map[string]interface{}) error {
	homework, err := dao.HomeworkDao.GetHomeworkByID(hwID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return pkg.ErrHomeworkNotFound
		}
		return err
	}

	if homework.Department != userDept {
		return pkg.ErrWrongDepartment
	}

	return dao.HomeworkDao.UpdateHomework(hwID, updates)
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
func (s *homeworkService) GetHomeworkId(id uint) (*model.Homework, error) {
	h, err := dao.HomeworkDao.GetHomeworkByID(id)
	if err != nil {
		return nil, pkg.ErrHomeworkNotFound
	}
	return h, nil
}
func (s *homeworkService) DetectSub(homework *model.Homework, user uint) (*model.Submission, error) {
	sub, err := dao.SubDao.DetectSub(homework, user)
	if err != nil {
		return nil, pkg.ErrorPkg.WithCause(err)
	}
	return sub, nil
}
