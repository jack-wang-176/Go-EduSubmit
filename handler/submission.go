package handler

import (
	"homework_submit/model"
	"homework_submit/pkg"
	"homework_submit/service"
	"strconv"
	"time"

	"github.com/jack-wang-176/Maple/web"
)

type submission struct{}

var Sub submission

func (s *submission) CreateSub(c *web.Context) {
	var req struct {
		ID      int    `json:"homework_id"`
		Content string `json:"content"`
	}
	if err := c.BindJson(&req); err != nil {
		SendResponse(c, nil, pkg.ParamError)
		return
	}

	user, flag := c.Get("user")
	if flag {
		sub, err := service.SubService.CreateSub(user.(string), req.Content, uint(req.ID))
		if err != nil {
			SendResponse(c, nil, err)
			return
		}
		SendResponse(c, map[string]interface{}{
			"id":           sub.ID,
			"homework_id":  sub.HomeworkID,
			"is_late":      sub.IsLate,
			"submitted_at": sub.SubmittedAt,
		}, nil, "提交成功")
	}
}

func (s *submission) MySub(c *web.Context) {

	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	val, exists := c.Get("userID")
	if !exists {
		SendResponse(c, nil, pkg.ServerError)
		return
	}
	userID := val.(uint)

	subs, err := service.SubService.MySub(userID, page, pageSize)
	if err != nil {
		SendResponse(c, nil, err)
		return
	}

	type MySubItem struct {
		ID       uint `json:"id"`
		Homework struct {
			ID              uint   `json:"id"`
			Title           string `json:"title"`
			Department      string `json:"department"`
			DepartmentLabel string `json:"department_label"`
		} `json:"homework"`
		Score       *int   `json:"score"` // 指针，支持返回 null
		Comment     string `json:"comment"`
		IsExcellent bool   `json:"is_excellent"`
		SubmittedAt string `json:"submitted_at"` // 格式化后的时间
	}

	resList := make([]MySubItem, 0)

	if subs != nil && subs.ListSub != nil {
		for _, item := range *subs.ListSub {
			if &item != nil {
				elem := MySubItem{
					ID:          item.ID,
					Score:       item.Score,
					Comment:     item.Comment,
					IsExcellent: item.IsExcellent,
					SubmittedAt: item.SubmittedAt.Format("2006-01-02 15:04:05"),
				}

				if item.Homework.ID != 0 {
					elem.Homework.ID = item.Homework.ID
					elem.Homework.Title = item.Homework.Title
					elem.Homework.Department = model.DeptNameMap[item.Homework.Department]
					elem.Homework.DepartmentLabel = model.DeptLabelMap[item.Homework.Department]
				} else {
					elem.Homework.Title = "未知作业"
				}
				resList = append(resList, elem)
			}
		}
	}
	var total int64 = 0
	if subs != nil {
		total = subs.Total
	}
	SendResponse(c, map[string]interface{}{
		"list":      resList,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	}, nil)
}

