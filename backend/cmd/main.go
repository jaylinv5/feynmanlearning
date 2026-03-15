package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	"github.com/jaylinv5/feynmanlearning/internal/controller"
	"github.com/jaylinv5/feynmanlearning/internal/pkg/config"
	"github.com/jaylinv5/feynmanlearning/internal/pkg/database"
)

func main() {
	// 加载配置
	if err := config.LoadConfig("config/config.yaml"); err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 初始化数据库
	if err := database.InitMySQL(&config.GlobalConfig.MySQL); err != nil {
		log.Fatalf("初始化MySQL失败: %v", err)
	}
	defer database.CloseMySQL()

	// 初始化Redis
	if err := database.InitRedis(&config.GlobalConfig.Redis); err != nil {
		log.Fatalf("初始化Redis失败: %v", err)
	}
	defer database.CloseRedis()

	// 设置Gin模式
	gin.SetMode(config.GlobalConfig.Server.Mode)

	// 创建Gin引擎
	r := gin.Default()

	// 配置CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 健康检查接口
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "费曼学习平台服务运行正常",
			"time":    time.Now().Format(time.RFC3339),
		})
	})

	// API路由组
	v1 := r.Group("/api/v1")
	{
		// 注册知识点路由
		knowledgeController := controller.NewKnowledgePointController()
		knowledgeController.RegisterRoutes(v1)

		// 后续可以注册其他控制器路由
		// userController.RegisterRoutes(v1)
		// learningController.RegisterRoutes(v1)
		// feynmanController.RegisterRoutes(v1)
	}

	// 启动服务器
	serverAddr := fmt.Sprintf(":%d", config.GlobalConfig.Server.Port)
	srv := &http.Server{
		Addr:         serverAddr,
		Handler:      r,
		ReadTimeout:  time.Duration(config.GlobalConfig.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.GlobalConfig.Server.WriteTimeout) * time.Second,
	}

	// 启动服务
	go func() {
		log.Printf("服务器启动成功，监听地址: %s", serverAddr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("服务器启动失败: %v", err)
		}
	}()

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("正在关闭服务器...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("服务器关闭异常: %v", err)
	}

	log.Println("服务器已正常关闭")
}
