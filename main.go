package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/config"
	"main/model"
	"strings"
)

func main() {
	if config.Export {
		model.Export(config.PostsPath, config.ExportOut)
		return
	}
	if !config.Dev {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Recovery())

	/* Load html */
	r.LoadHTMLFiles(model.WalkHtml(config.Html)...)

	/* Static file */
	r.Static("/"+config.StaticPath, config.StaticPath)
	if !strings.HasPrefix(config.PostsPath, config.StaticPath) {
		r.Static(config.PostsPath, config.PostsPath)
	}
	r.StaticFile("/favicon.ico", config.StaticPath+"/favicon.ico")

	r.StaticFile("/robots.txt", config.StaticPath+"/robots.txt")

	/* Readme */
	r.StaticFile("/README.md", "./README.md")
	r.StaticFile("/README.zh.md", "./README.zh.md")
	r.StaticFile("/README.ru.md", "./README.ru.md")

	/* Root */
	r.GET("/", func(c *gin.Context) {
		post := model.ReadMarkdown("/README.md")
		c.HTML(200, "page.html", gin.H{"Markdown": post})
	})

	/* Main */
	model.I18Ninit(r)
	/* Img */
	r.GET("/img", model.Img)
	err := r.Run(":" + config.Port)
	if err != nil {
		fmt.Printf("Run failed at :%s", config.Port)
	}
}
