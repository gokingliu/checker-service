package utils

import (
	"os"
	"os/exec"
	"strings"
)

// CheckProcessExists 检查进程是否存在
func CheckProcessExists(process string) bool {
	// 使用命令行读取进程
	result, err := exec.Command("/bin/sh", "-c", `ps ux | awk '/`+process+`/ && !/awk/ {print $2}'`).Output()
	// 执行错误返回 false
	if err != nil {
		return false
	}
	// 读取命令行结果中的 pid
	pid := strings.TrimSpace(string(result))

	return pid != ""
}

// CheckPathExists 检查目录是否存在
func CheckPathExists(path string) bool {
	// 读取文件信息
	_, err := os.Stat(path)
	// 读取错误或文件不存在，则返回 false
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

// CheckPathUpdateTime 检查目录文件更新时间
func CheckPathUpdateTime(path string) int64 {
	// 读取文件信息
	info, err := os.Stat(path)
	// 读取错误或文件不存在，则返回 false
	if err != nil && os.IsNotExist(err) {
		return 0
	}

	return info.ModTime().Unix()
}
