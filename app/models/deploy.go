package models

import (
	"time"
)

const Undeployed = 0
const Deploying = 1
const SuccessfullyDeployed = 2
const DeploymentFailed = 3
const Abandoned = 4

type Deploy struct {
	Id           int       `gorm:"primary_key" json:"id"`
	Name         string    `gorm:"type:varchar(50) not null comment '发布单名称'" form:"name" json:"name" binding:"required"`
	ProjectId    int       `gorm:"type:int not null comment '项目ID'" form:"project_id" json:"project_id" binding:"required,exists=projects:id"`
	DeployStatus int       `gorm:"type:tinyint not null comment '部署状态 0 未部署 1 正在部署 2 部署成功 3 部署失败'" json:"deploy_status"`
	DeployLog    JSON      `gorm:"type:json comment '部署日志'" json:"deploy_log"`
	BuildStatus  int       `gorm:"type:tinyint not null comment '状态 0 未构建 1 正在构建 2 构建成功 3 构建失败 4 废弃'" json:"build_status"`
	BuildLog     JSON      `gorm:"type:json comment '构建日志'" json:"build_log"`
	BuildPath    string    `gorm:"type:varchar(255)" json:"build_path"`
	UserId       int       `gorm:"type:int not null comment '申请人ID'" form:"user_id" json:"user_id"`
	CreatedAt    time.Time `gorm:"type:timestamp not null" json:"created_at"`
	User         User      `gorm:"save_associations:false:false;association_save_reference:false" json:"user" binding:"-"`
	Project      Project   `gorm:"save_associations:false:false;association_save_reference:false" json:"project" binding:"-"`
}

func (d *Deploy) IsAbandoned() bool {
	if d.BuildStatus == Abandoned {
		return true
	}
	return false
}

func (d *Deploy) IsSuccessfullyDeployed() bool {
	if d.DeployStatus == SuccessfullyDeployed {
		return true
	}
	return false
}
