// creator: zacyuan
// date: 2019-12-28

package controllers

import (
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/yuanzhangcai/microgin/common"
	"github.com/yuanzhangcai/microgin/errors"
	"github.com/yuanzhangcai/microgin/log"
)

// ControllerInterface Controller接口定义
type ControllerInterface interface {
	Init(*gin.Context)
	Prepare() bool
	Finish()
}

// Controller 逻辑控制处理器基类组件
type Controller struct {
	ctx      *gin.Context
	PostForm *url.Values
}

// Prepare 在主逻辑处理之前的前置操作
// return true 续继后面的操作
//        false 逻辑处理提前结束
func (c *Controller) Prepare() bool {
	return true
}

// Finish 在主逻辑处理之前的收尾操作
func (c *Controller) Finish() {
}

// Init 设置Context
func (c *Controller) Init(ctx *gin.Context) {
	c.ctx = ctx
	err := c.ctx.Request.ParseForm()
	if err != nil {
		log.Panic("parse from failed")
	}
	c.PostForm = &c.ctx.Request.PostForm
}

// Anything 路由默认方法
func (c *Controller) Anything() {
	log.Debug("Received Anything API request")
	c.Output(0, "OK", map[string]string{
		"message": "Hi, this is the nnc API ",
	})
}

// Version 返回当前版本信息
func (c *Controller) Version() {
	log.Debug("Receive version API request")
	c.Output(0, "OK", common.GetVersion())
}

// Output 输入出json
func (c *Controller) Output(ret errors.ErrorCode, msg string, data ...interface{}) {
	params := make(map[string]interface{})
	params["ret"] = ret
	params["msg"] = msg
	if len(data) > 0 {
		params["data"] = data[0]
	}

	c.OutputJSON(params)
	return
}

// OutputJSON 将参数直接输出为json
func (c *Controller) OutputJSON(params interface{}) {
	c.ctx.JSON(200, params)
	c.ctx.Set("response", params)
	return
}
