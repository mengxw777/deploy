package models

import (
	"time"
)

type Server struct {
	Id        int       `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"type:varchar(50) not null comment '服务器名称'" json:"name"`
	Ip        uint64    `gorm:"type:bigint unsigned not null comment '服务器登录IP'" json:"ip"`
	Port      int       `gorm:"type:tinyint not null comment '服务器登录端口'" json:"port"`
	GroupId   int       `gorm:"type:int default 0 not null comment '集群ID'" json:"group_id"`
	CreatedAt time.Time `gorm:"type:timestamp not null" json:"created_at"`
}