func (s *submission) ChangeSub(c *web.Context) {
	subIDStr, _ := c.Param("id")
	subID, err := strconv.ParseUint(subIDStr, 10, 64)
	if err != nil {
		SendResponse(c, nil, pkg.ParamError)
		return
	}

	var req struct {
		Score       int    `json:"score"`
		Comment     string `json:"comment"`
		IsExcellent bool   `json:"is_excellent"`
	}
	if err := c.BindJson(&req); err != nil {
		SendResponse(c, nil, pkg.ParamError)
		return
	}

	reviewerVal, exists := c.Get("userID")
	if !exists {
		SendResponse(c, nil, pkg.ServerError)
		return
	}
	reviewerID := reviewerVal.(uint)

	sub, err := service.SubService.ChangeSub(uint(subID), reviewerID, req.Comment, req.Score, req.IsExcellent)
	if err != nil {

		SendResponse(c, nil, err)
		return
	}

	data := map[string]interface{}{
		"id":           sub.ID,
		"score":        *sub.Score,
		"comment":      sub.Comment,
		"is_excellent": sub.IsExcellent,
		"reviewed_at":  time.Now().Format("2006-01-02 15:04:05"),
	}

	SendResponse(c, data, nil, "批改成功")
}
func (s *submission) GetExcellentList(c *web.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	list, err := service.SubService.GetExcellentList(page, pageSize)
	if err != nil {
		SendResponse(c, nil, err)
		return
	}
	var resSubs []model.SubmissionResponse
	for _, sub := range *list.ListSub {
		resSubs = append(resSubs, *sub.ToResponse())
	}
	SendResponse(c, resSubs, nil)
}
func (s *submission) GetWorkSubs(c *web.Context) {
	page := 1
	pageSize := 10

	param, err := c.Param("id")
	if err != nil {
		SendResponse(c, nil, pkg.ParamError)
		return
	}
	homeworkID, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		SendResponse(c, nil, pkg.ParamError)
		return
	}
	homework, err := service.HomeworkService.GetHomeworkId(uint(homeworkID))
	if err != nil {
		SendResponse(c, nil, pkg.ErrHomeworkNotFound)
		return
	}
	subs, err := service.SubService.GetWorkSubs(homeworkID, page, pageSize, homework.Department)
	if err != nil {
		SendResponse(c, nil, err)
		return
	}
	type StudentInfo struct {
		ID              uint   `json:"id"`
		Nickname        string `json:"nickname"`
		Department      string `json:"department"`
		DepartmentLabel string `json:"department_label"`
	}
	type SubmissionItem struct {
		ID          uint        `json:"id"`
		Student     StudentInfo `json:"student"`
		Content     string      `json:"content"`
		Score       *int        `json:"score"`
		Comment     string      `json:"comment"`
		IsExcellent bool        `json:"is_excellent"`
		IsLate      bool        `json:"is_late"`
		SubmittedAt string      `json:"submitted_at"`
	}

	resList := make([]SubmissionItem, 0)
	if subs != nil && subs.ListSub != nil {
		for _, item := range *subs.ListSub {
			if &item != nil {
				stu := StudentInfo{
					ID:       item.StudentID,
					Nickname: "未知用户",
				}
				if item.Student.ID != 0 {
					stu.ID = item.Student.ID
					stu.Nickname = item.Student.Nickname
					stu.Department = model.DeptNameMap[item.Student.Department]
					stu.DepartmentLabel = model.DeptLabelMap[item.Student.Department]
				}
				resList = append(resList, SubmissionItem{
					ID:          item.ID,
					Student:     stu,
					Content:     item.Content,
					Score:       item.Score,
					Comment:     item.Comment,
					IsExcellent: item.IsExcellent,
					IsLate:      item.IsLate,
					SubmittedAt: item.SubmittedAt.Format("2006-01-02 15:04:05"),
				})
			}
		}
	}

	var total int64 = 0
	if subs != nil {
		total = subs.Total
	}

	SendResponse(c, map[string]interface{}{
		"list":      resList,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	}, nil)
}
func (s *submission) MarkExcellent(c *web.Context) {
	subIDStr, _ := c.Param("id")
	subID, err := strconv.ParseUint(subIDStr, 10, 64)
	if err != nil {
		SendResponse(c, nil, pkg.ParamError)
		return
	}
	var req struct {
		IsExcellent bool `json:"is_excellent"`
	}
	if err := c.BindJson(&req); err != nil {
		SendResponse(c, nil, pkg.ParamError)
		return
	}

	val, exists := c.Get("userID")
	if !exists {
		SendResponse(c, nil, pkg.ServerError)
		return
	}
	reviewerID := val.(uint)

	err = service.SubService.SetExcellent(uint(subID), req.IsExcellent, reviewerID)
	if err != nil {
		SendResponse(c, nil, err)
		return
	}

	data := map[string]interface{}{
		"id":           uint(subID),
		"is_excellent": req.IsExcellent,
	}

	SendResponse(c, data, nil, "标记成功")
}
