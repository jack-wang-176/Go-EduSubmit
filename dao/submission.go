package dao

import "homework_submit/model"

type Submission struct{}

var Sub Submission

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

//TODO 考虑将下面这几个函数进行集成

func (Sub *Submission) ChangeScore(s *model.Submission, score int) error {
	*s.Score = score
	return DB.Save(&s).Error
}
func (Sub *Submission) ChangeComment(s *model.Submission, comment string) error {
	s.Comment = comment
	return DB.Save(&s).Error
}
func (Sub *Submission) ChangeExcellent(s *model.Submission) error {
	s.IsExcellent = true
	return DB.Save(&s).Error
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
