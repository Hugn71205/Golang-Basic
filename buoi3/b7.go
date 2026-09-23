package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// type sinhVien struct {
// 	Name     string `json:"Ho va ten:"`
// 	Age      int 	`json:"So tuoi:"`
// 	Class    string `json:"Lop dang hoc:"`
// 	NganhHoc string `json:"Nganh dang hoc:"`
// 	// => dung `json:"Ho va ten"` (structtag) de hien thi thong tin du lieu ban muon
// }
// func hienThiThongTinSV(sv *sinhVien){
// func (sv sinhVien) hienThiThongTinSV(){
// 	fmt.Println("Ho va ten:",sv.Name)
// 	fmt.Println("Tuoi:",sv.Age)
// 	fmt.Println("Lop:",sv.Class)
// 	fmt.Println("Nganh hoc:",sv.NganhHoc)
// }
// 	// =>de tranh tao cac du lieu tren vung bien moi, dung pointer
// 	// =>neu la bien thi phai khai bao dau (*) nhung neu la sttruct thi co the bo qua
// func (sv *sinhVien)xoaDuLieu(){
// 	sv.Name = ""
// 	sv.Age = 0
// 	sv.Class = ""
// 	sv.NganhHoc = ""

// }
// bai6
type hcn struct {
	chieuDai  float32 `decs:"Chieu dai hcn"`
	chieuRong float32 `decs:"Chieu rong hcn"`
}

func (cv *hcn) chuViHCN() float32 {
	return (cv.chieuDai + cv.chieuRong) * 2
	// ketQuacv := (cv.chieuDai+cv.chieuRong)*2
	// fmt.Println("Chu vi hcn:",ketQuacv)
}
func (dt *hcn) dienTichHCN() float32 {
	return dt.chieuDai * dt.chieuRong
	// ketQuadt := dt.chieuDai*dt.chieuRong
	// fmt.Println("Dien tich hcn:",ketQuadt)
}
//ham cu chua chuyen string to float
func inputHcn() hcn {
	var kichThuochcn hcn
	for {
		fmt.Print("Vui long nhap chieu dai hcn:")
		_, err := fmt.Scan(&kichThuochcn.chieuDai)
		if err == nil && kichThuochcn.chieuDai > 0 {
			break
		}
		fmt.Println("Chieu dai phai lon hon 0!")
	}
	for {
		fmt.Print("Vui long nhap chieu rong hcn:")
		_, err := fmt.Scan(&kichThuochcn.chieuRong)
		if err == nil && kichThuochcn.chieuRong > 0 {
			break
		}
		fmt.Println("Chieu rong phai lon hon 0!")
	}
	return kichThuochcn
}
//Ham kiem tra va chuyen du lieu dau vao tu string sang float
func readFloat(prompt string) (float32, error) {
	//thong bao nguoi dung nhap lieu
	fmt.Print(prompt)
	//Doc du lieu tu ban phim
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("Loi doc du lieu dau vao: %v", err)
	}
	//Xoa bo khoang trang truoc va sau chuoi
	input = strings.TrimSpace(input)
	//chuyen tu chuoi sang so
	// numb, err := strconv.Atoi(input
	//chuyen tu chuoi sang float (dulieu, kichthuoc)
	numb, err := strconv.ParseFloat(input, 32)
	if err != nil {
		return 0, fmt.Errorf("Invalid number: %v", err)
	}
	return float32(numb), nil
	//vi du lieu minh nhap va mong muon tra ve 32 nen phai ghep du lieu: float32(numb)
}
func inputV2() hcn {
	var cd, cr float32
	for {
		var err error
		cd, err = readFloat("Vui long nhap chieu dai hcn:")
		if err != nil || cd <= 0 {
			fmt.Println("Du lieu khong hop le, vui long nhap lai!")
		}else{
			break
		}
	}
	for {
		var err error
		cr, err = readFloat("Vui long nhap chieu rong hcn:")
		if err != nil || cr <= 0 {
			fmt.Println("Du lieu khong hop le, vui long nhap lai!")
		}else{
			break
		}
	}

	return hcn{
		chieuDai:  cd,
		chieuRong: cr,
	}
}
func buoi7() {
	// quocHung := sinhVien{
	// 	Name:     "Bui Quoc Hung",
	// 	Age:      21,
	// 	Class:    "CNTT.C2",
	// 	NganhHoc: "Cong nghe thong tin",
	// }
	// hienThiThongTinSV(&quocHung)
	// quocHung.hienThiThongTinSV()
	// fmt.Println()
	// quocHung.xoaDuLieu()
	// fmt.Println()
	// quocHung.hienThiThongTinSV()
	// fmt.Println()
	// thanhTung := sinhVien{
	// 	"Dau Thanh Tung",
	// 	21,
	// 	"CNTT.C1",
	// 	"Dien-Dien tu",
	// }
	// // hienThiThongTinSV(&thanhTung)
	// thanhTung.hienThiThongTinSV()
	// // =>khong nen dung khai bao kieu nay, de bi lan lon hoac nham cac du lieu
	// output,err := json.Marshal(quocHung)
	//can phai export de dung cac gia tri: name=>Name
	// if err != nil{
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println(string(output))
	// =>vi output dang dau ra la byte nen gan du lieu cua no thanh string: string(output)
	//bai6
	// thongTin := hcn{
	// 	chieuDai: 5,
	// 	chieuRong: 4,
	// }
	// fmt.Println("Chu vi hinh chu nhat:",thongTin.chuViHCN())
	// fmt.Println()
	// fmt.Println("Dien tich hinh chu nhat:",thongTin.dienTichHCN())
	kichThuoc := inputV2()
	fmt.Println()
	fmt.Println("Chieu dai hinh chu nhat:", kichThuoc.chieuDai)
	fmt.Println("Chieu rong hinh chu nhat:", kichThuoc.chieuRong)
	fmt.Println()
	fmt.Println("Chu vi hinh chu nhat:", kichThuoc.chuViHCN())
	fmt.Println()
	fmt.Println("Dien tich hinh chu nhat:", kichThuoc.dienTichHCN())
	fmt.Println()
}
