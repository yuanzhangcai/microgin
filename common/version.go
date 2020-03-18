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

var (
	// Version 程序版本号
	Version string
	// Env 程序运行环境
	Env string = ProdEnv
	// Commit 最后一次提交的id
	Commit string
	// BuildTime 编译时间
	BuildTime string
	// BuildUser 编译人
	BuildUser string
	// GoVersion go编译版本
	GoVersion string
)

// GetVersion 获取版本信息
func GetVersion() map[string]string {
	return map[string]string{
		"version":    Version,   // 程序版本号
		"env":        Env,       // 程序运行环境
		"commit":     Commit,    // 最后一次提交的id
		"build_time": BuildTime, // 编译时间
		"build_user": BuildUser, // 编译人
		"go_version": GoVersion, // 版本号
	}
}
