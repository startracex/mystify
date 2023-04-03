---
title: Type
date: 2023-01-04
update: 2023-01-04
---

## Catalog

```go
type Cat struct {
 Title string
 Level int8
 Href  string
}
type Cata struct {
 Name  string
 Level string
 Super string
}
```

`Cat`and`Cat`are structures used to render directories, and they have the same meaning

## Post

```go
type Post struct {
 Title      string
 Date       string
 Update     string
 Summary    string
 Author     template.HTML
 Body       template.HTML
 OriginFile string // Origin File Path
}
```

`Post`is a structure used to render articles. Some of the content comes from `Meta`

## Meta

```go
type Meta struct {
 Title   string   `yaml:"title,omitempty"`
 Author  string   `yaml:"author,omitempty"`
 Authors []string `yaml:"authors,omitempty"`
 Summary string   `yaml:"summary,omitempty"`
 Date    string   `yaml:"date,omitempty"`
 Update  string   `yaml:"update,omitempty"`
}
```

`Meta`is the metadata of the document
