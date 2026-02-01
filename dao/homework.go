package dao

import (
	"homework_submit/model"
	"time"
)

type homeworkDao struct{}

var HomeworkDao = new(homeworkDao)

func (d *homeworkDao) LaunchHomework(h *model.Homework) error {
	return DB.Create(h).Error
}

func (d *homeworkDao) UpdateHomework(h *model.Homework, title, des string, department model.Department, deadline time.Time, allow bool) error {
	tx := DB.Model(h).Where("Version = ?", h.Version).Updates(map[string]interface{}{
		"Title":       title,
		"Description": des,
		"AllowLate":   allow,
		"Deadline":    deadline,
		"Department":  department,
	})
	if tx.Error != nil {
		//TODO
	}
	if tx.RowsAffected == 0 {
		//TODO
	}
	return nil
}
func (d *homeworkDao) DeleteHomework(h *model.Homework) error {
	return DB.Delete(h).Error
}

func (d *homeworkDao) GetHomeworkByTitle(title string) (model.Homework, error) {
	var homeworks model.Homework
	tx := DB.Where("title = ?", title).First(&homeworks)
	if tx.Error != nil {
		return homeworks, tx.Error
	}
	return homeworks, nil
}
func (d *homeworkDao) GetHomeworkByID(id uint) (*model.Homework, error) {
	var h model.Homework
	tx := DB.First(&h, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &h, nil
}
func (d *homeworkDao) GetHomeworkByDepartment(department model.Department) ([]model.Homework, error) {
	var homeworks []model.Homework
	tx := DB.Where("Department = ?", department).Find(&homeworks)
	if tx.Error != nil {
		//TODO
	}
	return homeworks, nil
}
