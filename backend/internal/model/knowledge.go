package model

// KnowledgePoint 知识点模型
type KnowledgePoint struct {
	BaseModel
	Subject      string    `gorm:"size:20;not null;index;comment:学科:math/chinese/english" json:"subject"`
	Grade        int       `gorm:"not null;index;comment:年级" json:"grade"`
	Chapter      string    `gorm:"size:100;not null;comment:章节名称" json:"chapter"`
	ChapterOrder int       `gorm:"not null;comment:章节排序" json:"chapter_order"`
	Name         string    `gorm:"size:100;not null;comment:知识点名称" json:"name"`
	Code         string    `gorm:"size:50;uniqueIndex;not null;comment:知识点编码" json:"code"`
	Difficulty   int       `gorm:"not null;default:2;comment:难度:1-简单,2-中等,3-困难" json:"difficulty"`
	EstimatedTime int      `gorm:"not null;default:15;comment:预计学习时长(分钟)" json:"estimated_time"`
	Content      string    `gorm:"type:text;not null;comment:知识点讲解内容" json:"content"`
	Examples     string    `gorm:"type:text;comment:例题(JSON数组)" json:"-"`
	Exercises    string    `gorm:"type:text;comment:练习题(JSON数组)" json:"-"`
	FeynmanGuide string    `gorm:"type:text;comment:费曼引导问题(JSON数组)" json:"-"`
	PreRequires  string    `gorm:"type:text;comment:前置知识点ID列表(JSON数组)" json:"-"`
	Status       int       `gorm:"default:1;comment:1-草稿,2-审核中,3-上线,4-下线" json:"status"`
	Tags         string    `gorm:"size:255;comment:标签，逗号分隔" json:"tags"`
	CreatedBy    uint64    `gorm:"not null;comment:创建人ID" json:"created_by"`
	UpdatedBy    uint64    `comment:更新人ID" json:"updated_by"`
}

// Example 例题结构
type Example struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Analysis    string `json:"analysis"`
	Answer      string `json:"answer"`
	Difficulty  int    `json:"difficulty"`
}

// Exercise 练习题结构
type Exercise struct {
	ID          string   `json:"id"`
	Type        int      `json:"type"` // 1-单选,2-多选,3-判断,4-填空,5-简答
	Question    string   `json:"question"`
	Options     []string `json:"options,omitempty"`
	Answer      string   `json:"answer"`
	Analysis    string   `json:"analysis"`
	Difficulty  int      `json:"difficulty"`
	Score       int      `json:"score"`
}

// FeynmanGuide 费曼引导问题
type FeynmanGuide struct {
	ID          string `json:"id"`
	Question    string `json:"question"`
	Difficulty  int    `json:"difficulty"`
	KeyPoint    string `json:"key_point"` // 考察的核心要点
}

// KnowledgePointListItem 知识点列表项
type KnowledgePointListItem struct {
	ID           uint64 `json:"id"`
	Subject      string `json:"subject"`
	Grade        int    `json:"grade"`
	Chapter      string `json:"chapter"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	Difficulty   int    `json:"difficulty"`
	EstimatedTime int   `json:"estimated_time"`
	Status       int    `json:"status"`
	Tags         string `json:"tags"`
	Learned      bool   `json:"learned"`  // 是否已学习
	Mastered     bool   `json:"mastered"` // 是否已掌握
}

// KnowledgePointDetail 知识点详情
type KnowledgePointDetail struct {
	KnowledgePoint
	Examples     []Example     `json:"examples"`
	Exercises    []Exercise    `json:"exercises"`
	FeynmanGuide []FeynmanGuide `json:"feynman_guide"`
	PreRequires  []uint64      `json:"pre_requires"`
}

// CreateKnowledgePointRequest 创建知识点请求
type CreateKnowledgePointRequest struct {
	Subject      string         `json:"subject" binding:"required"`
	Grade        int            `json:"grade" binding:"required,min=1,max=12"`
	Chapter      string         `json:"chapter" binding:"required"`
	ChapterOrder int            `json:"chapter_order" binding:"required,min=1"`
	Name         string         `json:"name" binding:"required"`
	Difficulty   int            `json:"difficulty" binding:"required,min=1,max=3"`
	EstimatedTime int           `json:"estimated_time" binding:"required,min=5,max=60"`
	Content      string         `json:"content" binding:"required"`
	Examples     []Example      `json:"examples"`
	Exercises    []Exercise     `json:"exercises"`
	FeynmanGuide []FeynmanGuide `json:"feynman_guide"`
	PreRequires  []uint64       `json:"pre_requires"`
	Tags         string         `json:"tags"`
}

// UpdateKnowledgePointRequest 更新知识点请求
type UpdateKnowledgePointRequest struct {
	ID           uint64         `json:"id" binding:"required"`
	Subject      string         `json:"subject"`
	Grade        int            `json:"grade" binding:"omitempty,min=1,max=12"`
	Chapter      string         `json:"chapter"`
	ChapterOrder int            `json:"chapter_order" binding:"omitempty,min=1"`
	Name         string         `json:"name"`
	Difficulty   int            `json:"difficulty" binding:"omitempty,min=1,max=3"`
	EstimatedTime int           `json:"estimated_time" binding:"omitempty,min=5,max=60"`
	Content      string         `json:"content"`
	Examples     []Example      `json:"examples"`
	Exercises    []Exercise     `json:"exercises"`
	FeynmanGuide []FeynmanGuide `json:"feynman_guide"`
	PreRequires  []uint64       `json:"pre_requires"`
	Tags         string         `json:"tags"`
	Status       int            `json:"status" binding:"omitempty,min=1,max=4"`
}
