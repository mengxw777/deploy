package route

import (
	authController "deploy/app/controllers/auth"
	deployController "deploy/app/controllers/deploy"
	projectController "deploy/app/controllers/project"
	serverController "deploy/app/controllers/server"
	serverGroupController "deploy/app/controllers/serverGroup"
	"deploy/app/models"
	_ "deploy/app/validators"
	"deploy/database"
	"deploy/helper/render"
	"io"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var (
	db          = database.DB
	identityKey = "id"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// 	设备日志文件
	f, _ := os.Create("runtime.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 	中间件
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	// 	JWT
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "gin blog",
		Key:           []byte(os.Getenv("JWT_SECRET")),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		IdentityKey:   identityKey,
		TokenLookup:   "header: Authorization, query: token, cookie: token",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.Id,
				}
			}
			return jwt.MapClaims{}
		},
		// 	验证动作
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var user models.User

			if err := c.ShouldBind(&user); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}

			db.Where("account = ?", user.Account).Find(&user)
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.PostForm("password")))
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return &user, nil
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			render.Data(c, map[string]string{"token": token})
		},
	})

	{

		api := router.Group("api")
		api.Use(authMiddleware.MiddlewareFunc())

		// 	权限
		router.POST("/api/auth/register", authController.Register)
		auth := router.Group("api/auth")
		{
			auth.POST("login", authMiddleware.LoginHandler)
			auth.GET("/refresh_token", authMiddleware.RefreshHandler)
		}

		// 	服务器集群
		serverGroup := api.Group("server_group")
		{
			serverGroup.GET("/", serverGroupController.Index)
			serverGroup.PUT("/:id", serverGroupController.Update)
			serverGroup.POST("/store", serverGroupController.Store)
			serverGroup.DELETE("/:id", serverGroupController.Destroy)
		}

		// 	服务器
		server := api.Group("server")
		{
			server.GET("/", serverController.Index)
			server.PUT("/:id", serverController.Update)
			server.POST("/store", serverController.Store)
			server.DELETE("/:id", serverController.Destroy)
		}

		// 	项目
		project := api.Group("project")
		{
			project.GET("/", projectController.Index)
			project.PUT("/:id", projectController.Update)
			project.POST("/store", projectController.Store)
			project.DELETE("/:id", projectController.Destroy)
		}

		// 	部署
		deploy := api.Group("deploy")
		{
			deploy.GET("/", deployController.Index)
			deploy.POST("/store", deployController.Store)
			deploy.DELETE("/:id", deployController.Destroy)
			deploy.POST("/build/:id", deployController.Build)
			deploy.POST("/deploy/:id", deployController.Deploy)
		}
	}

	return router
}
