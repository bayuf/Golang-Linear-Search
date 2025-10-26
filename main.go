package main

import "fmt"

func main() {
	data := []int{100, 200, 50, 500, 400, 300}

	cari := 300
	for i, value := range data {
		if value == cari {
			fmt.Println("Data :", value, "Berada pada indeks", i)
		}
	}
}
