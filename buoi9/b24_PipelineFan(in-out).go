package main

import (
	"fmt"
	"sync"
	"time"
)

// BƯỚC 1: GENERATE (Máy phát hàng)
// Bơm các con số (đơn hàng) vào băng chuyền đầu vào
func mayPhatHang(danhSachSo []int) <-chan int {
	bangChuyenVao := make(chan int)
	go func() {
		for _, so := range danhSachSo {
			bangChuyenVao <- so
		}
		close(bangChuyenVao) // Hết hàng thì tắt băng chuyền
	}()
	return bangChuyenVao
}

// BƯỚC 2: FAN-OUT (Công nhân xử lý)
// Từng công nhân đứng ở băng chuyền, lấy hàng ra xử lý (nhân 10)
func congNhan(ten string, bangChuyenVao <-chan int) <-chan int {
	bangChuyenRa := make(chan int)
	go func() {
		for so := range bangChuyenVao {
			// Giả vờ công nhân này mất 1 giây để xử lý đơn hàng
			time.Sleep(1 * time.Second)
			kq := so * 10
			fmt.Printf("%s vừa xử lý xong số %d thành %d\n", ten, so, kq)

			bangChuyenRa <- kq // Làm xong thì vứt vào băng chuyền của riêng mình
		}
		close(bangChuyenRa)
	}()
	return bangChuyenRa
}

// BƯỚC 3: FAN-IN (Cái phễu gom hàng)
// Gom hàng từ băng chuyền của công nhân A và công nhân B vào chung 1 thùng
func phieuGom(bangChuyenA <-chan int, bangChuyenB <-chan int) <-chan int {
	thungGom := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2) // Báo trước là có 2 công nhân đang làm

	// Luồng gom hàng của công nhân A
	go func() {
		for kq := range bangChuyenA {
			thungGom <- kq
		}
		wg.Done() // Công nhân A báo cáo: "Tôi đã xả hết hàng"
	}()

	// Luồng gom hàng của công nhân B
	go func() {
		for kq := range bangChuyenB {
			thungGom <- kq
		}
		wg.Done() // Công nhân B báo cáo: "Tôi đã xả hết hàng"
	}()

	// Luồng đứng canh: Khi cả 2 công nhân đều Done thì đóng thùng gom
	go func() {
		wg.Wait()
		close(thungGom)
	}()

	return thungGom
}

func pipeline() {
	danhSachDonHang := []int{1, 2, 3, 4, 5}
	fmt.Println("Bắt đầu xử lý 5 đơn hàng...")

	// 1. Bật máy phát hàng
	bangChuyenVao := mayPhatHang(danhSachDonHang)

	// 2. Chia cho 2 công nhân (Fan-out)
	// Dù có 5 đơn, nhưng nhờ 2 người làm song song nên sẽ tốn ít thời gian hơn
	kqCongNhanA := congNhan("Nam", bangChuyenVao)
	kqCongNhanB := congNhan("Nữ", bangChuyenVao)

	// 3. Đặt cái phễu gom kết quả của Nam và Nữ lại (Fan-in)
	thungGom := phieuGom(kqCongNhanA, kqCongNhanB)

	// In kết quả cuối cùng ra màn hình
	for ketQua := range thungGom {
		fmt.Printf("Đã thu thập được thành phẩm: %d\n", ketQua)
	}
	fmt.Println("Hoàn thành toàn bộ công việc!")
}
