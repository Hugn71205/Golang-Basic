package main

import (
	"fmt"
	"sync"
)

// 1. Giấu đồ vật đi: Chữ 'e' viết thường để không ai tự khởi tạo được
type eraser struct {
	Color string
}

// 2. Két sắt cất giữ và Người bảo vệ
var theEraser *eraser // Biến toàn cục để lưu chiếc giẻ lau duy nhất
var once sync.Once    // Ổ khóa đảm bảo việc tạo ra chỉ xảy ra 1 lần

// 3. Hàm cấp phát: Ai cần giẻ lau thì phải gọi hàm này
func GetEraser() *eraser {
	// once.Do sẽ chạy đúng 1 lần duy nhất trong toàn bộ chương trình
	once.Do(func() {
		fmt.Println("=> Đang đi mua chiếc giẻ lau màu Đỏ!")
		theEraser = &eraser{Color: "Đỏ"}
	})

	// Trả về chiếc giẻ lau đang nằm trong két sắt
	return theEraser
}

func singleton() {
	// Học sinh A lên mượn giẻ lau (Là người đầu tiên nên chương trình sẽ đi mua)
	fmt.Println("Học sinh A mượn:")
	eraserA := GetEraser()

	// Học sinh B lên mượn (Lúc này đã có giẻ lau, chương trình không đi mua nữa mà đưa luôn cái cũ)
	fmt.Println("Học sinh B mượn:")
	eraserB := GetEraser()

	// In địa chỉ bộ nhớ để chứng minh A và B đang cầm cùng 1 chiếc giẻ lau
	fmt.Printf("Địa chỉ giẻ lau của A: %p\n", eraserA)
	fmt.Printf("Địa chỉ giẻ lau của B: %p\n", eraserB)
}
