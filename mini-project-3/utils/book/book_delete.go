package book

import (
	"fmt"
	"sekolahbeta/mini-project-3/config"
	"sekolahbeta/mini-project-3/model"
	"sekolahbeta/mini-project-3/utils/helpers"
)

func BookDelete() {
	var modelBook model.Book
	bookID := 0

	helpers.Line()
	fmt.Println("Hapus Buku")
	helpers.Line()

	helpers.LineInput("Masukan ID Buku yang ingin dihapus : ", &bookID)

	modelBook.Model = model.Model{
		ID: uint(bookID),
	}

	book, err := modelBook.GetByID(config.Mysql.DB)
	if err != nil {
		fmt.Println("Buku tidak ditemukan!")
	}

	err = book.DeleteByID(config.Mysql.DB)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Berhasil Menghapus Buku!")
}
