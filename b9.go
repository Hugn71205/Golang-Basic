package main

import (
	"cmp"
	"fmt"
)

func PrintValue[T any](val T) {
	fmt.Println(val)
}

// func Toantubang[bang comparable](a,b bang)bool{
// neu nhu khai bao comparable chi su dung duoc cho cac bai toan so sanh ==, !=, khong lam duoc cac bai < > <= >=
func Toantubang[bang comparable](a, b bang) bool {
	return a == b
}

// neu muon so sanh < > <= >= bang generic thi su dung thu vien cmp.Ordered:
func Toantusosanh[H cmp.Ordered](a, b H) H {
	if a > b {
		return a
	}
	return b
}
func buoi9() {
	PrintValue("Tuan")
	PrintValue(2.5)
	fmt.Println()
	pheptinh1 := Toantubang(4, 5)
	fmt.Println(pheptinh1)
	fmt.Println()
	pheptinh2 := Toantubang("Tung", "Trung") //So sanh chu cai dau qua bang chu cai ASCII
	fmt.Println(pheptinh2)
	fmt.Println()
	PrintValue(Toantusosanh(5,9))
	PrintValue(Toantusosanh(5,15.6))
}
