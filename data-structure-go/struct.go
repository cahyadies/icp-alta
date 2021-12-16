package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// Bisa menggunakan tipe data yang berbeda
	var orang1 Person
	orang1.Name = "Eko Sugeng Cahyadi"
	orang1.Age = 22

	var orang2 Person
	orang2.Name = "Fulan bin Fulan"
	orang2.Age = 25

	fmt.Println(orang1.Name)
	fmt.Println(orang2.Name)

}
