package main

import (
	"fmt"
	"sync"
)

var (
	c1    = make(chan struct{})
	c2    = make(chan struct{})
	c3    = make(chan struct{})
	numWg sync.WaitGroup
)

func num1() {
	defer numWg.Done()
	for i := 1; i <= 100; i = i + 3 {
		<-c3
		fmt.Println(i)
		c1 <- struct{}{}
	}
}

func num2() {
	defer numWg.Done()
	for i := 2; i <= 101; i = i + 3 {
		<-c1
		fmt.Println(i)
		c2 <- struct{}{}
	}
}

func num3() {
	defer numWg.Done()
	for i := 3; i <= 102; i = i + 3 {
		<-c2
		fmt.Println(i)
		if i+3 <= 102 { // 避免最后一次发送信号，防止死锁
			c3 <- struct{}{} // 通知 Goroutine 1
		}
	}
}
