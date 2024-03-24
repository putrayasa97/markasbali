package menu

import (
	"fmt"
	"os"
	"sekolahbeta/mini-project-3/utils/book"
	"sekolahbeta/mini-project-3/utils/helpers"
)

func PrintMenu() {
	var opsiPrint int
	fmt.Println("Silakan pilih opsi print : ")
	fmt.Println("1. Print Berdasarkan ID Buku")
	fmt.Println("2. Print Semua Buku")
	fmt.Println("3. Kembali ke menu utama")
	helpers.Line()

	os.Mkdir("pdf", 0777)

	helpers.LineInput("Masukan Pilihan : ", &opsiPrint)

	switch opsiPrint {
	case 1:
		book.BookPrintByID()
	case 2:
		book.BookPrintAll()
	case 3:
		MainMenu()
	case 4:
		os.Exit(0)
	}
}
