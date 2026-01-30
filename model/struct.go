package model

import (
	"time"

	"gorm.io/gorm"
)

type Role int8
type Department int8

const (
	Student Role = iota + 1
	Admin
)
const (
	Backend Department = iota + 1
	Frontend
	Sre
	Product
	Design
	Android
	Ios
)

type User struct {
	gorm.Model
	Name       string     `gorm:"type:varchar(50);not null;unique" json:"name"`
	Email      string     `gorm:"type:varchar(100);not null;unique" json:"email"`
	Password   string     `gorm:"type:varchar(255);not null" json:"-"`
	Nickname   string     `gorm:"type:varchar(50);not null" json:"nickname"`
	Role       Role       `gorm:"type:tinyint;not null;default:1;comment:1=Student,2=Admin" json:"role"`
	Department Department `gorm:"type:tinyint;not null;default:1;comment:1=Backend..." json:"department"`
}

type Homework struct {
	gorm.Model
	Title       string     `gorm:"type:varchar(200);not null" json:"title"`
	Description string     `gorm:"type:text;default:nil" json:"description"`
	CreatorID   uint       `gorm:"not null" json:"creator_id"`
	Deadline    time.Time  `gorm:"not null" json:"deadline"`
	AllowLate   bool       `gorm:"not null;default:false" json:"allow_late"`
	Department  Department `gorm:"type:tinyint;not null;comment:展示所属部门" json:"department"`
}

type Submission struct {
	gorm.Model
	HomeworkID  uint      `gorm:"not null;index" json:"homework_id"`
	StudentID   uint      `gorm:"not null;index" json:"student_id"`
	SubmittedAt time.Time `gorm:"not null" json:"submitted_at"`
	IsLate      bool      `gorm:"not null;default:false" json:"is_late"`
	ReviewerID  *uint     `gorm:"default:null" json:"reviewer_id"`
	Score       *int      `gorm:"default:null" json:"score"`
	Comment     string    `gorm:"type:text" json:"comment"`
	FileUrl     string    `gorm:"type:varchar(500)" json:"file_url"`
	IsExcellent bool      `gorm:"not null;default:false" json:"is_excellent"`
	Version     *int      `gorm:"default:1" json:"version"`
}
