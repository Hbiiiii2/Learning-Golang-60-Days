package main

import "fmt"

func main(){
	var name string // 1. -> deklarasi varaible dan harus di inisialisasi terlebih dahulu

	name = "Endulita mamah"
	fmt.Println("Hello", name)

	name = "Golang"
	fmt.Println("Aku pakai", name)

	var age = 24  // 2. -> deklarasi variable tanpa inisialisasi atau otomatis age = 24
	fmt.Println("Umur saya", age)

	varShort := "Abdul gany" // 3. -> deklarasi variable dengan cara singkat (hanya bisa di dalam function)
	fmt.Println("Nama Saya", varShort)

	varInt := 30 
	fmt.Println("Umur saya", varInt)

	varInt = 35 // kalau mengubah value variable dan tipe data yg sama, tidak perlu pakai := cukup pakai =
	fmt.Println("Umur saya sekarang", varInt)

	// Multiple Variable

	var (
		firstName = "Udin"
		LastName = "Awaludin"
	)
	fmt.Println(firstName, LastName)
}
