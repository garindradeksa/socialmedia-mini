package main

import (
	"log"

	"github.com/garindradeksa/socialmedia-mini/config"
	cmtD "github.com/garindradeksa/socialmedia-mini/features/comment/data"
	cmtH "github.com/garindradeksa/socialmedia-mini/features/comment/handler"
	cmtS "github.com/garindradeksa/socialmedia-mini/features/comment/services"
	cntD "github.com/garindradeksa/socialmedia-mini/features/content/data"
	cntH "github.com/garindradeksa/socialmedia-mini/features/content/handler"
	cntS "github.com/garindradeksa/socialmedia-mini/features/content/services"
	usrD "github.com/garindradeksa/socialmedia-mini/features/user/data"
	usrH "github.com/garindradeksa/socialmedia-mini/features/user/handler"
	usrS "github.com/garindradeksa/socialmedia-mini/features/user/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	userData := usrD.New(db)
	userSrv := usrS.New(userData)
	userHdl := usrH.New(userSrv)

	contentData := cntD.New(db)
	contentSrv := cntS.New(contentData)
	contentHdl := cntH.New(contentSrv)

	commentData := cmtD.New(db)
	commentSrv := cmtS.New(commentData)
	commentHdl := cmtH.New(commentSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.POST("/register", userHdl.Register())
	e.POST("/login", userHdl.Login())

	user := e.Group("/users")

	content := e.Group("")
	content.Use(middleware.JWT([]byte(config.JWTKey)))

	comment := e.Group("")
	comment.Use(middleware.JWT([]byte(config.JWTKey)))

	user.GET("/:username", contentHdl.GetProfile())
	user.GET("/profile", userHdl.Profile(), middleware.JWT([]byte(config.JWTKey)))
	user.PUT("/profile", userHdl.Update(), middleware.JWT([]byte(config.JWTKey)))
	user.DELETE("", userHdl.Deactivate(), middleware.JWT([]byte(config.JWTKey)))

	content.PUT("/contents/:id", contentHdl.Update())
	content.DELETE("/contents/:id", contentHdl.Delete())
	e.GET("/contents/:id", contentHdl.ContentDetail())
	e.GET("/contents", contentHdl.ContentList())

	comment.POST("/comments/:id", commentHdl.Add())
	comment.DELETE("/comments/:id", commentHdl.Delete())

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
