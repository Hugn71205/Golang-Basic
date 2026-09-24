package main

import (
	"fmt"
	"sync"
	"time"
)

func heavytask(wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for i := 0; i < 100e8; i++ {
		sum += i
	}
}
func main() {
	var wg sync.WaitGroup
	start := time.Now()

	// numbCPU := runtime.NumCPU() //ham kiem tra xem may co bao nhieu CPU
	numbCPU := 1
	// runtime.GOMAXPROCS(numbCPU)//ham khoa cpu toi da de chay chuong trinh
	fmt.Println("So CPU hien tai trong may la:", numbCPU)
	// wg.Add(1)
	// go heavytask(&wg)
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go heavytask(&wg)
	}
	wg.Wait()
	fmt.Println("Tong so thoi gian chay la:", time.Since(start))
}
