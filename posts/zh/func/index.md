---
title: 用于go的功能
date: 2022-09-30
update: 2023-01-04
author: "[STARTRACEX](/)"
---

## 概览

这些功能位于./model

```go
func RenderPost(c *gin.Context) {}
```

`RenderPost` 将会合并语言与请求路径为地址并解析文档，并渲染至HTML模板

```go
func ReadMarkdown(source string) Post {}
```

`ReadMarkdown` 解析source并返回Post结构,将会尝试以yaml解析Meta数据,失败将使用以行标准读取Meta

```go
func AutoPosts(r *gin.RouterGroup) {}
```

`AutoPosts` 会将路由组r之后的任何路由作为文档地址

```go
func WalkCata(ymlpath string) []Cat {}
```

`WalkCata` 将会读取ymlpath和其中包含的所有目录文件并返回Cat切片

```go
func Img(c *gin.Context) {}
```

`Img` 将会通过搜索路径`pathname`和`filepath`来发送非`/public`起始的相对路径的图片

## 已弃用

```go
func RenderMarkdown(mdFilePath string) Post {}
```

`RenderMarkdown` 将传入的字符串作为地址返回解析的目标文档

```go
func RenderCatalog(DirectoryFilePath string) []Cata {}
```

`RenderCatalog` 将读取传入的DirectoryFilePath并以"\\\\"为分割符号,分割缩进等级，标题，URL路径，返回分割完毕的数据切片
