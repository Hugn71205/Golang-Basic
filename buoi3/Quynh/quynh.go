package quynh

import (
	"errors"
	"strings"
)

type Quynh struct {
	Name string
}

// tao constructor cho Quynh
func New(name string) (*Quynh, error) {
	//Trimspace kiem tra khoang trong truoc va sau cua ten
	name = strings.TrimSpace(name)
	if name == ""{
	return nil, errors.New("Ten khong duoc de trong!")
	//len kiem tra so ky tu nhap vao
	}else if len(name) > 50 {
	return nil, errors.New("Vuot qua so luong ky tu cho phep!")

	}
	return &Quynh{
		Name: name,
	},nil
}
func (q *Quynh) Getname() string {
	return q.Name
}
func (q *Quynh) Ho() string {
	return "Ngo Nhu"
}
func (q *Quynh) Gioitinh() string {
	return "Nu"
}