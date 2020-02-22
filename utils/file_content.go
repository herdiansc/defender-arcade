package utils

import (
	"strings"
)

// ReadFile creates readfile boilerplate.
type ReadFile func(filename string) ([]byte, error)

// FileUtil holds ReadFile factory method.
type FileUtil struct {
	ReadFile
}

// NewFileUtil initiates FileUtil.
func NewFileUtil(rf ReadFile) FileUtil {
	return FileUtil{
		ReadFile: rf,
	}
}

// ContentToList loads content of file to list.
func (fu FileUtil) ContentToList(filename string) ([]string, error) {
	content, err := fu.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}

	var list []string
	list = strings.Split(strings.TrimRight(string(content), "\n"), "\n")
	return list, nil
}
