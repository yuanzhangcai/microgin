// creator: zacyuan
// date: 2019-12-28

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yuanzhangcai/microgin/common"
)

// BaseCtl 逻辑控制处理器基类组件
type BaseCtl struct {
}

// Anything 路由默认方法
func (c *BaseCtl) Anything(ctx *gin.Context) {
	logrus.Debug("Received Anything API request")
	common.ReturnJSON(ctx, 0, "OK", map[string]string{
		"message": "Hi, this is the nnc API ",
	})
}

// Version 返回当前版本信息
func (c *BaseCtl) Version(ctx *gin.Context) {
	logrus.Debug("Receive version API request")
	common.ReturnJSON(ctx, 0, "OK", common.GetVersion())
}
