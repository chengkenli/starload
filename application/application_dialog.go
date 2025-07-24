/*
 * Copyright (c) 2025 @chengkenli. All rights reserved.
 * Alternative more formal version for the second sentence:
 * "The applicable license is available at https://github.com/chengkenli."
 *
 * @Project: starload
 * @File: dialog
 * @Author: chengkenli
 * @Created: 2025/7/13 20:21
 * @Description:
 */

package application

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"os"
)

func dialogs(window fyne.Window) {
	var labelText string
	var exit bool
	itemStore = auther()
	if itemStore.Segment <= 0 {
		exit = true
		labelText = "版本无法使用，请更新版本！"
	} else {
		labelText = itemStore.Dialog
	}
	// 创建声明对话框
	declarationDialog := dialog.NewCustom(
		"使用声明",
		"同意",
		widget.NewLabel(labelText),
		window,
	)
	// 设置对话框大小
	declarationDialog.Resize(fyne.NewSize(400, 300))

	// 用户点击按钮后的处理
	declarationDialog.SetOnClosed(func() {
		starmain(window)
		if exit {
			os.Exit(-1)
		}
	})
	// 显示声明对话框
	declarationDialog.Show()
}
