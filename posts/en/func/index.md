# Functions for go

2022-09-30
2022-11-01
[STARTRACE](/)

These functions can be found in./model

```go
func RenderPost(c *gin.Context) {}
```

`RenderPost` will merge the language and request path as the address, parse the document, and render it to the HTML template

```go
func ReadMarkdown(source string) Post {}
```

`ReadMarkdown` parses the source and returns the Post structure. It will attempt to parse the meta data in yaml. If it fails, it will use the row standard to read the meta

```go
func AutoPosts(r *gin.RouterGroup) {}
```

`AutoPosts` will take any route after routing group r as the document address

```go
func WalkCata(ymlpath string) []Cat {}
```

`WalkCata` will read ymlpath and all directory files contained in it and return Cat slices

```go
func Img(c *gin.Context) {}
```

`Img`will send pictures of relative paths other than`/public`by searching the paths `pathname` and `filepath`

## 已弃用

```go
func RenderMarkdown(mdFilePath string) Post {}
```

`RenderMarkdown` Return the parsed target document with the passed in string as the address

```go
func RenderCatalog(DirectoryFilePath string) []Cata {}
```

`RenderCatalog` will read the incoming DirectoryFilePath with "\\\\" as the dividing symbol, split the indent level, title, URL path, and return the split data slice
