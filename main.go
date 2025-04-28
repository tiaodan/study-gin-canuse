package main

import (
	"log"
	"study-spider-manhua-gin/business/order"
	"study-spider-manhua-gin/config"
	"study-spider-manhua-gin/db"
	"study-spider-manhua-gin/logger"
	"study-spider-manhua-gin/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 初始化, 默认main会自动调用本方法
func init() {
	// 设置go自带log框架, 日志格式：日期时间 + 短文件名 + 行号
	// log.SetFlags(log.LstdFlags | log.Lshortfile)  // 不会影响自己写的 logger打印, 因此注释调

	// 获取配置实例（首次调用时触发初始化）
	cfg := config.GetConfig(".", "config", "yaml")

	// 读取配置文件，并设置为日志级别, 默认info
	switch cfg.Log.Level {
	case "debug":
		logger.SetLogLevel(logger.LevelDebug)
	case "info":
		logger.SetLogLevel(logger.LevelInfo)
	case "warn":
		logger.SetLogLevel(logger.LevelWarn)
	case "error":
		logger.SetLogLevel(logger.LevelError)
	default:
		logger.SetLogLevel(logger.LevelInfo)
	}

	// 打印配置
	logger.Debug("network.ximalayaIIp_ip: ", cfg.Network.XimalayaIIp)
	logger.Debug("db.name: ", cfg.DB.Name)
	logger.Debug("gin模式: ", cfg.Gin.Mode)
	log.Println("------------ gin模式: ", cfg.Gin.Mode)

	// 初始化数据库连接
	db.InitDB("mysql", cfg.DB.Name, cfg.DB.User, cfg.DB.Password)

	// 自动迁移表结构
	db.DB.AutoMigrate(&models.Website{}, &models.Country{}, &models.Category{}, &models.Type{}, &models.Order{}) // 有几个表, 写几个参数

	// 插入默认数据
	db.InsertDefaultData()
}

/*
思路:
 1. 读取配置文件， (如果配置文件不填, 自动会有默认值)
 2. 设置日志级别, 默认info
 3. 统一调用错误打印, 封装函数
 4. 封装restful api
*/
func main() {
	// 1. 读取配置文件， (如果配置文件不填, 自动会有默认值)
	// 2. 设置日志级别, 默认info
	// 3. 统一调用错误打印, 封装函数
	// 4. 封装restful api

	gin.SetMode(gin.ReleaseMode) // 关键代码：切换到 release 模式
	r := gin.Default()
	r.Use(cors.Default()) // 允许所有跨域

	// 封装api
	//---------------------------- 一会再弄这个
	r.POST("/orders", order.OrderAdd)
	r.DELETE("/orders/:id", order.OrderDelete)
	r.PUT("/orders", order.OrderUpdate)
	r.GET("/orders", order.OrdersPageQuery) // 分页查询

	r.Run(":8888") // 启动服务

}
