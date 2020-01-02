package project

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
	var projects []models.Project
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	_, brief := c.GetQuery("brief")

	if brief == true {
		db.Model(&projects).Select("id, name").Order("created_at desc").Find(&projects)

		render.Data(c, projects)
	} else {
		builder := db.Set("gorm:auto_preload", true).
			Model(&projects).Order("created_at desc").
			Preload("ServerGroup")

		paginator := pagination.Paging(&pagination.Param{
			DB:   builder,
			Page: page,
		}, &projects)

		render.Page(c, *paginator)
	}
}

func Show(c *gin.Context) {
	var project models.Project
	projectId := c.Param("id")

	if err := db.Model(&project).Where("id = ?", projectId).Find(&project).Error; err != nil {
		render.Fail(c, err.Error())
		return
	}

	render.Data(c, project)
}

func Store(c *gin.Context) {
	var project models.Project

	if err := c.ShouldBind(&project); err != nil {
		render.Fail(c, render.NewValidatorError(err))
		return
	}

	if err := db.Create(&project).Error; err != nil {
		render.Fail(c, err.Error())
		return
	}

	render.Success(c, "添加成功")
}

func Update(c *gin.Context) {
	var project models.Project
	projectId := c.Param("id")

	if err := c.ShouldBind(&project); err != nil {
		render.Fail(c, render.NewValidatorError(err))
		return
	}

	if err := db.Model(&project).Where("id = ?", projectId).Update(&project).Error; err != nil {
		render.Fail(c, "更新失败")
		return
	}

	render.Success(c, "更新成功")

}

func Destroy(c *gin.Context) {
	var project models.Project
	projectId := c.Param("id")

	if err := db.Where("id = ?", projectId).Delete(&project).Error; err != nil {
		render.Fail(c, "删除失败")
		return
	}

	render.Success(c, "删除成功")
}
