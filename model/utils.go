package model

import (
	"os"
	"path/filepath"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// 返回dir下的所有文件
func WalkHtmlTemplate(dir string) []string {
	fileList := []string{}
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if path != dir && !f.IsDir() {
			fileList = append(fileList, path)
		}
		return nil
	})
	return fileList
}
