package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jaylinv5/feynmanlearning/internal/model"
	"github.com/jaylinv5/feynmanlearning/internal/repository"
)

// KnowledgePointService 知识点服务
type KnowledgePointService struct {
	knowledgeRepo *repository.KnowledgePointRepository
}

// NewKnowledgePointService 创建知识点服务实例
func NewKnowledgePointService() *KnowledgePointService {
	return &KnowledgePointService{
		knowledgeRepo: repository.NewKnowledgePointRepository(),
	}
}

// CreateKnowledgePoint 创建知识点
func (s *KnowledgePointService) CreateKnowledgePoint(req *model.CreateKnowledgePointRequest, createdBy uint64) (*model.KnowledgePoint, error) {
	// 生成知识点编码
	code := fmt.Sprintf("%s-%d-%04d", req.Subject, req.Grade, time.Now().UnixNano()%10000)

	// 序列化JSON字段
	examplesJSON, err := json.Marshal(req.Examples)
	if err != nil {
		return nil, err
	}

	exercisesJSON, err := json.Marshal(req.Exercises)
	if err != nil {
		return nil, err
	}

	feynmanGuideJSON, err := json.Marshal(req.FeynmanGuide)
	if err != nil {
		return nil, err
	}

	preRequiresJSON, err := json.Marshal(req.PreRequires)
	if err != nil {
		return nil, err
	}

	knowledge := &model.KnowledgePoint{
		Subject:      req.Subject,
		Grade:        req.Grade,
		Chapter:      req.Chapter,
		ChapterOrder: req.ChapterOrder,
		Name:         req.Name,
		Code:         code,
		Difficulty:   req.Difficulty,
		EstimatedTime: req.EstimatedTime,
		Content:      req.Content,
		Examples:     string(examplesJSON),
		Exercises:    string(exercisesJSON),
		FeynmanGuide: string(feynmanGuideJSON),
		PreRequires:  string(preRequiresJSON),
		Tags:         req.Tags,
		Status:       1, // 草稿状态
		CreatedBy:    createdBy,
	}

	if err := s.knowledgeRepo.Create(knowledge); err != nil {
		return nil, err
	}

	return knowledge, nil
}

// UpdateKnowledgePoint 更新知识点
func (s *KnowledgePointService) UpdateKnowledgePoint(req *model.UpdateKnowledgePointRequest, updatedBy uint64) (*model.KnowledgePoint, error) {
	knowledge, err := s.knowledgeRepo.GetByID(req.ID)
	if err != nil {
		return nil, errors.New("知识点不存在")
	}

	// 更新字段
	if req.Subject != "" {
		knowledge.Subject = req.Subject
	}
	if req.Grade > 0 {
		knowledge.Grade = req.Grade
	}
	if req.Chapter != "" {
		knowledge.Chapter = req.Chapter
	}
	if req.ChapterOrder > 0 {
		knowledge.ChapterOrder = req.ChapterOrder
	}
	if req.Name != "" {
		knowledge.Name = req.Name
	}
	if req.Difficulty > 0 {
		knowledge.Difficulty = req.Difficulty
	}
	if req.EstimatedTime > 0 {
		knowledge.EstimatedTime = req.EstimatedTime
	}
	if req.Content != "" {
		knowledge.Content = req.Content
	}
	if req.Examples != nil {
		examplesJSON, err := json.Marshal(req.Examples)
		if err != nil {
			return nil, err
		}
		knowledge.Examples = string(examplesJSON)
	}
	if req.Exercises != nil {
		exercisesJSON, err := json.Marshal(req.Exercises)
		if err != nil {
			return nil, err
		}
		knowledge.Exercises = string(exercisesJSON)
	}
	if req.FeynmanGuide != nil {
		feynmanGuideJSON, err := json.Marshal(req.FeynmanGuide)
		if err != nil {
			return nil, err
		}
		knowledge.FeynmanGuide = string(feynmanGuideJSON)
	}
	if req.PreRequires != nil {
		preRequiresJSON, err := json.Marshal(req.PreRequires)
		if err != nil {
			return nil, err
		}
		knowledge.PreRequires = string(preRequiresJSON)
	}
	if req.Tags != "" {
		knowledge.Tags = req.Tags
	}
	if req.Status > 0 {
		knowledge.Status = req.Status
	}

	knowledge.UpdatedBy = updatedBy

	if err := s.knowledgeRepo.Update(knowledge); err != nil {
		return nil, err
	}

	return knowledge, nil
}

