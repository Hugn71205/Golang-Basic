package mangsv

import "fmt"

type SinhVien struct {
	Masv    int
	Hovaten string
	Lop     string
}

func Buoi11MangSV() {
	//cach1: khai bao tung phan tu
	// var character string
	// var characters [5]string
	// var numb int
	// var numbs [5]int
	// fmt.Println(character)
	// fmt.Println(characters)
	// fmt.Println()
	// fmt.Println(numb)
	// fmt.Println(numbs)
	// var stt [6] int
	// //so thu tu cua arry se la kich thuoc hien tai -1 vd: 0 0 1 0 0 => vtri:2
	// stt[2] = 1
	// stt[3] = 2
	// fmt.Println(stt)

	//cach khai bao thu 2:Khai bao truc tiep
	// numbers := [5]int{7, 1, 2, 0}
	// numbers[4]=5
	// fmt.Println(numbers)
	//cach khai bao thu 3:Khai bao khong gioi han kich thuoc
	//tuy la khai bao khong gioi han kich thuoc nhung khong the gian tiep tang gioi han kich thuoc ma chi co the tang truc tiep
	// numbers := [...]int{1, 2, 3, 4, 5, 6}
	// fmt.Println(numbers)
	// fmt.Printf("%T",numbers)

	// var matrix = [2][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// }
	// fmt.Println(matrix)
	// //lay gia tri theo thu tu [] truoc: la so dong, [] sau: la so cot
	// fmt.Println(matrix[1][1])
	// fmt.Println()

	// matrix[1][1] = 9
	// fmt.Println(matrix[1][1])
	// fmt.Println(matrix)
	// fmt.Println()
	// fmt.Println(numbers[2])
	// matrix := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// // fmt.Println(len(matrix))
	// for i := 0; i < len(matrix); i++ {
	// 	fmt.Println(matrix[i])
	// }
	//duyet mang bang for len quet bang do dai mang
	// matrix := [2][10]int{
	// 	{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	// 	{66, 77, 33, 46, 57, 86, 37, 78, 29, 130},
	// }
	// for l := 0; l < len(matrix); l++ {
	// 	fmt.Println(matrix[l])
	// 	fmt.Println()
	// 	for k:=0;k<len(matrix[l]);k++{
	// 		fmt.Println(matrix[l][k])
	// 	}
	// }
	//duyet mang bang for range
	// matrix := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// for viTri, giaTri := range matrix {
	// 	fmt.Printf("Gia tri cua o %d la: %d \n", viTri, giaTri)
	// }

	// matrix := [2][3]int{{1, 2, 3},{4, 5, 6}}
	// for vtri1, giaTri := range matrix {
	// 	// fmt.Printf("Gia tri cua day %d la: \n", giaTri)
	// 	for vtri2,giaTri2 := range giaTri{
	// 		fmt.Printf("Gia tri cua o [%d] [%d] la: %d \n", vtri1,vtri2, giaTri2)
	// 	}
	// }
	var sinhVien = [...]SinhVien{
		{Masv: 71205, Hovaten: "Bui Quoc Hung", Lop: "CNTT.C2"},
		{Masv: 23105, Hovaten: "Do Thanh Tung", Lop: "CNTT.C1"},
		{Masv: 23107, Hovaten: "Ngo Nhu Quynh", Lop: "CNTT.C2"},
	}
	for _, val := range sinhVien {
		fmt.Printf("Ma sinh vien: %d Ho va ten: %s Lop:%s\n", val.Masv, val.Hovaten, val.Lop)
	}
	fmt.Println()
	fmt.Println(sinhVien[2])
	fmt.Println(sinhVien[2].Hovaten)
}
