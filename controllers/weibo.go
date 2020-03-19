// creator: zacyuan
// date: 2019-12-28

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yuanzhangcai/microgin/models"
)

// WeiboCtl 微博逻辑主件
type WeiboCtl struct {
	BaseCtl
	weiboModel models.WeiboModel
}

// GetTop 获取微博置顶数据
func (c *WeiboCtl) GetTop(ctx *gin.Context) {
	logrus.Debug("Received Anything API request")
	list, err := c.weiboModel.GetWeiboTop()
	if err != nil {
		ctx.JSON(200, map[string]string{
			"message": err.Error(),
		})
		c.Output(ctx, -1, err.Error())
		return
	}
	c.Output(ctx, 0, "OK", list)
}
