package main

import "fmt"

func main() {
	var nama string = "adit"
	// variable pointer referensi dari variable nama dengan alamat memory
	var namaPointer *string = &nama

	fmt.Println("nama asal :", nama)
	fmt.Println("ini string dengan &: ", &nama)
	fmt.Println("ini variable namaPointer: ", namaPointer)

	// nama = "adit"
	// merubah value dari memory
	*namaPointer = "pgpy"
	fmt.Println(*namaPointer) // mengambil value
	fmt.Println(namaPointer)  // memanggil alamat memory

}
