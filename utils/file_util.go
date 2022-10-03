package utils

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/LCiZY/go-utils/logs"
)

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFile(path string) bool {
	return !IsDir(path)
}

func ReName(src, dst string) bool {
	return os.Rename(src, dst) == nil
}

func Remove(path string) bool {
	err := os.Remove(path)
	if err != nil {
		logs.Error("remove file failed, err: %v", err)
		return false
	}
	return true
}

func GetFileLastModifyTime(path, format string) string {
	//获取文件修改时间 返回unix时间戳
	f, err := os.Open(path)
	if err != nil {
		logs.Error("open file failed, err: %v", err)
		return time.Now().Format(format)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		logs.Error("stat fileinfo error: %v", err)
		return time.Now().Format(format)
	}
	return fi.ModTime().Format(format)
}

func SaveToDisk(file multipart.File, path string) bool {
	out, err := os.Create(path)
	if err != nil {
		logs.Error("failed to open the file %s for writing, err: %v", path, err)
		return false
	}
	defer file.Close()
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		logs.Error("copy file failed, err:%s", err)
		return false
	}
	return true
}

func IsImageExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			logs.Error("IsImageExist: path: %s, error: %v", path, err)
		}
		return false
	}
	return true
}

func IsSameImage(path string, bytes []byte) bool {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return false
	}
	if len(f) != len(bytes) {
		return false
	}
	h1 := sha1.New()
	bs1 := h1.Sum(f)
	src := fmt.Sprintf("%x", bs1)

	h2 := sha1.New()
	bs2 := h2.Sum(bytes)
	curr := fmt.Sprintf("%x", bs2)

	return src == curr
}

var fileTypeMap sync.Map
var supportedImageExtension = make(map[string]interface{})

func init() {
	supportedImageExtension["jpg"] = struct{}{}
	supportedImageExtension["png"] = struct{}{}
	supportedImageExtension["gif"] = struct{}{}
	supportedImageExtension["bmp"] = struct{}{}
	supportedImageExtension["tif"] = struct{}{}
	supportedImageExtension["ico"] = struct{}{}
	supportedImageExtension["webp"] = struct{}{}

	fileTypeMap.Store("ffd8", "jpg")      //JPEG (jpg)
	fileTypeMap.Store("89504e47", "png")  //PNG (png)
	fileTypeMap.Store("47494638", "gif")  //GIF (gif)
	fileTypeMap.Store("4949", "tif")      //TIFF (tif)
	fileTypeMap.Store("4d4d", "tif")      //TIFF (tif)
	fileTypeMap.Store("424d", "bmp")      //16色位图(bmp)
	fileTypeMap.Store("00000100", "ico")  //ico
	fileTypeMap.Store("52494646", "webp") //webp
}

// 获取前面结果字节的16进制
func bytesToHexString(src []byte) string {
	res := bytes.Buffer{}
	if src == nil || len(src) <= 0 {
		return ""
	}
	temp := make([]byte, 0)
	for _, v := range src {
		sub := v & 0xFF
		hv := hex.EncodeToString(append(temp, sub))
		if len(hv) < 2 {
			res.WriteString(strconv.FormatInt(int64(0), 10))
		}
		res.WriteString(hv)
	}
	return res.String()
}

// GetFileType
// 用文件前面几个字节来判断
// fSrc: 文件字节流（就用前面几个字节）
// 返回文件后缀类型，不包含点
func GetFileType(fSrc []byte) string {
	var fileType string
	fileCode := bytesToHexString(fSrc)
	fileTypeMap.Range(func(key, value interface{}) bool {
		k := key.(string)
		v := value.(string)
		if strings.HasPrefix(fileCode, strings.ToLower(k)) {
			fileType = v
			return false
		}
		return true
	})
	return fileType
}

func IsImage(bytes []byte) bool {
	if len(bytes) <= 256 {
		return false
	}
	_, exist := supportedImageExtension[GetFileType(bytes[:10])]
	return exist
}

// IsSupportedImageExt 返回是否支持此图片后缀
func IsSupportedImageExt(ext string) bool {
	if ext == "" {
		return false
	}
	if ext[:1] == "." {
		ext = ext[1:]
	}
	_, e := supportedImageExtension[ext]
	return e
}
