package main

import (
	"context"
	"fmt"
	"time"
)

func employee(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Cong viec da bi huy", ctx.Err())
			return
		default:
			priority := ctx.Value("priority")
			fmt.Println("Dang lam viec voi muc do uu tien la:", priority)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
func ctx() {
	//tao task chi duoc lam trong 2s neu sau 2s du chua xong se huy de tranh leak memory gay break chuong trinh
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	//truyen du lieu va ghi chu thich
	ctx = context.WithValue(ctx, "priority", "high")
	//giao cong viec cho ham lam viec
	go employee(ctx)
	//tranh truong hop main chay nhanh hon khien goroutine bi dung
	time.Sleep(3 * time.Second)
}
