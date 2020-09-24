package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(50);not null; comment:'用户名'"`
	Password string `json:"password" gorm:"type:varchar(128);not null; comment:'密码'"`
	RoleId   int    `json:"role_id" gorm:"type:int;not null; comment:'角色id'"`
	Status   int    `json:"status" gorm:"type:TINYINT(1);not null; comment:'状态：1启用 0禁用';default:1"`
}
