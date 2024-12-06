package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_readBytes(t *testing.T) {
	//index read not equal, i:67125141,
	//index read not equal, i:80739846,
	offset := int64(67125141)
	//var filePath = "/Users/ym/fs/sealed/s-t03252730-1"
	var filePath1 = "/Users/ym/Desktop/sealed/s-t03252730-1"
	var filePath2 = "/Users/ym/fs/sealed/s-t03252730-1"
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

	everySize := int64(333)
	b1, err1 := readBytes(file1, offset, everySize)
	if err1 != nil {
		fmt.Printf("err1:%v \n", err1)
		return
	}
	b2, err2 := readBytes(file2, offset, everySize)

	if err2 != nil {
		fmt.Printf("err2:%v \n", err2)
		return
	}

	if string(b1) != string(b2) {
		fmt.Printf("index read not equal, offset:%d, \n b1:%v, \n b2:%v \n", offset, b1, b2)
		return
	}
	fmt.Printf("index read is equal, offset:%d \n", offset)

}

func Test_Main(t *testing.T) {
	//offset := int64(3467247618)
	//var filePath = "/Users/ym/fs/sealed/s-t03252730-1"
	var filePath1 = "/Users/ym/Desktop/sealed/s-t03252730-2"
	var filePath2 = "/Users/ym/fs/sealed/s-t03252730-2"
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

	fileInfo, err := file1.Stat()
	if err != nil {
		fmt.Printf("获取文件信息失败: %v", err)
		return
	}
	everySize := int64(333)
	for i := int64(0); i < fileInfo.Size(); i = i + everySize {
		b1, err1 := readBytes(file1, i, everySize)
		if err1 != nil {
			fmt.Printf("err1:%v \n", err1)
			break
		}
		b2, err2 := readBytes(file2, i, everySize)

		if err2 != nil {
			fmt.Printf("err2:%v \n", err2)
			break
		}

		if string(b1) != string(b2) {
			fmt.Printf("index read not equal, i:%d, \n b1:%v, \n b2:%v \n", i, b1, b2)
			break
		}
		fmt.Printf("index read is equal, i:%d \n", i)
	}
}
