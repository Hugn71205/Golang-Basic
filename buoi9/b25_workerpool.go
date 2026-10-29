package main

import (
	"fmt"
	"time"
)

// Khai báo hàm cho 1 Worker
// jobs: Channel chỉ nhận (để worker rút việc ra làm)
// results: Channel chỉ gửi (để worker nhét kết quả vào sau khi làm xong)
func worker(id int, jobs <-chan int, results chan<- int) {
	// Worker liên tục chực chờ bốc việc từ channel jobs
	for job := range jobs {
		fmt.Printf("Worker %d bắt đầu xử lý job số %d\n", id, job)

		// Giả lập thời gian làm việc tốn 1 giây
		time.Sleep(1 * time.Second)

		fmt.Printf("Worker %d đã xong job số %d\n", id, job)

		// Xong việc, gửi kết quả vào channel results
		results <- job * 2
	}
}

func workerpool() {
	const soLuongJob = 5
	const soLuongWorker = 3

	// Tạo 2 channels có buffer (bộ đệm) chứa công việc và kết quả
	jobs := make(chan int, soLuongJob)
	results := make(chan int, soLuongJob)

	// BƯỚC 1: Khởi động Worker Pool (Thuê 3 công nhân)
	// Lúc này 3 worker đã chạy ngầm (Goroutines), nhưng đang tạm dừng (block) vì chưa có job
	for w := 1; w <= soLuongWorker; w++ {
		go worker(w, jobs, results)
	}

	// BƯỚC 2: Giao việc (Đẩy 5 jobs vào đường ống)
	for j := 1; j <= soLuongJob; j++ {
		jobs <- j
	}
	// Đẩy xong thì khoá ống jobs lại, báo cho worker biết là không còn việc mới đâu
	close(jobs)

	// BƯỚC 3: Thu thập kết quả
	// Biết trước có 5 jobs, nên ta sẽ chờ rút đủ 5 kết quả ra từ ống results
	for a := 1; a <= soLuongJob; a++ {
		ketQua := <-results
		fmt.Printf("Đã nhận kết quả: %d\n", ketQua)
	}
}
