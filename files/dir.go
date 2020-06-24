package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func IsDirExist(dir string) bool {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateDir(dir string) error {
	if IsDirExist(dir) {
		return nil
	}

	err := os.MkdirAll(dir, os.ModePerm)
	if nil != err {
		return fmt.Errorf("create path %s failed %v", dir, err)
	}

	return nil
}

func ListDirByPrefix(dir string, prefix string) (files []string, err error) {
	files = make([]string, 0, 10)

	dirList, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	prefix = strings.ToLower(prefix) //忽略后缀匹配的大小写

	for _, fi := range dirList {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasPrefix(strings.ToLower(fi.Name()), prefix) { //匹配文件
			files = append(files, fi.Name())
		}
	}

	return files, nil
}
