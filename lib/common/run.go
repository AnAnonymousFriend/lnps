package common

import (
	"os"
	"path/filepath"
	"runtime"
)

//Get the currently selected configuration file directory
//For non-Windows systems, select the /etc/nps as config directory if exist, or select ./
//windows system, select the C:\Program Files\nps as config directory if exist, or select ./
func GetRunPath() string {
	var path string
	if path = GetInstallPath(); !FileExists(path) {
		return GetAppPath()
	}
	return path
}

//Different systems get different installation paths
func GetInstallPath() string {
	var path string
	if IsWindows() {
		// 如果是windows 安装路径默认为C盘
		path = `C:\Program Files\lnps`
	} else {
		// 如果不是windows 则为 UNIX 系统 默认为/etc/路径
		path = "/etc/lnps"
	}
	return path
}

//Get the absolute path to the running directory
// 获取程序运行时文件夹路径
func GetAppPath() string {
	if path, err := filepath.Abs(filepath.Dir(os.Args[0])); err == nil {
		return path
	}
	return os.Args[0]
}

// 判断是否为windows
func IsWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	}
	return false
}

// 判断文件是否存在
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}