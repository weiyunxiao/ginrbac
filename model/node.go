package model

import (
	"github.com/jinzhu/gorm"
)

type Node struct {
	gorm.Model
	Name   string `json:"name" gorm:"type:VARCHAR(50);not null;comment:'节点名称'"`
	Path   string `json:"path" gorm:"type:VARCHAR(50);not null;comment:'节点路径'"`
	Pid    int    `json:"pid" gorm:"type:INT(11);not null;comment:'所属节点id'"`
	Sort   int    `json:"sort" gorm:"type:INT(4);not null;default:0;comment:'排序'"`
	Icon   string `json:"icon" gorm:"type:VARCHAR(50);default:'';comment:'图标'"`
	IsMenu int    `json:"is_menu" gorm:"type:TINYINT(1);not null;default:1;comment:'是否是菜单项 1 不是 2 是'"`
	Status int    `json:"status" gorm:"type:TINYINT(1);not null;default:1;comment:'状态：1 启用 0 禁用'"`
}
