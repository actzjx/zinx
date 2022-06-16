package main

import "zinx/znet"

func main() {
	// 1 使用zinx的api，创建一个server句柄
	s := znet.NewServer("[zinx v0.1]")
	// 2 启动sever
	s.Serve()
}
