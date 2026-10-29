package main

import (
	"fmt"
	"log"
	"sync"
)

type TanCong interface {
	Danh(muctieu string)
}
type PhongThu interface {
	Dodon(muctieu string)
}
type Counter interface {
	PhongThu
	TanCong
}
type DauSi struct {
	Ten        string
	Dobenkhien int
}

func (d *DauSi) Danh(muctieu string) {
	fmt.Printf("%v da tan cong %v\n", d.Ten, muctieu)
}
func (d *DauSi) Dodon(muctieu string) {
	fmt.Printf("%v da do don cua %v\n", d.Ten, muctieu)
	d.Dobenkhien = d.Dobenkhien - 10
	if d.Dobenkhien <= 50 {
		fmt.Printf("Khien sap hong!\n")
	}
	fmt.Printf("Do ben khien con lai sau khi do don:%v\n", d.Dobenkhien)
}

type PhapSu struct {
	Ten string
}

func (p *PhapSu) Danh(muctieu string) {
	fmt.Printf("%v da tan cong %v\n", p.Ten, muctieu)
}
func main() {
	log.SetPrefix("Loi cua ban la: ")
	log.SetFlags(0)
	message, err := Hello("Tunbeg")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
	fmt.Println()
	names := []string{
		"Tunbeg",
		"Diablo",
	}
	messages, err := Hellos(names)
	fmt.Println(messages)
	fmt.Println()
	minNumb := []int{1, 2, 3, 4, 5}
	minVal := minNumb[0]
	for i := 1; i < len(minNumb); i++ {
		if minNumb[i] < minVal {
			minVal = minNumb[i]
		}
	}
	fmt.Printf("So nho nhat trong day la: %v", minVal)
	fmt.Println()
	maxNumb := []int{1, 2, 3, 4, 5}
	maxVal := maxNumb[0]
	for i := 1; i < len(maxNumb); i++ {
		if maxNumb[i] > maxVal {
			maxVal = maxNumb[i]
		}
	}
	fmt.Printf("So lon nhat trong day la: %v", maxVal)
	fmt.Println()
	pho := "ngon"
	switch pho {
	case "ngon":
		fmt.Println("Pho ngon")
	case "te":
		fmt.Println("Pho khong ngon")
	default:
		fmt.Println("Chua an chua biet duoc")
	}
	ch := make(chan int, 10)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Printf("So trong channel la:%v\n", i)
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 17
		ch <- 24
		ch <- 28
		ch <- 22
		ch <- 21
		ch <- 50
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
	Hung := &DauSi{
		Ten:        "Diablo",
		Dobenkhien: 70,
	}
	Tung := &PhapSu{
		Ten: "Tunbeg",
	}
	var tc1 TanCong = Hung
	tc1.Danh("Goblin")
	var pt1 PhongThu = Hung
	pt1.Dodon("Goblin")
	var tc2 TanCong = Tung
	tc2.Danh("Elf")
	var ct Counter = Hung
	ct.Dodon("Dragon")
	ct.Danh("Dragon")
}
