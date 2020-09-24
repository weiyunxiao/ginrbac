package model

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Name   string `json:"name" gorm:"type:VARCHAR(30);not null;comment:'角色名称'"`
	Rules  string `json:"rules" gorm:"type:VARCHAR(255);not null;comment:'角色拥有的权限节点'"`
	Status int    `json:"status" gorm:"type:TINYINT(1);not null;default:1;comment:'状态：1 启用 0 禁用'"`
}
