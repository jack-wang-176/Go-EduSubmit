package dao

import "homework_submit/model"

type homeworkDao struct{}

var HomeworkDao = new(homeworkDao)

func (d *homeworkDao) LaunchHomework(h *model.Homework) error {
	return DB.Create(h).Error
}

// UpdateHomework TODO 处理并发
func (d *homeworkDao) UpdateHomework(h *model.Homework) error {
	return DB.Save(h).Error
}
func (d *homeworkDao) DeleteHomework(h *model.Homework) error {
	return DB.Delete(h).Error
}

func (d *homeworkDao) GetHomeworkByTitle(title string) ([]model.Homework, error) {
	var homeworks []model.Homework
	tx := DB.Where("title = ?", title).Find(&homeworks)
	if tx.Error != nil {
		return nil, tx.Error
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
