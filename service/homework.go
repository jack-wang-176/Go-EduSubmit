package service

import "homework_submit/model"

type homeworkService struct {
}

var HomeworkService = homeworkService{}

func (s *homeworkService) LaunchHomework(title, desp, creator string, late bool) (*model.Homework, error) {

}
