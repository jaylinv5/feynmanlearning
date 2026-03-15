package repository

import (
	"encoding/json"

	"gorm.io/gorm"

	"github.com/jaylinv5/feynmanlearning/internal/model"
	"github.com/jaylinv5/feynmanlearning/internal/pkg/database"
)

// KnowledgePointRepository 知识点仓库
type KnowledgePointRepository struct {
	db *gorm.DB
}

// NewKnowledgePointRepository 创建知识点仓库实例
func NewKnowledgePointRepository() *KnowledgePointRepository {
	return &KnowledgePointRepository{
		db: database.GetDB(),
	}
}

// Create 创建知识点
func (r *KnowledgePointRepository) Create(knowledge *model.KnowledgePoint) error {
	return r.db.Create(knowledge).Error
}

// Update 更新知识点
func (r *KnowledgePointRepository) Update(knowledge *model.KnowledgePoint) error {
	return r.db.Save(knowledge).Error
}

// Delete 删除知识点(软删除)
func (r *KnowledgePointRepository) Delete(id uint64) error {
	return r.db.Delete(&model.KnowledgePoint{}, id).Error
}

// GetByID 根据ID获取知识点
func (r *KnowledgePointRepository) GetByID(id uint64) (*model.KnowledgePoint, error) {
	var knowledge model.KnowledgePoint
	if err := r.db.First(&knowledge, id).Error; err != nil {
		return nil, err
	}
	return &knowledge, nil
}

// GetByCode 根据编码获取知识点
func (r *KnowledgePointRepository) GetByCode(code string) (*model.KnowledgePoint, error) {
	var knowledge model.KnowledgePoint
	if err := r.db.Where("code = ?", code).First(&knowledge).Error; err != nil {
		return nil, err
	}
	return &knowledge, nil
}

// List 分页查询知识点列表
func (r *KnowledgePointRepository) List(page, pageSize int, subject string, grade int, status int) ([]*model.KnowledgePoint, int64, error) {
	var list []*model.KnowledgePoint
	var total int64

	query := r.db.Model(&model.KnowledgePoint{})

	if subject != "" {
		query = query.Where("subject = ?", subject)
	}
	if grade > 0 {
		query = query.Where("grade = ?", grade)
	}
	if status > 0 {
		query = query.Where("status = ?", status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Order("subject asc, grade asc, chapter_order asc, id asc").Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

// ListBySubjectAndGrade 根据学科和年级获取知识点列表
func (r *KnowledgePointRepository) ListBySubjectAndGrade(subject string, grade int) ([]*model.KnowledgePoint, error) {
	var list []*model.KnowledgePoint
	if err := r.db.Where("subject = ? AND grade = ? AND status = 3", subject, grade).
		Order("chapter_order asc, id asc").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

// GetDetail 获取知识点详情(包含反序列化的JSON字段)
func (r *KnowledgePointRepository) GetDetail(id uint64) (*model.KnowledgePointDetail, error) {
	knowledge, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}

	detail := &model.KnowledgePointDetail{
		KnowledgePoint: *knowledge,
	}

	// 反序列化例题
	if knowledge.Examples != "" {
		var examples []model.Example
		if err := json.Unmarshal([]byte(knowledge.Examples), &examples); err == nil {
			detail.Examples = examples
		}
	}

	// 反序列化练习题
	if knowledge.Exercises != "" {
		var exercises []model.Exercise
		if err := json.Unmarshal([]byte(knowledge.Exercises), &exercises); err == nil {
			detail.Exercises = exercises
		}
	}

	// 反序列化费曼引导问题
	if knowledge.FeynmanGuide != "" {
		var guide []model.FeynmanGuide
		if err := json.Unmarshal([]byte(knowledge.FeynmanGuide), &guide); err == nil {
			detail.FeynmanGuide = guide
		}
	}

	// 反序列化前置知识点
	if knowledge.PreRequires != "" {
		var preRequires []uint64
		if err := json.Unmarshal([]byte(knowledge.PreRequires), &preRequires); err == nil {
			detail.PreRequires = preRequires
		}
	}

	return detail, nil
}

// BatchCreate 批量创建知识点
func (r *KnowledgePointRepository) BatchCreate(knowledges []*model.KnowledgePoint) error {
	return r.db.CreateInBatches(knowledges, 50).Error
}

// UpdateStatus 批量更新知识点状态
func (r *KnowledgePointRepository) UpdateStatus(ids []uint64, status int) error {
	return r.db.Model(&model.KnowledgePoint{}).Where("id IN ?", ids).Update("status", status).Error
}
