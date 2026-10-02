package greetings

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func Hello(name string) (string, error) {
	if name == "" {
		//khoi tao loi moi bang errors.New("")
		return "", errors.New("Vui long nhap ten")
	}
	//sử dụng random greetings để trả về câu chào ngẫu nhiên
	message := fmt.Sprintf(randomGreeting(), name)
	return message, nil
}
func randomGreeting() string {
	random := []string{
		"Xin chao, %v\n",
		"Kinh chao quy khach mang ten %v\n",
		"Have a great day mr.%v\n",
	}
	//dùng thư viện math phương thức rand để gọi random theo số ở đây chúng ta random và return số ngẫu nhiên từ 0 đến n-1 để gọi ra số thứ tự của các phần tử trong slice:n ở đây là độ lài slice(3)
	return random[rand.IntN(len(random))]
}
func Hellos(names []string) (map[string]string, error) {
	loichaos := make(map[string]string)
	for _, name := range names {
		loichao, err := Hello(name)
		if err != nil {
			return nil, err
		}
		loichaos[name] = loichao
	}
	return loichaos, nil
}
