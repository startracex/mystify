package model

import (
	"github.com/gin-gonic/gin"
	"md/config"
)

type Lang struct {
	Url string
	Desc string
}

var ZH = Lang{Url: "/zh", Desc: "zh"}
var EN = Lang{Url: "/en", Desc: "en"}
var RU = Lang{Url: "/ru", Desc: "ru"}

func I18ninit(r *gin.Engine) {
	EN.i18n(r)
	ZH.i18n(r)
	RU.i18n(r)
}

func (l Lang) i18n(r *gin.Engine) {
	g := r.Group(l.Url, func(c *gin.Context) {
		ResetCookie(c, "lang", l.Desc)
		c.Set("lang", l.Desc)
	})

	g.GET("/", func(c *gin.Context) {
		src := "/README.md"
		if l.Desc != config.DefaultLang {
			src = "/README." + l.Desc + ".md"
		}
		c.HTML(200, "posts.html", gin.H{"Markdown": ReadMarkdown(src)})
	})
	AutoPosts(g)
}

func AutoPosts(r *gin.RouterGroup) {
	docs := r.Group(config.PostsURL)
	docs.GET("/*url", func(c *gin.Context) { RenderPost(c) })
}

