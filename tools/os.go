package tools

import (
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
	"serveit/config"
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

func ParseToml(path string) (*config.Profile, error) {
	profile := config.Profile{}
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = toml.Unmarshal(file, profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}
