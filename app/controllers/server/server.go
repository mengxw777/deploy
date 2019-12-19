package server

import (
	"deploy/app/models"
	"deploy/database"
	"deploy/helper"
	"deploy/helper/pagination"
	"deploy/helper/render"
	"github.com/gin-gonic/gin"
	"strconv"
)

var db = database.DB

type serverForm struct {
	Name    string `form:"name"  binding:"required,min=1,max=50"`
	Ip      string `form:"ip"  binding:"required,ipv4"`
	Port    int    `form:"port"  binding:"required,numeric"`
	GroupId int    `form:"group_id"  binding:"numeric,exists=server_groups:id"`
}

func Index(c *gin.Context) {
	var servers []models.Server
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	paginator := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    page,
		OrderBy: []string{"created_at desc"},
	}, &servers)

	render.Page(c, *paginator)
}

func Store(c *gin.Context) {
	var serverForm serverForm
	var server models.Server

	if err := c.ShouldBind(&serverForm); err != nil {
		render.Fail(c, render.NewValidatorError(err))
		return
	}

	if serverForm.GroupId != 0 {
		server.GroupId = serverForm.GroupId
	}
	server.Name = serverForm.Name
	server.Ip = helper.Ip2int(serverForm.Ip)
	server.Port = serverForm.Port

	if err := db.Create(&server).Error; err != nil {
		render.Fail(c, err.Error())
		return
	}

	render.Success(c, "添加成功")
}

func Update(c *gin.Context) {
	var serverForm serverForm
	var server models.Server

	serverId := c.Param("id")

	if err := c.ShouldBind(&serverForm); err != nil {
		render.Fail(c, render.NewValidatorError(err))
		return
	}

	if serverForm.GroupId != 0 {
		server.GroupId = serverForm.GroupId
	}
	server.Name = serverForm.Name
	server.Ip = helper.Ip2int(serverForm.Ip)
	server.Port = serverForm.Port

	if err := db.Model(&server).Where("id = ?", serverId).Update(&server).Error; err != nil {
		render.Fail(c, "更新失败")
		return
	}

	render.Success(c, "更新成功")
}

func Destroy(c *gin.Context) {
	var server models.Server
	serverId := c.Param("id")

	if err := db.Where("id = ?", serverId).Delete(&server).Error; err != nil {
		render.Fail(c, "删除失败")
		return
	}

	render.Success(c, "删除成功")
}
