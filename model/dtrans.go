package model

import (
	"database/sql"
	"time"
)

var (
	DeptNameMap = map[Department]string{
		Backend:  "backend",
		Frontend: "frontend",
		Sre:      "sre",
		Product:  "product",
		Design:   "design",
		Android:  "android",
		Ios:      "ios",
	}
	DeptLabelMap = map[Department]string{
		Backend:  "后端",
		Frontend: "前端",
		Sre:      "SRE",
		Product:  "产品",
		Design:   "视觉设计",
		Android:  "Android",
		Ios:      "iOS",
	}
)

type UserResponse struct {
	ID              uint `json:"id"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       sql.NullTime `gorm:"index"`
	Username        string       `json:"username"`
	Nickname        string       `json:"nickname"`
	Role            string       `json:"role"`
	Department      string       `json:"department"`
	DepartmentLabel string       `json:"department_label"`
	Email           string       `json:"email,omitempty"`
}

func (u *User) ToResponse() *UserResponse {
	roleStr := "student"
	if u.Role == Admin {
		roleStr = "admin"
	}

	return &UserResponse{
		ID:              u.ID,
		Username:        u.Name,
		Nickname:        u.Nickname,
		Role:            roleStr,
		Department:      DeptNameMap[u.Department],
		DepartmentLabel: DeptLabelMap[u.Department],
		Email:           u.Email,
	}
}

type HomeworkResponse struct {
	ID              uint   `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	Department      string `json:"department"`
	DepartmentLabel string `json:"department_label"`
	CreatorID       uint   `json:"creator_id,omitempty"`
	CreatorName     string `json:"creator_name,omitempty"`
	Deadline        string `json:"deadline"`
	AllowLate       bool   `json:"allow_late"`
	SubmissionCount int64  `json:"submission_count,omitempty"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

func (h *Homework) ToResponse() *HomeworkResponse {
	return &HomeworkResponse{
		ID:              h.ID,
		Title:           h.Title,
		Description:     h.Description,
		Department:      DeptNameMap[h.Department],
		DepartmentLabel: DeptLabelMap[h.Department],
		CreatorID:       h.CreatorID,
		Deadline:        h.Deadline.Format("2006-01-02 15:04:05"),
		AllowLate:       h.AllowLate,
	}
}

type SubmissionResponse struct {
	ID              uint      `json:"id"`
	HomeworkID      uint      `gorm:"not null;index" json:"homework_id"`
	Homework        Homework  `gorm:"foreignKey:HomeworkID" json:"homework"`
	StudentID       uint      `gorm:"not null;index" json:"student_id"`
	Student         User      `gorm:"foreignKey:StudentID" json:"student"`
	SubmittedAt     time.Time `gorm:"not null" json:"submitted_at"`
	IsLate          bool      `gorm:"not null;default:false" json:"is_late"`
	ReviewerID      *uint     `gorm:"default:null" json:"reviewer_id"`
	Score           *int      `gorm:"default:null" json:"score"`
	Comment         string    `gorm:"type:text" json:"comment"`
	FileUrl         string    `gorm:"type:varchar(500)" json:"file_url"`
	Department      string    `json:"department"`
	DepartmentLabel string    `json:"department_label"`
	IsExcellent     bool      `gorm:"not null;default:false" json:"is_excellent"`
	Version         *int      `gorm:"default:1" json:"version"`
	CreatedAt       string    `json:"created_at"`
	UpdatedAt       string    `json:"updated_at"`
}

func (s *Submission) ToResponse() *SubmissionResponse {
	return &SubmissionResponse{
		ID:              s.ID,
		HomeworkID:      s.HomeworkID,
		StudentID:       s.StudentID,
		SubmittedAt:     s.SubmittedAt,
		IsLate:          s.IsLate,
		ReviewerID:      s.ReviewerID,
		Score:           s.Score,
		Comment:         s.Comment,
		FileUrl:         s.FileUrl,
		IsExcellent:     s.IsExcellent,
		Version:         s.Version,
		Department:      DeptNameMap[s.Department],
		DepartmentLabel: DeptLabelMap[s.Department],
		CreatedAt:       s.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:       s.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
