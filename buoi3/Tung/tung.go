package tung

import (
	"Demo/buoi3/service"
	"errors"
	"strings"
)

type Tung struct {
	Name string
}

// tao constructor cho Tung
// khi chi thang den interface thi chi goi cac du lieu trong interface do con goi tat ca thi goi den struct *Tung
func New(name string) (service.Tendacdiem, error) {
	//Trimspace kiem tra khoang trong truoc va sau cua ten
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("Ten khong duoc de trong!")
		//len kiem tra so ky tu nhap vao
	} else if len(name) > 50 {
		return nil, errors.New("Vuot qua so luong ky tu cho phep!")

	}
	return &Tung{
		Name: name,
	}, nil
}
func (t *Tung) Getname() string {
	return t.Name
}
func (t *Tung) Ho() string {
	return "Do Thanh"
}
func (t *Tung) Gioitinh() string {
	return "Nam"
}
func (t *Tung) Dacdiem() string {
	return "Da den"
}
