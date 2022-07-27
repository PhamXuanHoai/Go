package main

import "fmt"

func main()  {

	for true {
		var n int

		fmt.Print("Mời bạn nhập số: ")
		fmt.Scan(&n)

		if n % 2 == 0 {
			fmt.Println(n, " là số chẵn")
		} else {
			fmt.Println(n, " là số lẻ")
		}
	}
}
