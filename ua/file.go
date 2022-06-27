package ua

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type File struct {
	Dir          string
	Name         string
	CompletePath string
}

func NewFileCache(name string) *File {
	return &File{
		Name:         name,
		CompletePath: name,
	}
}

func (f *File) Read() ([]byte, error) {
	return ioutil.ReadFile(f.CompletePath)
}

func (f *File) Write(data []byte) error {
	return ioutil.WriteFile(f.CompletePath, data, 0644)
}

func (f *File) WriteJson(v interface{}) error {
	uasJson, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return f.Write(uasJson)
}

func (f *File) Remove() error {
	err := os.Remove(f.CompletePath)
	if err != nil {
		return err
	}

	return nil
}

func (f *File) IsExist() (bool, error) {
	_, err := os.Stat(f.CompletePath)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func GetTempDir() string {
	tempDir := os.TempDir()
	if exist := strings.HasSuffix(tempDir, "/"); exist == false {
		tempDir = tempDir + "/"
	}

	return tempDir
}
