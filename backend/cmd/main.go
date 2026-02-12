package main

import (
	"fmt"
	"time"
)

func main() {

	// 强制设置时区
	time.Local = time.FixedZone("CST", 8*3600)

	// 【关键】加上这行打印！启动时必须看到它！
	fmt.Println("----------------------------------------")
	fmt.Println(">>> 正在启动... 时区强制修正模式: ON")
	fmt.Println(">>> 当前系统时间:", time.Now())
	fmt.Println("----------------------------------------")

	composer()

}
