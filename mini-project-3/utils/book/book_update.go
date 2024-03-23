package book

import (
	"fmt"
	"sekolahbeta/mini-project-3/config"
	"sekolahbeta/mini-project-3/model"
	"sekolahbeta/mini-project-3/utils/helpers"
	"strconv"
)

var modelBook model.Book

func BookUpdate() {
	bookID := 0
	bookISBN := ""
	bookJudul := ""
	bookPenulis := ""
	bookTahun := 0
	bookGambar := ""
	bookStok := 0

	helpers.Line()
	fmt.Println("Ubah Buku")
	helpers.Line()

	helpers.LineInput("Masukan ID Buku yang ingin diubah : ", &bookID)
	modelBook.Model = model.Model{
		ID: uint(bookID),
	}
	book, err := modelBook.GetByID(config.Mysql.DB)
	if err != nil {
		fmt.Println("Buku tidak ditemukan !")
	}

	helpers.Line()

	countChange := 0
	confirm := helpers.LineConfirm("Apa anda ingin merubah ISBN Buku ?")
	if confirm {
		helpers.LineInput("Judul Buku Sebelumnya '"+book.ISBN+"' : ", &bookISBN)
		modelBook.ISBN = bookISBN
		countChange += 1
	}
	confirm = helpers.LineConfirm("Apa anda ingin merubah Judul Buku ?")
	if confirm {
		helpers.LineInput("Judul Buku Sebelumnya '"+book.Judul+"' : ", &bookJudul)
		modelBook.Judul = bookJudul
		countChange += 1
	}
	confirm = helpers.LineConfirm("Apa anda ingin merubah Penulis Buku ?")
	if confirm {
		helpers.LineInput("Penulis Buku Sebelumnya '"+book.Penulis+"' : ", &bookPenulis)
		modelBook.Penulis = bookPenulis
		countChange += 1
	}
	confirm = helpers.LineConfirm("Apa anda ingin merubah Tahun Buku ?")
	if confirm {
		helpers.LineInput("Tahun Buku Sebelumnya '"+strconv.Itoa(int(book.Tahun))+"' : ", &bookTahun)
		modelBook.Tahun = uint(bookTahun)
		countChange += 1
	}
	confirm = helpers.LineConfirm("Apa anda ingin merubah Jumlah Gambar Buku ?")
	if confirm {
		helpers.LineInput("Jumlah Halaman Buku Sebelumnya '"+book.Gambar+"' : ", &bookGambar)
		modelBook.Gambar = bookGambar
		countChange += 1
	}
	confirm = helpers.LineConfirm("Apa anda ingin merubah Stok Buku ?")
	if confirm {
		helpers.LineInput("Tahun Terbit Buku Sebelumnya '"+strconv.Itoa(int(book.Stok))+"' : ", &bookStok)
		modelBook.Stok = uint(bookStok)
		countChange += 1
	}

	if countChange == 0 {
		fmt.Println("Tidak ada perubahan data buku!")
	}

	bookUpdate()
	fmt.Println("Berhasil Merubah Buku!")
}

func bookUpdate() {
	err := modelBook.UpdateByID(config.Mysql.DB)
	if err != nil {
		fmt.Println(err)
	}
}
