// creator: zacyuan
// date: 2019-12-28

package models

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/config"
	"github.com/sirupsen/logrus"
)

var dzsns *gorm.DB

// BaseModel 数据库操作组件基类
type BaseModel struct {
}

type dbLogger struct {
}

func (c *dbLogger) Print(v ...interface{}) {
	logrus.Info(v...)
}

// Init 初始化顾
func Init() error {
	// 如果已初始化过，就直接返回
	if dzsns != nil {
		return nil
	}

	// 初始化连接
	var err error
	dbInfo := config.Get("db", "dzsns").String("")
	if dbInfo == "" {
		return errors.New("没有获取到数据库配置")
	}
	dzsns, err = gorm.Open("mysql", dbInfo)
	if err != nil {
		fmt.Println("数据库初始化失败。错误信息：" + err.Error())
		logrus.Error("数据库初始化失败。错误信息：" + err.Error())
		return err
	}

	// 取消DB复数
	dzsns.SingularTable(true)

	if config.Get("db", "write_log").Bool(false) {
		// 设置sql语句输出到日志文件中
		dzsns.LogMode(true)
		logger := &dbLogger{}
		dzsns.SetLogger(logger)
	}

	return nil
}
