package user

import (
	"deploy/app/models"
	"deploy/database"
	"deploy/helper/pagination"
	"deploy/helper/render"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

var db = database.DB

func Index(c *gin.Context) {
	var users []models.User
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	builder := db.Model(&users).Select("id,account,created_at").Order("created_at desc")

	paginator := pagination.Paging(&pagination.Param{
		DB:   builder,
		Page: page,
	}, &users)

	render.Page(c, *paginator)
}

func Destroy(c *gin.Context) {
	var user models.User
	userID := c.Param("id")

	if err := db.Where("id = ?", userID).Delete(&user).Error; err != nil {
		render.Fail(c, "删除失败")
		return
	}

	render.Success(c, "删除成功")
}

func Update(c *gin.Context) {
	var user models.User
	userID := c.Param("id")
	password := c.PostForm("password")

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		render.Fail(c, err.Error())
		return
	}

	if err := db.Model(&user).Where("id = ?", userID).Update("password", hash).Error; err != nil {
		render.Fail(c, "修改失败")
		return
	}

	render.Success(c, "修改成功")
}
