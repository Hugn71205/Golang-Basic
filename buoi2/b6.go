package main

import "fmt"

func lyThuyetPointer() {
	name := "Diablo"
	diaChi := &name
	diaChi2 := &diaChi
	fmt.Println("Infomation")
	fmt.Printf("Type: %T \n", name)
	fmt.Printf("Type: %T \n", diaChi)
	fmt.Printf("Value: %v \n", name)
	// fmt.Printf("Address: %v ", &name)
	fmt.Printf("Address: %v \n", diaChi)
	fmt.Printf("Find name by address of variable: %v \n", *diaChi)
	fmt.Printf("Find name by address of variable of variable: %v ", **diaChi2)

}

// func upDateName(name string) {
// 	name = "Bui Quoc Hung"
// 	fmt.Printf("Type: %T \n", name)
// 	fmt.Printf("Address: %v \n", &name)
// }
func upDateName(name *string) {
	*name = "Bui Quoc Hung"
	fmt.Printf("Name: %v \n", *name)
	fmt.Printf("Address: %v \n", name)
}
func buoi6() {
	name := "Diablo"
	fmt.Printf("Name: %v \n", name)
	fmt.Printf("Address: %v \n", &name)
	fmt.Println()
	name = "HungBui"
	fmt.Printf("Name: %v \n", name)
	fmt.Printf("Address: %v \n", &name)
	//=>Khi thay doi du lieu bien name khong trong func thi chi thay doi du lieu chu khong thay doi dia chi bien
	fmt.Println()
	// upDateName(name)
	//=>khi thay doi du lieu bien trong func thi la dang tao vung nho moi va se sinh ra dia chi moi
	upDateName(&name)
	fmt.Println()
	fmt.Printf("Name: %v \n", name)
	fmt.Printf("Address: %v \n", &name)
	//=>goi dia chi cu trong func de thay doi du lieu cua bien thi se giu nguyen dia chi bien cu
}
