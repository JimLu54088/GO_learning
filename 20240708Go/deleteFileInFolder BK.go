// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"path/filepath"
// 	"strings"
// 	"time"
// )

// func main() {
// 	// 提示用户输入目录路径
// 	fmt.Println("Please enter the directory path:")

// 	// 从标准输入读取目录路径
// 	reader := bufio.NewReader(os.Stdin)
// 	dir, err := reader.ReadString('\n')
// 	if err != nil {
// 		fmt.Printf("Error reading input: %v\n", err)
// 		return
// 	}

// 	// 移除路径字符串的前导和末尾空格（包括换行符）
// 	dir = strings.TrimSpace(dir)

// 	err = deleteFilesInDir(dir)
// 	if err != nil {
// 		fmt.Printf("Error deleting files: %v\n", err)
// 		time.Sleep(3 * time.Second)
// 	} else {
// 		fmt.Println("All files deleted successfully.")
// 		time.Sleep(2 * time.Second)
// 	}
// }

// func deleteFilesInDir(dir string) error {
// 	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}

// 		// 如果是文件，则删除它
// 		if !info.IsDir() {
// 			err := os.Remove(path)
// 			if err != nil {
// 				return fmt.Errorf("failed to delete file %s: %w", path, err)
// 			}
// 			fmt.Printf("Deleted file: %s\n", path)
// 		}

// 		return nil
// 	})
// }
