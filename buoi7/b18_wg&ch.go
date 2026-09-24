package main

import (
	"fmt"
	"sync"
	"time"
)

func task1(id int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Nhiem vu %d bat dau \n", id)
	time.Sleep(1 * time.Second)
	ch <- fmt.Sprintf("Nhiem vu %d ket thuc \n", id)
	ch <- fmt.Sprintf("End task %d \n", id) //them buffer
	ch <- fmt.Sprintf("End task %d \n", id) //test khi de close trong go routine va khong dung buffer
	ch <- fmt.Sprintf("End task %d \n", id) //test khi de close trong go routine va khong dung buffer
	ch <- fmt.Sprintf("End task %d \n", id) //test khi de close trong go routine va khong dung buffer
	ch <- fmt.Sprintf("End task %d \n", id) //test khi de close trong go routine va khong dung buffer
	ch <- fmt.Sprintf("End task %d \n", id) //test khi de close trong go routine va khong dung buffer
	ch <- fmt.Sprintf("End task %d \n", id) //test khi de close trong go routine va khong dung buffer
}
func wg_ch() {
	start := time.Now()
	var wg sync.WaitGroup
	// ch := make(chan string, 4)
	// ch := make(chan string, 8) //them so luong buffer
	ch := make(chan string)

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go task1(i, ch, &wg)
	}
	//khi su dung buffer thi phai nam duoc ro so luong buffer vi neu qua buffer se loi
	// wg.Wait()
	// close(ch)

	//nhung neu khong su dung buffer thi co 1 cach do chinh la bỏ phần close channel vao trong 1 go routine de no khong the chay ngoai main ma dam bao phai chay duoc het go routine
	go func() {
		wg.Wait()
		close(ch)
	}()
	for val := range ch {
		fmt.Print(val)
	}
	fmt.Println("Thoi gian thuc hien nhiem vu la: ", time.Since(start))
}
