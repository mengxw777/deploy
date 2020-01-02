package server

import (
	"deploy/app/models"
	"deploy/database"
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
	GroupId int    `form:"group_id"  binding:"omitempty,numeric,exists=server_groups:id"`
}

func Index(c *gin.Context) {
	var servers []models.Server
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	builder := db.Set("gorm:auto_preload", true).
		Model(&servers).Order("created_at desc")

	paginator := pagination.Paging(&pagination.Param{
		DB:   builder,
		Page: page,
	}, &servers)

	render.Page(c, *paginator)
}

func Show(c *gin.Context) {
	var server models.Server
	serverId := c.Param("id")

	if err := db.Set("gorm:auto_preload", true).Model(&server).First(&server, serverId).Error; err != nil {
		render.Fail(c, err.Error())
		return
	}

	render.Data(c, server)
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
	server.Ip = serverForm.Ip
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
	server.Ip = serverForm.Ip
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
