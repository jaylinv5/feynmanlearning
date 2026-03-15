#!/bin/bash

echo "🚀 启动费曼学习平台开发环境"

# 检查是否安装了docker-compose
if ! command -v docker-compose &> /dev/null; then
    echo "❌ 请先安装 docker-compose"
    exit 1
fi

# 检查是否安装了go
if ! command -v go &> /dev/null; then
    echo "❌ 请先安装 Go 1.22+"
    exit 1
fi

# 检查是否安装了node
if ! command -v node &> /dev/null; then
    echo "❌ 请先安装 Node.js 18+"
    exit 1
fi

echo "📦 启动依赖服务..."
docker-compose up -d mysql redis minio

echo "⏳ 等待数据库启动..."
sleep 10

echo "🔧 下载后端依赖..."
cd backend
go mod tidy

echo "🚀 启动后端服务..."
go run cmd/main.go &
BACKEND_PID=$!
echo "✅ 后端服务已启动，PID: $BACKEND_PID，监听端口: 8080"

cd ../frontend

echo "🔧 下载前端依赖..."
npm install

echo "🚀 启动前端服务..."
npm run dev &
FRONTEND_PID=$!
echo "✅ 前端服务已启动，PID: $FRONTEND_PID，访问地址: http://localhost:3000"

echo ""
echo "🎉 费曼学习平台开发环境已启动成功！"
echo ""
echo "📱 访问地址："
echo "  前端页面: http://localhost:3000"
echo "  后端API:  http://localhost:8080"
echo "  健康检查: http://localhost:8080/health"
echo "  MinIO控制台: http://localhost:9001 (账号: minioadmin / 密码: minioadmin)"
echo ""
echo "⌨️  按 Ctrl+C 停止所有服务"

# 等待用户中断
trap "echo '🛑 正在停止服务...'; kill $BACKEND_PID $FRONTEND_PID; docker-compose down; exit 0" INT

wait
