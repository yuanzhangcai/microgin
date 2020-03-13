package tools

import (
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	hasInit       bool   = false
	logTimeFormat string = "2006-01-02 15:04:05.999999" //日志时间输入出格式
)

// InitLogrus 初始化log组件
func InitLogrus(logdir, logfile string, level uint32, maxDay time.Duration) {
	if hasInit {
		return
	}

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

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: logTimeFormat,
	}) //日志格式
	logrus.SetLevel(logrus.Level(level)) //设置日志等级

	hasInit = true
	return
}
