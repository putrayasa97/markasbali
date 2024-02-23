package main

import "fmt"

func main() {
	// Tipe data
	var angka int8 = 5
	var angka1 float32 = -89.7
	var belajar bool = true
	nama := "putu"

	// %s string
	// %d integer
	// %f float
	// %v bool
	fmt.Printf(`nama saya: %s, angka saya adalah: %d %f saya belajar golang %v`, nama, angka, angka1, belajar)

	fmt.Println("Test")

	// deklarasi multi variabel

	// const variable

	// array & slice
	// [...] jumlah array tidak terhingga slice
	// [5] jumlah array bisa ditentukan array

	var angkaArray [2]int
	angkaArray[0] = 2

	angkaArray1 := [2]int{1, 2}
	angkaSlice := []string{"saya", "mantap"}
	angkaSlice1 := []string{"saya1", "mantap1"}

	// array multi dimensi
	arrayMulti := [2][2]string{{"saya", "mantap"}, {"saya1", "mantap1"}}
	arrayMulti1 := [][]string{{"saya", "mantap", "0"}, {"saya1", "mantap1", "2"}}

	fmt.Println(angkaArray[0])
	fmt.Println(angkaArray1[0], angkaArray1[1])
	fmt.Println(angkaSlice)
	fmt.Println(arrayMulti)
	fmt.Println(arrayMulti1)

	// operator slice
	// - len untuk mengetahui jumlah data pada array
	fmt.Println(len(angkaSlice))
	// - append untuk menambahkan nilai ke array
	angkaSlice = append(angkaSlice, "pgpy")
	fmt.Println(angkaSlice)
	// - make untuk membuat slice
	makeSlice := make([]string, 3)
	// - copy untuk menduplikasi array
	copy(makeSlice, angkaSlice1)
	fmt.Println(makeSlice)

	// map
	// merupakan tipe data asosiatif
	// terdapat key dan value
	// cara deklrasi
	// var varMap map[string]string
	// varMaps := map[string]string{"IDR": "Rp", "USD": "$"}
	// varMaps := map[string]string{
	//	"PERSETA_1": "ANI",
	//	"PERSETA_2": "Budi"
	// }

	// operasi map
	// - delete untuk menghapus data berdasarkan key
	// - merubah data
	// - isExist untuk mengecek value data pada map berdasarkan key

	mataUang := map[string]string{
		"USD": "$",
		"IDR": "Rp",
	}
	fmt.Println(mataUang)

	// merubah data
	mataUang = map[string]string{
		"JPY": "Y",
		"USD": "Dollar",
	}
	fmt.Println(mataUang)
	fmt.Println(mataUang["USD"], mataUang["IDR"], mataUang["JPY"])

	// delete data
	delete(mataUang, "JPY")
	fmt.Println(mataUang)

	// isExist
	value, isExist := mataUang["USD"]
	fmt.Println(value, isExist)

	// menampilkan value
	testing1 := []map[string]int{
		{
			"testing1": 1,
			"testing2": 2,
		},
		{"testing3": 3},
		{"testing4": 4},
	}
	fmt.Println(testing1[1]["testing3"])

	// Operator Golang
	// operator aritmatika +,/ ,-,*
	// operator permbanding ==, !=, <, >
	// operator logika &&, ||
	// conditional if else, switch case
	// perulangan for loop

	// style 1
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// style 2
	var i int = 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// style 3
	var a int = 0
	for {
		fmt.Println(a)
		a++
		if a > 5 {
			break
		} else {
			continue
		}
	}
}
