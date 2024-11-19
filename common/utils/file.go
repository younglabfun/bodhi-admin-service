package utils

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func CreateDir(dir string) error {
	err := os.MkdirAll(dir, os.ModeDir)
	if err != nil {
		return err
	}
	return nil
}

func CreateFile(file string) error {
	_, err := os.Stat(file)
	if os.IsExist(err) {
		return nil
	}
	err = os.WriteFile(file, []byte(""), 0666)
	//f, err := os.Create(file)
	if err != nil {
		return err
	}
	//defer f.Close()
	return nil
}

func CreateFileInPath(fullFilePath string) error {
	info := strings.Split(fullFilePath, "/")
	if len(info) > 1 {
		path := strings.Join(info[:len(info)-1], "/")
		err := CreateDir(path)
		if err != nil {
			return err
		}
		err = CreateFile(fullFilePath)
		//fmt.Println("create file", err)
		if err != nil {
			return err
		}
		return nil
	} else {
		return CreateFile(fullFilePath)
	}
}

func WriteFile(filePath, content string) error {
	// if file doesn't exist, create it
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return errors.New("open file failed")
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(content)
	write.Flush()
	return nil
}
