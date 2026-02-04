// Tipe data Number !

// Integer :
// int8 : -128 - 127
// int16 : -32768 - 32767
// int32 : -2147483648 - 2147483647
// int64 : -9223372036854775808 - 9223372036854775807

// !note : int dan uint adalah tipe data yang sama, tapi int adalah tipe data yang signed (bernilai negatif dan positif) dan uint adalah tipe data yang unsigned (bernilai positif)
// uint8 : 0 - 255
// uint16 : 0 - 65535
// uint32 : 0 - 4294967295
// uint64 : 0 - 18446744073709551615

// Floating Point :

// float32 : 1.18e-38 - 3.402823e+38
// float64 : 2.23e-308 - 1.79e+308

// !note : float32 dan float64 adalah tipe data yang sama, tapi float32 adalah tipe data yang presisi rendah dan float64 adalah tipe data yang presisi tinggi

//!note : complex64 dan complex128 adalah tipe data yang sama, tapi complex64 adalah tipe data yang presisi rendah dan complex128 adalah tipe data yang presisi tinggi
// complex64
// complex128

// Alias tipe data :
// byte : uint8
// rune : int32
// int : minimal int32 -> berdasarkan sistem operasi (32 bit atau 64 bit)
// uint : minimal uint32 -> berdasarkan sistem operasi (32 bit atau 64 bit)

package main

import "fmt"

func main() {
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("tiga koma lima = ", 3.5)
}
