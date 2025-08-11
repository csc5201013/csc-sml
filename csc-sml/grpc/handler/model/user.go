package model

import (
	"grpc/basic/global"
	"time"
)

type User struct {
	Id         uint64    `gorm:"column:id;type:bigint UNSIGNED;comment:用户ID(主键) 自增;primaryKey;not null;" json:"id"`                             // 用户ID(主键) 自增
	Name       string    `gorm:"column:name;type:varchar(32);comment:姓名 非空;not null;" json:"name"`                                                // 姓名 非空
	Mobile     string    `gorm:"column:mobile;type:varchar(11);comment:手机号 非空，唯一;not null;" json:"mobile"`                                     // 手机号 非空，唯一
	Password   string    `gorm:"column:password;type:varchar(128);comment:登录密码 ;not null;" json:"password"`                                       // 登录密码
	Status     uint32    `gorm:"column:status;type:int UNSIGNED;comment:账号状态 0=待审批；1=正常；2=驳回；3=冻结;not null;default:1;" json:"status"`    // 账号状态 0=待审批；1=正常；2=驳回；3=冻结
	CreateTime time.Time `gorm:"column:create_time;type:datetime(3);comment:创建时间 非空;not null;default:CURRENT_TIMESTAMP(3);" json:"create_time"` // 创建时间 非空
	UpdateTime time.Time `gorm:"column:update_time;type:datetime(3);comment:更新时间 非空;not null;default:CURRENT_TIMESTAMP(3);" json:"update_time"` // 更新时间 非空
}

func (u *User) TableName() string {

	return "user"
}

func (u *User) UserList() ([]*User, error) {
	var user []*User
	err := global.DB.Order("create_time desc").Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
