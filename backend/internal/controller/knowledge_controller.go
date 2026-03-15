package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/jaylinv5/feynmanlearning/internal/model"
	"github.com/jaylinv5/feynmanlearning/internal/middleware"
	"github.com/jaylinv5/feynmanlearning/internal/service"
)

// KnowledgePointController 知识点控制器
type KnowledgePointController struct {
	knowledgeService *service.KnowledgePointService
}

// NewKnowledgePointController 创建知识点控制器实例
func NewKnowledgePointController() *KnowledgePointController {
	return &KnowledgePointController{
		knowledgeService: service.NewKnowledgePointService(),
	}
}

// RegisterRoutes 注册路由
func (c *KnowledgePointController) RegisterRoutes(r *gin.RouterGroup) {
	knowledgeGroup := r.Group("/knowledge")
	{
		// 公开接口
		knowledgeGroup.GET("/list", c.List)
		knowledgeGroup.GET("/detail/:id", c.Detail)
		knowledgeGroup.GET("/subject/:subject/grade/:grade", c.GetBySubjectAndGrade)

		// 需要管理员权限的接口
		adminGroup := knowledgeGroup.Group("")
		adminGroup.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		{
			adminGroup.POST("/create", c.Create)
			adminGroup.PUT("/update", c.Update)
			adminGroup.DELETE("/delete/:id", c.Delete)
			adminGroup.PUT("/batch/status", c.BatchUpdateStatus)
			adminGroup.POST("/batch/import", c.BatchImport)
		}
	}
}

// Create 创建知识点
func (c *KnowledgePointController) Create(ctx *gin.Context) {
	var req model.CreateKnowledgePointRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Fail(400, "参数错误: "+err.Error()))
		return
	}

	userID := middleware.GetUserID(ctx)
	knowledge, err := c.knowledgeService.CreateKnowledgePoint(&req, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Fail(500, "创建失败: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, model.Success(knowledge))
}

// Update 更新知识点
func (c *KnowledgePointController) Update(ctx *gin.Context) {
	var req model.UpdateKnowledgePointRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Fail(400, "参数错误: "+err.Error()))
		return
	}

	userID := middleware.GetUserID(ctx)
	knowledge, err := c.knowledgeService.UpdateKnowledgePoint(&req, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Fail(500, "更新失败: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, model.Success(knowledge))
}

// Delete 删除知识点
func (c *KnowledgePointController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Fail(400, "无效的知识点ID"))
		return
	}

	if err := c.knowledgeService.DeleteKnowledgePoint(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Fail(500, "删除失败: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, model.Success(nil))
}

// Detail 获取知识点详情
func (c *KnowledgePointController) Detail(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Fail(400, "无效的知识点ID"))
		return
	}

	detail, err := c.knowledgeService.GetKnowledgePointDetail(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Fail(500, "获取详情失败: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, model.Success(detail))
}

// List 分页查询知识点列表
func (c *KnowledgePointController) List(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))
	subject := ctx.Query("subject")
	grade, _ := strconv.Atoi(ctx.Query("grade"))
	status, _ := strconv.Atoi(ctx.Query("status"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	list, total, err := c.knowledgeService.ListKnowledgePoints(page, pageSize, subject, grade, status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Fail(500, "查询失败: "+err.Error()))
		return
	}

	result := model.PageResult{
		Total:    total,
		Page:     page,
		PageSize: pageSize,
		Data:     list,
	}

	ctx.JSON(http.StatusOK, model.Success(result))
}

// GetBySubjectAndGrade 根据学科和年级获取知识点列表
func (c *KnowledgePointController) GetBySubjectAndGrade(ctx *gin.Context) {
	subject := ctx.Param("subject")
	gradeStr := ctx.Param("grade")
	grade, err := strconv.Atoi(gradeStr)
	if err != nil || grade < 1 || grade > 12 {
		ctx.JSON(http.StatusBadRequest, model.Fail(400, "无效的年级参数"))
		return
	}

	list, err := c.knowledgeService.GetKnowledgePointsBySubjectAndGrade(subject, grade)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Fail(500, "查询失败: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, model.Success(list))
}

// BatchUpdateStatus 批量更新知识点状态
func (c *KnowledgePointController) BatchUpdateStatus(ctx *gin.Context) {
	var req struct {
		IDs    []uint64 `json:"ids" binding:"required"`
		Status int      `json:"status" binding:"required,min=1,max=4"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Fail(400, "参数错误: "+err.Error()))
		return
	}

	if err := c.knowledgeService.BatchUpdateStatus(req.IDs, req.Status); err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Fail(500, "更新失败: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, model.Success(nil))
}

// BatchImport 批量导入知识点
func (c *KnowledgePointController) BatchImport(ctx *gin.Context) {
	var req []*model.CreateKnowledgePointRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Fail(400, "参数错误: "+err.Error()))
		return
	}

	userID := middleware.GetUserID(ctx)
	count, err := c.knowledgeService.BatchImport(req, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Fail(500, "导入失败: "+err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, model.Success(gin.H{
		"import_count": count,
	}))
}
