// creator: zacyuan
// date: 2019-12-30
// 登录中间简

package middleware

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/config"
	"github.com/sirupsen/logrus"
	"github.com/yuanzhangcai/microgin/controllers"
	"github.com/yuanzhangcai/microgin/errors"
)

const (
	sessionKey   = "PHPSESSID"
	loginUserKey = "opensns_m_OX_LOGGED_USER"
)

// Auth 生成登录中间件
func Auth() func(*gin.Context) {
	baseCtl := controllers.BaseCtl{}
	checkURL := config.Get("login", "check_url").String("")

	return func(ctx *gin.Context) {

		client := &http.Client{}
		reqest, err := http.NewRequest("GET", checkURL, nil)
		if err != nil {
			ctx.Abort()
			logrus.Error(err.Error())
			baseCtl.Output(ctx, errors.ErrorCheckLoginCreateRequestFailed, "登录验证，创建请求失败")
			return
		}

		// 获取当前登录session ID
		sessionID, _ := ctx.Cookie(sessionKey)
		if sessionID != "" {
			cookie := &http.Cookie{Name: sessionKey, Value: sessionID, HttpOnly: true}
			reqest.AddCookie(cookie)
		}

		// 获取当前登录用户ID
		userID, _ := ctx.Cookie(loginUserKey)
		if userID != "" {
			cookie := &http.Cookie{Name: loginUserKey, Value: userID, HttpOnly: true}
			reqest.AddCookie(cookie)
		}

		response, err := client.Do(reqest)
		if err != nil {
			ctx.Abort()
			logrus.Error(err.Error())
			baseCtl.Output(ctx, errors.ErrorCheckLoginSendRequestFailed, "登录验证，发送请求失败")
			return
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			ctx.Abort()
			logrus.Error(err.Error())
			baseCtl.Output(ctx, errors.ErrorCheckLoginReadResponseFailed, "登录验证，读取登录验证返回数据失败。")
			return
		}

		type checkLoginRet struct {
			ErrorCode int    `json:"error_code"`
			Info      string `json:"info"`
		}

		ret := checkLoginRet{}
		err = json.Unmarshal(body, &ret)
		if err != nil {
			ctx.Abort()
			logrus.Error("获取session id 失败。")
			baseCtl.Output(ctx, errors.ErrorCheckLoginUnmarshalJSONailed, "登录验证，解析登录验证返回数据失败。")
			return
		}

		if ret.ErrorCode == 1 {
			ctx.Next()
		} else {
			ctx.Abort()
			logrus.Error("当前没有登录")
			baseCtl.Output(ctx, errors.ErrorNoLogin, "当前没有登录")
			return
		}
	}
}
