/*
 *@author  chengkenli
 *@project starload
 *@package tools
 *@file    object
 *@date    2025/7/17 16:11
 */

package tools

import (
	"fmt"
	"gorm.io/gorm"
	"starload/util"
	"strings"
)

func Cancelload(db *gorm.DB, schema, label string) {
	sign := strings.Split(strings.NewReplacer(" ", "").Replace(schema), ".")
	stmt := fmt.Sprintf("show load from %s where label='%s'", sign[0], label)
	util.Logger.Info(stmt)
	r := db.Exec(stmt)
	if r.Error != nil {
		util.Logger.Error(r.Error.Error())
		return
	}
}
