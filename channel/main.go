package main

import "fmt"

//func main() {
//	wg.Add(3)
//	go Cat()
//	go Dog()
//	go Pig()
//
//	// 初始化，让 pig 先开始
//	pig <- struct{}{}
//
//	wg.Wait()
//	fmt.Println("打印结束")
//}

func main() {
	numWg.Add(3)
	go num1()
	go num2()
	go num3()
	// 启动
	c3 <- struct{}{}
	numWg.Wait()
	fmt.Print("打印结束")
}
