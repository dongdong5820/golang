package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"ranlay.com/sql/learnKu/file"
	"strings"
)

//
// 目录描述
// /home/vagrant/code/gowork/netdisk/img/supplier-temp/54226/6642e/e42f7/bb68b/542266642ee42f7bb68bc43d.jpg
// /home/vagrant/code/gowork/images/public/company/G1011/G101176/logo/logo.jpg
// /home/vagrant/code/gowork/images/thumbnail/company/G1011/G101176/logo/logo_200x100.jpg  | logo_200x200.jpg
//
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

//
// 通过文件类型，文件用途，文件ID获取文件名
//
func getFileNameByFileId(fileType, use, fileId string) string {
	//parentPath := "/www/target/netdisk"
	parentPath := "/home/vagrant/code/gowork/netdisk"
	filePath := fmt.Sprintf("%s/%s/%s", fileType, use, getFileSavePath(fileId))

	return fmt.Sprintf("%s/%s/%s",
		parentPath,
		filePath,
		fileId+".jpg")
}

//
// upload文件保存路径
//
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

//
// 根据公司ID获取logo文件名
//
func getFileNameByCompanyId(companyId string) (string, error) {
	logoFileName := ""
	fileName := "logo.jpg"
	companyId = strings.ToUpper(companyId)
	//imageOrRoot := "/www/target/images/public"
	imageOrRoot := "/home/vagrant/code/gowork/images/public"
	filePath := fmt.Sprintf("%s/company/%s/%s/logo/",
		imageOrRoot,
		companyId[0:5],
		companyId)

	// 创建目录
	err := os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		return logoFileName, err
	}

	// 删除缩略图
	err = removeThumbnail(filePath, fileName)
	if err != nil {
		return logoFileName, err
	}
	logoFileName = fmt.Sprintf("%s%s", filePath, fileName)

	return logoFileName, nil
}

//
// 删除缩略图
//
func removeThumbnail(oriImagePath, fileName string) error {
	//imageOrRoot := "/www/target/images/public"
	imageOrRoot := "/home/vagrant/code/gowork/images/public"
	//imageThumbnaiRoot := "/www/target/images/thumbnail"
	imageThumbnailRoot := "/home/vagrant/code/gowork/images/thumbnail"
	filePath := strings.Replace(oriImagePath, imageOrRoot, imageThumbnailRoot, -1)
	i := strings.Index(fileName, ".")
	if i <= -1 {
		return errors.New("logo file type error")
	}
	fileName = fileName[:i]

	dir, err := ioutil.ReadDir(filePath)
	if err != nil {
		return err
	}
	for _, f := range dir {
		fName := f.Name()
		if strings.HasPrefix(fName, fileName+"_") {
			err = os.Remove(filePath + fName)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
