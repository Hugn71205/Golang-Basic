package main

import (
	"fmt"
	"strconv"
)

func buoi4() {
	// var diem float64
	// fmt.Print("Nhap diem cua ban: ")
	// fmt.Scanf("%f", &diem)
	// if diem >= 8.5 {
	// 	fmt.Println("Xuat sac")
	// } else if diem < 8.5 && diem >= 7 {
	// 	fmt.Println("Gioi")
	// } else {
	// 	fmt.Println("Trung binh")
	// }
	// monAn := "banh"
	// switch monAn {
	// case "com":
	// 	fmt.Println("Mon an la com")
	// case "pho bo":
	// 	fmt.Println("Mon an la pho")
	// default:
	// 	fmt.Println("Mon an khong xac dinh")
	// }
	// switch monAn {
	// case "com", "bun":
	// 	fmt.Println("Mon an la com hoac la bun")
	// case "pho bo":
	// 	fmt.Println("Mon an la pho")
	// default:
	// 	fmt.Println("Mon an khong xac dinh")
	// }
	// for i := 1; i <= 10; i++ {
	// 	//diembatdau;diemketthuc;buocnhay
	// 	fmt.Printf("Gia tri cua i: %d\n", i)
	// }
	//break va continue
	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		// break //dung vong lap
	// 		continue //bo qua gia tri hien tai va tiep tuc vong lap
	// 	}
	// 	fmt.Printf("Gia tri cua i: %d\n", i)
	// }
	// fmt.Println("Ket thuc chuong trinh")
	// bt1:bo qua cac so 6,48,75,89 va in ra cac so con lai tu 1 den 100,moi so cach nhau boi dau phay,sau so cuoi cung khong co dau phay
	// for i := 1; i <= 100; i++ {
	// 	if i == 6 || i == 48 || i == 75 || i == 89 {
	// 		continue
	// 	}
	// 	fmt.Printf("%d", i)
	// 	if i != 100 {
	// 		fmt.Print(",")
	// 	}
	// }
	//cach2:khong dung fmt.Printf qua nhieu lan
	// var ketqua = ""
	// for i := 1; i <= 100; i++ {
	// 	if i == 6 || i == 48 || i == 75 || i == 89 {
	// 		continue
	// 	}
	// 	ketqua += fmt.Sprintf("%d", i)
	// 	if i != 100 {
	// 		ketqua += ","
	// 	} else {
	// 		ketqua += "\n"
	// 	}
	// }
	// fmt.Print(ketqua)
	//cach3:int=>string
	// var ketqua = ""
	// for i := 1; i <= 100; i++ {
	// 	if i == 6 || i == 48 || i == 75 || i == 89 {
	// 		continue
	// 	}
	// 	ketqua += strconv.Itoa(i)
	// 	if i != 100 {
	// 		ketqua += ","
	// 	} else {
	// 		ketqua += "\n"
	// 	}
	// }
	// fmt.Print(ketqua)
	// bai2:chi in ra cac so le tu 1 den 100, moi so cach nhau boi dau phay, moi dong in ra 3 so le,bo qua cac so chan, sau so cuoi cung khong co dau phay
	// cach1:
	// 	count := 0
	// 	for i := 1; i <= 100; i++ {
	// 		if i%2 == 0 {
	// 			continue
	// 		}
	// 		if count > 0 {
	// 			fmt.Print(",")
	// 		}
	// 		fmt.Printf("%d", i)
	// 		count++
	// 		if count == 3 {
	// 			fmt.Print("\n")
	// 			count = 0
	// 		}
	// 	}
	// cach2:
	// ketqua := ""
	// count := 0
	// for i := 1; i <= 100; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	if count > 0 {
	// 		ketqua += ","
	// 	}
	// 	ketqua += strconv.Itoa(i)
	// 	count++
	// 	if count == 3 {
	// 		ketqua += "\n"
	// 		count = 0
	// 	}
	// }
	// fmt.Print(ketqua)
	// bai3
	var batDau, ketThuc int
	fmt.Print("Nhap gia tri bat dau: ")
	fmt.Scanf("%d \n", &batDau)
	fmt.Print("Nhap gia tri ket thuc: ")
	fmt.Scanf("%d", &ketThuc)
	if batDau == 0 || ketThuc == 0 {
		fmt.Println("Gia tri bat dau va ket thuc phai khac 0")
		return
	} else if batDau > ketThuc {
		fmt.Println("Gia tri bat dau phai nho hon gia tri ket thuc")
		return
	} else {
		ketQua := ""
		for k := batDau; k <= ketThuc; k++ {
			ketQua += "Bang cuu chuong" + strconv.Itoa(k) + "\n"
			// fmt.Printf("%d x %d = %d\n", k, 1, k*1)
			// fmt.Printf("%d x %d = %d\n", k, 2, k*2)
			// fmt.Printf("%d x %d = %d\n", k, 3, k*3)
			//cach2
			for i := 1; i <= 10; i++ {
				ketQua += strconv.Itoa(k) + "x" + strconv.Itoa(i) + "=" + strconv.Itoa(k*i) + "\n"
				// fmt.Printf("%d x %d = %d\n", k, i, k*i)
			}
		}
		fmt.Print(ketQua)
	}
}
