package server

import (
	"fmt"

	"zhihu.com/m/middleware"
)

func construct() {
	middleware.InitRuntime()
	middleware.InitConfig()
 	fmt.Println(middleware.GetConfig("test.url"))
}