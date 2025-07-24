/*
 *@author  chengkenli
 *@project starload
 *@package servi
 *@file    engine_StreamData2xlsx
 *@date    2025/7/16 15:08
 */

package servi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"net/http"
	"os"
	"starload/util"
	"strings"
	"time"
)

// toCsv 慢查询落表
func toXlsx(citem *util.ConnectData, dataFile, spacer, filterRatio string) error {
	util.Logger.Info("textfile.")
	stime := time.Now()
	defer func() {
		util.Logger.Info("Total Time:" + time.Now().Sub(stime).String())
	}()
	if len(citem.Schema) == 0 {
		return errors.New("schema is nil")
	}
	tb := strings.Split(citem.Schema, ".")
	util.Logger.Info(dataFile)

	file, err := os.Open(dataFile)
	if err != nil {
		util.Logger.Error(err)
		return err
	}
	defer file.Close()

	//创建Resty客户端
	client := resty.New().SetDisableWarn(true)
	//发送POST请求并处理响应
	uri := fmt.Sprintf("http://%s:%d/api/%s/%s/_stream_load", citem.Host, citem.Port, tb[0], tb[1])
	response, err := client.SetRedirectPolicy(resty.RedirectPolicyFunc(func(req *http.Request, via []*http.Request) error {
		// 这里你可以根据需要添加自定义逻辑，比如保留headers等
		for key, values := range via[0].Header {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}
		// 如果想要完全信任所有重定向，只需返回nil
		return nil
	})).R().
		SetHeaders(map[string]string{
			"label":            "starload-" + uuid.New().String(), /*label*/
			"Content-Type":     "application/octet-stream",        /*指定文本类型*/
			"Expect":           "100-continue",                    /*在服务器拒绝导入作业请求的情况下，避免不必要的数据传输，减少不必要的资源开销。*/
			"max_filter_ratio": filterRatio,                       /*指定导入作业的最大容错率 取值范围：0~1*/
			"column_separator": spacer,                            /*用于指定源数据文件中的列分隔符*/
		}).SetBasicAuth(citem.User, citem.Password).
		SetBody(file).
		SetContentLength(true).
		Put(uri)
	if err != nil {
		util.Logger.Error(err.Error())
		return err
	}
	var result util.JesultData
	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		util.Logger.Error(err.Error())
		return err
	}
	if result.Status == "Success" {
		util.Logger.Info(uri)
		util.Logger.Info(string(response.Body()))
		util.Logger.Info("导入成功.")
		return nil
	} else {
		util.Logger.Error(uri)
		util.Logger.Error(string(response.Body()))
		util.Logger.Error("导入失败.")
		return errors.New("load failed")
	}
}
