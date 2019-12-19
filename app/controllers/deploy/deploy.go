package deploy

import (
	"deploy/app/models"
	"deploy/database"
	"deploy/helper/auth"
	"deploy/helper/pagination"
	"deploy/helper/render"
	"github.com/gin-gonic/gin"
	"strconv"
)

var db = database.DB

func Index(c *gin.Context) {
	var deploy []models.Deploy
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	paginator := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    page,
		OrderBy: []string{"created_at desc"},
	}, &deploy)

	render.Page(c, *paginator)
}

func Store(c *gin.Context) {
	var deploy models.Deploy

	if err := c.ShouldBind(&deploy); err != nil {
		render.Fail(c, render.NewValidatorError(err))
		return
	}

	deploy.Status = models.UNDEPLOYED
	deploy.UserId = auth.GetUserId(c)

	if err := db.Create(&deploy).Error; err != nil {
		render.Fail(c, err.Error())
		return
	}

	render.Success(c, "添加成功")

}

func Build(c *gin.Context) {
	var deploy models.Deploy
	var servers []models.Server
	var project models.Project
	deployId := c.Param("id")

	if err := db.Where("id = ?", deployId).Find(&deploy).Error; err != nil {
		render.Fail(c, err.Error())
		return
	}

	if deploy.IsAbandoned() {
		render.Fail(c, "申请单已废弃")
		return
	}

	/**
	TODO
	1 查询项目信息
	2 查询服务器信息
	3 生成脚本
	*/

	if err := db.Where("id = ?", deploy.ProjectId).Find(&project).Error; err != nil {
		render.Fail(c, err.Error())
		return
	}

	if err := db.Where("group_id = ?", deploy.GroupId).Find(&servers).Error; err != nil {
		render.Fail(c, err.Error())
		return
	}
}

func Deploy(c *gin.Context) {

}

// 	废弃申请单
func Destroy(c *gin.Context) {
	var deploy models.Deploy
	deployId := c.Param("id")

	if err := db.Model(&deploy).Where("id = ?", deployId).Update("status", models.ABANDONED).Error; err != nil {
		render.Fail(c, "废弃失败")
		return
	}

	render.Success(c, "已废弃")
}
