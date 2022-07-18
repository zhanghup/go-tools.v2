package tools

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
	"xorm.io/xorm"
)

var __fileEngine *xorm.Engine
var ErrFileNotExist = errors.New("[uploader] 读取文件不存在")

func FileInfo(id string) (io.Reader, error) {
	uuids := strings.Split(id, "-")
	date := uuids[0][0:1]
	date = uuids[4][0:1] + date
	date = uuids[3][0:1] + date
	date = uuids[2][0:1] + date
	date = uuids[1][0:1] + date
	date = uuids[0][1:2] + date

	ftype := ""
	// 文件后缀解析
	if n := strings.LastIndex(id, "."); n > -1 {
		ftype = strings.Replace(id[n:], ".", "", -1)
	} else {
		return nil, ErrFileNotExist
	}
	path := fmt.Sprintf("./upload/%s/%s/%s/%s", date[:4], date[4:], ftype, id)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}

/*
	FileUploadIO 文件上传
	reader: 文件流
	name: 文件名，例如1.jpg
*/
func FileUploadIO(reader io.Reader, filename string) (string, error) {
	now := time.Now()
	date := fmt.Sprintf("%04d%02d", now.Year(), now.Month())

	ftype := ""
	// 文件后缀解析
	if n := strings.LastIndex(filename, "."); n > -1 {
		ftype = strings.Replace(filename[n:], ".", "", -1)
	} else {
		ftype = "stream"
	}

	uuids := strings.Split(UUID(), "-")
	for i := 0; i < 5; i++ {
		if i == 4 {
			uuids[i] = date[i:] + uuids[i]
		} else {
			uuids[i] = date[i:i+1] + uuids[i]
		}
	}
	for i, s := range date {
		idx := i % len(uuids)
		uuids[idx] = string(s) + uuids[idx]
	}

	id := strings.Join(uuids, "-")
	if ftype != "" {
		id += "." + ftype
	}

	// 资源文件路径定义
	path := ""
	{
		dir := fmt.Sprintf("upload/%d/%02d/%s", now.Year(), now.Month(), ftype)
		if _, err := os.Stat(dir); err != nil {
			if os.IsNotExist(err) {
				err = os.MkdirAll(dir, os.ModePerm)
				if err != nil {
					return "", err
				}
			} else {
				return "", err
			}
		}

		path = fmt.Sprintf("%s/%s", dir, id)
	}

	f, err := os.Create(path)
	if err != nil {
		return "", err
	}

	// 写入文件数据
	_, err = io.Copy(f, reader)
	if err != nil {
		return "", err
	}

	return id, nil
}
