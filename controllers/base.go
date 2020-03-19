// creator: zacyuan
// date: 2019-12-28

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yuanzhangcai/microgin/common"
	"github.com/yuanzhangcai/microgin/errors"
)

// BaseCtl 逻辑控制处理器基类组件
type BaseCtl struct {
}

// Anything 路由默认方法
func (c *BaseCtl) Anything(ctx *gin.Context) {
	logrus.Debug("Received Anything API request")
	c.Output(ctx, 0, "OK", map[string]string{
		"message": "Hi, this is the nnc API ",
	})
}

// Version 返回当前版本信息
func (c *BaseCtl) Version(ctx *gin.Context) {
	logrus.Debug("Receive version API request")
	c.Output(ctx, 0, "OK", common.GetVersion())
}

// Output 输入出json
func (c *BaseCtl) Output(ctx *gin.Context, ret errors.ErrorCode, msg string, data ...interface{}) {
	params := make(map[string]interface{})
	params["iRet"] = ret
	params["sMsg"] = msg
	if len(data) > 0 {
		params["jData"] = data[0]
	}

	c.OutputData(ctx, params)
	return
}

// OutputData 奖参数直接输出为json
func (c *BaseCtl) OutputData(ctx *gin.Context, params interface{}) {
	ctx.JSON(200, params)
	ctx.Set("response", params)
	return
}
