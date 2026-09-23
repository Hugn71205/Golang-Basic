package hung

import (
	"errors"
	"strings"
)

type Hung struct {
	Name string
}

// tao constructor cho Hung
func New(name string) (*Hung, error) {
	//Trimspace kiem tra khoang trong truoc va sau cua ten
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("Ten khong duoc de trong!")
		//len kiem tra so ky tu nhap vao
	} else if len(name) > 50 {
		return nil, errors.New("Vuot qua so luong ky tu cho phep!")

	}
	return &Hung{
		Name: name,
	}, nil
}
func (h *Hung) Getname() string {

	return h.Name
}
func (h *Hung) Ho() string {
	return "Bui Quoc"
}