// creator: zacyuan
// date: 2019-12-28

package common

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/yuanzhangcai/microgin/log"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// TimeToStr 时间戳转日期
func TimeToStr(fmt string, value interface{}) string {
	str := ""
	var sec int64
	switch value.(type) {
	case int:
		sec = int64(value.(int))
	case int64:
		sec = value.(int64)
	case string:
		sec, _ = strconv.ParseInt(value.(string), 10, 64)
	}

	str = time.Unix(sec, 0).Format(fmt)
	return str
}

// StrToTime 日期转时间戳
func StrToTime(fmt string, value string) int64 {
	tm, _ := time.ParseInLocation(fmt, value, time.Local)
	return tm.Unix()
}

// ParseInt64 字符串转int64
func ParseInt64(str string) int64 {
	value, _ := strconv.ParseInt(str, 10, 64)
	return value
}

// Md5Str 计算md5，返回字符串
func Md5Str(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Md5Byte 计算md5，返回字节
func Md5Byte(str string) []byte {
	h := md5.New()
	h.Write([]byte(str))
	return h.Sum(nil)
}

// Decimal 保留几位小数
func Decimal(value float64, num int) float64 {
	format := "%." + strconv.Itoa(num) + "f"
	value, _ = strconv.ParseFloat(fmt.Sprintf(format, value), 64)
	return value
}

// GetRandomString 生成随机字符串
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// GbkToUtf8 GBK转UTF8
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Utf8ToGbk UTF8转GBK
func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CloneObject 深度clone对像
func CloneObject(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = CloneObject(v)
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = CloneObject(v)
		}

		return newSlice
	}

	return value
}

// GetFileNameWithoutSuffix 获取不带后缀文件名
func GetFileNameWithoutSuffix(fullFilename string) string {
	var filenameWithSuffix string
	filenameWithSuffix = path.Base(fullFilename) //获取文件名带后缀

	var fileSuffix string
	fileSuffix = path.Ext(filenameWithSuffix) //获取文件后缀

	var filenameOnly string
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix) //获取文件名
	return filenameOnly
}

// GetRunInfo 获取程序运行信息
func GetRunInfo() {
	ex, err := os.Executable()
	if err != nil {
		log.Error("获取当前程序执行目录失败。")
		os.Exit(-1)
	}

	// 获取当前程序运行文件目录与文件名
	CurrRunPath = filepath.Dir(ex)
	CurrRunFileName = GetFileNameWithoutSuffix(ex)
}

// ToString interface转string
func ToString(value interface{}) string {
	str := ""
	switch value.(type) {
	case string:
		str = value.(string)
	case int:
		intVal := value.(int)
		str = strconv.Itoa(intVal)
	case int64:
		str = strconv.FormatInt(value.(int64), 10)
	case float64:
		str = strconv.FormatFloat(value.(float64), 'f', -1, 64)
	case error:
		str = value.(error).Error()
	}
	return str
}

var transport *http.Transport // 全局变变更，用于保存长链接缓存。

// HTTPParam HTTP请求参数
type HTTPParam struct {
	Method   string                  // http请求方法，POST/GET
	URL      string                  // 请求URL
	Data     string                  // 请求数据
	Headers  *map[string]interface{} // header
	Cookies  *map[string]interface{} // cookie
	UseShort bool                    //使用短链接
}

// HTTP 发送http请求
func HTTP(params HTTPParam) (string, int, error) {
	var t *http.Transport
	if params.UseShort {
		t = &http.Transport{
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives: true,
		}
	} else {
		if transport == nil {
			transport = &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
		}
		t = transport
	}

	body := bytes.NewBuffer([]byte(params.Data))
	client := &http.Client{Transport: t}
	reqest, err := http.NewRequest(params.Method, params.URL, body)
	if err != nil {
		return "", 0, err
	}

	// 设置header
	if params.Headers != nil {
		for key, value := range *params.Headers {
			strValue := ToString(value)
			reqest.Header.Set(key, strValue)
		}
	}

	// 设置cookie
	if params.Cookies != nil {
		for key, value := range *params.Cookies {
			reqest.AddCookie(&http.Cookie{Name: key, Value: ToString(value), HttpOnly: true})
		}
	}

	response, err := client.Do(reqest)
	if err != nil {
		return "", 0, err
	}
	defer response.Body.Close()

	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", 0, err
	}
	return string(resBody), response.StatusCode, nil
}
