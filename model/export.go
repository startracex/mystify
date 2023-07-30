package model

import (
	"fmt"
	"html/template"
	"io"
	"main/config"
	"os"
	"path"
	"strings"
)

var c = ""
var cc []Cat

func Export(mdpath, out string) {
	export(mdpath, out, mdpath)

}

func export(mdpath, out, root string) {
	dirEntry, _ := os.ReadDir(mdpath)
	for _, fi := range dirEntry {
		thispath := path.Join(mdpath, fi.Name())
		if fi.IsDir() {
			export(thispath, out, root)
		} else {
			if path.Ext(fi.Name()) == ".md" {
				Post := ReadMarkdown(thispath)
				thispath = strings.ReplaceAll(thispath, root, "")
				// 创建out下的同名html
				htmldir := path.Join(out, path.Dir(thispath))
				fmt.Println(htmldir, thispath)
				os.MkdirAll(htmldir, 0777)
				f, _ := os.Create(path.Join(out, thispath[:len(thispath)-3]+".html"))
				// 将views下的layou和page中的模板进行解析
				t := template.Must(template.ParseFiles(
					WalkHtml(config.Html)...,
				))
				// catapath的值是thispath在第二个/后的值
				catapath := strings.Split(thispath, "/")[1]
				// 储存当前值,防止重复
				if catapath != c {
					c = catapath
					cc = ExportCata(path.Join(config.PostsPath, catapath, config.Cat))
				}
				lang := strings.Split(thispath, "/")[0]
				t.ExecuteTemplate(f, "page.html",
					map[string]interface{}{
						"Markdown": Post,
						"Catalog":  cc,
						"BaseURL":  config.BaseURL,
						"Lang":     lang,
					})
				f.Close()
			} else {
				// 如果不是md文件,直接复制到out
				if fi.Name() != config.Cat {
					CopyFile(path.Join(out, thispath), thispath)
				}
			}
		}
	}
	CopyDir(config.StaticPath, path.Join(out, config.StaticPath))
}
func CopyFile(dstName, srcName string) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		return
	}
}
func CopyDir(src, out string) {
	dirEntry, _ := os.ReadDir(src)
	for _, fi := range dirEntry {
		thispath := path.Join(src, fi.Name())
		if fi.IsDir() {
			// 创建out下的同名文件夹
			err := os.MkdirAll(path.Join(out, fi.Name()), 0777)
			if err != nil {
				return
			}
			CopyDir(thispath, path.Join(out, fi.Name()))
		} else {
			CopyFile(path.Join(out, fi.Name()), thispath)
		}
	}
}
