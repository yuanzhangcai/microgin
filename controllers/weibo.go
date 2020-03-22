// creator: zacyuan
// date: 2019-12-28

package controllers

import (
	"fmt"

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

	name := c.PostForm.Get("name")
	for k, v := range *c.PostForm {
		fmt.Printf("k:%v", k)
		fmt.Printf("v:%v\n", v)

		fmt.Printf("k:%T", k)
		fmt.Printf("v:%T\n", v)
	}

	c.Output(0, "OK", name)
}
