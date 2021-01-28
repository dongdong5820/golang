package main

import (
	"fmt"
	"os"
	"ranlay.com/sql/learnKu/file"
	"strings"
)

func main() {
	fmt.Println("文件拷贝")
	companyId := "G101176"
	fileId := "542266642ee42f7bb68bc43d"
	dstName, err := getFileNameByCompanyId(companyId)
	if err != nil {
		fmt.Printf("err: %+v\n", err)
	}
	srcName := getFileNameByFileId("img", "supplier-temp", fileId)

	fmt.Printf("dstName: %s\nsrcName: %s\n", dstName, srcName)
	//dstName := "/home/vagrant/code/gowork/images/p1.jpg"
	//srcName := "/home/vagrant/code/gowork/netdisk/p1.jpg"
	num, err := file.CopyFile(dstName, srcName)
	if err != nil {
		fmt.Printf("err: %+v\n", err)
	}
	fmt.Printf("num: %d\n", num)
}

func getFileNameByFileId(fileType, use, fileId string) string {
	//parentPath := "/www/target/netdisk/"
	parentPath := "/home/vagrant/code/gowork/netdisk/"
	filePath := fmt.Sprintf("%s/%s/%s", fileType, use, getFileSavePath(fileId))

	return fmt.Sprintf("%s%s/%s",
		parentPath,
		filePath,
		fileId+".jpg")
}

func getFileSavePath(fileId string) string {
	if len(fileId) <= 20 {
		return fileId
	}

	return fmt.Sprintf("%s/%s/%s/%s",
		fileId[0:5],
		fileId[5:10],
		fileId[10:15],
		fileId[15:20])
}

func getFileNameByCompanyId(companyId string) (string, error) {
	fileName := ""
	companyId = strings.ToUpper(companyId)
	//basePath := "/www/target/images/public/company/"
	basePath := "/home/vagrant/code/gowork/images/public/company/"
	filePath := fmt.Sprintf("%s%s/%s/logo/",
		basePath,
		companyId[0:5],
		companyId)
	// 创建目录
	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		return fileName, err
	}
	fileName = fmt.Sprintf("%s/logo.jpg", filePath)

	return fileName, nil
}
