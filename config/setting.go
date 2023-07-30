package config

import (
	"flag"
)

var Html = "views"
var Port = "8080"
var Domain = "localhost"
var Dev = false
var Export = false
var ExportOut = "docs"
var Cat = ".yml"
var DefaultLang = "en"
var StaticPath = "public"
var PostsPath = "posts"
var PostsURL = "/docs"
var BaseURL = "/"

const ExpirationTime = 2592000

func init() {
	flag.StringVar(&Port, "p", Port, "Port")
	flag.StringVar(&Domain, "d", Domain, "Domain")
	flag.StringVar(&Html, "html", Html, "Html template path")
	flag.StringVar(&ExportOut, "o", ExportOut, "Export out path")
	flag.StringVar(&ExportOut, "out", ExportOut, "Export out path")
	flag.StringVar(&PostsPath, "posts-path", PostsPath, "Posts path")
	flag.StringVar(&PostsURL, "posts-url", PostsURL, "Posts url")
	flag.StringVar(&StaticPath, "static-path", StaticPath, "Static path")
	flag.StringVar(&BaseURL, "base-url", BaseURL, "Base url")
	flag.StringVar(&Cat, "c", Cat, "Catalog basic")
	flag.BoolVar(&Dev, "dev", Dev, "Development mode")
	flag.BoolVar(&Export, "export", Export, "Export html")
	flag.Parse()
	if PostsPath == "" || PostsPath[0] == '/' {
		panic("posts-path set failed")
	}
	if BaseURL[0] != '/' {
		panic("base-url set failed")
	}else{
		if BaseURL[len(BaseURL)-1] != '/' {
			BaseURL = BaseURL + "/"
		}
	}
	if StaticPath == "" || StaticPath[0] == '/' {
		panic("static-path set failed")
	}
	if PostsURL != "" && len(PostsURL) > 1 {
		if PostsURL[0] != '/' {
			panic("posts-url set failed")
		}
	}
}
