package request

type RegisterReq struct {
	Name     string `form:"name" json:"name"`
	Mobile   string `form:"mobile" json:"mobile"`
	Password string `form:"password" json:"password" `
}

type UserListsReq struct {
}
