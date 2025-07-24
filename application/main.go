/*
 * Copyright (c) 2025 @chengkenli. All rights reserved.
 * Alternative more formal version for the second sentence:
 * "The applicable license is available at https://github.com/chengkenli."
 *
 * @Project: starload
 * @File: interface
 * @Author: chengkenli
 * @Created: 2025/7/13 11:04
 * @Description:
 */

package application

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/patrickmn/go-cache"
	_ "net/http/pprof"
	"starload/servi"
	"starload/util"
	"strings"
	"time"
)

var (
	itemCache = cache.New(24*time.Hour, 24*time.Hour)
	itemStore util.DataMeta
)

func App() {
	//go func() { log.Println(http.ListenAndServe("localhost:6060", nil)) }()
	util.Software.Settings().SetTheme(&MacOSTheme{})
	//util.Software.Settings().SetTheme(theme.LightTheme())
	util.Software.SetIcon(icos())
	window := util.Software.NewWindow(fmt.Sprintf("Starload%0.1f", util.Segment))
	window.Resize(fyne.NewSize(900, 900))

	go dialogs(window)
	showapp(window)

	window.ShowAndRun()
}

func starmain(window fyne.Window) {

	envOptions := strings.Split(itemStore.Envoptions, ",")
	envSelector := widget.NewSelect(envOptions, func(selected string) {
		util.Logger.Info(fmt.Sprintf("load environment:[%s]", selected))
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
	spacerOptions := strings.Split(itemStore.Sproptions, "|")
	spacerSelector := NewSpacerSelector(spacerOptions, func(selected string) {
		util.Logger.Info(fmt.Sprintf("load spacers:[%s]", selected))
		// 这里可以添加处理间隔符选择的逻辑
	})
	spacerSelector.SetSelected(spacerOptions[0]) // 默认选择第一个选项

	// 新增文件格式选择组件
	fileFormatOptions := strings.Split(itemStore.Txtoptions, ",")
	fileFormatSelector := NewSpacerSelector(fileFormatOptions, func(selected string) {
		util.Logger.Info(fmt.Sprintf("load format:[%s]", selected))
		// 这里可以添加处理间隔符选择的逻辑
	})
	fileFormatSelector.SetSelected(spacerOptions[0]) // 默认选择第一个选项

	// 新增指定json格式是否数组选择组件
	stripOuterArrayOptions := []string{"true", "false"}
	stripOuterArraySelector := NewSpacerSelector(stripOuterArrayOptions, func(selected string) {
		util.Logger.Info(fmt.Sprintf("load strip_outer_array:[%s]", selected))
		// 这里可以添加处理间隔符选择的逻辑
	})
	stripOuterArraySelector.SetSelected(spacerOptions[1]) // 默认选择第一个选项

	// 数据绑定
	filePath := binding.NewString()
	filePath.Set("")

	var filename, schema string
	// 文件选择组件
	fileSelectButton := widget.NewButtonWithIcon("选择数据文件(csv,txt,json)", theme.FileIcon(), func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err == nil && reader != nil {
				filePath.Set(reader.URI().Path())
				//util.Logger("已选择文件: " + reader.URI().Path())
				util.Logger.Info(fmt.Sprintf("load datafile:[%s]", reader.URI().Path()))
				filename = reader.URI().Path()
			}
		}, window)
	})
	filePathLabel := widget.NewLabelWithData(filePath)

	// 新增容错率输入框
	faultToleranceEntry := widget.NewEntry()
	faultToleranceEntry.SetText("0") // 设置默认值为0
	faultToleranceEntry.SetPlaceHolder("最大容错率(0-1)")

	// 新增跳过行头
	skipheaderEntry := widget.NewEntry()
	skipheaderEntry.SetText("0") // 设置默认值为0
	skipheaderEntry.SetPlaceHolder("用于指定跳过csv文件最开头的几行数据")

	// 新增数据预览按钮
	widget.NewButton("预览", func() {

	}).Importance = widget.MediumImportance

	previewButton := widget.NewButtonWithIcon("预览", theme.ViewFullScreenIcon(), func() {
		if path, _ := filePath.Get(); path == "" {
			//util.Logger("请先选择文件再启动")
			util.Logger.Warn("please select a datafile")
			return
		}
		ft, err := GetFileType(filename)
		if err != nil {
			util.Logger.Error(err.Error())
			return
		}
		util.Logger.Info(ft)
		if ft == "xlsx" {
			util.Logger.Info("暂不支持xlsx格式文件")
			return
		}
		// 这里实现数据预览的逻辑
		util.Logger.Info("preview top 10")
		val, err := firstlines(filename, 10)
		if err != nil {
			util.Logger.Error(err.Error())
			return
		}
		util.Logger.Info("\n" + strings.Join(val, "\n"))
	})

	// 启动按钮
	startButton := widget.NewButtonWithIcon("启动", theme.MediaPlayIcon(), func() {
		schema = innerTableEntry.Text
		spacerVal := spacerSelector.Selected() // 获取当前选择的间隔符
		faultTolerance := faultToleranceEntry.Text
		skipVal := skipheaderEntry.Text
		fileformatVal := fileFormatSelector.Selected()
		stripOuterArrayVal := stripOuterArraySelector.Selected()

		if path, _ := filePath.Get(); path == "" {
			//util.Logger("请先选择文件再启动")
			util.Logger.Warn("please select a datafile")
			return
		}
		if schema == "" {
			util.Logger.Warn("please fill in the table name")
			return
		}

		// 判断文件格式
		ft, err := GetFileType(filename)
		if err != nil {
			util.Logger.Error(err.Error())
			return
		}
		util.Logger.Info(ft)
		if ft == "xlsx" {
			util.Logger.Info("暂不支持xlsx格式文件")
			return
		}

		// 在日志中输出选择的间隔符
		util.Logger.Info(fmt.Sprintf("spacers: %s", spacerVal))
		util.Logger.Info(fmt.Sprintf("max filter ratio: %s", faultTolerance))

		valu, oku := itemCache.Get("username")
		valp, okp := itemCache.Get("password")

		// 账号密码已经存在
		if oku && okp {
			util.Logger.Info("the account password is loaded from the cache")
			go func() {
				err := servi.ImportFile(
					&util.ImportData{
						User:            valu.(string),
						Password:        valp.(string),
						Host:            envSet(envSelector.Selected),
						Port:            8030,
						Schema:          replacer(schema),
						File:            filename,
						FileFormat:      fileformatVal,
						Spacer:          replacer(spacerVal),
						FilterRatio:     replacer(faultTolerance),
						StripOuterArray: stripOuterArrayVal,
						ItemStore:       itemStore,
						SkipHeader:      skipVal,
					})
				if err != nil {
					util.Logger.Error(err.Error())
					return
				}
			}()
			return
		}
		// 如果账号密码不对，好吧，展开组件让用户输入进行校验
		// 创建输入组件
		accountEntry := widget.NewEntry()
		passwordEntry := widget.NewPasswordEntry()
		// 创建对话框内容
		loginForm := widget.NewForm(
			widget.NewFormItem("username", accountEntry),
			widget.NewFormItem("password", passwordEntry),
		)
		content := container.NewVBox(
			widget.NewLabel("please enter your StarRocks login credentials(native/ldap):"),
			loginForm,
		)
		// 显示对话框
		dialog.ShowCustomConfirm("authentication", "login", "cancel",
			content,
			func(confirmed bool) {
				if !confirmed {
					//util.Logger("用户取消登录")
					util.Logger.Info("cancels the login")
					return
				}
				account := accountEntry.Text
				password := passwordEntry.Text
				if account == "" || password == "" {
					util.Logger.Warn("the account password cannot be empty")
					//util.Logger("错误：账号密码不能为空")
					return
				}
				fileSelectButton.Disable()
				//util.Logger(fmt.Sprintf("正在使用账号 %s 启动程序...", account))
				util.Logger.Info("login...")
				// 实际启动逻辑
				go func() {
					util.Logger.Info("on going...")
					// 继续其他启动逻辑
					util.Logger.Info(schema)
					// stream load
					go func() {
						err := servi.ImportFile(
							&util.ImportData{
								User:            account,
								Password:        password,
								Host:            envSet(envSelector.Selected),
								Port:            8030,
								Schema:          replacer(schema),
								File:            filename,
								FileFormat:      fileformatVal,
								Spacer:          replacer(spacerVal),
								FilterRatio:     replacer(faultTolerance),
								StripOuterArray: stripOuterArrayVal,
								ItemStore:       itemStore,
								SkipHeader:      skipVal,
							})
						if err != nil {
							util.Logger.Error(err.Error())
							return
						}
						// 账号密码先放进缓存中
						itemCache.Set("username", account, cache.DefaultExpiration)
						itemCache.Set("password", password, cache.DefaultExpiration)

					}()
				}()
			},
			window,
		)
	})

	// 创建取消按钮
	cancelButton := widget.NewButtonWithIcon("取消", theme.MediaStopIcon(), func() {
		if path, _ := filePath.Get(); path == "" {
			return
		}
		fileSelectButton.Enable()
		util.Logger.Info("cancel")
		// 清空库表填写框
		innerTableEntry.SetText("")
		// 清空已经选择的文件
		filePath.Set("")
		// 清空账号和密码
		itemCache.DeleteExpired()
		// 这里可以添加实际的取消逻辑
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
		container.NewHBox( // 新增用于指定跳过 CSV 文件最开头的几行数据
			widget.NewIcon(theme.WarningIcon()),
			widget.NewLabel("略过首行:"),
			skipheaderEntry,
			labelHint("(非必填，默认0，用于指定跳过csv文件最开头的几行数据)"),
		),
		container.NewHBox(
			widget.NewIcon(theme.SearchIcon()),
			widget.NewLabel("文件路径:"),
			fileSelectButton,
			filePathLabel,
			previewButton,
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
	util.Logger.Info("ready.")
}

func replacer(val string) string {
	return strings.NewReplacer(" ", "").Replace(val)
}

func labelHint(val string) *widget.RichText {
	tableHint := widget.NewRichTextWithText(val)
	tableHint.Segments[0].(*widget.TextSegment).Style.SizeName = theme.SizeNameCaptionText
	return tableHint
}
