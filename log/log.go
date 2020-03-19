package log

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	once          sync.Once
	logTimeFormat string = "2006-01-02 15:04:05.999999" //日志时间输入出格式
)

// InitLog 初始化log组件
func InitLog(logdir, logfile string, level uint32, maxDay time.Duration) {
	// 只做一次
	once.Do(func() {
		maxDay = time.Hour * 24 * maxDay
		interval := time.Hour
		format := "%Y-%m-%d_%H"

		baseLogPath := logdir + logfile
		writer, _ := rotatelogs.New(
			baseLogPath+"."+format,
			//rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件
			rotatelogs.WithMaxAge(maxDay),         // 文件最大保存时间
			rotatelogs.WithRotationTime(interval), // 日志切割时间间隔
		)

		pathMap := lfshook.WriterMap{
			logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
			logrus.InfoLevel:  writer,
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: writer,
			logrus.FatalLevel: writer,
			logrus.PanicLevel: writer,
		}

		logrus.AddHook(lfshook.NewHook(pathMap, &logrus.JSONFormatter{
			TimestampFormat: logTimeFormat,
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime: "atime",
			},
		}))

		//logrus.SetReportCaller(true)
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: logTimeFormat,
		}) //日志格式
		logrus.SetLevel(logrus.Level(level)) //设置日志等级
	})

}

func getFileInfo() string {
	pc, filename, line, _ := runtime.Caller(2)
	filename = filepath.Base(filename)
	funcname := runtime.FuncForPC(pc).Name()
	funcname = filepath.Ext(funcname)
	funcname = strings.TrimPrefix(funcname, ".")
	return fmt.Sprintf("%v:%v:%v", filename, funcname, line)
}

// func TdeError(serial string, nid uint64, msg string) {
// 	file := getFileInfo()
// 	logrus.WithFields(logrus.Fields{
// 		"bfile":   file,
// 		"bserial": serial,
// 		"bnid":    nid,
// 	}).Error(msg)
// }

// Trace Trace级别日志
func Trace(args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"bfile": getFileInfo(),
	}).Trace(args...)
}

// Debug Debug级别日志
func Debug(args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"bfile": getFileInfo(),
	}).Debug(args...)
}

// Info Info级别日志
func Info(args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"bfile": getFileInfo(),
	}).Info(args...)
}

// Warn Warn级别日志
func Warn(args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"bfile": getFileInfo(),
	}).Warn(args...)
}

// Error Error级别日志
func Error(args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"bfile": getFileInfo(),
	}).Error(args...)
}

// Fatal Fatal级别日志
func Fatal(args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"bfile": getFileInfo(),
	}).Fatal(args...)
}

// Panic Panic级别日志
func Panic(args ...interface{}) {
	logrus.WithFields(logrus.Fields{
		"bfile": getFileInfo(),
	}).Panic(args...)
}
