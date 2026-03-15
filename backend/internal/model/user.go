package model

// User 用户模型
type User struct {
	BaseModel
	Username     string `gorm:"size:50;uniqueIndex;not null" json:"username"`
	Password     string `gorm:"size:100;not null" json:"-"`
	RealName     string `gorm:"size:50" json:"real_name"`
	Email        string `gorm:"size:100;uniqueIndex" json:"email"`
	Phone        string `gorm:"size:20;uniqueIndex" json:"phone"`
	Avatar       string `gorm:"size:255" json:"avatar"`
	Role         int    `gorm:"default:1;comment:1-学生,2-教师,3-管理员" json:"role"`
	Grade        int    `gorm:"comment:年级，7表示七年级" json:"grade"`
	Class        string `gorm:"size:50;comment:班级" json:"class"`
	School       string `gorm:"size:100;comment:学校" json:"school"`
	Status       int    `gorm:"default:1;comment:1-正常,2-禁用" json:"status"`
	LastLoginAt  *time.Time `json:"last_login_at"`
	LastLoginIP  string `gorm:"size:50" json:"last_login_ip"`
}

// UserProfile 用户个人信息
type UserProfile struct {
	User
	TotalStudyTime   int     `json:"total_study_time"`   // 总学习时长(分钟)
	Learned知识点数  int     `json:"learned_knowledge_count"` // 已学习知识点数
	Mastered知识点数 int     `json:"mastered_knowledge_count"` // 已掌握知识点数
	AverageScore     float64 `json:"average_score"`      // 平均得分
	Level            int     `json:"level"`              // 等级
	Experience       int     `json:"experience"`         // 经验值
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6,max=50"`
	RealName string `json:"real_name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Grade    int    `json:"grade" binding:"required,min=1,max=12"`
	Class    string `json:"class"`
	School   string `json:"school"`
}

// LoginResponse 登录返回
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
