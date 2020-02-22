package utils

import (
	"strings"
)

type ReadFile func(filename string) ([]byte, error)
type FileUtil struct {
	ReadFile
}

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
