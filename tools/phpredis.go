// creator: zacyuan
// date: 2019-12-31
// 兼容thinkphp的redis使用习惯的redis组件

package tools

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	"github.com/micro/go-micro/config"
	"github.com/sirupsen/logrus"
)

// Phpredis php使用习惯的redis组件
type Phpredis struct {
	client *redis.Client
	prefix string
}

var redisClient *Phpredis

func (c *Phpredis) connect(server, password string) error {
	c.client = redis.NewClient(&redis.Options{
		Addr:     server,
		Password: password,
	})

	pong, err := c.client.Ping().Result()
	if err != nil {
		logrus.Error(pong)
		logrus.Error(err)
		return err
	}

	c.prefix = config.Get("redis", "prefix").String("")

	return nil
}

// Set redis set 方法
func (c *Phpredis) Set(key string, value interface{}, expire time.Duration) (*redis.StatusCmd, error) {
	key = c.prefix + key
	data := []interface{}{true, value}
	buf, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return c.client.Set(key, string(buf), expire), nil
}

// Get redis get 方法
func (c *Phpredis) Get(key string) *redis.StringCmd {
	key = c.prefix + key
	return c.client.Get(key)
}

// InitPhpRedis 初始化php redis
func InitPhpRedis(server, password string) error {
	if redisClient != nil {
		return nil
	}
	redisClient = &Phpredis{}
	return redisClient.connect(server, password)
}

// GetPhpRedisClient 获了php redis 客户端
func GetPhpRedisClient() *Phpredis {
	return redisClient
}
