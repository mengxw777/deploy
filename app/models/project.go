package models

import "time"

type Project struct {
	Id              int         `gorm:"primary_key"  json:"id"`
	Name            string      `gorm:"type:varchar(50) not null comment '项目名称'" form:"name" json:"name" binding:"required"`
	Description     string      `gorm:"type:text comment '项目描述'" form:"description" json:"description"`
	GitAddr         string      `gorm:"type:text not null comment '仓库地址'" form:"git_addr" json:"git_addr" binding:"required"`
	DeployMode      int         `gorm:"type:tinyint(1) default 1 not null comment '上线模式 1 分支 2 标签'" form:"deploy_mode" json:"deploy_mode" binding:"required,numeric,min=1,max=2"`
	ServerGroupId   int         `gorm:"type:int not null comment '集群ID'" form:"group_id" json:"group_id" binding:"required,numeric,exists=server_groups:id"`
	DeployUser      string      `gorm:"type:varchar(50) not null comment '部署用户'" form:"deploy_user" json:"deploy_user" binding:"required,max=50"`
	DeployPath      string      `gorm:"type:text comment '部署目录'" form:"deploy_path" json:"deploy_path" binding:"required" `
	DeployBeforeCmd string      `gorm:"type:text comment '部署前服务器上执行命令'" form:"deploy_before_cmd" json:"deploy_before_cmd"`
	DeployAfterCmd  string      `gorm:"type:text comment '部署后服务器上执行命令'" form:"deploy_after_cmd" json:"deploy_after_cmd"`
	Branch          string      `gorm:"type:varchar(50) default '' comment '分支名称，在部署模式为1时存在'" form:"branch" json:"branch" binding:"required_if=DeployMode:1"`
	BuildCmd        string      `gorm:"type:text comment '构建命令'"  form:"build_cmd" json:"build_cmd"`
	CreatedAt       time.Time   `gorm:"type:timestamp not null" json:"created_at"`
	ServerGroup     ServerGroup `gorm:"save_associations:false:false;association_save_reference:false" json:"server_group" binding:"-"`
}
