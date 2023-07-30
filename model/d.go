package model

import "html/template"

// md文件的元数据结构
type Meta struct {
	Title   string   `yaml:"title,omitempty"`
	Author  string   `yaml:"author,omitempty"`
	Authors []string `yaml:"authors,omitempty"`
	Summary string   `yaml:"summary,omitempty"`
	Date    string   `yaml:"date,omitempty"`
	Update  string   `yaml:"update,omitempty"`
}

// 渲染文档的元数据结构
type Post struct {
	Title      string
	Date       string
	Update     string
	Summary    string
	Author     template.HTML
	Body       template.HTML
	OriginFile string // Origin File Path
}

// 目录文件数据结构
type Cata struct {
	Name  string
	Level string
	Super string
}

// 侧边目录
type Cat struct {
	Title string
	Level int8
	Href  string
}
