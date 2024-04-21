package main

import (
	"fmt"
	"os"
)

type vaulter interface {
	AllFiles() ([]string, error)
	FileContent(file string) (string, error)
	FileLinks(file string) ([]string, error)
	EditFile(file, content string) error
}

type vault struct {
}

func NewVault() vault {
	return vault{}
}

func (v vault) AllFiles() ([]string, error) {
	files, err := os.ReadDir("./vault")
	if err != nil {
		return []string{}, err
	}
	res := make([]string, 0, len(files))
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		res = append(res, file.Name())
	}
	return res, nil
}

func (v vault) FileContent(fileName string) (string, error) {
	fileName = string(fmt.Sprintf("./vault/%s", fileName))

	content, err := os.ReadFile(fileName)
	return string(content), err
}

func (v vault) FileLinks(file string) ([]string, error) {
	panic("not implemented") // TODO: Implement
}

func (v vault) EditFile(fileName string, content string) error {
	fileName = string(fmt.Sprintf("./vault/%s", fileName))

	file, err := os.OpenFile(fileName, os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(content))
	if err != nil {
		return err
	}

	return nil
}
