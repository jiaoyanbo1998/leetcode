package main

import (
	"fmt"
	"sync"
)

var (
	cat = make(chan struct{})
	dog = make(chan struct{})
	pig = make(chan struct{})
	wg  *sync.WaitGroup
)

func Cat() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		<-pig // 等待 pig 的信号
		fmt.Print("cat-")
		cat <- struct{}{} // 发送信号给 dog
	}
}

func Dog() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		<-cat // 等待 cat 的信号
		fmt.Print("dog-")
		dog <- struct{}{} // 发送信号给 pig
	}
}

func Pig() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		<-dog // 等待 dog 的信号
		fmt.Println("pig")
		if i < 99 { // 避免最后一次发送信号，防止死锁
			pig <- struct{}{} // 发送信号给 cat
		}
	}
}
