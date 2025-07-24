/*
 * Copyright (c) 2025 @chengkenli. All rights reserved.
 * Alternative more formal version for the second sentence:
 * "The applicable license is available at https://github.com/chengkenli."
 *
 * @Project: starload
 * @File: main
 * @Author: chengkenli
 * @Created: 2025/7/13 10:53
 * @Description:
 */

package main

import (
	"starload/application"
	_ "starload/init"
	"starload/util"
)

func main() {
	util.Init()
	application.App()
}
