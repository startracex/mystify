# функция go

2022-09-30
2022-11-01
[STARTRACE](/)

Эти функции находятся в ./model

```go
func RenderPost(c *gin.Context) {}
```

`RenderPost` объединяет язык с маршрутом запроса в адрес, анализирует документ и отображает его в HTML - шаблон

```go
func ReadMarkdown(source string) Post {}
```

`ReadMarkdown` aнализируйте источник и возвращайтесь в структуру Post, попробуйте интерпретировать данные Meta с помощью Yaml, а неудача будет использована для чтения Meta со стандартами строк

```go
func AutoPosts(r *gin.RouterGroup) {}
```

`AutoPosts` будет использовать любой маршрут после группы R в качестве адреса документа

```go
func WalkCata(ymlpath string) []Cat {}
```

`WalkCata` прочитает ymlpath и все содержащиеся в нем файлы каталога и вернется в раздел Cat

```go
func Img(c *gin.Context) {}
```

`Img` отправит изображение относительного пути, не начинающегося с `/public`, по пути поиска `pathname` и `filepath`.

## Заброшено

```go
func RenderMarkdown(mdFilePath string) Post {}
```

`RenderMarkdown` возвращает входящую строку в качестве адреса к целевому документу

```go
func RenderCatalog(DirectoryFilePath string) []Cata {}
```

`RenderCatalog` будет читать входящий DirectoryFilePath и разделить символ "\\\\", разделить уровень отступа, заголовок, путь URL, вернуть фрагмент разделенных данных
