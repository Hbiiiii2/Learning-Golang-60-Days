package main

import "fmt"

func main(){

	// var name string

	// name = "Cihuy"
	// fmt.Println("Hello",name)

	// name = "Golang"
	// fmt.Println("aku pakai",name)

	test()
}

func coba(){
	 umur := 20
	fmt.Println("umur saya",umur)
}

func test(){
	var umur int
	fmt.Println("Masukan Umur anda")
	fmt.Scanln(&umur)
	fmt.Println("Umur anda adalah",umur)
}