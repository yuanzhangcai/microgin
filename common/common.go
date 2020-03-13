// creator: zacyuan
// date: 2019-12-28

package common

const (
	// DevEnv 开发环境
	DevEnv = "dev"
	// TestEnv 测试环境
	TestEnv = "test"
	// ProdEnv 生产环境
	ProdEnv = "prod"
)

// 日期格式化模版
const (
	Y      string = "2006"
	YM     string = "2006-01"
	YMD    string = "2006-01-02"
	YMD2   string = "20060102"
	YMDH   string = "2006-01-02 15"
	YMDHI  string = "2006-01-02 15:04"
	YMDHI2 string = "200601021504"
	YMDHIS string = "2006-01-02 15:04:05"
	HI     string = "15:04"
	HI2    string = "1504"
)

// Env 用地保存当前程序运行环境类型
var Env string
