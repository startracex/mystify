package model

import (
	"github.com/gin-gonic/gin"
	setting "md/config"
)

// 设置cookie name:value
func SetCookie(c *gin.Context, name, value string) {
	c.SetCookie(name, value, setting.ExpirationTime, "/", setting.Domain, false, false)
}

// 返回name的值，不存在返回nil,存在返回获取值
func GetCookie(c *gin.Context, name string) interface{} {
	get, err := c.Cookie(name)
	if err != nil {
		return nil
	}
	return get
}

// 移除name
func RemoveCookie(c *gin.Context, name string) {
	v := GetCookie(c, name)
	if v != nil {
		c.SetCookie(name, "", -1, "/", setting.Domain, false, false)
	}
}

// 如果name的值不为reset,重设name的值到reset
func ResetCookie(c *gin.Context, name, reset string) {
	get, err := c.Cookie(name)
	if err != nil || get != reset {
		SetCookie(c, name, reset)
	}
}

// 设置安全cookie
func SetCookieSecurity(c *gin.Context, name, value string) {
	c.SetCookie(name, value, setting.ExpirationTime, "/", setting.Domain, false, true)
}

// 重设安全cookie
func ResetCookieSecurity(c *gin.Context, name, reset string) {
	get, err := c.Cookie(name)
	if err != nil || get != reset {
		SetCookieSecurity(c, name, reset)
	}
}