package main

func buoi13() {
	// TÌM HIỂU HÀM SLICES
	/** Clone: Tạo bản sao của slice **/
	// copied := slices.Clone([]int{1, 2, 3})
	// fmt.Println(copied)
	// fmt.Println()

	/** So sánh 2 slice có giống nhau không **/
	// compareSlice := slices.Equal([]int{1, 2, 3}, []int{1, 2})
	// fmt.Println(compareSlice)
	// fmt.Println()

	/** Tìm vị trí dau tien của phần tử **/
	// findFirstPosition := slices.Index([]int{1, 2, 3, 4, 2}, 2)
	// fmt.Println(findFirstPosition)
	// fmt.Println()

	/** Kiểm tra phần tử có trong slice không **/
	// itemExist := slices.Contains([]int{1, 2, 3}, 5)
	// fmt.Println(itemExist)
	// fmt.Println()

	/** slices.Insert(s, i, v...) Chèn phần tử vào vị trí i  **/
	// insertValueAnyPostion := slices.Insert([]int{1, 3}, 1, 2)
	// fmt.Println(insertValueAnyPostion)
	// fmt.Println()

	/** slices.Delete(s, i, j) Xóa phần tử từ vị trí i đến j-1 **/
	// deleteValueAnyPostion := slices.Delete([]int{1, 2, 3, 4}, 1, 3)
	// fmt.Println(deleteValueAnyPostion)
	// fmt.Println()

	/** Đảo ngược slice **/
	// s := []int{1, 2, 3}
	// slices.Reverse(s)
	// fmt.Println(s)
	// fmt.Println()

	/** Sắp xếp slice tăng dần **/
	// td := []int{3, 1, 2}
	// slices.Sort(td)
	// fmt.Println(td)
	// fmt.Println()

	/** Sắp xếp theo điều kiện tùy chỉnh **/
	// tc := []int{3, 1, 2}
	// slices.SortFunc(tc, func(a, b int) int {
	// 	// return a - b
	// 	return b - a
	// })
	// fmt.Println(tc)
	// fmt.Println()

	/** Lấy giá trị lớn nhất **/
	// itemMax := slices.Max([]int{1, 3, 2})
	// fmt.Println(itemMax)
	// fmt.Println()

	/** Lấy giá trị nhỏ nhất **/
	// itemMin := slices.Min([]int{1, 3, 2})
	// fmt.Println(itemMin)
}