package deploy

import (
	"deploy/app/controllers/websocket"
	"deploy/app/models"
	"deploy/app/service"
	"deploy/database"
	"deploy/helper/auth"
	"deploy/helper/command"
	"deploy/helper/pagination"
	"deploy/helper/render"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
	"time"
)

var db = database.DB

func Index(c *gin.Context) {
	var deploy []models.Deploy
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	builder := db.
		Model(&deploy).
		Select("id, name, user_id, project_id, build_status, deploy_status, created_at").
		Order("created_at desc").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, account")
		}).
		Preload("Project", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		})

	paginator := pagination.Paging(&pagination.Param{
		DB:   builder,
		Page: page,
	}, &deploy)

	render.Page(c, *paginator)
}

func Show(c *gin.Context) {
	var deploy models.Deploy
	deployId := c.Param("id")

	db.Model(&deploy).Preload("Project").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, account")
	}).Find(&deploy, deployId)

	render.Data(c, deploy)
}

func Store(c *gin.Context) {
	var deploy models.Deploy

	if err := c.ShouldBind(&deploy); err != nil {
		render.Fail(c, render.NewValidatorError(err))
		log.Printf("%v", deploy)
		return
	}

	deploy.BuildStatus = models.Undeployed
	deploy.UserId = auth.GetUserId(c)

	if err := db.Create(&deploy).Error; err != nil {
		render.Fail(c, err.Error())
		return
	}

	render.Success(c, "添加成功")
}

func Build(c *gin.Context) {
	auth.GetUserInfo(c)
	var deploy models.Deploy
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

	if deploy.IsSuccessfullyDeployed() {
		render.Fail(c, "已部署成功的申请单，无法重新构建")
		return
	}

	if err := db.Where("id = ?", deploy.ProjectId).Find(&project).Error; err != nil {
		render.Fail(c, err.Error())
		return
	}

	service.Build(project, deploy, func(buildPath string, task *command.Task) {
		jsonMsg, _ := json.Marshal(task.Result())
		err := task.GetError()
		if err != nil {
			db.Model(&deploy).Update(map[string]interface{}{"build_log": jsonMsg, "build_status": models.DeploymentFailed})
			time.Sleep(1 * time.Second)
			websocket.SendByUserID(deploy.UserId, websocket.Data{
				Action: "build fail",
				Data:   map[string]time.Time{"time": time.Now()},
			})
			return
		}

		db.Model(&deploy).Update(map[string]interface{}{"build_log": jsonMsg, "build_status": models.SuccessfullyDeployed, "BuildPath": buildPath})
		time.Sleep(1 * time.Second)
		websocket.SendByUserID(deploy.UserId, websocket.Data{
			Action: "build success",
			Data:   map[string]time.Time{"time": time.Now()},
		})
	})

	render.Success(c, "正在构建中")
}

func Deploy(c *gin.Context) {
	var deploy models.Deploy
	var project models.Project
	var serverList []models.Server
	deployID := c.Param("id")

	if err := db.Where("id = ?", deployID).Find(&deploy).Error; err != nil {
		render.Fail(c, err.Error())
		return
	}

	if deploy.BuildStatus != models.SuccessfullyDeployed {
		render.Fail(c, "此申请单未构建成功,无法部署")
		return
	}

	if deploy.DeployStatus == models.SuccessfullyDeployed {
		render.Fail(c, "此申请单已部署成功")
		return
	}

	if err := db.Where("id = ?", deploy.ProjectId).Find(&project).Error; err != nil {
		render.Fail(c, err.Error())
		return
	}

	if err := db.Where("group_id = ?", project.ServerGroupId).Find(&serverList).Error; err != nil {
		render.Fail(c, err.Error())
		return
	}

	service.Deploy(deploy, project, serverList, func(_ string, task *command.Task) {
		jsonMsg, _ := json.Marshal(task.Result())
		err := task.GetError()
		if err != nil {
			db.Model(&deploy).Update(map[string]interface{}{"deploy_log": jsonMsg, "deploy_status": models.DeploymentFailed})
			time.Sleep(1 * time.Second)
			websocket.SendByUserID(deploy.UserId, websocket.Data{
				Action: "deploy fail",
				Data:   map[string]time.Time{"time": time.Now()},
			})
			return
		}

		db.Model(&deploy).Update(map[string]interface{}{"deploy_log": jsonMsg, "deploy_status": models.SuccessfullyDeployed})
		time.Sleep(1 * time.Second)
		websocket.SendByUserID(deploy.UserId, websocket.Data{
			Action: "deploy success",
			Data:   map[string]time.Time{"time": time.Now()},
		})
	})

	render.Success(c, "正在部署中")
}

// 	废弃申请单
func Destroy(c *gin.Context) {
	var deploy models.Deploy
	deployID := c.Param("id")
	log.Print(deployID)

	if err := db.Model(&deploy).Where("id = ?", deployID).Update("build_status", models.Abandoned).Error; err != nil {
		render.Fail(c, err.Error())
		return
	}

	render.Success(c, "已废弃")
}

/*	ws test */
func Test(c *gin.Context) {
	websocket.SendByUserID(1, websocket.Data{
		Action: "build success",
		Data:   map[string]time.Time{"time": time.Now()},
	})
}
