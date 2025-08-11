package api

import (
	"api/basic/global"
	__ "api/basic/proto"
	"api/handler/request"

	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Register(c *gin.Context) {

	var req request.RegisterReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "传参失败",
			"data": err.Error(),
		})
		return
	}
	r, err := global.UserClient.Register(context.Background(), &__.RegisterReq{
		Name:     req.Name,
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": r.Code,
		"msg":  r.Msg,
	})

}

func UserList(c *gin.Context) {
	var req request.UserListsReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 10000,
			"msg":  "传参失败",
			"data": err.Error(),
		})
		return
	}
	var r, err = global.UserClient.UserList(context.Background(), &__.UserListReq{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "展示成功",
		"data": r.List,
	})
}
