/*
 *@author  chengkenli
 *@project srload
 *@package app
 *@file    object
 *@date    2025/7/11 17:42
 */

package servi

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"starload/util"
)

var Jobdone = make(chan struct{})

type fileInfo struct {
	Name    string
	Size    int64
	ModTime string
	Model   string
	Count   int64
}

// 获取文件基本信息
func getInfo(filename string) fileInfo {
	util.Logger.Info(filename)
	// 获取文件基本信息
	info, err := os.Stat(filename)
	if err != nil {
		util.Logger.Error(err.Error())
	}
	file, err := os.Open(filename)
	if err != nil {
		util.Logger.Error(err.Error())
		return fileInfo{
			Name:    info.Name(),
			Size:    info.Size(),
			ModTime: info.ModTime().Format("2006-01-02 15:04:05"),
			Model:   "",
			Count:   0,
		}
	}
	// 只读取前512字节用于判断类型
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		util.Logger.Error(err.Error())
	}
	// 获取行数
	defer file.Close()
	buf := make([]byte, 32*1024*1024) // 32KB 缓冲区
	count := 0
	lineSep := []byte{'\n'}
	for {
		c, err := file.Read(buf)
		if err != nil && err != io.EOF {
			util.Logger.Error(err.Error())
			return fileInfo{
				Name:    info.Name(),
				Size:    info.Size(),
				ModTime: info.ModTime().Format("2006-01-02 15:04:05"),
				Model:   http.DetectContentType(buffer),
				Count:   int64(count),
			}
		}
		count += bytes.Count(buf[:c], lineSep)
		if err == io.EOF {
			break
		}
	}

	return fileInfo{
		Name:    info.Name(),
		Size:    info.Size(),
		ModTime: info.ModTime().Format("2006-01-02 15:04:05"),
		Model:   http.DetectContentType(buffer),
		Count:   int64(count),
	}
}
