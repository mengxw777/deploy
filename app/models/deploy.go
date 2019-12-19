package models

import "time"

const UNDEPLOYED = 0
const DEPLOYING = 1
const SUCCESSFULLY_DEPLOYED = 2
const DEPLOYMENT_FAILED = 3
const ABANDONED = 4

type Deploy struct {
	Id        int       `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"type:varchar(50) not null comment '发布单名称'" form:"name" json:"name" binding:"required"`
	ProjectId int       `gorm:"type:int not null comment '项目ID'" form:"project_id" json:"project_id" binding:"required,exists=projects:id"`
	GroupId   int       `gorm:"type:int not null comment '集群ID'" form:"group_id" json:"group_id" binding:"required,exists=server_groups:id"`
	Status    int       `gorm:"type:tinyint not null comment '状态 0 未部署 1 正在部署 2 部署成功 3 部署失败 4 废弃'" form:"status" json:"status"`
	UserId    int       `gorm:"type:int not null comment '申请人ID'" form:"user_id" json:"user_id"`
	CreatedAt time.Time `gorm:"type:timestamp not null" json:"created_at"`
}

func (d *Deploy) IsAbandoned() bool {
	if d.Status == ABANDONED {
		return true
	}
	return false
}
