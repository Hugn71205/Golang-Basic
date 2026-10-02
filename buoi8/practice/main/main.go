package main

import (
	"buoi8/greetings"
	"fmt"
	"log"
)

func main() {
	// đặt nhãn cho lỗi sai bằng cách đặt tên "greetings: "để biết lỗi ipmort từ greetings
	log.SetPrefix("greetings: ")
	//nếu như mặc định thì log sẽ in ra lỗi kèm ngày giờ và để 0 để giá trị không trả về ngày giờ nữa
	log.SetFlags(0)
	// message, err := greetings.Hello("Diablo")
	names := []string{
		"Hung Bui",
		"Tung Do",
		"Quynh Ngo",
	}
	message, err := greetings.Hellos(names)

	if err != nil {
		//Nếu phát hiện có lỗi (do truyền vào chuỗi rỗng ở dòng trên):log.Fatal sẽ in lỗi ra màn hình (kèm tiền tố "greetings: ")
		// và lập tức thoát chương trình với mã lỗi (exit status 1).
		// Bất kỳ đoạn code nào bên dưới dòng này đều bị hủy và không được chạy.
		log.Fatal(err)
	}
	fmt.Println(message)
}
