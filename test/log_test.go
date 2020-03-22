package test_test

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/yuanzhangcai/microgin/log"
)

func init() {
	log.InitLog("/Users/zacyuan/logs/microgin/", "test", 5, 7)
}

// go test -run ^TestLogrus$   #正则表达式匹配，所以要加:^ $
func TestLogrus(t *testing.T) {
	logrus.Debug("test log")
	fmt.Println("logrus finished")
}

// go test -run ^TestLog$
func TestLog(t *testing.T) {
	log.Debug("test log")
	fmt.Println("log finished")
}

func TestLogAll(t *testing.T) {
	t.Run("logrus", func(t *testing.T) {
		logrus.Debug("run test log")
		fmt.Println("run logrus finished")
	})

	// go test -run TestLogAll/log$
	t.Run("log", func(t *testing.T) {
		logrus.Debug("run test log")
		fmt.Println("run log finished")
	})
}

// go test -bench=BenchmarkLogrus -benchtime=20s
func BenchmarkLogrus(b *testing.B) {
	// for i := 0; i < b.N; i++ {
	// 	logrus.Debug("test log")
	// }

	// 内存统计
	b.ReportAllocs()

	// 并发测试
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logrus.Debug("test log")
		}
	})
}

// go test -bench=BenchmarkLog -benchtime=20s
func BenchmarkLog(b *testing.B) {
	// for i := 0; i < b.N; i++ {
	// 	log.Debug("test log")
	// }

	// 内存统计
	b.ReportAllocs()

	// 并发测试
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Debug("test log")
		}
	})
}

// go test -bench=BenchmarkLogAll -benchtime=20s
func BenchmarkLogAll(b *testing.B) {
	b.Run("logrus", func(b *testing.B) {
		// 内存统计
		b.ReportAllocs()

		// 并发测试
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				logrus.Debug("test log")
			}
		})
	})

	b.Run("log", func(b *testing.B) {
		// 内存统计
		b.ReportAllocs()

		// 并发测试
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				log.Debug("test log")
			}
		})
	})
}
