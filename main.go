// creator: zacyuan
// date: 2019-12-28

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/web"
	"github.com/sirupsen/logrus"
	"github.com/yuanzhangcai/microgin/common"
	"github.com/yuanzhangcai/microgin/middleware"
	"github.com/yuanzhangcai/microgin/models"
	"github.com/yuanzhangcai/microgin/routers"
	"github.com/yuanzhangcai/microgin/tools"
)

const (
	// ConfigFile 通用配置文件路径
	ConfigFile = "./etc/microgin.toml"

	// ProdConfigFile 正式环境专有配置文件路径
	ProdConfigFile = "./etc/prod.toml"

	// TestConfigFile 测试环境专有配置文件路径
	TestConfigFile = "./etc/test.toml"

	// DevConfigFile 本地开发环境专用配置文件路径
	DevConfigFile = "./etc/dev.toml"
)

var (
	_commit    string // 最后一次提交的id
	_buildtime string // 编译时间
	_buildby   string // 编译人
	_env       string // 程序运行环境
	_version   string // 版本号
)

// InitConfig 载配置文件
func InitConfig(filename string) {
	// 加载配置文件
	err := config.LoadFile(ConfigFile)
	if err != nil {
		logrus.Error("读取配置文件" + ConfigFile + "失败。")
		os.Exit(-1)
	}

	// 加载正式环境配置文件
	err = config.LoadFile(ProdConfigFile)
	if err != nil {
		logrus.Error("读取配置文件" + ProdConfigFile + "失败。")
		os.Exit(-1)
	}

	if _env == common.TestEnv {
		// 加载测试环境配置文件
		err := config.LoadFile(TestConfigFile)
		if err != nil {
			logrus.Error("读取配置文件" + TestConfigFile + "失败。")
			os.Exit(-1)
		}
	}

	if _env == common.DevEnv {
		// 加载测试环境配置文件
		err := config.LoadFile(DevConfigFile)
		if err != nil {
			logrus.Error("读取配置文件" + DevConfigFile + "失败。")
			os.Exit(-1)
		}
	}

	logrus.Info("mode = " + config.Get("common", "mode").String(""))
}

// InitLog 初始化log
func InitLog() {
	logSetting := struct {
		FileDir string        `json:"filedir"`
		MaxDays time.Duration `json:"maxdays"`
		Level   uint32        `json:"level"`
	}{}
	config.Get("log").Scan(&logSetting)

	fileName := filepath.Base(os.Args[0]) + ".log"

	_, err := os.Stat(logSetting.FileDir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(logSetting.FileDir, 0755)
			if err != nil {
				fmt.Println("创建日志目录失败。")
				os.Exit(-1)
			}
		} else {
			fmt.Println("目录：" + logSetting.FileDir + "，stat失败")
			os.Exit(-1)
		}
	}

	tools.InitLogrus(logSetting.FileDir, fileName, logSetting.Level, logSetting.MaxDays)

	return
}

func init() {
	if _env == "" {
		// 如果编译时没有指定运行环境，则看运行是是否有加“--env=”参数
		flag.StringVar(&_env, "env", "", "Running environment.")
		flag.Parse()
	}

	// 设置版本信息
	common.SetVersion(map[string]string{
		"commit":    _commit,    // 最后一次提交的id
		"buildtime": _buildtime, // 编译时间
		"buildby":   _buildby,   // 编译人
		"env":       _env,       // 程序运行环境
		"version":   _version,   // 版本号
	})

	// 设置当前运行环境
	common.Env = _env

	// 加载配置文件
	InitConfig(ConfigFile)

	// 初始化log
	InitLog()

	// 初始化DB
	if err := models.Init(); err != nil {
		logrus.Error(err.Error())
		os.Exit(-1)
	}

	// 初始化php redis
	err := tools.InitPhpRedis(config.Get("redis", "server").String(""), config.Get("redis", "password").String(""))
	if err != nil {
		logrus.Error(err.Error())
		os.Exit(-1)
	}

	showInfo()
}

func main() {

	if common.Env == common.ProdEnv {
		// 正式环境时，将gin的模式，设置成ReleaseMode
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// 添加日志中间件
	router.Use(
		middleware.UsedTime(), // 耗时中间件
		gin.Recovery(),        // 异常恢复中间件
		middleware.Accessor(), //跨域访问中间件
		middleware.Auth(),     // 登录验证中间件
	)

	// 创建路由规则
	routers.CreateRouters(router)

	if config.Get("common", "mode").String("") == "micro" {
		// 以微服务型式启动

		// 创建微服务
		service := web.NewService(
			web.Name("go.micro.api.microgin"),
		)

		// 服务初始化
		service.Init()

		// 注册Handler事件
		service.Handle("/", router)

		// Run server
		if err := service.Run(); err != nil {
			logrus.Error(err)
		}
	} else {
		// 以传统web服务启动
		router.Run(config.Get("gin", "address").String(""))
	}
}

func showInfo() {
	fmt.Println("=======================================================================")
	fmt.Println("     build commit : " + _commit)
	fmt.Println("     build time : " + _buildtime)
	fmt.Println("     build user : " + _buildby)
	fmt.Println("     version : " + _version)
	fmt.Println("     run env : " + _env)
	fmt.Println("=======================================================================")
}
