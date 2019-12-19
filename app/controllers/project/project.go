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

	paginator := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    page,
		OrderBy: []string{"created_at desc"},
	}, &projects)

	render.Page(c, *paginator)
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
