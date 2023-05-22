package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

//func main() {
//	scriptPath := "gopython/createpyspark.py"
//	arg1 := "Hello"
//	arg2 := "World!"
//	//exec.Command()
//
//	cmd := exec.Command("python", scriptPath, arg1, arg2)
//	output, err := cmd.Output()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 处理Python脚本的输出
//	result := string(output)
//	log.Println(result)
//}

func main() {
	// 创建一个通道来接收信号
	signalCh := make(chan os.Signal, 1)

	// 告诉程序在收到 SIGINT 或 SIGTERM 信号时发送到 signalCh 通道
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	// 启动一个 goroutine 监听信号
	go func() {
		// 阻塞等待信号
		sig := <-signalCh
		fmt.Println("Received signal:", sig)
		_ = exec.Command("python", "gopython/stopsession.py")

		// 执行程序退出逻辑
		os.Exit(0)
	}()
	//运行你的程序逻辑
	//...
	scriptPath := "gopython/socketstreming.py"
	//arg1 := "Hello"
	//arg2 := "World!"
	//exec.Command()

	_ = exec.Command("python", scriptPath)
	//cmd := exec.Command("python", "gopython/stopsession.py")
	//output, err := cmd.Output()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 处理Python脚本的输出
	//result := string(output)
	//log.Println(result)
	////fmt.Printf("helloworld!%d", 111)

	// 程序逻辑执行完成后，等待信号
	<-signalCh
}
