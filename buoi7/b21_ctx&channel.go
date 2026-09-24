package main

import (
	"context"
	"fmt"
	"time"
)

func cookPho(ctx context.Context, phoCh chan string) {
	fmt.Println("Bat dau nau pho")
	select {
	case <-time.After(1 * time.Second):
		phoCh <- "Pho da nau xong"
	case <-ctx.Done():
		fmt.Println("Huy nau pho")
		return
	}
}
func cookCom(ctx context.Context, comCh chan string) {
	fmt.Println("Bat dau nau com")
	select {
	case <-time.After(100 * time.Millisecond):
		comCh <- "Com da nau xong"
	case <-ctx.Done():
		fmt.Println("Huy nau com")
		return
	}
}
func cookPizza(ctx context.Context, ch chan string) {
	fmt.Println("Bat dau nau pizza")
	select {
	case <-time.After(2 * time.Second):
		ch <- "Pizza da nau xong"
	case <-ctx.Done():
		fmt.Println("Huy nau pizza")
		return
	}

}
func democtxchan() {
	phoCh := make(chan string)
	pizzaCh := make(chan string)
	comCh := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()
	go cookPho(ctx, phoCh)
	go cookPizza(ctx, pizzaCh)
	go cookCom(ctx, comCh)
	for i := 1; i <= 3; i++ {
		select {
		case pho := <-phoCh:
			fmt.Println("Nhan mon", pho)
		case pizza := <-pizzaCh:
			fmt.Println("Nhan mon", pizza)
		case com := <-comCh:
			fmt.Println("Nhan mon", com)
		case <-ctx.Done():
			fmt.Println("Dung nhan mon")
			return
		}
	}
}
