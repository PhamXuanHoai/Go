package main

import "fmt"

func main()  {

	var a = 0
	var b = 0

	fmt.Print("Nhập vào số a: ")
	fmt.Scan(&a)

	fmt.Print("Nhập vào số b: ")
	fmt.Scan(&b)

	fmt.Println("Chưa hoán đổi: ", a, " - ", b)
	hoanDoi(&a, &b)
	fmt.Print("Đã hoán đổi: ", a, " - ", b)

}

func hoanDoi(a *int, b *int)  {
	*a = *a * *b
	*b = *a / *b
	*a = *a / *b
}