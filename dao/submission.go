package dao

import (
	"errors"
	"fmt"
	"homework_submit/model"

	"gorm.io/gorm"
)

type submission struct{}

var SubDao = new(submission)

func (Sub *submission) CreateSub(s *model.Submission) error {
	return DB.Create(&s).Error
}
func (Sub *submission) DeleteSub(s *model.Submission) error {
	return DB.Delete(&s).Error
}
func (Sub *submission) MySubs(id uint, page, pageSize int) (*[]model.Submission, int64, error) {
	var s []model.Submission
	var total int64

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	query := DB.Model(&model.Submission{}).Where("student_id = ?", id)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Preload("Homework").
		Preload("Student").
		Offset(offset).
		Limit(pageSize).
		Find(&s).Error

	if err != nil {
		return nil, 0, err
	}
	return &s, total, nil
}
func (Sub *submission) DepartmentSubs(department model.Department) (*[]model.Submission, error) {
	var s []model.Submission
	tx := DB.Where("DepartmentID = ?", department).Find(&s)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &s, nil
}

func (Sub *submission) UpdateSubmissionOptimistic(sub *model.Submission, updates map[string]interface{}) error {

	updates["version"] = gorm.Expr("version + 1")

	result := DB.Model(&model.Submission{}).
		Where("id = ? AND version = ?", sub.ID, sub.Version).
		Updates(updates)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("数据已被修改，请刷新后重试")
	}

	return nil
}
func (Sub *submission) MarkExcellent(s *model.Submission, reviewer string) error {

	result := DB.Model(s).Where("Version = ?", s.Version).Updates(map[string]interface{}{
		"excellent": true,
		"version":   gorm.Expr("Version + 1"),
		"reviewer":  reviewer,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("Sub Already be changed ")
	}
	return nil
}

func (Sub *submission) GetSub(title, name string) (*model.Submission, error) {
	var sub model.Submission
	u, err := UserDao.GetUserByName(name)
	if err != nil {
		return nil, err
	}
	h, err := HomeworkDao.GetHomeworkByTitle(title)
	if err != nil {
		return nil, err
	}

	if err := DB.Where("student_id = ? AND homework_id = ?", u.ID, h.ID).First(&sub).Error; err != nil {
		return nil, err
	}
	return &sub, nil
}
func (Sub *submission) GetExcellentList(page, pageSize int) ([]model.Submission, int64, error) {
	var submissions []model.Submission
	var total int64
	query := DB.Model(&model.Submission{}).Where("is_excellent = ?", true)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize

	err := query.Preload("Homework").Preload("Student").
		Offset(offset).Limit(pageSize).
		Find(&submissions).Error

	return submissions, total, err
}
func (Sub *submission) GetSubByID(id uint) (*model.Submission, error) {
	var sub model.Submission
	err := DB.Preload("Student").Preload("Homework").First(&sub, id).Error
	return &sub, err
}
func (Sub *submission) GetSubByHomeId(id uint64, page, pageSize int) ([]model.Submission, int64, error) {
	var subs []model.Submission
	var total int64
	query := DB.Model(&model.Submission{}).Where("homework_id = ?", id)
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * pageSize
	err := query.Preload("Student").Offset(offset).Limit(pageSize).Find(&subs).Error
	if err != nil {
		return nil, 0, err
	}
	return subs, total, nil
}
func (Sub *submission) DetectSub(homework *model.Homework, user uint) (*model.Submission, error) {

	var sub model.Submission
	tx := DB.Model(&model.Submission{}).Where("homework_id = ? AND student_id = ?", homework.ID, user).Find(&sub)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &sub, nil
}
