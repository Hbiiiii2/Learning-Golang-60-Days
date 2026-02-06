package main

import "fmt"

func main(){
	const name  = "Nama Konstanta"
	const lastname  = "Nama Akhir"
	const firstName  = "Nama Awal"
	


	// constanta -> tidak akan bisa di ubah isi value nya dan tidak akan di complaine ketika tidak di pakai 

	fmt.Println(name)
	fmt.Println(firstName)
	fmt.Println(lastname)

	// multiple Constnta
	const (
		contry = "Indonesia"
		city = "Jakarta"
		Postalcode = "122201"
	)

	fmt.Println("\nAnda berasal dari negara :",contry)
	fmt.Println("Kota :",city)
	fmt.Println("Kode Pos :",Postalcode)
}