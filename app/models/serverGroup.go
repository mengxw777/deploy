package models

import (
	"time"
)

type ServerGroup struct {
	Id        int       `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"type:varchar(50) not null comment '集群名称'" form:"name" json:"name"`
	CreatedAt time.Time `gorm:"type:timestamp not null" json:"created_at"`
}
