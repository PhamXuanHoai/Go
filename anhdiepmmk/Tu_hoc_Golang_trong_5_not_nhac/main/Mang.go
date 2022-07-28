package main

import "fmt"

func main()  {

	var n int
	var arr []int

	fmt.Print("Nhập n phần từ: ")
	fmt.Scan(&n)

	XuatMang(NhapMang(n, arr))

}

func NhapMang(n int, arr []int) (result []int) {

	for i := 0; i < n; i++ {
		var tmpN int
		fmt.Print("arr[", i, "] = ")
		fmt.Scan(&tmpN)
		arr = append(arr, tmpN)
	}

	return arr
}

func XuatMang(arr []int)  {

	doDai := len(arr)

	fmt.Print("Mảng vừa nhập: ")
	for i := 0; i < doDai; i++ {
		fmt.Print(arr[i], " ")
	}
}