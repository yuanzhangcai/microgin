package tools

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

// Redis redis组件
type Redis struct {
	*redis.Client
	prefix string
}

// GenerateScoreKey 生成积分redis的key
func (c *Redis) GenerateScoreKey(scoreID, nid uint64) string {
	return c.prefix + strconv.FormatUint(scoreID, 10) + "_" + strconv.FormatUint(nid, 10)
}

// ScoreSet 写入积分redis
func (c *Redis) ScoreSet(scoreID, nid uint64, value interface{}, expire time.Duration) error {
	key := c.GenerateScoreKey(scoreID, nid)
	buf, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.Set(key, string(buf), expire).Err()
}

// ScoreGet 获取积分redis
func (c *Redis) ScoreGet(scoreID, nid uint64, value interface{}) error {
	key := c.GenerateScoreKey(scoreID, nid)
	ret := c.Get(key)
	if ret.Err() != nil {
		return ret.Err()
	}

	err := json.Unmarshal([]byte(ret.Val()), value)
	if err == nil {
		return err
	}

	return nil
}

// ScoreDel 删除积分redis
func (c *Redis) ScoreDel(scoreID, nid uint64) *redis.IntCmd {
	key := c.GenerateScoreKey(scoreID, nid)
	return c.Del(key)
}

var client *Redis

// InitRedis 初始化redis
func InitRedis(server, password, prefix string) error {
	if client != nil {
		return nil
	}

	if server == "" {
		return errors.New("redis服务器地址为空。")
	}

	client = &Redis{
		Client: redis.NewClient(&redis.Options{
			Addr:     server,
			Password: password,
		}),
		prefix: prefix,
	}

	pong, err := client.Ping().Result()
	if err != nil {
		logrus.Error(pong)
		logrus.Error(err)
		return err
	}
	return nil
}

// GetRedis 获取redis实例
func GetRedis() *Redis {
	return client
}
