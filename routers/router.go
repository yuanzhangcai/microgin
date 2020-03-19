// creator: zacyuan
// date: 2019-12-28

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuanzhangcai/microgin/controllers"
)

// CreateRouters 创建路由规则
func CreateRouters(router *gin.Engine) {

	baseCtl := new(controllers.BaseCtl)
	weiboCtl := new(controllers.WeiboCtl)

	// 根路由
	nnc := router.Group("/microgin")
	{
		// 设置获取版本信息接口路由
		HandleGroup(nnc, "/version", []string{"GET", "POST"}, baseCtl.Version)

		// 动态相关路由
		weibo := nnc.Group("/weibo")
		{
			HandleGroup(weibo, "/GetTop", []string{"GET", "POST"}, weiboCtl.GetTop) // 数据库查读demo
		}

		// 动态相关路由
		news := nnc.Group("/news")
		{
			HandleGroup(news, "/login", []string{"GET", "POST"}, baseCtl.Anything) // demo
		}
	}
}

// Handle 批量设置路由
func Handle(r *gin.Engine, relativePath string, httpMethods []string, handlers ...gin.HandlerFunc) {
	for _, httpMethod := range httpMethods {
		r.Handle(httpMethod, relativePath, handlers...)
	}
}

// HandleGroup 批量设置路由
func HandleGroup(r *gin.RouterGroup, relativePath string, httpMethods []string, handlers ...gin.HandlerFunc) {
	for _, httpMethod := range httpMethods {
		r.Handle(httpMethod, relativePath, handlers...)
	}
}
