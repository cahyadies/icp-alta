package main

import "fmt"

func main() {

	// Tidak spesifik value index yang dipasang pada []
	bulan := []string{"Januari", "Februari", "Maret", "April"}
	fmt.Println(bulan)
	// Menambahkan data pada Slices
	// Cocok untuk data yang dinamis, beda dengan array yang datanya statis
	bulan = append(bulan, "Mei")
	fmt.Println(bulan)

}
