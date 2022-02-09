package middleware

import (
	"fmt"
	"runtime"
)

// InitRuntime ...
func InitRuntime() {
	// 设置程序使用的CPU最大数目
	cpuNum := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuNum)

	fmt.Printf("当前使用的cpu的数目为%v",cpuNum)
}