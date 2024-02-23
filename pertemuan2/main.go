package main

import "fmt"

func main() {
	urutanAngka := map[string]int{
		"satu":  1,
		"dua":   2,
		"tiga":  3,
		"empat": 4,
	}
	// Looping
	// looping map dengan range
	for key, value := range urutanAngka {
		fmt.Println("Key :", key)
		fmt.Println("Value :", value)
	}

	// Function
	hello()
	hasilKonversi := konversiMataUangUSD(2)
	fmt.Println(hasilKonversi)

	hasilKonversi1, _ := konversiMataUang(5000, "JPY")
	fmt.Println(hasilKonversi1)

	angkaInput := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	hasilHitungKalkulator := kalkulator("+", angkaInput...)
	fmt.Println(hasilHitungKalkulator)

	// Struct

}

func kalkulator(operator string, angka ...int) int {
	var hasil int
	for i := 0; i < len(angka); i++ {
		if operator == "+" {
			hasil += angka[i]
		} else if operator == "-" {
			hasil -= angka[1]
		}
	}
	return hasil
}

// fungsi bersisi return atau keluaran dengan tipe integer
// dan meiliki parameter uang dengan type int
func konversiMataUangUSD(uang int) int {
	var hasil int
	hasil = uang * 15000
	return hasil

}

func konversiMataUang(uang int, currency string) (int, string) {
	var hasil int
	switch currency {
	case "USD":
		hasil = uang * 15000
	case "JPY":
		hasil = uang * 300
	default:
		hasil = 0
	}
	return hasil, currency
}

// fungsi bisa dibuat di atas atau dibawah fungsi main
// fungsi void tidak ada return
func hello() {
	fmt.Println("Hello World")
}

// function
// struct
// pointer
