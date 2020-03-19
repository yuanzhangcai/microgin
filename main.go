// creator: zacyuan
// date: 2019-12-28

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
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

// InitConfig 载配置文件
func InitConfig() {
	// 加载通用配置文件
	filepath := common.CurrRunPath + "/etc/"
	configFile := filepath + common.CurrRunFileName + ".toml"
	err := config.LoadFile(configFile)
	if err != nil {
		logrus.Error("读取配置文件" + configFile + "失败。")
		os.Exit(-1)
	}

	// 加载通用配置文件
	envFile := filepath + common.Env + ".toml"
	err = config.LoadFile(envFile)
	if err != nil {
		logrus.Error("读取当前运行环境配置文件" + envFile + "失败。")
		os.Exit(-1)
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
				fmt.Println("创建日志目录(" + logSetting.FileDir + ")失败。")
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
	// 如果编译时没有指定运行环境，则看运行是是否有加“--env=”参数
	env := ""
	flag.StringVar(&env, "env", "", "Running environment[dev test prod].")
	flag.Parse()

	if env != "" {
		common.Env = env
	}

	showInfo()

	if common.Env != common.DevEnv && common.Env != common.TestEnv && common.Env != common.ProdEnv {
		fmt.Println("当前运行环境不正确。")
		os.Exit(-1)
	}

	// 获取当前程序运行信息
	common.GetRunInfo()

	// 加载配置文件
	InitConfig()

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
}

func main() {

	if common.Env == common.ProdEnv {
		// 正式环境时，将gin的模式，设置成ReleaseMode
		gin.SetMode(gin.ReleaseMode)
	} else {
		go func() {
			log.Println(http.ListenAndServe(":6063", nil))
		}()
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
	fmt.Println("     Version   : " + common.Version)
	fmt.Println("     Env       : " + common.Env)
	fmt.Println("     Commit    : " + common.Commit)
	fmt.Println("     BuildTime : " + common.BuildTime)
	fmt.Println("     BuildUser : " + common.BuildUser)
	fmt.Println("     GoVersion : " + common.GoVersion)
	fmt.Println("=======================================================================")
}
