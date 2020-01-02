package serverGroup

import (
	"deploy/app/models"
	"deploy/database"
	"deploy/helper/pagination"
	"deploy/helper/render"
	"github.com/gin-gonic/gin"
	"strconv"
)

var db = database.DB

func Index(c *gin.Context) {
	var ServerGroup []models.ServerGroup
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	_, brief := c.GetQuery("brief")

	if brief {
		db.Model(ServerGroup).Order("created_at desc").Find(&ServerGroup)
		render.Data(c, ServerGroup)
	} else {
		builder := db.Set("gorm:auto_preload", true).
			Model(&ServerGroup).Order("created_at desc")

		paginator := pagination.Paging(&pagination.Param{
			DB:   builder,
			Page: page,
		}, &ServerGroup)

		render.Page(c, *paginator)
	}
}

func Store(c *gin.Context) {
	var ServerGroup models.ServerGroup

	if err := c.ShouldBind(&ServerGroup); err != nil {
		render.Fail(c, render.NewValidatorError(err))
		return
	}

	if err := db.Create(&ServerGroup).Error; err != nil {
		render.Fail(c, err.Error())
		return
	}

	render.Success(c, "添加成功")
}

func Update(c *gin.Context) {
	var ServerGroup models.ServerGroup

	groupId := c.Param("id")

	if err := c.ShouldBind(&ServerGroup); err != nil {
		render.Fail(c, render.NewValidatorError(err))
		return
	}

	if err := db.Model(&ServerGroup).Where("id = ?", groupId).Update(&ServerGroup).Error; err != nil {
		render.Fail(c, "更新失败")
		return
	}

	render.Success(c, "更新成功")
}

func Destroy(c *gin.Context) {
	var ServerGroup models.ServerGroup
	var server models.Server
	var count int
	groupId := c.Param("id")

	if err := db.Model(&server).Where("group_id = ?", groupId).Count(&count).Error; err != nil {
		render.Fail(c, err.Error())
		return
	}

	if count > 0 {
		render.Fail(c, "请将绑定此集群的服务器解绑后重试")
		return
	}

	if err := db.Where("id = ?", groupId).Delete(&ServerGroup).Error; err != nil {
		render.Fail(c, "删除失败")
		return
	}

	render.Success(c, "删除成功")
}
