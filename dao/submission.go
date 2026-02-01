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
func (Sub *submission) MySubs(my string) (*[]model.Submission, error) {
	var s []model.Submission
	me, err := UserDao.GetUserByName(my)
	if err != nil {
		return nil, err
	}
	tx := DB.Where("CreatorID = ?", me).Find(&s)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &s, nil
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

func (Sub *submission) DetectExcellent() (*[]uint, error) {
	var s []model.Submission
	var homeworks []uint
	tx := DB.Where("isExcellent = ?", true).Find(&s)
	if tx.Error != nil {
		return nil, tx.Error
	}
	for _, sub := range s {
		homeworks = append(homeworks, sub.HomeworkID)
	}
	return &homeworks, nil
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
