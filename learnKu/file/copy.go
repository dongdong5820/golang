package file

import (
	"io"
	"os"
)

//
// 拷贝文件
//
func CopyFile(dstName, srcName string) (wirtten int64, err error) {
	srcFile, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}
