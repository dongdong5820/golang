package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"time"
)

// 定义常量
const (
	IMG_PATH   = "./"
	IMG_URL    = "https://img.huxiucdn.com/event/202012/11/094122988093.jpg"
	TARGET_URL = "https://upload.xxx.com/xxx.do"
)

type DownLoadAndUploadMsgBo struct {
	Jsonrpc int
	Message string
	Id      string
}

func main() {
	//_, err := download(IMG_URL)
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//fileName := "a.jpg"
	params := map[string]string{
		"usage": "supplier-temp",
		"type":  "img",
	}
	/*
		cont, err := upload(fileName, TARGET_URL, params)
		if err != nil {
			panic(err)
		}

		m := UploadMsgBo{}
		err = json.Unmarshal(cont, &m)
		if err != nil {
			fmt.Printf("Unmarshal err: %+v\n", err)
		} else {
			fmt.Printf("resp: %+v\n", m)
		}
	*/
	//err := postFile(fileName, TARGET_URL, params)
	fileId, err := downLoadAndUploadImg(IMG_URL, TARGET_URL, params)
	if err != nil {
		panic(err)
	}
	fmt.Printf("fileId: %s\n", fileId)
}

//
// 下载远程图片，上传至upload项目
// imgUrl: 远程图片地址
// targetUrl: 上传地址
// params: 其他参数, 如 usage: supplier-temp, type:img
// fileId: 生成的文件ID
//
func downLoadAndUploadImg(imgUrl string, targetUrl string, params map[string]string) (fileId string, err error) {
	if imgUrl == "" {
		return
	}

	// 支持的图片格式
	validFileExt := map[string]string{
		"bmp":  "bmp",
		"png":  "png",
		"jpg":  "jpg",
		"gif":  "gif",
		"jpeg": "jpeg",
	}
	fileExt := path.Ext(imgUrl)
	if validFileExt[fileExt] == "" {
		fmt.Printf("img url:%s ,img ext err: %+v\n", imgUrl, fileExt)
		return
	}
	fileName := path.Base(imgUrl)

	// 下载图片
	res, err := http.Get(imgUrl)
	if err != nil {
		fmt.Printf("download file err: %+v\n", err)
		return
	}
	defer res.Body.Close()

	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	// 上传图片
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// 关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("file", fileName)
	if err != nil {
		fmt.Printf("writing to buffer err : %+v\n", err)
		return
	}

	// iocopy
	_, err = io.Copy(fileWriter, reader)
	if err != nil {
		fmt.Printf("io copy err: %+v\n", err)
		return
	}

	// 设置参数
	for key, val := range params {
		_ = bodyWriter.WriteField(key, val)
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		fmt.Printf("upload file err : %+v\n", err)
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read response err : %+v\n", err)
		return
	}

	fmt.Println(resp.Status)
	fmt.Println(string(b))

	m := DownLoadAndUploadMsgBo{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		fmt.Printf("unmarshar err : %+v\n", err)
		return
	}
	fileId = m.Id

	return
}

//
// 下载远程图片
//
func download(imgUrl string) (fileName string, err error) {
	fileName = path.Base(imgUrl)
	fmt.Printf("file name: %s\n", fileName)

	res, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("a error occurred!", err.Error())
		return
	}
	defer res.Body.Close()

	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32*1024)

	file, err := os.Create(IMG_PATH + fileName)
	if err != nil {
		panic(err)
	}

	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	written, _ := io.Copy(writer, reader)
	fmt.Printf("total length: %d\n", written)
	return
}

//
// 上传本地图片至远程服务器
//
func upload(fileName string, targetUrl string, params map[string]string) ([]byte, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	// 关键操作
	formFile, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return nil, err
	}

	// 打开文件
	file, err := os.Open(IMG_PATH + fileName)
	if err != nil {
		fmt.Printf("opening file err: %+v", err)
		return nil, err
	}
	defer file.Close()

	// ioCopy
	_, err = io.Copy(formFile, file)
	if err != nil {
		return nil, err
	}

	// 设置参数
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	// 构造request
	req, err := http.NewRequest("POST", targetUrl, body)
	if err != nil {
		return nil, err
	}
	// 添加header
	req.Header.Add("Content-type", writer.FormDataContentType())
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	// 接收响应
	httpClient := http.Client{Timeout: 3 * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status, resp.StatusCode)
	// 解析内容
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}

//
// 上传本地图片
//
func postFile(filename string, targetUrl string, params map[string]string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	//打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	// 设置参数
	for key, val := range params {
		_ = bodyWriter.WriteField(key, val)
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}
