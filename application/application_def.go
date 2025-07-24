/*
 * Copyright (c) 2025 @chengkenli. All rights reserved.
 * Alternative more formal version for the second sentence:
 * "The applicable license is available at https://github.com/chengkenli."
 *
 * @Project: starload
 * @File: def
 * @Author: chengkenli
 * @Created: 2025/7/13 14:21
 * @Description:
 */

package application

import (
	"bufio"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"io"
	"os"
	"starload/util"
)

type SizedEntry struct {
	widget.Entry
	customSize fyne.Size
}

func NewSizedEntry(width, height float32) *SizedEntry {
	e := &SizedEntry{customSize: fyne.NewSize(width, height)}
	e.ExtendBaseWidget(e)
	return e
}

func (e *SizedEntry) MinSize() fyne.Size {
	return e.customSize
}

func showapp(window fyne.Window) {
	envOptions := []string{"adhoc"}
	envSelector := widget.NewSelect(envOptions, func(selected string) {
		// 这里可以添加处理环境选择的逻辑
	})
	envSelector.SetSelected(envOptions[0]) // 默认选择 adhoc
	envSelectorWrapper := container.NewGridWrap(
		fyne.NewSize(372, 30), // 强制设置宽度 372，高度 30
		envSelector,
	)

	innerTableEntry := NewSizedEntry(372, 30)
	innerTableEntry.SetPlaceHolder("请填写StarRocks库表名称")
	innerTableEntry.SetText("")

	// 新增间隔符选择组件
	spacerOptions := []string{",", ";"}
	spacerSelector := NewSpacerSelector(spacerOptions, func(selected string) {
		// 这里可以添加处理间隔符选择的逻辑
	})
	spacerSelector.SetSelected(spacerOptions[0]) // 默认选择第一个选项

	// 新增文件格式选择组件
	fileFormatOptions := []string{"textfile", "json"}
	fileFormatSelector := NewSpacerSelector(fileFormatOptions, func(selected string) {
		// 这里可以添加处理间隔符选择的逻辑
	})
	fileFormatSelector.SetSelected(fileFormatOptions[0]) // 默认选择第一个选项

	// 新增指定json格式是否数组选择组件
	stripOuterArrayOptions := []string{"true", "false"}
	stripOuterArraySelector := NewSpacerSelector(stripOuterArrayOptions, func(selected string) {
		// 这里可以添加处理间隔符选择的逻辑
	})
	stripOuterArraySelector.SetSelected(stripOuterArrayOptions[1]) // 默认选择第一个选项

	// 数据绑定
	filePath := binding.NewString()
	filePath.Set("")

	// 文件选择组件
	fileSelectButton := widget.NewButtonWithIcon("选择本地数据文件(csv,txt,json)", theme.FileIcon(), func() {
	})
	filePathLabel := widget.NewLabelWithData(filePath)

	// 新增容错率输入框
	faultToleranceEntry := widget.NewEntry()
	faultToleranceEntry.SetText("0") // 设置默认值为0
	faultToleranceEntry.SetPlaceHolder("最大容错率(0-1)")

	// 启动按钮
	startButton := widget.NewButtonWithIcon("启动", theme.MediaPlayIcon(), func() {
	})

	// 创建取消按钮
	cancelButton := widget.NewButtonWithIcon("取消", theme.MediaStopIcon(), func() {
	})

	// 布局调整
	fileSelectionRow := container.NewVBox(
		container.NewHBox(
			widget.NewIcon(theme.ComputerIcon()),
			widget.NewLabel("运行环境:"),
			envSelectorWrapper,
			labelHint("(默认：adhoc集群)"),
		),
		container.NewHBox(
			widget.NewIcon(theme.DesktopIcon()),
			widget.NewLabel("导入库表:"),
			innerTableEntry,
			labelHint("(必填，库表需要存在的，并且是有写入权限的)"),
		),
		container.NewHBox( // 新增的间隔符选择行
			widget.NewIcon(theme.ListIcon()),
			widget.NewLabel("文本格式:"),
			fileFormatSelector,
			labelHint("(非必填，默认textfile)"),
		),
		container.NewHBox( // 新增的间隔符选择行
			widget.NewIcon(theme.InfoIcon()),
			widget.NewLabel("Json数组:"),
			stripOuterArraySelector,
			labelHint("(非必填，默认false，选择json格式的文件时可用)"),
		),
		container.NewHBox( // 新增的间隔符选择行
			widget.NewIcon(theme.FileTextIcon()),
			widget.NewLabel("字段间隔:"),
			spacerSelector,
			labelHint("(非必填，默认逗号',')"),
		),
		container.NewHBox( // 新增的容错率行
			widget.NewIcon(theme.WarningIcon()),
			widget.NewLabel("容错极值:"),
			faultToleranceEntry,
			labelHint("(非必填，默认0，最大容错率，范围0~1)"),
		),
		container.NewHBox(
			widget.NewIcon(theme.SearchIcon()),
			widget.NewLabel("文件路径:"),
			fileSelectButton,
			filePathLabel,
			startButton,
			cancelButton,
		),
	)

	content := container.NewBorder(
		container.NewVBox(
			fileSelectionRow,
			widget.NewSeparator(),
			widget.NewLabel("操作日志:"),
		),
		nil, nil, nil,
		util.LogScroll,
	)

	window.SetContent(NewWatermarkContainer(content))
}

// SpacerSelector 自定义间隔符选择组件
type SpacerSelector struct {
	widget.BaseWidget
	options  []string
	selected string
	onChange func(string)
}

// NewSpacerSelector 创建新实例，限制选项为1-4个
func NewSpacerSelector(options []string, onChange func(string)) *SpacerSelector {
	if len(options) < 1 || len(options) > 4 {
		util.Logger.Warn("SpacerSelector only supports 1-4 options")
	}
	s := &SpacerSelector{
		options:  options,
		onChange: onChange,
	}
	s.ExtendBaseWidget(s)
	return s
}

// CreateRenderer 实现 WidgetRenderer 接口
func (s *SpacerSelector) CreateRenderer() fyne.WidgetRenderer {
	radio := widget.NewRadioGroup(s.options, func(selected string) {
		s.selected = selected
		if s.onChange != nil {
			s.onChange(selected)
		}
	})
	radio.Horizontal = true // 水平排列选项
	return widget.NewSimpleRenderer(radio)
}

// Selected 获取当前选中的值
func (s *SpacerSelector) Selected() string {
	return s.selected
}

// SetSelected 设置选中项
func (s *SpacerSelector) SetSelected(option string) {
	s.selected = option
	s.Refresh()
}

func envSet(val string) string {
	var serverHost string
	switch val {
	case "StarRocks(Tencent Cloud) SR-ADHOC":
		serverHost = util.StarRocksServer
	case "StarRocks(Tencent Cloud) SR-QA":
		serverHost = util.StarRocksServerQa
	}
	return serverHost
}

func firstlines(filePath string, n int) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	lines := make([]string, 0, n)

	for i := 0; i < n; i++ {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		lines = append(lines, string(line))
	}
	return lines, nil
}
