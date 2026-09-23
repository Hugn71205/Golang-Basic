package main

import "fmt"

// func phepToan(num1 int, num2 int) (int, int, int, float32) {
// 	if num2 == 0 {
// 		num2 = 6
// 	}
// 	phepCong := num1 + num2
// 	phepTru := num1 - num2
// 	phepNhan := num1 * num2
// 	phepChia := float32(num1) / float32(num2)
// 	// fmt.Println("Ket qua:", phepCong)
// 	return phepCong, phepTru, phepNhan, phepChia
// }
// func countDowm(numb int) {
// 	fmt.Println(numb)
// 	if numb > 0 {
// 		countDowm(numb - 1)
// 	}
// }
// bai4:xay dung 1 ham de quy va tinh tong day so N Vd: nhap 10 thi tinh tong tu 1 den 10
func tongDaySoN(n int) int {
	if n == 1 {
		return 1
	}
	return n + tongDaySoN(n-1)
}

// bai5: tinh tong day so fibonacci
func fibo(n int) {
	term1, term2 := 0, 1
	if n == 1 {
		fmt.Printf("Day so fibonacci: %d", term1)
		fmt.Printf("Tong day so fibo:%d", term1)
	} else if n == 2 {
		fmt.Printf("Day so fibonacci: %d %d", term1, term2)
		fmt.Printf("Tong day so fibo:%d", term1+term2)
	} else {
		tongDayFibo := term1 + term2
		fmt.Printf("Day so fibonacci: %d %d", term1, term2)
		for i := 3; i <= n; i++ {
			term3 := term1 + term2
			fmt.Printf(" %d", term3)
			tongDayFibo += term3
			term1, term2 = term2, term3
		}
		fmt.Println()
		fmt.Printf("Tong day so fibo:%d \n", tongDayFibo)

	}
}
func buoi5() {
	for {
		fmt.Println("")
		fmt.Println("=======MENU CHUC NANG=======")
		fmt.Println("1. Bai 4: Tinh tong day so")
		fmt.Println("2. Bai 5: Day so fibonacci")
		fmt.Println("0. Thoat chuong trinh")
		fmt.Println("============================")

		var choice int
		for {
			fmt.Println("Vui long chon truong trinh tuong ung voi cac so o tren!")
			_, err := fmt.Scan(&choice)
			if err != nil || choice < 0 || choice > 3 {
				fmt.Printf("Vui long nhap lai!")
			} else {
				break
			}
		}

		// countDowm(10)
		// cong, tru, nhan, chia := phepToan(5, 0)
		// fmt.Println("Day chinh la ket luan cuoi cung:", cong)
		// fmt.Println("Day chinh la ket luan cuoi cung:", tru)
		// fmt.Println("Day chinh la ket luan cuoi cung:", nhan)
		// fmt.Println("Day chinh la ket luan cuoi cung:", chia)
		// b4:
		switch choice {
		case 1:
			var numb int
			for {
				fmt.Print("Vui long nhap so N:")
				_, err := fmt.Scan(&numb)
				if err != nil || numb <= 0 {
					fmt.Println("Vui long nhap chu so nguyen duong!")
					// return
					continue
				} else {
					break
				}
			}
			ketQua := tongDaySoN(numb)
			fmt.Printf("Ket qua day so %d la:%d \n", numb, ketQua)
		case 2:
			// bai5
			var numb int
			for {
				fmt.Printf("Vui long nhap so luong day fibonacci ban muon:")
				_, err := fmt.Scan(&numb)
				if err != nil || numb <= 0 {
					fmt.Println("Ky tu ban nhap khong hop le!")
					// return
					// continue
				} else {
					break
				}
			}
			fibo(numb)
		case 0:
			fmt.Println("Chao tam biet!")

			return
		default:
			fmt.Println("Vui long chon so chuc nang hien co!")
		}
	}
}
