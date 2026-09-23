package main

import "fmt"
type SinhVien struct{
	name string
	age int
	class string
}
func buoi15() {
	//cac cach khai bao map

	//khai bao truc tiep
	drink := map[string]int{
		"tea":    300,
		"coffee": 350,
	}
	//%+v: in ca ten field cung du lieu
	fmt.Printf("%+v\n",drink)

	//dung make
	m := make(map[string]int)
		m["apple"] = 10
		m["pineapple"] = 15
		fmt.Println(m)
	
	//khai bao map rong(zero value)
	//gia tri mac dinh la `nil`, khong the su dung luon duoc
	var n map[string]int
	fmt.Println(n) //map[]
	//cần cấp phát bộ nhớ trước khi sử dụng
	n = make(map[string]int)
	n["apple"] = 10
	fmt.Println(n)

	//kiem tra key co ton tai trong map hay khong
	value, exists := drink["tea"]
	if exists{
		fmt.Println(value)
	}else{
		fmt.Println("khong ton tai!")
	}
	 
	//duyet map bang for range
	for key,val := range drink{
		fmt.Printf("Key: %s, value: %d\n",key,val)
	}
	
	//ket hop map va struct
	sv := map[string]SinhVien{
	"sv1" : {"Hung",21,"CNTT.C2"},
	"sv2" : {"Quynh",19,"09H"},
	}
	fmt.Printf("%+v",sv)
	//hien thi tung key ,value 
	fmt.Printf("Ho va ten: %v, tuoi: %d, lop: %s \n",sv["sv1"].name,sv["sv1"].age,sv["sv1"].class)
	//hien thi tat ca
	for key,val := range sv{
		fmt.Printf("Gia tri cua %v\n",key)
		fmt.Printf("Ten: %s \n",val.name)
		fmt.Printf("Tuoi: %d \n",val.age)
		fmt.Printf("Lop: %s \n",val.class)
	}

	//map let hop voi slice
	monHoc := map[string][]string{
		"Hung":{"Toan","Anh","Golang"},
		"Tung":{"Van","Toan"},
	}
	fmt.Printf("%+v",monHoc)
	fmt.Println()
	fmt.Printf("Hung dang hoc mon: %v",monHoc["Hung"])
	//co the goi ra tung phan tu trong slice
	fmt.Println()
	fmt.Printf("Hung dang hoc mon: %v",monHoc["Hung"][0])
	fmt.Println()
	fmt.Printf("Hung dang hoc mon: %v",monHoc["Hung"][1])
	//hien thi toan bo
	fmt.Println()
	for key1,val1 := range monHoc{
		// fmt.Printf("Hoc sinh ten la: %v \n",key)
		// fmt.Printf("Cac mon dang hoc hien tai la: %v \n",val)
		// fmt.Printf("Cac mon dang hoc gioi nhat la: %v \n",val[1])
		for _,val := range val1{
			fmt.Printf("Mon hoc cua %v la %v \n",key1,val)
		}
	}
}