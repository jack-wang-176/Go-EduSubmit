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
	// 1. 获取参数 (并处理默认值)
	pageStr := c.Query("page")
	// ⚠️ 关键修正：前端发的是 pageSize，后端之前只读 page_size
	// 这里做个兼容，先读 pageSize，读不到再读 page_size
	pageSizeStr := c.Query("pageSize")
	if pageSizeStr == "" {
		pageSizeStr = c.Query("page_size")
	}

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	// 2. 🛡️ 容错处理：如果参数没传或者转数字失败，给默认值
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10 // 默认每页 10 条
	}

	// 3. 调用 Service
	list, err := service.SubService.GetExcellentList(page, pageSize)
	if err != nil {
		SendResponse(c, nil, err)
		return
	}

	// 定义响应结构体 (建议移到 handler 外面或 model 包里，但放这里也能用)
	type HomeworkInfo struct {
		ID              uint   `json:"id"`
		Title           string `json:"title"`
		Department      string `json:"department"`
		DepartmentLabel string `json:"department_label"`
	}

	type StudentInfo struct {
		ID       uint   `json:"id"`
		Nickname string `json:"nickname"`
	}

	type ExcellentItem struct {
		ID        uint         `json:"id"`
		Homework  HomeworkInfo `json:"homework"`
		Student   StudentInfo  `json:"student"`
		Score     int          `json:"score"`
		Comment   string       `json:"comment"`
		CreatedAt string       `json:"created_at"` // 建议加上时间
	}

	resList := make([]ExcellentItem, 0)

	// 4. 数据转换 (Model -> ViewModel)
	if list != nil && list.ListSub != nil {
		for _, item := range *list.ListSub {
			// 这里不需要 if &item != nil，range 出来的 item 是结构体值拷贝，永远不会是 nil

			elem := ExcellentItem{
				ID:      item.ID,
				Comment: item.Comment,
				Score:   0, // 默认 0
				// 格式化时间
				CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
			}

			if item.Score != nil {
				elem.Score = *item.Score
			}

			// 填充作业信息
			if item.Homework.ID != 0 {
				elem.Homework.ID = item.Homework.ID
				elem.Homework.Title = item.Homework.Title
				// 映射部门名称
				if val, ok := model.DeptNameMap[item.Homework.Department]; ok {
					elem.Homework.Department = val
				} else {
					elem.Homework.Department = strconv.Itoa(int(item.Homework.Department))
				}
				if val, ok := model.DeptLabelMap[item.Homework.Department]; ok {
					elem.Homework.DepartmentLabel = val
				}
			} else {
				elem.Homework.Title = "作业已被删除"
			}

			// 填充学生信息
			if item.Student.ID != 0 {
				elem.Student.ID = item.Student.ID
				elem.Student.Nickname = item.Student.Nickname
			} else {
				elem.Student.Nickname = "未知用户"
			}

			resList = append(resList, elem)
		}
	}

	var total int64 = 0
	if list != nil {
		total = list.Total
	}

	// 5. 构造返回数据
	data := map[string]interface{}{
		"list":      resList,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	}

	SendResponse(c, data, nil)
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
