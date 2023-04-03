---
title: Тип
date: 2023-01-04
update: 2023-01-04
---

## Каталог

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

`Cat` и `Cata` это структуры, используемые для отображения каталогов, оба из которых имеют одинаковое значение

## Статья

```go
type Post struct {
 Title      string
 Date       string
 Update     string
 Summary    string
 Author     template.HTML
 Body       template.HTML
 OriginFile string
}
```

`Post` структура, используемая для рендеринга статей, частично из`Meta`

## Метаданные

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

`Meta` метаданные для документа.
