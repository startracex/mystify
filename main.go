package main

import (
	"md/config"
	"md/model"
	"strings"

	"github.com/gin-gonic/gin"
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
	var htmlg []string
	for _, k := range model.WalkHtmlTemplate(config.Html) {
		if strings.HasSuffix(k, ".html") {
			htmlg = append(htmlg, k)
		}
	}
	r.LoadHTMLFiles(htmlg...)

	/* Static file */
	r.Static("/"+config.StaticPath, config.StaticPath)
	if !strings.HasPrefix(config.PostsPath, config.StaticPath) {
		r.Static(config.PostsPath, config.PostsPath)
	}
	r.StaticFile("/favicon.ico", config.StaticPath+"/favicon.ico")
	r.StaticFile("/robots.txt", config.StaticPath+"/robots.txt")
	r.StaticFile("/README.md", "./README.md")
	r.StaticFile("/README.zh.md", "./README.zh.md")
	r.StaticFile("/README.ru.md", "./README.ru.md")

	/* root */
	r.GET("/", func(c *gin.Context) {
		post := model.ReadMarkdown("/README.md")
		c.HTML(200, "posts.html", gin.H{"Markdown": post})
	})
	model.AutoPosts(r.Group("/"))
	/* main */
	model.I18ninit(r)
	/* img */
	r.GET("/img", model.Img)
	r.Run(":" + config.Port)
}
