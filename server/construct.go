package server

import (
	"fmt"

	"zhihu.com/m/middleware"
)

func construct() {
	middleware.InitRuntime()
	middleware.InitConfig()
	middleware.InitLog("/tmp/zhihu", 32)
	middleware.LogBuesiness(map[string]interface{}{
		"test": "json",
	}, "test")
	fmt.Println(middleware.GetConfig("test.url"))
}
