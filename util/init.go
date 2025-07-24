/*
 * Copyright (c) 2025 @chengkenli. All rights reserved.
 * Alternative more formal version for the second sentence:
 * "The applicable license is available at https://github.com/chengkenli."
 *
 * @Project: starload
 * @File: init
 * @Author: chengkenli
 * @Created: 2025/7/13 11:30
 * @Description:
 */

package util

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"net"
)

func init() {
	Software = app.New()
	//创建日志显示组件
	logDisplay := widget.NewLabel("")
	logDisplay.Wrapping = fyne.TextWrapOff // 禁止自动换行
	LogScroll = container.NewScroll(logDisplay)
	// 绑定日志内容
	logContent.AddListener(binding.NewDataListener(func() {
		text, _ := logContent.Get()
		logDisplay.SetText(text)
		LogScroll.ScrollToTop() // 改为滚动到顶部
	}))
}

func Init() {
	/*获取当前主机的IP*/
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		logger.Error(err.Error())
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				H.Ip = ipnet.IP.String()
			}
		}
	}
}
