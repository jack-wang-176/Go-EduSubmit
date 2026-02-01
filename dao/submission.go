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
	query := DB.Model(&model.Submission{}).Where("CreatorID = ?", me)
	if query.Count(&total).Error != nil {
		//TODO
	}
	offset := (page - 1) * pageSize
	query.Preload("Homework").Preload("Student").Offset(offset).Limit(pageSize).Find(&s)
	if query.Error != nil {
		return nil, 0, query.Error
	}
	return &s, 0, nil
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

func (Sub *submission) GetSub(title, name string) (model *model.Submission, err error) {
	var sub model.Submission
	u, err := UserDao.GetUserByName(name)
	if err != nil {
		//TODO
	}
	h, err := HomeworkDao.GetHomeworkByTitle(title)
	if err != nil {
		//TODO
	}
	DB.Where("StudentID = ?, HomeworkID = ?", u.ID, h.ID).First(&sub)
	return &sub, nil
}
func (Sub *submission) GetExcellentList(page, pageSize int) ([]model.Homework, int64, error) {
	var s []model.Homework
	var total int64
	query := DB.Model(&model.Homework{}).Where("isExcellent = ?", true)
	tx := query.Count(&total)
	if tx.Error != nil {
		//TODO
	}
	offset := (page - 1) * pageSize
	err := query.Preload("Homework").Preload("Student").Offset(offset).Limit(pageSize).Find(&s).Error
	if err != nil {
		//TODO
	}
	return s, total, nil
}
