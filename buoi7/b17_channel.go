package main

import (
	"fmt"
	"time"
)

func channel() {
	//gui thong tin qua goroutine khong co buffered
	// cn := make(chan int)
	//goroutine anonymous:goroutine khong co ten
	// go func() {
	// 	defer close(cn)
	// 	cn <- 1
	// 	cn <- 10
	// 	cn <- 100
	// }()

	//gui thong tin bang buffered
	cn := make(chan int, 3) //them buffer bang cach ghi them so luong trong make chan
	cn <- 1
	cn <- 10
	cn <- 100
	// close(cn) //chi dung cho for range

	// go func() {
	// 	fmt.Println(<-cn)
	// }()
	// fmt.Println(<-cn)
	// fmt.Println(<-cn)
	// fmt.Println(<-cn)

	for i := 1; i <= 3; i++ {
		fmt.Println(<-cn)

	}

	//chi dung truong hop nay khi su dung close channel
	// for val := range cn {
	// 	fmt.Println(val)

	// }
	time.Sleep(1 * time.Second)
}
