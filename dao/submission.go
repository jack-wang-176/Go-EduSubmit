package dao

import (
	"fmt"
	"homework_submit/model"

	"gorm.io/gorm"
)

type Submission struct{}

var Sub = new(Submission)
var Excellent = 1

func (Sub *Submission) CreateSub(s *model.Submission) error {
	return DB.Create(&s).Error
}
func (Sub *Submission) UpdateSub(s *model.Submission) error {
	return DB.Save(&s).Error
}
func (Sub *Submission) DeleteSub(s *model.Submission) error {
	return DB.Delete(&s).Error
}
func (Sub *Submission) MySubs(my int) (*[]model.Submission, error) {
	var s []model.Submission
	tx := DB.Where("CreatorID = ?", my).Find(&s)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &s, nil
}
func (Sub *Submission) DepartmentSubs(department int) (*[]model.Submission, error) {
	var s []model.Submission
	tx := DB.Where("DepartmentID = ?", department).Find(&s)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &s, nil
}

func (Sub *Submission) ChangeSub(s *model.Submission, score int, comment string, excellent int) error {
	ex := excellent == Excellent
	result := DB.Model(s).Where("Version = ?", s.Version).Updates(map[string]interface{}{
		"Score":     score,
		"Comment":   comment,
		"excellent": ex,
		"version":   gorm.Expr("Version + 1"),
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("Sub Already be changed ")
	}
	return nil
}

func (Sub *Submission) DetectExcellent() (*[]uint, error) {
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
