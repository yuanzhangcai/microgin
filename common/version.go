// creator: zacyuan
// date: 2019-12-28

package common

var verInfo map[string]string

// SetVersion 设置版本信息
func SetVersion(version map[string]string) {
	verInfo = version
}

// GetVersion 获取版本信息
func GetVersion() map[string]string {
	return verInfo
}
