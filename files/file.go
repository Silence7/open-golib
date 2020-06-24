package files

import (
	"io/ioutil"
	"os"
)

func ReadFile(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

//func ReadFromFile(fileName string) (string, error) {
//	var str strings.Builder
//
//	fp, err := os.Open(fileName)
//	if err != nil {
//		return "", fmt.Errorf("read %s error %v", fileName, err)
//	}
//	defer fp.Close()
//
//	reader := bufio.NewReader(fp)
//	if nil == reader {
//		return "", fmt.Errorf("new reader error")
//	}
//
//	for {
//		line, err := reader.ReadString('\n')
//		if nil != err {
//			return "", fmt.Errorf("read string error %v", err)
//		}
//
//		str.WriteString(line)
//	}
//
//	return str.String(), nil
//}

func AddFile(file string, Content string) error {
	data := []byte(Content)
	fd, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0644)
	if nil != err {
		return err
	}
	defer fd.Close()
	fd.Write(data)

	return nil
}

func DeleteFile(file string) error {
	if !IsDirExist(file) {
		return nil
	}

	//os.Truncate(file, 0)

	err := os.Remove(file)
	if nil != err {
		return err
	}

	return nil
}
