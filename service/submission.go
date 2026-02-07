package service

import (
	"homework_submit/dao"
	"homework_submit/model"
	"homework_submit/pkg"
	"time"
)

type submission struct{}

var SubService submission

func (Sub *submission) CreateSub(creator, content string, id uint) (*model.Submission, error) {
	student, err := dao.UserDao.GetUserByName(creator)
	if err != nil {
		return nil, pkg.ErrUserNotFound
	}
	homework, err := dao.HomeworkDao.GetHomeworkByID(id)
	if err != nil {
		return nil, pkg.ErrHomeworkNotFound
	}
	if !homework.AllowLate && homework.Deadline.Before(time.Now()) {
		return nil, pkg.ErrAlreadyLate
	}
	if student.Department != homework.Department {
		return nil, pkg.ErrWrongDepartment
	}
	sub := model.Submission{
		HomeworkID:  homework.ID,
		StudentID:   student.ID,
		SubmittedAt: time.Now(),
		IsLate:      homework.Deadline.Before(time.Now()),
		Department:  student.Department,
		Content:     content,
	}
	err = dao.SubDao.CreateSub(&sub)
	if err != nil {
		return nil, pkg.ErrorPkg.WithCause(err)
	}
	return &sub, nil
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
	}, nil
}
func (Sub *submission) DepartmentSub(department model.Department) (*[]model.Submission, error) {
	subs, err := dao.SubDao.DepartmentSubs(department)
	if err != nil {
		return nil, pkg.ErrDepartmentSubNotFound
	}
	return subs, nil
}

func (Sub *submission) ChangeSub(subID uint, reviewer, comment string, score, excellent, version int) error {
	sub, err := dao.SubDao.GetSubByID(subID)
	if err != nil {
		return pkg.ErrNoSuchSub
	}
	if sub.Version != &version {
		return pkg.ErrSubBeChanged
	}
	user, err := dao.UserDao.GetUserByName(reviewer)
	if err != nil {
		return pkg.ErrUserNotFound
	}
	if user.Department != sub.Department {
		return pkg.ErrWrongDepartment
	}
	err = dao.SubDao.ChangeSub(sub, reviewer, score, comment, excellent)
	return err
}
func (Sub *submission) GetExcellentList(page, pageSize int) (*model.PageResponse, error) {
	subs, total, err := dao.SubDao.GetExcellentList(page, pageSize)
	if err != nil {
		return nil, pkg.ErrorPkg.WithCause(err)
	}
	return &model.PageResponse{
		ListSub:  &subs,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}
func (Sub *submission) GetWorkSubs(id uint64, page, pageSize int, dept model.Department) (*model.PageResponse, error) {

	subs, total, err := dao.SubDao.GetSubByHomeId(id, page, pageSize)
	if err != nil {
		return nil, pkg.ErrorPkg.WithCause(err)
	}
	if len(subs) == 0 {
		return nil, pkg.ErrWrongHomeID
	}
	var flag = subs[0].Department == dept
	if flag {
		return &model.PageResponse{
			ListSub:  &subs,
			Page:     page,
			PageSize: pageSize,
			Total:    total,
		}, nil
	} else {
		return nil, pkg.ErrWrongDepartment
	}

}
func (Sub *submission) MarkExcellent(subID uint, reviewer string) error {
	sub, err := dao.SubDao.GetSubByID(subID)
	if err != nil {
		return pkg.ErrNoSuchSub
	}
	user, err := dao.UserDao.GetUserByName(reviewer)
	if err != nil {
		return pkg.ErrUserNotFound
	}
	if user.Department != sub.Department {
		return pkg.ErrWrongDepartment
	}
	err = dao.SubDao.MarkExcellent(sub, reviewer)
	return err
}
