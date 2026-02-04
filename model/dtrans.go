package model

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
	ID              uint   `json:"id"`
	Username        string `json:"username"`
	Nickname        string `json:"nickname"`
	Role            string `json:"role"`
	Department      string `json:"department"`
	DepartmentLabel string `json:"department_label"`
	Email           string `json:"email,omitempty"`
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
