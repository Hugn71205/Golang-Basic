package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Box[T any] struct {
	Contain T
	Mota    T
}
//Dinh nghia du lieu cho du lieu dau vao
type Hamdinhnghia interface {
	constraints.Integer | constraints.Float
}
func MaxLenghtString(a,b string) string{
	if len(a)>len(b){
		return a
	}
	return b
}
func Tinhtoan[Numb Hamdinhnghia](a, b Numb) Numb {
	return a + b
}
func buoi10() {
	intBox := Box[int]{Contain: 57, Mota: 55}

	stringBox := Box[string]{Contain: "Hugn", Mota: "Noi dung ben trong"}
	PrintValue(stringBox.Contain)
	PrintValue(intBox.Mota)

	PrintValue(stringBox.Mota)
	PrintValue(intBox.Contain)
	fmt.Println()
	fmt.Println(MaxLenghtString("Trung", "Tung"))
	fmt.Println()
	fmt.Println(Tinhtoan(5,7.6))
	fmt.Println(Tinhtoan(5.5,7.2))
}