/*
 *@author  chengkenli
 *@project starload
 *@package tools
 *@file    object
 *@date    2025/7/14 14:32
 */

package application

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/xuri/excelize/v2"
	"net/http"
	"os"
	"starload/util"
	"strings"
	"time"
)

func Xlsx2Csv(xlsxPath, csvPath string) error {
	done := make(chan struct{})
	defer func() {
		done <- struct{}{}
	}()
	go func() {
		for {
			select {
			case <-done:
				util.Logger.Info("done.")
				return
			default:
				util.Logger.Info("xlsx2csv converting...")
				time.Sleep(time.Second)
			}
		}
	}()
	// 1. 打开 XLSX 文件
	f, err := excelize.OpenFile(xlsxPath)
	if err != nil {
		return fmt.Errorf("failed to open XLSX file: %v", err)
	}
	defer f.Close()
	// 2. 创建 CSV 文件
	csvFile, err := os.Create(csvPath)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %v", err)
	}
	defer csvFile.Close()
	csvFile.WriteString("\xEF\xBB\xBF") // UTF-8 BOM
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()
	// 3. 获取第一个工作表名
	sheetName := f.GetSheetName(0)
	// 4. 读取所有行数据
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return fmt.Errorf("failed to get rows: %v", err)
	}
	// 5. 写入 CSV
	for _, row := range rows {
		// 处理空单元格（Excelize 会跳过空单元格，导致行长度不一致）
		fullRow := make([]string, 0)
		for _, cell := range row {
			fullRow = append(fullRow, strings.TrimSpace(cell))
		}
		if err := writer.Write(fullRow); err != nil {
			return fmt.Errorf("failed to write CSV row: %v", err)
		}
	}
	return nil
}

func GetFileType(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 只读取文件前几个字节用于判断
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}
	// 使用http.DetectContentType作为基础检测
	mimeType := http.DetectContentType(buffer)
	// 更精确的检测
	switch {
	case bytes.HasPrefix(buffer, []byte{0x50, 0x4B, 0x03, 0x04}):
		// ZIP格式开头，可能是XLSX/DOCX等
		return "xlsx", nil
	case strings.Contains(mimeType, "text/plain"):
		return "txt", nil
	case strings.Contains(mimeType, "text/csv"):
		return "csv", nil
	default:
		return "unknown", nil
	}
}
