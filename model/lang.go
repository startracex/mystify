package model

import (
	"github.com/gin-gonic/gin"
	"main/config"
	"os"
)

var LangMap map[string]string = make(map[string]string)

func I18Ninit(r *gin.Engine) {
	dirs, _ := os.ReadDir(config.PostsPath)
	for _, v := range dirs {
		name := v.Name()
		LangMap[name] = config.BaseURL + name
		if v.IsDir() {
			WithLang(r, config.BaseURL+name, name)
		}
	}
}

func AutoPosts(r *gin.RouterGroup) {
	docs := r.Group(config.PostsURL)
	docs.GET("/*url", func(c *gin.Context) { RenderPost(c) })
}

func WithLang(r *gin.Engine, url, desc string) {
	g := r.Group(url, func(c *gin.Context) {
		ResetCookie(c, "lang", desc)
		c.Set("lang", desc)
	})
	g.GET("/", func(c *gin.Context) {
		src := "/README.md"
		if desc != config.DefaultLang {
			src = "/README." + desc + ".md"
		}
		c.HTML(200, "page.html", gin.H{
			"Markdown": ReadMarkdown(src),
			"BaseURL":  config.BaseURL,
		})
	})
	AutoPosts(g)
}
