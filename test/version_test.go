package test_test

import (
	"encoding/json"
	"flag"
	"testing"

	"github.com/yuanzhangcai/microgin/common"
)

var server string = "http://127.0.0.1:8080"

func init() {
	var _ = func() bool {
		testing.Init()
		return true
	}()

	//解析命令行参数
	flag.StringVar(&server, "server", "http://127.0.0.1:8080", "请求域名")
	flag.Parse()
}

type TestingFun func(...interface{})

func send(t interface{}, url, data string) {
	var Fatal TestingFun
	//var Log TestingFun
	switch t.(type) {
	case *testing.T:
		Fatal = t.(*testing.T).Fatal
		// Log = t.(*testing.T).Log
	case *testing.B:
		Fatal = t.(*testing.B).Fatal
		// Log = t.(*testing.B).Log
	}

	params := common.HTTPParam{
		URL:  url,
		Data: data,
	}

	result, statusCode, err := common.HTTP(params)
	if err != nil {
		Fatal(err.Error())
	}
	if statusCode != 200 {
		Fatal("status code is not 200")
	}
	// Log(result)
	response := struct {
		Msg string `json:"msg"`
		Ret int    `json:"ret"`
	}{}

	err = json.Unmarshal([]byte(result), &response)
	if err != nil {
		Fatal("Unmarshal failed")
	}

	if response.Ret != 0 {
		Fatal("response ret is not 0")
	}
}

func TestVersion(t *testing.T) {
	send(t, server+"/microgin/version", "")
	t.Log("Success")
}

// go test -benchmem -run version_test.go -bench ^BenchmarkVersion$
func BenchmarkVersion(b *testing.B) {
	b.Run("gintest", func(b *testing.B) {
		// for i := 0; i < b.N; i++ {
		// 	send(b, "http://127.0.0.1:11000/microgin/version", "")
		// }

		// 并发测试
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				send(b, "http://127.0.0.1:11000/microgin/version", "")
			}
		})
	})

	b.Run("beegotest", func(b *testing.B) {
		// for i := 0; i < b.N; i++ {
		// 	send(b, "http://127.0.0.1:11001/beegotest", "")
		// }

		// 并发测试
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				send(b, "http://127.0.0.1:11001/beegotest", "")
			}
		})
	})

	b.Run("microgin", func(b *testing.B) {
		// for i := 0; i < b.N; i++ {
		// 	send(b, server+"/microgin/version", "")
		// }

		// 并发测试
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				send(b, server+"/microgin/version", "")
			}
		})
	})
}
