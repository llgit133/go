package main

import "time"

func main() {

	// go 语言的多线程 - 协程 天然支持并发
	// 不建议锁，用管道通信 channel
	// 结果交替执行
	go testGo()
	for i := 0; i < 1000; i++ {
		println("main", i)
	}
	time.Sleep(time.Second * 3)
}

func testGo() {
	for i := 0; i < 1000; i++ {
		println("========>>>", i)
	}
}
