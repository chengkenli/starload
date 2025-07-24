/*
 *@author  chengkenli
 *@project starload
 *@package application
 *@file    application_backend
 *@date    2025/7/17 13:48
 */

package application

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

// MacOSTheme 实现一个精美的 macOS 风格主题
type MacOSTheme struct{}

// Color 定义 macOS 风格颜色
func (MacOSTheme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return &color.NRGBA{R: 242, G: 242, B: 247, A: 255} // 系统背景色
	case theme.ColorNameForeground:
		return &color.NRGBA{R: 28, G: 28, B: 30, A: 255} // 主文本色
	case theme.ColorNamePrimary:
		return &color.NRGBA{R: 0, G: 122, B: 255, A: 255} // macOS 标志性蓝色
	case theme.ColorNameHover:
		return &color.NRGBA{R: 0, G: 122, B: 255, A: 30} // 悬停效果
	case theme.ColorNameFocus:
		return &color.NRGBA{R: 0, G: 122, B: 255, A: 102} // 聚焦光环
	case theme.ColorNameSelection:
		return &color.NRGBA{R: 204, G: 228, B: 255, A: 255} // 选中项背景
	case theme.ColorNameDisabled:
		return &color.NRGBA{R: 199, G: 199, B: 204, A: 255} // 禁用状态
	case theme.ColorNamePlaceHolder:
		return &color.NRGBA{R: 174, G: 174, B: 178, A: 255} // 占位文本
	case theme.ColorNameInputBackground:
		return &color.NRGBA{R: 255, G: 255, B: 255, A: 255} // 输入框背景
	case theme.ColorNameButton:
		return &color.NRGBA{R: 0, G: 122, B: 255, A: 255} // 按钮颜色
	case theme.ColorNameOverlayBackground:
		return &color.NRGBA{R: 249, G: 249, B: 249, A: 230} // 弹窗背景
	default:
		return theme.DefaultTheme().Color(name, theme.VariantLight)
	}
}

// Font 返回 macOS 风格字体
func (MacOSTheme) Font(style fyne.TextStyle) fyne.Resource {
	// 实际应用中应该替换为 San Francisco 字体文件
	if style.Monospace {
		return theme.DefaultTheme().Font(style)
	}
	return theme.DefaultTheme().Font(style)
}

// Icon 返回 macOS 风格图标
func (MacOSTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	// 可以替换为自定义的 macOS 风格图标
	return theme.DefaultTheme().Icon(name)
}

// Size 定义 macOS 风格尺寸
func (MacOSTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNamePadding:
		return 12 // 更大的内边距
	case theme.SizeNameInlineIcon:
		return 20 // 图标尺寸
	case theme.SizeNameScrollBar:
		return 10 // 更细的滚动条
	case theme.SizeNameScrollBarSmall:
		return 5
	case theme.SizeNameText:
		return 14 // 基础字体大小
	case theme.SizeNameHeadingText:
		return 20 // 标题大小
	case theme.SizeNameSubHeadingText:
		return 16
	case theme.SizeNameInputBorder:
		return 1.5 // 更细的边框
	default:
		return theme.DefaultTheme().Size(name)
	}
}
