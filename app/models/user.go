package models

import "time"

type User struct {
	Id        int       `gorm:"primary_key" json:"id"`
	Account   string    `gorm:"type:varchar(50) not null comment '账号'" form:"account" json:"account" binding:"required"`
	Password  string    `gorm:"type:varchar(255) not null comment '密码'" form:"password" json:"password" binding:"required"`
	CreatedAt time.Time `gorm:"type:timestamp not null" json:"created_at"`
}
