# 部署指南

## 环境要求
- Docker & Docker Compose
- Go 1.22+
- Node.js 18+
- MySQL 8.0+
- Redis 7.0+

## 快速启动

### 方式一：一键启动（推荐）
```bash
# 克隆项目
git clone <仓库地址>
cd feynman-learning-platform

# 执行启动脚本
./start.sh
```

### 方式二：手动启动

#### 1. 启动依赖服务
```bash
docker-compose up -d mysql redis minio
```

#### 2. 启动后端服务
```bash
cd backend

# 下载依赖
go mod tidy

# 运行服务
go run cmd/main.go
```
后端服务启动后访问：http://localhost:8080/health

#### 3. 启动前端服务
```bash
cd frontend

# 下载依赖
npm install

# 启动开发服务
npm run dev
```
前端服务启动后访问：http://localhost:3000

## 访问地址
- 前端页面: http://localhost:3000
- 后端API:  http://localhost:8080
- API文档:  http://localhost:8080/swagger/index.html (后续集成)
- MinIO控制台: http://localhost:9001 (账号: minioadmin / 密码: minioadmin)

## 测试账号
- 管理员账号: admin / admin123

## 接口测试示例

### 1. 健康检查
```bash
curl http://localhost:8080/health
```

响应：
```json
{
  "status": "ok",
  "message": "费曼学习平台服务运行正常",
  "time": "2026-03-15T15:00:00+08:00"
}
```

### 2. 查询七年级数学知识点列表
```bash
curl http://localhost:8080/api/v1/knowledge/subject/math/grade/7
```

### 3. 获取知识点详情
```bash
curl http://localhost:8080/api/v1/knowledge/detail/1
```

### 4. 创建知识点（需要管理员token）
```bash
curl -X POST http://localhost:8080/api/v1/knowledge/create \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-token>" \
  -d '{
    "subject": "math",
    "grade": 7,
    "chapter": "有理数",
    "chapterOrder": 1,
    "name": "有理数的概念",
    "difficulty": 2,
    "estimatedTime": 15,
    "content": "有理数是整数（正整数、0、负整数）和分数的统称...",
    "examples": [],
    "exercises": [],
    "feynmanGuide": [],
    "preRequires": [],
    "tags": "有理数,七年级,数学"
  }'
```

## 生产部署

### 环境变量配置
```bash
# MySQL配置
FEYNMAN_MYSQL_HOST=your-mysql-host
FEYNMAN_MYSQL_PORT=3306
FEYNMAN_MYSQL_DATABASE=feynman
FEYNMAN_MYSQL_USERNAME=feynman
FEYNMAN_MYSQL_PASSWORD=your-password

# Redis配置
FEYNMAN_REDIS_HOST=your-redis-host
FEYNMAN_REDIS_PORT=6379
FEYNMAN_REDIS_PASSWORD=your-password

# JWT配置
FEYNMAN_JWT_SECRET=your-secret-key

# AI配置
FEYNMAN_AI_PROVIDER=tongyi
FEYNMAN_AI_API_KEY=your-api-key
```

### 后端编译部署
```bash
cd backend
go build -o feynman-server cmd/main.go
./feynman-server
```

### 前端编译部署
```bash
cd frontend
npm run build
# 将dist目录部署到Nginx等静态服务器
```

## 常见问题

### 1. 数据库连接失败
- 检查MySQL服务是否正常启动
- 确认配置文件中的数据库账号密码正确
- 检查防火墙是否开放3306端口

### 2. 前端无法访问后端接口
- 确认后端服务正常启动
- 检查CORS配置是否正确
- 确认代理配置正确

### 3. AI功能不可用
- 检查AI API Key是否配置正确
- 确认网络可以访问AI服务接口
- 查看服务日志排查具体错误
