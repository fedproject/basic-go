package main

import "fmt"

func main(){
	// 定义一个channel
	c := make(chan int)

	go func(){
		defer fmt.Println("goroutine End")
		fmt.Println("Goroutine running currently")

		c <- 666 // send 666 to variable c
	}()

	num := <-c // accept data from c, and assign to num


	fmt.Println("num = ", num)
	//fmt.Print

}
