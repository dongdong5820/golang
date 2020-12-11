package main

import (
	"encoding/json"
	"fmt"
	"path"
	"strings"
)

type UploadMsgBo struct {
	Jsonrpc int
	Message string
	Id      string
}

func main() {
	str := `{"jsonrpc":2,"message":"文件上传成功","id":"5fd31538a1bbe92fcc650aa2"}`

	b := []byte(str)
	m := UploadMsgBo{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Printf("unmarshal err: %+v\n", err)
	} else {
		fmt.Printf("fileId: %s\n", m.Id)
	}

	//  获取扩展名
	//imgUrl := "https://prod-czsaas.oss-cn-hangzhou.aliyuncs.com/ycg/cz_5b6d4b0486db4/common/7b803682_1878_11e9_946d_7cd30aeb75e4.jpg"
	imgUrl := "https://prod-czsupplier.oss-cn-hangzhou.aliyuncs.com/SaaS_CZ/cz_5c05e071c7a46/15548816544535/%E5%9B%BD%E5%8D%8E%E8%B5%84%E8%B4%A8.pdf"
	fileExt := path.Ext(imgUrl)
	fileExt = strings.ToLower(fileExt[1:])
	validFileExt := map[string]string{
		"bmp":  "bmp",
		"png":  "png",
		"jpg":  "jpg",
		"gif":  "gif",
		"jpeg": "jpeg",
	}
	if validFileExt[fileExt] == "" {
		fmt.Printf("invalid file ext: %s\n", fileExt)
	} else {
		fmt.Printf("file ext: %s\n", fileExt)
	}
}
