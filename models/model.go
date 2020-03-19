// creator: zacyuan
// date: 2019-12-28

package models

import (
	"errors"
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/config"
	"github.com/yuanzhangcai/microgin/log"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Model 数据库操作组件基类
type Model struct {
}

type dbLogger struct {
}

func (c *dbLogger) Print(v ...interface{}) {
	log.Info(v...)
}

// Init 初始化顾
func Init() error {
	var err error
	once.Do(func() {
		// 初始化连接
		dbInfo := config.Get("db", "server").String("")
		if dbInfo == "" {
			err = errors.New("没有获取到数据库配置")
			return
		}
		db, err = gorm.Open("mysql", dbInfo)
		if err != nil {
			log.Error("数据库初始化失败。错误信息：" + err.Error())
			return
		}

		// 取消DB复数
		db.SingularTable(true)

		if config.Get("db", "write_log").Bool(false) {
			// 设置sql语句输出到日志文件中
			db.LogMode(true)
			logger := &dbLogger{}
			db.SetLogger(logger)
		}
	})

	return err
}
