/*
 *@author  chengkenli
 *@project starload
 *@package application
 *@file    application_auth
 *@date    2025/7/17 11:02
 */

package application

import "starload/util"

// 校验当前版本是否已经开启使用
func auther() util.DataMeta {
	for _, item := range util.MetaData {
		if util.Segment == item.Segment {
			return item
		}
	}
	return util.DataMeta{}
}
