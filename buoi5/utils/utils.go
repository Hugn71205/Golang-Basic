package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// Đọc dữ liệu từ bàn phím và trả về chuỗi đã loại bỏ khoảng trắng thừa
func ReadInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Nhập và kiểm tra số nguyên dương (> 0)
func GetPostiveInt(prompt string) int {
	for {
		input := ReadInput(prompt)
		value, err := strconv.Atoi(input)
		if err == nil && value > 0 {
			return value
		}

		fmt.Println("❌ Gia tri khong hop le! Vui long thu lai.")
	}
}

// Nhập và kiểm tra số thực dương hoặc bằng 0 (>= 0)
func GetPostiveFloat(prompt string) float64 {
	for {
		input := ReadInput(prompt)
		value, err := strconv.ParseFloat(input, 64)
		if err == nil && value >= 0 {
			return value
		}

		fmt.Println("❌ Gia tri khong hop le! Hay nhap so thuc duong.")
	}
}

// Nhập và kiểm tra chuỗi không được để trống
func GetNonEmptyString(prompt string) string {
	for {
		input := ReadInput(prompt)
		if input != "" {
			return input
		}

		fmt.Println("❌ Gia tri khong duoc bo trong! Hay nhap it nhat mot ky tu.")
	}
}

// Nhập chuỗi tùy chọn, nếu bỏ trống thì giữ nguyên giá trị cũ
func GetOptionalString(prompt string, oldValue string) string {
	input := ReadInput(prompt)
	if input == "" {
		return oldValue
	}

	return input
}

// Nhập số thực tùy chọn, nếu bỏ trống thì giữ nguyên giá trị cũ
func GetOptionalPostiveFloat(prompt string, oldValue float64) float64 {
	input := ReadInput(prompt)
	if input == "" {
		return oldValue
	}

	value, err := strconv.ParseFloat(input, 64)
	if err != nil && value < 0 {
		fmt.Println("❌ Gia tri khong hop le! Giu nguyen gia tri cu")
		return oldValue
	}

	return value
}

// Xóa dữ liệu hiển thị cũ trên màn hình và chỉ hiển thị dữ liệu mới
func ClearScreen() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Error clearing screen: ", err)
	}
}

// Interface yêu cầu kiểu dữ liệu phải có hàm GetId() trả về ID
type HasId interface {
	GetId() int
}

// Kiểm tra ID có bị trùng trong danh sách hay chưa
func IsIdUnique[T HasId](id int, list []T) bool {
	for _, item := range list {
		if item.GetId() == id {
			return false
		}
	}

	return true
}