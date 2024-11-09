package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	// 提示用户输入目录路径
	fmt.Println("Please enter the directory path: (ex: D:\\test\\test)")
	os.Stdout.Sync() // 手動刷新標準輸出

	// 从标准输入读取目录路径
	reader := bufio.NewReader(os.Stdin)
	dir, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	// 移除路径字符串的前导和末尾空格（包括换行符）
	dir = strings.TrimSpace(dir)

	// 提示用户是否打印日志
	fmt.Println("Do you want to print log? It means that print out all the deleted file list. (y/n):")
	os.Stdout.Sync() // 手動刷新標準輸出

	strIsPrintLog, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	// 移除换行符
	strIsPrintLog = strings.TrimSpace(strIsPrintLog)

	// 判斷並設置布爾變量
	isPrintLog := strIsPrintLog == "y"

	// 调用删除文件的函数
	err = deleteFilesInDir(dir, isPrintLog)
	if err != nil {
		fmt.Printf("Error deleting files: %v\n", err)
		time.Sleep(3 * time.Second)
	} else {
		fmt.Println("All files deleted successfully.")
		time.Sleep(3 * time.Second)
	}
}

func deleteFilesInDir(dir string, isPrintLog bool) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 如果是文件，则删除它
		if !info.IsDir() {
			err := os.Remove(path)
			if err != nil {
				return fmt.Errorf("failed to delete file %s: %w", path, err)
			}

			// 打印删除日志
			if isPrintLog {
				fmt.Printf("Deleted file: %s\n", path)
			}
		}

		return nil
	})
}
