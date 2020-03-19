// creator: zacyuan
// date: 2019-12-28

package models

import "github.com/yuanzhangcai/microgin/log"

// OWeibo 微博数据结构体
type OWeibo struct {
	ID            int64
	CommentCount  int64
	Content       string
	CreateTime    int64
	CrowdID       int64
	Data          string
	From          string
	GeolocationID int64
	IsCrowdTop    int8
	IsTop         int8
	Pos           string
	ReplyTime     int64
	RepostCount   int64
	Status        int64
	Type          string
	UID           int64
}

// WeiboModel 微博表组件
type WeiboModel struct {
	Model
}

// GetWeiboTop 获取微博置顶数据
func (c *WeiboModel) GetWeiboTop() ([]OWeibo, error) {
	var list []OWeibo
	err := db.Limit(10).Find(&list).Error
	if err != nil {
		log.Error("SQL执行失败。")
		return list, err
	}
	return list, nil
}
