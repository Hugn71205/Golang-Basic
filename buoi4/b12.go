package main

import (
	"fmt"
	"reflect"
)

func buoi12() {
	// //slice thi k co mo ta gi ve kich thuoc
	// latcat :=
	// 	[]int{1,2,3,4,5}
	// //array thi luon di cung mo ta ve kick thuoc
	// mang := [5]int{1, 2, 3, 4, 5}
	// fmt.Println("Slice:",latcat)
	// fmt.Println()
	// fmt.Println("Array:",mang)
	// fmt.Println()

	// fmt.Println("(slice) Co phai Slice khong?",reflect.TypeOf(latcat).Kind()==reflect.Slice)
	// fmt.Println()

	// fmt.Println("(array) Co phai Array khong?",reflect.TypeOf(mang).Kind()==reflect.Array)
	//mang co tong phan tu la 5
	//Vay[1:4]=>
	//1=>tinh tu vi tri dau la 0->n-1(n la tong so phan tu => 5-1)
	//=>0->4
	//4=>tinh tu vi tri dau la 1->n(n la tong so phan tu => 5)
	//=>1->5
	// arr := [5]int{10, 20, 30, 40, 50}
	// fmt.Println(arr)
	// sli := arr[1:4]
	// fmt.Println(sli)
	// fmt.Println("(sli) Co phai Slice khong?",reflect.TypeOf(sli).Kind()==reflect.Slice)
	//tao slice bang phuong thuc makeco3phan(kieu du lieu cua slice, chieu dai cua slice(lenght), dung luong toi da cua slice(capacity))
	// slice := make([]int, 4, 5) // [0 0 0 0]
	// slice = append(slice, 5,6) // [0 0 0 0 5 6] //append la ham them du lieu vao slice va khi nhap qua dung luong toi da cua slice hien tai thi go se cho x2 dung luong hien tai
	// slice[0] = 1 //[1 0 0 0 5 6]
	// slice[1] = 2 //[1 2 0 0 5 6]
	// slice[2] = 3 //[1 2 3 0 5 6]
	// fmt.Println(slice)
	// fmt.Println("Do dai cua slice:",len(slice))
	// fmt.Println("Dung luong toi da cua slice:",cap(slice))
	// hung := []string{"Diablo", "Hugn7125", "HungBui", "radaii"}
	// // for _,i := range hung {
	// // 	fmt.Println("Ten cua Hung la:",i)
	// // }
	// for i := 0; i < len(hung); i++ {
	// 	fmt.Println(hung[i])

	// }
	//slice da chieu
	// truong:= [][]string{
	// {"Hung", "Tung", "Quynh"},
	// {"Nhung","Hue","Hang"},
	// }
	// fmt.Println(truong)
	// // fmt.Println(so[1][2])
	// for key1,lop :=range truong{
	// 	fmt.Println(lop)
	// 	for key2,hocsinh := range lop{
	// 		fmt.Printf("Phan tu cua lop %d thuoc %d la: %s\n",key2,key1,hocsinh)
	// 	}
	// }
	//them du lieu tu mang nay vao mang kia
	// so := []int{1, 2, 3}
	// chu := []int{5,6,7}
	// fmt.Println(so)
	// fmt.Println(chu)
	// fmt.Println()
	// so = append(so,chu...)
	// fmt.Println(so)
	//tao ra slice tu slice
	so := []int{1, 2, 3, 4, 5}
	fmt.Println(so)
	fmt.Println("Co phai slice khong? ",reflect.TypeOf(so).Kind() == reflect.Slice)
	fmt.Println("Do dai cua slice:",len(so))
	fmt.Println("Dung luong toi da cua slice:",cap(so))
	fmt.Println()
	subslice := so[1:4] // [2 3 4]
	// subslice := so[1:] // [2 3 4 5]
	// subslice := so[:4] // [1 2 3 4]
	fmt.Println(subslice)
	fmt.Println("Co phai slice khong? ",reflect.TypeOf(subslice).Kind() == reflect.Slice)
	fmt.Println("Do dai cua slice:",len(subslice))
	fmt.Println("Dung luong toi da cua slice:",cap(subslice))
}
