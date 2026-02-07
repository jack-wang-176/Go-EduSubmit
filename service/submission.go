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
func (Sub *submission) MySub(id uint, page, pageSize int) (*model.PageResponse, error) {
	subs, total, err := dao.SubDao.MySubs(id, page, pageSize)
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

func (Sub *submission) ChangeSub(subID uint, reviewerID uint, comment string, score int, isExcellent bool) (*model.Submission, error) {
	sub, err := dao.SubDao.GetSubByID(subID)
	if err != nil {
		return nil, pkg.ErrHomeworkNotFound
	}

	updates := map[string]interface{}{
		"score":        score,
		"comment":      comment,
		"is_excellent": isExcellent,
		"reviewer_id":  reviewerID,
		"reviewed_at":  time.Now(),
	}

	err = dao.SubDao.UpdateSubmissionOptimistic(sub, updates)
	if err != nil {
		return nil, err
	}

	sub.Score = &score
	sub.Comment = comment
	sub.IsExcellent = isExcellent
	sub.ReviewerID = &reviewerID

	return sub, nil
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

// SetExcellent 设置或取消优秀作业
func (Sub *submission) SetExcellent(subID uint, isExcellent bool, reviewerID uint) error {

	sub, err := dao.SubDao.GetSubByID(subID)
	if err != nil {
		return pkg.ErrHomeworkNotFound
	}
	user, err := dao.UserDao.GetUserById(reviewerID)
	if err != nil {
		return pkg.ErrUserNotFound
	}
	if user.Department != sub.Department {
		return pkg.ErrWrongDepartment
	}

	updates := map[string]interface{}{
		"is_excellent": isExcellent,
	}

	return dao.SubDao.UpdateSubmissionOptimistic(sub, updates)
}
