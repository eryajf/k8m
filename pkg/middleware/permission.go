package middleware

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/weibaohui/k8m/pkg/comm/utils/amis"
	"github.com/weibaohui/k8m/pkg/models"
)

func RolePlatformOnly(handler interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, role := amis.GetLoginUser(c)
		if role == "" {
			role = "guest"
		}

		// 通过反射获取方法名
		handlerValue := reflect.ValueOf(handler)
		// handlerType := handlerValue.Type()

		// 获取 struct tag
		// requiredRole := handlerType.Name()

		// 权限检查
		if role != models.RolePlatformAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access Denied for your role"})
			c.Abort()
			return
		}

		// 继续执行请求处理
		handlerValue.Call([]reflect.Value{reflect.ValueOf(c)})
	}
}
