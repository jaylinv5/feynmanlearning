# 费曼学习法自主学习平台

面向中小学生的AI驱动自主学习平台，基于费曼学习法核心原理，通过"学习-教学-验证"的闭环学习模式，帮助学生深度掌握知识点。

## 技术栈
- **后端**: Go 1.22 + Gin + GORM
- **前端**: React 18 + TypeScript + Tailwind CSS + shadcn/ui
- **数据库**: MySQL 8.0 + Redis 7.0
- **AI**: 大语言模型API(通义千问/文心一言/Claude) + 向量数据库
- **部署**: Docker + Kubernetes

## 项目结构
```
feynman-learning-platform/
├── backend/                 # 后端代码
│   ├── cmd/                 # 启动入口
│   ├── internal/            # 内部业务逻辑
│   │   ├── controller/      # 控制器层
│   │   ├── service/         # 业务逻辑层
│   │   ├── model/           # 数据模型层
│   │   ├── repository/      # 数据访问层
│   │   ├── middleware/      # 中间件
│   │   └── pkg/             # 公共工具包
│   ├── config/              # 配置文件
│   ├── deploy/              # 部署相关
│   └── go.mod
├── frontend/                # 前端代码
│   ├── src/
│   │   ├── components/      # 公共组件
│   │   ├── pages/           # 页面组件
│   │   ├── services/        # API服务
│   │   ├── store/           # 状态管理
│   │   └── utils/           # 工具函数
│   ├── package.json
│   └── tsconfig.json
├── docs/                    # 项目文档
└── docker-compose.yml       # 本地开发环境
```

## 核心功能模块
1. **知识体系模块**：知识点管理、分类、关联图谱
2. **AI交互学习模块**：智能讲解、实时提问、学习诊断
3. **费曼教学验证模块**：角色反转、智能体提问、讲授质量评估
4. **个人空间模块**：学习进度、掌握图谱、错题本
5. **系统管理模块**：用户、班级、权限、统计

## 快速开始
### 本地开发环境
```bash
# 启动依赖服务
docker-compose up -d mysql redis

# 启动后端
cd backend
go run cmd/main.go

# 启动前端
cd frontend
npm install
npm run dev
```

## 开发进度
- [ ] 第一周：基础架构 + 知识点系统
- [ ] 第二周：AI交互学习模块
- [ ] 第三周：费曼教学验证模块
- [ ] 第四周：个人空间 + 系统上线