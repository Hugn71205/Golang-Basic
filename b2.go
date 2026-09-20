package main

import "fmt"

func buoi2() {
	// fmt.Print("Hello Diablo. ")
	// fmt.Println("Goodbye Diablo. ")
	// name := "Diablo"
	// age := 20
	// fmt.Printf("Hello my name is %s and I am %d years old.", name, age)
	// var firstName, lastName string
	// fmt.Print("Vui long nhap ho va ten: ")
	// fmt.Scanln(&firstName)
	// fmt.Print("Vui long nhap ho va ten: ")
	// fmt.Scanln(&lastName)
	// fmt.Printf("Xin chao %s %s!", firstName, lastName)
	// var name string
	// var age int
	// fmt.Print("Vui long nhap ten: ")
	// fmt.Scanf("%s\n", &name)
	// fmt.Print("Vui long nhap tuoi: ")
	// fmt.Scanf("%d", &age)
	// fmt.Printf("Xin chao %s, ban co %d tuoi!", name, age)

	// message := fmt.Sprint("Hello ", "God")
	// fmt.Println(message)
	// name := "Diablo"
	// age := 20
	// message := fmt.Sprintf("Hello %s, you are now %d years old.", name, age)
	// fmt.Println(message)
	ten := "Diablo"
	canNang := 48.5
	tuoi := 20
	chuyenNganh := "CNTT"
	daTotNghiep := true
	tiLeTotNghiep := 85
	//%s: string
	//%f: float
	//%d: int
	//%v: any type
	//%t: boolean
	//%.f: float with 1 decimal
	//%c: character
	//%p: pointer
	//%%: neu can them %
	fmt.Printf("Kieu du lieu cua bien ten: %T \n", ten)
	fmt.Printf("Kieu du lieu cua bien canNang: %.1f \n", canNang)
	fmt.Printf("Kieu du lieu cua bien tuoi: %d \n", tuoi)
	fmt.Printf("Kieu du lieu cua bien chuyenNganh: %v \n", chuyenNganh)
	fmt.Printf("Kieu du lieu cua bien daTotNghiep: %t \n", daTotNghiep)
	fmt.Printf("Kieu du lieu cua bien tiLeTotNghiep: %v%% \n", tiLeTotNghiep)
}
