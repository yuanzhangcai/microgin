// creator: zacyuan
// date: 2019-12-28

package routers

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/yuanzhangcai/microgin/controllers"
)

// CreateRouters 创建路由规则
func CreateRouters(router *gin.Engine) {
	// 根路由
	nnc := router.Group("/microgin")
	{
		// 设置获取版本信息接口路由
		HandleGroup(nnc, "/version", []string{"GET", "POST"}, HadleMain(&controllers.Controller{}, "Version"))

		// 动态相关路由
		weibo := nnc.Group("/weibo")
		{
			HandleGroup(weibo, "/GetTop", []string{"GET", "POST"}, HadleMain(&controllers.WeiboCtl{}, "GetTop")) // 数据库查读demo
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

// HadleMain 主要处理逻构造方法
func HadleMain(ctl interface{}, method string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		// 初始一个新的处理器
		temp := initialize(ctl)
		if temp == nil {
			panic("controller is not ControllerInterface")
		}

		// 设置请求上下文
		temp.Init(ctx)

		// 逻辑提前结束
		if !temp.Prepare() {
			return
		}

		// 主处理逻辑
		value := reflect.ValueOf(temp)
		main := value.MethodByName(method)
		if main.IsValid() {
			main.Call(nil)
		} else {
			panic(method + " is not exist.")
		}

		// 最后收尾工作
		temp.Finish()
	}
}

func initialize(c interface{}) controllers.ControllerInterface {

	reflectVal := reflect.ValueOf(c)
	t := reflect.Indirect(reflectVal).Type()

	vc := reflect.New(t)
	execController, ok := vc.Interface().(controllers.ControllerInterface)
	if !ok {
		return nil
	}

	// 对象复制
	// elemVal := reflect.ValueOf(c).Elem()
	// elemType := reflect.TypeOf(c).Elem()
	// execElem := reflect.ValueOf(execController).Elem()

	// numOfFields := elemVal.NumField()
	// for i := 0; i < numOfFields; i++ {
	// 	fieldType := elemType.Field(i)
	// 	elemField := execElem.FieldByName(fieldType.Name)
	// 	if elemField.CanSet() {
	// 		fieldVal := elemVal.Field(i)
	// 		elemField.Set(fieldVal)
	// 	}
	// }

	return execController
}
