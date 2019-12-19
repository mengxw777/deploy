package controllers

import (
	"deploy/app/models"
	"deploy/database"
	"deploy/helper/render"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var db = database.DB

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		render.Fail(c, render.NewValidatorError(err))
		return
	}

	if !db.Select("account = ?", user.Account).Find(&user).RecordNotFound() {
		render.Fail(c, "账号已存在")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		render.Fail(c, err.Error())
		return
	}

	user.Password = string(hash)

	if err := db.Create(&user).Error; err != nil {
		render.Fail(c, err.Error())
		return
	}

	render.Success(c, "注册成功")
}
