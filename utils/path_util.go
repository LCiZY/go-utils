package utils

import (
	"strings"
)

// JoinPath 将路径与路径/文件名连接起来
func JoinPath(path string, name string) string {
	sep := ""
	if !strings.HasSuffix(path, "/") && !strings.HasSuffix(path, "\\") {
		sep = "/"
	}
	if strings.HasPrefix(name, "./") {
		name = name[2:]
	}
	return path + sep + name
}

func GetDirFromPath(path string) string {
	return path[:strings.LastIndexAny(path, "/")]
}

// GetActualPath 由相对路径得到绝对路径
func GetActualPath(dir, path string) string {
	var noteDir = dir
	if noteDir[len(noteDir)-1:] == "/" {
		noteDir = noteDir[:len(noteDir)-1]
	}
	if path[0] == '.' {
		return noteDir + path[1:]
	}
	if path[0] == '/' || path[0] == '\\' {
		return noteDir + path
	}
	return noteDir + path
}

func GetFileNameFromURL(url string) string {
	splits := strings.Split(url, "/")
	if len(splits) > 0 {
		return splits[len(splits)-1]
	}
	return ""
}
func GetExtensionFromURL(url string) string {
	idx := strings.LastIndex(url, ".")
	if -1 == idx {
		return ""
	}
	return url[idx:]
}

func IsPathLegal(path string) bool {
	if len(path) == 0 {
		return false
	}
	if path[0] == '-' {
		return false
	}
	for _, c := range path { // 以rune遍历 【index会跳着变（以字节为单位）】
		if c > 255 {
			continue
		}
		flag1 := c >= 45 && c <= 57
		flag2 := c >= 65 && c <= 90
		flag3 := c >= 97 && c <= 122
		if !(flag3 || flag2 || flag1) {
			return false
		}
	}
	return true
}

func AddExtensionFromUrl(filename string, url string) string {
	idx := strings.LastIndex(url, ".")
	if idx == -1 {
		return filename
	}
	return filename + url[idx:]
}

func AddSuffixToNameBeforeExtension(prefix, path string) string {
	idx := strings.LastIndex(path, ".")
	return path[:idx] + prefix + path[idx:]
}
