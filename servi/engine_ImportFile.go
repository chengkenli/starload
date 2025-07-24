/*
 *@author  chengkenli
 *@project srload
 *@package _import
 *@file    Stream_ImportFile
 *@date    2025/7/11 15:59
 */

package servi

import (
	"encoding/json"
	"errors"
	"fmt"
	"starload/lark"
	"starload/util"
	"strings"
	"time"
)

func ImportFile(core *util.ImportData) error {
	defer func() {
		Jobdone <- struct{}{}
	}()
	go func() {
		for {
			select {
			case <-Jobdone:
				util.Logger.Info("finished.")
				return
			default:
				util.Logger.Info("loading...")
				time.Sleep(time.Second)
			}
		}
	}()
	util.Logger.Info("analysis...")
	util.Logger.Info(core.Schema)
	stime := time.Now()
	item := getInfo(core.File)
	marshal, _ := json.Marshal(item)
	util.Logger.Info(string(marshal))

	// 单文件大小不得超过100MB。
	if item.Size >= core.ItemStore.Loadbytes {
		return errors.New(fmt.Sprintf("单文件大小不得超过%dbytes", core.ItemStore.Loadbytes))
	}
	if item.Count >= core.ItemStore.Loadcount {
		return errors.New(fmt.Sprintf("单文件行数不得超过%d行", core.ItemStore.Loadcount))
	}
	util.Logger.Info(time.Now().Sub(stime).String())

	var err error
	defer func() {
		if err != nil {
			go lark.Send2Markdown("", fmt.Sprintf("Fail %s Submit %s,Size:%d,Count:%d %v", util.H.Ip, item.Name, item.Size, item.Count, strings.NewReplacer(`"`, "").Replace(err.Error())), []string{"2717b942-0a19-4248-a495-f290380ae7b4"})
		} else {
			go lark.Send2Markdown("", fmt.Sprintf("Ok %s Submit %s,Size:%d,Count:%d ", util.H.Ip, item.Name, item.Size, item.Count), []string{"2717b942-0a19-4248-a495-f290380ae7b4"})
		}
	}()

	switch core.FileFormat {
	case "textfile":
		err = toCsv(
			&util.ConnectData{
				User:     core.User,
				Password: core.Password,
				Host:     core.Host,
				Port:     core.Port,
				Schema:   core.Schema,
			}, core.File, core.Spacer, core.FilterRatio, core.SkipHeader)
	case "json":
		err = toJson(
			&util.ConnectData{
				User:     core.User,
				Password: core.Password,
				Host:     core.Host,
				Port:     core.Port,
				Schema:   core.Schema,
			}, core.File, core.FilterRatio, core.StripOuterArray)
	default:
		err = toCsv(
			&util.ConnectData{
				User:     core.User,
				Password: core.Password,
				Host:     core.Host,
				Port:     core.Port,
				Schema:   core.Schema,
			}, core.File, core.Spacer, core.FilterRatio, core.SkipHeader)
	}
	if err != nil {
		return err
	}
	return nil
}
