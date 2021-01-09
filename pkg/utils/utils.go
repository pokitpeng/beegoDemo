package utils

import "os"

// Exists file or directory exists ?
func Exists(path string) bool {
	_, err := os.Stat(path) // os.Stat get file or dir info
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDir ...
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile ...
func IsFile(path string) bool {
	return !IsDir(path)
}
