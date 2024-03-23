package menu

import (
	"fmt"
	"os"
	"sekolahbeta/mini-project-3/utils/book"
	"sekolahbeta/mini-project-3/utils/helpers"
)

func MainMenu() {
	var pilihMenu int
	helpers.Line()
	fmt.Println("Aplikasi Manajemen Daftar Buku Perpustakaan")
	helpers.Line()
	fmt.Println("Silakan Pilih Menu : ")
	fmt.Println("1. Tambah Buku")
	fmt.Println("2. Lihat Daftar Buku")
	fmt.Println("3. Ubah Buku")
	fmt.Println("4. Hapus Buku")
	fmt.Println("5. Print Buku")
	fmt.Println("6. Keluar")
	helpers.Line()
	helpers.LineInput("Masukan Pilihan : ", &pilihMenu)
	switch pilihMenu {
	case 1:
		book.BookCreate()
	case 2:
		book.BookGet()
	case 3:
		book.BookUpdate()
	case 4:
		book.BookDelete()
	// case 5:
	// 	printBuku()
	case 6:
		os.Exit(0)
	}
	MainMenu()
}
