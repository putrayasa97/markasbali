package book

import (
	"fmt"
	"sekolahbeta/mini-project-3/config"
	"sekolahbeta/mini-project-3/model"
	"sekolahbeta/mini-project-3/utils/helpers"
	"sync"
)

var draftListBook []model.Book

func BookCreate() {
	bookISBN := ""
	bookJudul := ""
	bookPenulis := ""
	bookTahun := 0
	bookGambar := ""
	bookStok := 0

	draftListBook = []model.Book{}

	helpers.Line()
	fmt.Println("Tambah Buku")
	helpers.Line()

	for {
		helpers.LineInput("Masukan ISBN : ", &bookISBN)
		helpers.LineInput("Masukan Judul : ", &bookJudul)
		helpers.LineInput("Masukan Penulis : ", &bookPenulis)
		helpers.LineInput("Masukan Tahun : ", &bookTahun)
		helpers.LineInput("Masukan Gambar : ", &bookGambar)
		helpers.LineInput("Masukan Stok: ", &bookStok)

		draftListBook = append(draftListBook, model.Book{
			ISBN:    bookISBN,
			Penulis: bookPenulis,
			Tahun:   uint(bookTahun),
			Judul:   bookJudul,
			Gambar:  bookGambar,
			Stok:    uint(bookStok),
		})

		var nextInput = 0
		helpers.LineInput("Ketikan 1 untuk tambah buku, ketika 0 untuk keluar : ", &nextInput)
		if nextInput == 0 {
			break
		}
	}

	bookSave()
	fmt.Println("Berhasil Menambah Buku!")
}

func bookSave() {
	chSave := make(chan model.Book)
	wgSave := sync.WaitGroup{}
	jmlThread := 5

	for i := 0; i < jmlThread; i++ {
		wgSave.Add(1)
		go func(ch <-chan model.Book, wg *sync.WaitGroup) {
			for book := range ch {
				err := book.Create(config.Mysql.DB)
				if err != nil {
					fmt.Println(err)
				}
			}
			wgSave.Done()
		}(chSave, &wgSave)
	}

	for _, book := range draftListBook {
		chSave <- book
	}

	close(chSave)
	wgSave.Wait()
}
