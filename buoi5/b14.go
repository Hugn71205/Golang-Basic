package main

import (
	giangvien "Demo/giangVien"
	sinhvien "Demo/sinhVien"
	"Demo/utils"
	"fmt"
)

func buoi14() {
	// buoi5()

	// Vong lap menu
	for {
		utils.ClearScreen()
		fmt.Println()
		fmt.Println("=*=*=*=*=*=*=MENU CHUC NANG=*=*=*=*=*=*=")
		fmt.Println()
		fmt.Println("1.Quan Ly Sinh Vien")
		fmt.Println("2.Quan Ly Giang Vien")
		fmt.Println("3.Thoat Chuong Trinh")
		fmt.Println("=*=*=*=*=*=*=*=*=*=*=*=*=*==*=*=*=*=*=*=")
		fmt.Println()
		choice := utils.GetPostiveInt("👉 Chon chuc nang: ")
		switch choice {
		case 1:
			sinhvien.StudentMenu()
		case 2:
			giangvien.TeacherMenu()
		case 3:
			fmt.Println("Chao tam biet ban!")
			return
		default:
			fmt.Println("❌ Lua chon khong hop le!")
		}
		//neu chuc nang chua code thi auto chay lai vong lap => co the dung chuc nang sau:
		utils.ReadInput("\nNhan phim Enter de tiep tuc...")
	}
}
