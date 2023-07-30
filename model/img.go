package model

import (
	"github.com/gin-gonic/gin"
	"main/config"
	"path"
	"strings"
)

func Img(c *gin.Context) {
	pathname, _ := c.GetQuery("pathname")
	filepath, _ := c.GetQuery("filepath")
	if strings.Contains(pathname, config.PostsURL) {
		pathname = strings.Replace(pathname, config.PostsURL, "", 1)
		e := path.Join(path.Dir(pathname), filepath)
		c.File("." + config.PostsPath + e)
	} else {
		e := path.Join(path.Dir(pathname), filepath)
		c.File("." + e)
	}
}
