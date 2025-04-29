package main

import (
	"io"
	"os"
	"study-spider-manhua-gin/business/order"
	"study-spider-manhua-gin/config"
	"study-spider-manhua-gin/db"
	"study-spider-manhua-gin/log"
	"study-spider-manhua-gin/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 初始化, 默认main会自动调用本方法
func init() {
	// 1. 读取配置文件， (如果配置文件不填, 自动会有默认值)
	cfg := config.GetConfig(".", "config", "yaml")

	// 2. 根据配置文件,设置日志相关,现在用logrus框架
	log.InitLog()

	// 获取日志实例
	log := log.GetLogger()

	// 设置日志级别
	switch cfg.Log.Level {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}

	// 创建一个文件用于写入日志
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Failed to open log file: %v", err)
	}
	defer file.Close() // 关闭日志文件

	// 使用 io.MultiWriter 实现多写入器功能
	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)

	// 打印配置
	log.Debug("network.ximalayaIIp_ip: ", cfg.Network.XimalayaIIp)
	log.Debug("db.name: ", cfg.DB.Name)
	log.Debug("gin模式: ", cfg.Gin.Mode)
	log.Println("------------ gin模式: ", cfg.Gin.Mode)

	// 初始化数据库连接
	db.InitDB("mysql", cfg.DB.Name, cfg.DB.User, cfg.DB.Password)

	// 自动迁移表结构
	db.DB.AutoMigrate(&models.Website{}, &models.Country{}, &models.Category{}, &models.Type{}, &models.Order{})

	// 插入默认数据
	db.InsertDefaultData()
}

func main() {
	gin.SetMode(gin.ReleaseMode) // 关键代码：切换到 release 模式
	r := gin.Default()
	r.Use(cors.Default()) // 允许所有跨域

	// 封装api
	r.POST("/orders", order.OrderAdd)
	r.DELETE("/orders/:id", order.OrderDelete)
	r.PUT("/orders", order.OrderUpdate)
	r.GET("/orders", order.OrdersPageQuery) // 分页查询

	r.Run(":8888") // 启动服务
}
