/*
 * Copyright (c) 2025 @chengkenli. All rights reserved.
 * Alternative more formal version for the second sentence:
 * "The applicable license is available at https://github.com/chengkenli."
 *
 * @Project: starload
 * @File: logger
 * @Author: chengkenli
 * @Created: 2025/7/13 10:55
 * @Description:
 */

package util

import (
	"bufio"
	"fmt"
	"fyne.io/fyne/v2/data/binding"
	"os"
	"strings"
	"time"
)

type LogLevel string

const (
	INFO  LogLevel = "INFO"
	WARN  LogLevel = "WARN"
	ERROR LogLevel = "ERROR"
)

type Logs struct{}

var (
	logger = &Logs{}
)

func GetLogger() *Logs {
	return logger
}

func binglog(logLine string) {
	current, _ := logContent.Get()
	if current == "" {
		current = logLine
	} else {
		current = logLine + current // 新日志置顶
	}
	// 限制200行日志
	if lines := strings.Split(current, "\n"); len(lines) > 200 {
		current = strings.Join(lines[:200], "\n")
	}
	logContent.Set(current)
}

// 核心日志方法
func (l *Logs) log(level LogLevel, msg interface{}) {
	logLine := fmt.Sprintf("%s %v\n", fmt.Sprintf("[%s %s]", time.Now().Format("2006-01-02 15:04:05"), level), msg)
	go writeFile("starload.log", logLine)
	binglog(logLine)
}

func (l *Logs) Info(msg interface{})  { l.log(INFO, msg) }
func (l *Logs) Warn(msg interface{})  { l.log(WARN, msg) }
func (l *Logs) Error(msg interface{}) { l.log(ERROR, msg) }

func (l *Logs) GetLogContent() binding.String {
	return logContent
}

// WriteFile 文件落地
func writeFile(fname, msg string) {
	fileHandle, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer fileHandle.Close()
	// NewWriter 默认缓冲区大小是 4096
	// 需要使用自定义缓冲区的writer 使用 NewWriterSize()方法
	buf := bufio.NewWriterSize(fileHandle, len(msg))

	buf.WriteString(msg)

	err = buf.Flush()
	if err != nil {
		return
	}
}
