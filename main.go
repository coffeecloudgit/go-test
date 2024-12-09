package main

import (
	"fmt"
	"os"
)

func RandomReadFile(file *os.File, offset, size int64) ([]byte, error) {
	data := make([]byte, size)
	_, err := file.ReadAt(data, offset)
	if err != nil {
		return nil, fmt.Errorf("读取文件失败: %v", err)
	}
	return data, nil
}

func worker(file *os.File, offset, fileSize, readSize int64) ([]byte, error) {

	// 计算最大偏移量
	maxOffset := fileSize - readSize
	if maxOffset <= 0 {
		fmt.Printf("协程错误: 文件大小小于读取大小\n")
		return nil, fmt.Errorf("协程错误: 文件大小小于读取大小\n")
	}
	//offset := int64(200000)

	//offset := int64(5 * 1024 * 1024 * 1024)
	//200000
	//[14 24 87 35 143 127 212 171 50 148 0 131 132 209 198 60 60 43 67 71 65]
	//1024 * 1024 * 1024
	//[237 82 97 113 185 129 196 75 24 194 39 179 69 6 26 248 139 49 205 42 53]

	return RandomReadFile(file, offset, readSize)
}

func readBytes(file *os.File, offset, readSize int64) ([]byte, error) {
	//var filePath = "/Users/ym/fs/sealed/s-t03252730-1"
	// 打开文件

	// 获取文件大小
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("获取文件信息失败: %v", err)
		return nil, err
	}
	fileSize := fileInfo.Size()

	return worker(file, offset, fileSize, readSize)
}

func main() {
	offset := int64(3467247618)
	readSize := int64(300)
	var filePath1 = "/Users/ym/fs/sealed/s-t03252730-1"
	var filePath2 = "/Users/ym/Desktop/sealed/s-t03252730-2"
	file1, err := os.OpenFile(filePath1, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
		return
	}
	defer file1.Close()

	file2, err := os.OpenFile(filePath2, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
		return
	}
	defer file2.Close()

	b1, err1 := readBytes(file1, offset, readSize)
	if err1 != nil {
		fmt.Printf("err1:%v \n", err1)
		return
	}

	b2, err2 := readBytes(file2, offset, readSize)

	if err2 != nil {
		fmt.Printf("err2:%v \n", err2)
		return
	}

	if string(b1) != string(b2) {
		fmt.Printf("b1 b2 not equal, offset:%d, \n b1:%v, \n b2:%v \n", offset, b1, b2)
		return
	}

	fmt.Println("b1 b2 is equal")
}
