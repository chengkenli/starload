/*
 * Copyright (c) 2025 @chengkenli. All rights reserved.
 * Alternative more formal version for the second sentence:
 * "The applicable license is available at https://github.com/chengkenli."
 *
 * @Project: starload
 * @File: init
 * @Author: chengkenli
 * @Created: 2025/7/13 10:54
 * @Description:
 */

package init

import (
	"fmt"
	"os"
	"starload/conn"
	"starload/util"
)

func init() {
	os.Setenv("FYNE_SCALE", "1.0")              // 强制使用1.0缩放
	os.Setenv("FYNE_DISABLE_OPENGL_CHECK", "1") // 跳过GPU检测

	_init()
}

func _init() {
	connect, err := conn.ConnectMySQL()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	condb := "chengken.starrocks_starload_cf"
	r := connect.Raw(fmt.Sprintf("select * from %s where state >= 1", condb)).Scan(&util.MetaData)
	if r.Error != nil {
		fmt.Println(r.Error.Error())
		return
	}
}
