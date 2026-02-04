package dao

import (
	"fmt"
	"homework_submit/model"

	"gorm.io/gorm"
)

type homeworkDao struct{}

var HomeworkDao = new(homeworkDao)

func (d *homeworkDao) LaunchHomework(h *model.Homework) error {
	return DB.Create(h).Error
}

func (d *homeworkDao) UpdateHomework(h *model.Homework, updates map[string]interface{}) error {

	updates["version"] = gorm.Expr("version + 1")

	tx := DB.Model(h).
		Where("id = ? AND version = ?", h.ID, h.Version).
		Updates(updates)

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return fmt.Errorf("更新失败：作业可能已被其他人修改，请刷新后重试")
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
func (d *homeworkDao) GetHomeworkByDepartment(department model.Department, page, pageSize int) ([]model.Homework, int64, error) {
	var homeworks []model.Homework
	var total int64
	query := DB.Model(&model.Homework{}).Where("Department = ?", department)
	if query.Error != nil {
		return nil, 0, query.Error
	}
	tx := query.Count(&total)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	offset := (page - 1) * pageSize
	err := query.Preload("Homework").Preload("Student").Offset(offset).Limit(pageSize).Find(&homeworks).Error
	if err != nil {
		return nil, 0, err
	}
	return homeworks, total, nil
}
