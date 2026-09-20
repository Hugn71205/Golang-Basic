package service

type Ten interface {
	Ho() string
	Getname() string
}
type Tenplus interface {
	Ten
	Gioitinh() string
}
type Tendacdiem interface {
	Dacdiem() string
}