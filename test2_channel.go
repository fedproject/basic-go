package main

import (
	"fmt"
	"time"
)

func main(){
	c:=make(chan int, 3) // 带有缓冲的channel
	fmt.Println("len(c) = ", len(c), ", cap(c)", cap(c))


	//sub goroutine
	go func() {
		defer fmt.Println("sub go end")

		for i:=0; i < 3; i++{
			c <-i
			fmt.Println("Sub goroutine is running, element being transferred ", i, "len(c)=", len(c), ", cap(c)=", cap(c))
		}
	}()

	time.Sleep(2*time.Second)

	for i:=0; i <3; i++{
		num := <-c // accept from c, and assign to num
		fmt.Println("num = ", num)
	}


}

