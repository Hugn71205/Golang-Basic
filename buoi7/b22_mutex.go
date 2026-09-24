package main

import (
	"fmt"
	"sync"
)

func mutexx() {
	tokens := 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func() {
			// khi chay binh thuong ma khong dung mutex => 1000 goroutine chay cung luc va se chi in ra thang chay nhanh nhat chu khong hien ket qua chinh xac
			// dung mutex de khoa dap an lai sau khi thay doi thi moi mo ra
			mu.Lock()
			tokens++
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(tokens)
}
