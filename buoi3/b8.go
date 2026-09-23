package main

import (
	hung "Demo/buoi3/Hung"
	quynh "Demo/buoi3/Quynh"
	tung "Demo/buoi3/Tung"
	"Demo/buoi3/service"
	"fmt"
)

func Firstname(t service.Ten) {

	fmt.Printf("Day la ho cua %s \n", t.Getname())
	fmt.Println(t.Ho())

}
func Firstnameplus(p service.Tenplus) {

	fmt.Printf("Day la ho cua %s \n", p.Getname())
	fmt.Println(p.Ho())
	fmt.Println(p.Gioitinh())

}

// thay vi ghi interface{} thi co the ghi la any
//
//	func PrintValue (val interface{}){
//		// fmt.Println(val)
//		//cach dung interface{} nhung van kiem soat duoc du lieu dau vao:
//		str, ok := val.(string)
//		if ok{
//			fmt.Println("Gia tri la:",str)
//		}else{
//			fmt.Printf("Gia tri khong hop le, vui long nhap du lieu phu hop(%T)!",str)
//		}
//	}
//
// cach kiem soat da du lieu dau vao khi dung interface{}:
func checkDulieu(val interface{}) {
	switch val.(type) {
	case int:
		fmt.Printf("Gia tri la: %d\n", val)
	case string:
		fmt.Printf("Gia tri la: %s\n", val)
	case bool:
		fmt.Printf("Gia tri la: %t\n", val)
	default:
		fmt.Printf("Gia tri (%T) khong hop le!\n", val)
	}
}
func buoi8() {
	// hungBui := hung.Hung{
	// 	Name: "Hung",
	// }
	hungBui, err := hung.New("Hung")
	if err != nil {
		panic(err)
	}
	// fmt.Printf("Day la ho cua %s",hungBui.Name)
	fmt.Println()
	// fmt.Println(hungBui.Ho())
	Firstname(hungBui)

	//goi bang constructor
	quynhNgo, err := quynh.New("Quynh")
	if err != nil {
		panic(err)
	}
	// fmt.Printf("Day la ho cua %s",quynhNgo.Name)
	fmt.Println()
	// fmt.Println(quynhNgo.Ho())
	Firstnameplus(quynhNgo)
	fmt.Println()
	tungDau, err := tung.New("Tung")
	if err != nil {
		panic(err)
	}
	fmt.Println(tungDau.Dacdiem())
	//dung ham chua kieu du lieu la mot interface rong thi du lieu hien thi ben trong se la any
	// PrintValue(tungDau.Dacdiem())
	// PrintValue(23112005)
	checkDulieu(2.5)
}
