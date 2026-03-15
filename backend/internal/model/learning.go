package model

import "time"

// LearningRecord 学习记录模型
type LearningRecord struct {
	BaseModel
	UserID           uint64    `gorm:"not null;index:idx_user_knowledge,unique" json:"user_id"`
	KnowledgePointID uint64    `gorm:"not null;index:idx_user_knowledge,unique" json:"knowledge_point_id"`
	Status           int       `gorm:"default:1;comment:1-学习中,2-已学完,3-已掌握" json:"status"`
	StudyTime        int       `gorm:"default:0;comment:学习时长(秒)" json:"study_time"`
	Progress         int       `gorm:"default:0;comment:学习进度(0-100)" json:"progress"`
	FirstStudyAt     *time.Time `json:"first_study_at"`
	LastStudyAt      *time.Time `json:"last_study_at"`
	CompletedAt      *time.Time `json:"completed_at"` // 学完时间
	MasteredAt       *time.Time `json:"mastered_at"`  // 掌握时间(费曼验证通过)
	TotalAttempts    int       `gorm:"default:0;comment:费曼验证尝试次数" json:"total_attempts"`
	BestScore        int       `gorm:"default:0;comment:费曼验证最高得分" json:"best_score"`
	LastScore        int       `gorm:"default:0;comment:最近一次费曼验证得分" json:"last_score"`
}

// LearningSession 学习会话模型
type LearningSession struct {
	BaseModel
	UserID           uint64    `gorm:"not null;index" json:"user_id"`
	KnowledgePointID uint64    `gorm:"not null;index" json:"knowledge_point_id"`
	SessionType      int       `gorm:"not null;comment:1-学习模式,2-费曼教学模式" json:"session_type"`
	Status           int       `gorm:"default:1;comment:1-进行中,2-已完成,3-已中断" json:"status"`
	StartTime        time.Time `json:"start_time"`
	EndTime          *time.Time `json:"end_time"`
	Duration         int       `gorm:"default:0;comment:会话时长(秒)" json:"duration"`
	MessageCount     int       `gorm:"default:0;comment:交互消息数量" json:"message_count"`
	Score            int       `gorm:"default:0;comment:会话得分" json:"score"`
	Evaluation       string    `gorm:"type:text;comment:会话评估结果" json:"evaluation"`
}

// LearningMessage 学习会话消息模型
type LearningMessage struct {
	BaseModel
	SessionID   uint64 `gorm:"not null;index" json:"session_id"`
	UserID      uint64 `gorm:"not null;index" json:"user_id"`
	Role        int    `gorm:"not null;comment:1-用户,2-AI" json:"role"`
	Content     string `gorm:"type:text;not null" json:"content"`
	MessageType int    `gorm:"default:1;comment:1-文本,2-图片,3-语音,4-文件" json:"message_type"`
	MetaData    string `gorm:"type:text;comment:元数据(JSON)" json:"meta_data"`
	IsEvaluated bool   `gorm:"default:false;comment:是否已评估" json:"is_evaluated"`
}

// ExerciseRecord 练习题记录模型
type ExerciseRecord struct {
	BaseModel
	UserID           uint64    `gorm:"not null;index" json:"user_id"`
	KnowledgePointID uint64    `gorm:"not null;index" json:"knowledge_point_id"`
	ExerciseID       string    `gorm:"size:50;not null" json:"exercise_id"`
	UserAnswer       string    `gorm:"type:text" json:"user_answer"`
	CorrectAnswer    string    `gorm:"type:text;not null" json:"correct_answer"`
	IsCorrect        bool      `gorm:"default:false" json:"is_correct"`
	Score            int       `gorm:"default:0" json:"score"`
	TimeSpent        int       `gorm:"default:0;comment:用时(秒)" json:"time_spent"`
	AnsweredAt       time.Time `json:"answered_at"`
}

// FeynmanRecord 费曼教学记录模型
type FeynmanRecord struct {
	BaseModel
	UserID           uint64    `gorm:"not null;index" json:"user_id"`
	KnowledgePointID uint64    `gorm:"not null;index" json:"knowledge_point_id"`
	SessionID        uint64    `gorm:"not null;uniqueIndex" json:"session_id"`
	Score            int       `gorm:"not null;comment:总得分(0-100)" json:"score"`
	AccuracyScore    int       `gorm:"not null;comment:准确性得分(0-40)" json:"accuracy_score"`
	CompletenessScore int      `gorm:"not null;comment:完整性得分(0-30)" json:"completeness_score"`
	LogicScore       int       `gorm:"not null;comment:逻辑性得分(0-20)" json:"logic_score"`
	ExpressionScore  int       `gorm:"not null;comment:表达能力得分(0-10)" json:"expression_score"`
	Evaluation       string    `gorm:"type:text;not null;comment:详细评估" json:"evaluation"`
	Suggestions      string    `gorm:"type:text;comment:改进建议" json:"suggestions"`
	IsPassed         bool      `gorm:"default:false;comment:是否通过" json:"is_passed"`
	WeakPoints       string    `gorm:"type:text;comment:薄弱点(JSON数组)" json:"weak_points"`
}

// UserKnowledgeProgress 用户知识点进度
type UserKnowledgeProgress struct {
	KnowledgePointID uint64 `json:"knowledge_point_id"`
	Name             string `json:"name"`
	Subject          string `json:"subject"`
	Chapter          string `json:"chapter"`
	Status           int    `json:"status"` // 1-未学习,2-学习中,3-已学完,4-已掌握
	Progress         int    `json:"progress"`
	StudyTime        int    `json:"study_time"`
	BestScore        int    `json:"best_score"`
}
