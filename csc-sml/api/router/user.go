package router

import (
	"api/handler/api"
	"github.com/gin-gonic/gin"
)

func UserGroup(v1 *gin.RouterGroup) {

	user := v1.Group("/video")
	{
		user.POST("/register", api.Register) //注册
		user.POST("/userList", api.UserList) //列表/查询接口排序，按时间
	}

}
