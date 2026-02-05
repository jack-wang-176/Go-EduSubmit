package dao

import (
	"fmt"
	"homework_submit/model"

	"gorm.io/gorm"
)

type submission struct{}

var SubDao = new(submission)
var Excellent = 1

func (Sub *submission) CreateSub(s *model.Submission) error {
	return DB.Create(&s).Error
}
func (Sub *submission) DeleteSub(s *model.Submission) error {
	return DB.Delete(&s).Error
}
func (Sub *submission) MySubs(my string, page, pageSize int) (*[]model.Submission, int64, error) {
	var s []model.Submission
	var total int64
	me, err := UserDao.GetUserByName(my)
	if err != nil {
		return nil, 0, err
	}
	query := DB.Model(&model.Submission{}).Where("CreatorID = ?", me.ID)
	if query.Count(&total).Error != nil {
		return nil, 0, err
	}
	offset := (page - 1) * pageSize
	query.Preload("Homework").Preload("Student").Offset(offset).Limit(pageSize).Find(&s)
	if query.Error != nil {
		return nil, 0, query.Error
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

func (Sub *submission) ChangeSub(s *model.Submission, reviewer string, score int, comment string, excellent int) error {
	ex := excellent == Excellent
	result := DB.Model(s).Where("Version = ?", s.Version).Updates(map[string]interface{}{
		"Score":     score,
		"Comment":   comment,
		"excellent": ex,
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
	if err := DB.First(&sub, id).Error; err != nil {
		return nil, err
	}
	return &sub, nil
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
