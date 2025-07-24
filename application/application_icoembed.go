/*
 * Copyright (c) 2025 @chengkenli. All rights reserved.
 * Alternative more formal version for the second sentence:
 * "The applicable license is available at https://github.com/chengkenli."
 *
 * @Project: starload
 * @File: embed
 * @Author: chengkenli
 * @Created: 2025/7/13 21:15
 * @Description:
 */

package application

import (
	_ "embed"
	"fyne.io/fyne/v2"
)

//go:embed app.ico
var iconBytes []byte // 嵌入的二进制数据

func icos() *fyne.StaticResource {
	// 从内嵌数据加载图标
	icon := &fyne.StaticResource{
		StaticName:    "app.ico",
		StaticContent: iconBytes,
	}
	return icon
}
