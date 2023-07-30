package model

import (
	"github.com/gin-gonic/gin"
	"main/config"
	"os"
	"path/filepath"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// 返回dir下的所有html文件
func WalkHtml(dir string) []string {
	return Walk(dir, func(path string, fi os.FileInfo) bool {
		return !fi.IsDir() && filepath.Ext(path) == ".html"
	})
}

func Walk(dir string, f func(path string, f os.FileInfo) bool) []string {
	fileList := []string{}
	filepath.Walk(dir, func(path string, fi os.FileInfo, err error) error {
		if f(path, fi) {
			fileList = append(fileList, path)
		}
		return nil
	})
	return fileList
}

// 设置cookie name:value
func SetCookie(c *gin.Context, name, value string) {
	c.SetCookie(name, value, config.ExpirationTime, "/", config.Domain, false, false)
}

// 返回name的值，不存在返回nil,存在返回获取值
func GetCookie(c *gin.Context, name string) interface{} {
	get, err := c.Cookie(name)
	if err != nil {
		return nil
	}
	return get
}

// 如果name的值不为reset,重设name的值到reset
func ResetCookie(c *gin.Context, name, reset string) {
	get, err := c.Cookie(name)
	if err != nil || get != reset {
		SetCookie(c, name, reset)
	}
}