// DeleteKnowledgePoint 删除知识点
func (s *KnowledgePointService) DeleteKnowledgePoint(id uint64) error {
	_, err := s.knowledgeRepo.GetByID(id)
	if err != nil {
		return errors.New("知识点不存在")
	}
	return s.knowledgeRepo.Delete(id)
}

// GetKnowledgePointDetail 获取知识点详情
func (s *KnowledgePointService) GetKnowledgePointDetail(id uint64) (*model.KnowledgePointDetail, error) {
	return s.knowledgeRepo.GetDetail(id)
}

// ListKnowledgePoints 分页查询知识点列表
func (s *KnowledgePointService) ListKnowledgePoints(page, pageSize int, subject string, grade int, status int) ([]*model.KnowledgePoint, int64, error) {
	return s.knowledgeRepo.List(page, pageSize, subject, grade, status)
}

// GetKnowledgePointsBySubjectAndGrade 根据学科和年级获取知识点列表
func (s *KnowledgePointService) GetKnowledgePointsBySubjectAndGrade(subject string, grade int) ([]*model.KnowledgePointListItem, error) {
	knowledges, err := s.knowledgeRepo.ListBySubjectAndGrade(subject, grade)
	if err != nil {
		return nil, err
	}

	// 转换为列表项
	list := make([]*model.KnowledgePointListItem, len(knowledges))
	for i, k := range knowledges {
		list[i] = &model.KnowledgePointListItem{
			ID:           k.ID,
			Subject:      k.Subject,
			Grade:        k.Grade,
			Chapter:      k.Chapter,
			Name:         k.Name,
			Code:         k.Code,
			Difficulty:   k.Difficulty,
			EstimatedTime: k.EstimatedTime,
			Status:       k.Status,
			Tags:         k.Tags,
			Learned:      false, // 后续需要结合用户学习状态设置
			Mastered:     false, // 后续需要结合用户学习状态设置
		}
	}

	return list, nil
}

// BatchUpdateStatus 批量更新知识点状态
func (s *KnowledgePointService) BatchUpdateStatus(ids []uint64, status int) error {
	if len(ids) == 0 {
		return errors.New("知识点ID列表不能为空")
	}
	if status < 1 || status > 4 {
		return errors.New("无效的状态值")
	}
	return s.knowledgeRepo.UpdateStatus(ids, status)
}

// BatchImport 批量导入知识点
func (s *KnowledgePointService) BatchImport(knowledges []*model.CreateKnowledgePointRequest, createdBy uint64) (int, error) {
	if len(knowledges) == 0 {
		return 0, errors.New("导入数据不能为空")
	}

	var models []*model.KnowledgePoint
	for _, req := range knowledges {
		// 生成知识点编码
		code := fmt.Sprintf("%s-%d-%04d", req.Subject, req.Grade, time.Now().UnixNano()%10000)

		// 序列化JSON字段
		examplesJSON, _ := json.Marshal(req.Examples)
		exercisesJSON, _ := json.Marshal(req.Exercises)
		feynmanGuideJSON, _ := json.Marshal(req.FeynmanGuide)
		preRequiresJSON, _ := json.Marshal(req.PreRequires)

		model := &model.KnowledgePoint{
			Subject:      req.Subject,
			Grade:        req.Grade,
			Chapter:      req.Chapter,
			ChapterOrder: req.ChapterOrder,
			Name:         req.Name,
			Code:         code,
			Difficulty:   req.Difficulty,
			EstimatedTime: req.EstimatedTime,
			Content:      req.Content,
			Examples:     string(examplesJSON),
			Exercises:    string(exercisesJSON),
			FeynmanGuide: string(feynmanGuideJSON),
			PreRequires:  string(preRequiresJSON),
			Tags:         req.Tags,
			Status:       1, // 草稿状态
			CreatedBy:    createdBy,
		}
		models = append(models, model)
	}

	if err := s.knowledgeRepo.BatchCreate(models); err != nil {
		return 0, err
	}

	return len(models), nil
}
