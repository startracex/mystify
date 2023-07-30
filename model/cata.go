package model

import (
	"main/config"
	"os"
	"path"
	"strings"
)

func ExportCata(ymlpath string) []Cat {
	var c = new([]Cat)
	exportCata(ymlpath, 1, c)
	return *c
}

// 遍历目录
func WalkCata(ymlpath string) []Cat {
	var c = new([]Cat)
	walkCata(ymlpath, 1, c)
	return *c
}

func walkCata(ymlpath string, level int8, catList *[]Cat) *[]Cat {
	data, _ := os.ReadFile(ymlpath)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.ReplaceAll(line, "\r", "")
		s := strings.Split(line, ":")
		direntry, _ := os.ReadDir(path.Dir(ymlpath))
		// 全部文件
		for _, fi := range direntry {
			if !fi.IsDir() {
				// 获取后缀为.md的文件
				if path.Ext(fi.Name()) == ".md" {
					if s[0] == strings.TrimSuffix(fi.Name(), ".md") {
						*catList = append(*catList, Cat{
							Level: level,
							Title: s[1],
							Href:  MaselHref(path.Join(path.Dir(ymlpath), fi.Name())),
						})
					}
				}
			}
		}
		if IsExist(path.Join(path.Dir(ymlpath), s[0], config.Cat)) {
			walkCata(path.Join(path.Dir(ymlpath), s[0], config.Cat), level+1, catList)
		}
	}
	return catList
}

func MaselHref(str string) string {
	str = strings.TrimSuffix(str, "/index.md")
	str = strings.TrimSuffix(str, ".md")
	for _, v := range LangMap {
		if strings.HasPrefix(str, config.PostsPath+v) {
			str = strings.Replace(str, config.PostsPath+v, v+config.PostsURL, 1)
		}
	}
	return str
}
func exportCata(ymlpath string, level int8, catList *[]Cat) *[]Cat {
	data, _ := os.ReadFile(ymlpath)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		// line = strings.ReplaceAll(line, "\r", "")
		s := strings.Split(line, ":")
		direntry, _ := os.ReadDir(path.Dir(ymlpath))
		// 全部文件
		for _, fi := range direntry {
			if !fi.IsDir() {
				// 获取后缀为.md的文件
				if path.Ext(fi.Name()) == ".md" {
					noext := strings.TrimSuffix(fi.Name(), ".md")
					if s[0] == strings.TrimSuffix(fi.Name(), ".md") {
						*catList = append(*catList, Cat{
							Level: level,
							Title: s[1],
							Href:  exportMaselHref(path.Join(path.Dir(ymlpath), noext)),
						})
					}
				}
			}
		}
		if IsExist(path.Join(path.Dir(ymlpath), s[0], config.Cat)) {
			exportCata(path.Join(path.Dir(ymlpath), s[0], config.Cat), level+1, catList)
		}
	}
	return catList
}
func exportMaselHref(str string) string {
	str = strings.TrimSuffix(str, ".md") + ".html"
	if strings.HasPrefix(str, config.PostsPath) {
		str = strings.Replace(str, config.PostsPath, "", 1)
	}
	return path.Join(config.BaseURL, str)
}
