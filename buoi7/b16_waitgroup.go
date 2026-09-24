package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Nhiem vu %d bat dau \n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Nhiem vu %d ket thuc \n", id)
}
func waitGroup() {
	// start := time.Now()
	// go task(1)
	// go task(2)
	// go task(3)
	// go task(4)
	// time.Sleep(2 * time.Second)
	// fmt.Println("Thoi gian thuc hien nhiem vu la: ", time.Since(start))
	start := time.Now()
	var wg sync.WaitGroup
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait()
	fmt.Println("Thoi gian thuc hien nhiem vu la: ", time.Since(start))

}
