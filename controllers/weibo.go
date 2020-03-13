// creator: zacyuan
// date: 2019-12-28

package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yuanzhangcai/microgin/common"
	"github.com/yuanzhangcai/microgin/models"
	"github.com/yuanzhangcai/microgin/tools"
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
		common.ReturnJSON(ctx, -1, err.Error())
		return
	}
	common.ReturnJSON(ctx, 0, "OK", list)
}

// SetDemo php redis set demo
func (c *WeiboCtl) SetDemo(ctx *gin.Context) {
	phpredis := tools.GetPhpRedisClient()
	ret, err := phpredis.Set("demo", map[string]interface{}{
		"name": "zacyuan",
		"age":  18,
	}, time.Second*300)
	if err != nil {
		common.ReturnJSON(ctx, -1, err.Error())
		return
	}
	if ret.Err() != nil {
		common.ReturnJSON(ctx, -1, ret.Err().Error())
		return
	}
	common.ReturnJSON(ctx, 0, "OK")
}

// GetDemo php redis get demo
func (c *WeiboCtl) GetDemo(ctx *gin.Context) {
	phpredis := tools.GetPhpRedisClient()
	ret := phpredis.Get("demo")
	if ret.Err() != nil {
		common.ReturnJSON(ctx, -1, ret.Err().Error())
		return
	}
	common.ReturnJSON(ctx, 0, ret.Val())
}
