package service

import (
	"context"
	__ "grpc/basic/proto"
	"grpc/dao"
	"grpc/handler/model"
)

type Server struct {
	__.UnimplementedUserServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) Register(_ context.Context, in *__.RegisterReq) (*__.RegisterResp, error) {

	var user model.User

	// 校验手机号是否存在
	if !dao.GetOneByFields(&model.User{Mobile: in.Mobile}, &user) {
		return &__.RegisterResp{
			Code: 400,
			Msg:  "查询失败",
		}, nil
	}
	if user.Id != 0 {
		return &__.RegisterResp{
			Code: 500,
			Msg:  "手机号已存在",
		}, nil
	}

	user = model.User{
		Name:     in.Mobile,
		Mobile:   in.Mobile,
		Password: in.Password,
	}
	if !dao.Create(&user) {
		return &__.RegisterResp{
			Code: 500,
			Msg:  "注册失败",
		}, nil
	}

	return &__.RegisterResp{
		Code: 200,
		Msg:  "注册成功",
	}, nil
}
func (s *Server) UserList(_ context.Context, in *__.UserListReq) (*__.UserListResp, error) {

	var user model.User

	list, _ := user.UserList()

	var userList []*__.UserList
	for _, l := range list {
		userList = append(userList, &__.UserList{
			Id:         int64(l.Id),
			Name:       l.Name,
			Mobile:     l.Mobile,
			Status:     int64(l.Status),
			CreateTime: l.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: l.UpdateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &__.UserListResp{
		List: userList,
	}, nil
}
