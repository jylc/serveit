package tools

import (
	"os"
	"path/filepath"
)

// JudgeFile 判断文件路径是否正确
func JudgeFile(path string) (string, error) {
	path = filepath.Clean(path)
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return "", err
		} else {
			return path, nil
		}
	}
	return path, nil
}
