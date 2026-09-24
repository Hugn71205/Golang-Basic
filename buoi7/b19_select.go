package main

import (
	"fmt"
	"time"
)

func goselect() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- "Du lieu chuyen vao ch1"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Du lieu chuyen vao ch2"
	}()
	// du ch2 xong truoc nhung phai doi ch1 xong thi moi chay ch2
	// fmt.Println(<-ch1)
	// fmt.Println(<-ch2)

	// dung select de chon ra cai nao xong truoc thi chay truoc
	for i := 1; i <= 2; i++ {
		select {
		case mess1 := <-ch1:
			fmt.Println(mess1)
		case mess2 := <-ch2:
			fmt.Println(mess2)
		}
	}
}
