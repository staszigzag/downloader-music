package filestorage

import (
	"fmt"
	"io"
	"os"
)

type FileStorage interface {
	Create(name string, data io.ReadCloser) (answer string, err error)
}

type LocalStorage struct {
	path string
}

func NewLocalStorage(path string) *LocalStorage {
	return &LocalStorage{path}
}

func (s *LocalStorage) Create(name string, data io.ReadCloser) (answer string, err error) {
	defer data.Close()

	if s.path != "" {
		err := os.MkdirAll(s.path, 0o777)
		if err != nil {
			return "", fmt.Errorf("error creating dir : %v", err)
		}
	}
	filepath := s.path + "/" + name
	file, err := os.Create(filepath)
	if err != nil {
		return "", fmt.Errorf("error сreating file %s: %v", name, err)
	}
	defer file.Close()

	_, err = io.Copy(file, data)
	if err != nil {
		os.Remove(name)
		return "", fmt.Errorf("error read or write in file %s: %v", name, err)
	}

	return filepath, nil
}
