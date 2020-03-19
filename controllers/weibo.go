// creator: zacyuan
// date: 2019-12-28

package controllers

import (
	"github.com/yuanzhangcai/microgin/log"
	"github.com/yuanzhangcai/microgin/models"
)

// WeiboCtl 微博逻辑主件
type WeiboCtl struct {
	Controller
	weiboModel models.WeiboModel
}

// GetTop 获取微博置顶数据
func (c *WeiboCtl) GetTop() {
	log.Debug("Received Anything API request")

	name := c.ctx.Query("name")

	c.Output(0, "OK", name)
}
