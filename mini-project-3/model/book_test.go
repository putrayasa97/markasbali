package model_test

import (
	"fmt"
	"sekolahbeta/mini-project-3/config"
	"sekolahbeta/mini-project-3/model"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("env not found, using system env")
	}
	ConnDB()
}

func ConnDB() {
	config.OpenDB("testing")
}

func TestBookCreate(t *testing.T) {
	Init()

	bookData := model.Book{
		ISBN:    "316012823",
		Penulis: "Kate Atkinson",
		Tahun:   2006,
		Judul:   "One Good Turn",
		Gambar:  "https://s.gr-assets.com/assets/nophoto/book/111x148-bcc042a9c91a29c1d680899eff700a03.png",
		Stok:    64,
	}
	// create data buku
	err := bookData.Create(config.Mysql.DB)
	assert.Nil(t, err)
}

func TestBookGetByID(t *testing.T) {
	Init()

	bookData := model.Book{
		ISBN:    "316012823",
		Penulis: "Kate Atkinson",
		Tahun:   2006,
		Judul:   "One Good Turn",
		Gambar:  "https://s.gr-assets.com/assets/nophoto/book/111x148-bcc042a9c91a29c1d680899eff700a03.png",
		Stok:    64,
	}
	// create data buku baru
	err := bookData.Create(config.Mysql.DB)
	assert.Nil(t, err)

	bookData = model.Book{
		Model: model.Model{
			ID: bookData.ID,
		},
	}
	// get data buku yg telah dibuat
	book, err := bookData.GetByID(config.Mysql.DB)
	assert.Nil(t, err)
	assert.Equal(t, bookData.ID, book.ID)
}

func TestBookGetAll(t *testing.T) {
	Init()

	bookData := model.Book{
		ISBN:    "316012823",
		Penulis: "Kate Atkinson",
		Tahun:   2006,
		Judul:   "One Good Turn",
		Gambar:  "https://s.gr-assets.com/assets/nophoto/book/111x148-bcc042a9c91a29c1d680899eff700a03.png",
		Stok:    64,
	}

	err := bookData.Create(config.Mysql.DB)
	assert.Nil(t, err)

	books, err := bookData.GetAll(config.Mysql.DB)

	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(books), 1)
}

func TestBookUpdateByID(t *testing.T) {
	Init()

	// data buku pertama
	bookData1 := model.Book{
		ISBN:    "316012823",
		Penulis: "Kate Atkinson",
		Tahun:   2006,
		Judul:   "One Good Turn",
		Gambar:  "https://s.gr-assets.com/assets/nophoto/book/111x148-bcc042a9c91a29c1d680899eff700a03.png",
		Stok:    64,
	}

	// create data pertama
	err := bookData1.Create(config.Mysql.DB)
	assert.Nil(t, err)

	// data buku kedua
	bookData2 := model.Book{
		Model: model.Model{
			ID: bookData1.ID,
		},
		ISBN:    "525425888",
		Penulis: "Nina LaCour",
		Tahun:   2014,
		Judul:   "Everything Leads to You",
		Gambar:  "https://images.gr-assets.com/books/1389744233m/18667779.jpg",
		Stok:    6,
	}

	// update data pertama dengan data kedua
	err = bookData2.UpdateByID(config.Mysql.DB)
	assert.Nil(t, err)

	// panggil data yg sudah di update
	book, err := bookData2.GetByID(config.Mysql.DB)
	assert.Nil(t, err)

	// check data ke kedua dengan data yg telah ubah dari database
	assert.Equal(t, bookData2.ISBN, book.ISBN)
	assert.Equal(t, bookData2.Penulis, book.Penulis)
	assert.Equal(t, bookData2.Tahun, book.Tahun)
	assert.Equal(t, bookData2.Judul, book.Judul)
	assert.Equal(t, bookData2.Gambar, book.Gambar)
	assert.Equal(t, bookData2.Stok, book.Stok)
}

func TestDeleteByID(t *testing.T) {
	Init()

	bookData := model.Book{
		ISBN:    "525425888",
		Penulis: "Nina LaCour",
		Tahun:   2014,
		Judul:   "Everything Leads to You",
		Gambar:  "https://images.gr-assets.com/books/1389744233m/18667779.jpg",
		Stok:    6,
	}
	// create data buku
	err := bookData.Create(config.Mysql.DB)
	assert.Nil(t, err)

	// cek data buku yg dibuat dengan yg di get
	book, err := bookData.GetByID(config.Mysql.DB)
	assert.Nil(t, err)
	assert.Equal(t, bookData.ID, book.ID)

	// hapus data buku
	err = bookData.DeleteByID(config.Mysql.DB)
	assert.Nil(t, err)

	// cek data buku dengan yg sudah di delete
	// bandingkan id buku dan cek deleted_at nya
	book, err = bookData.GetByIDWithDelete(config.Mysql.DB)
	assert.Nil(t, err)
	assert.Equal(t, bookData.ID, book.ID)
	assert.NotNil(t, book.DeletedAt)
}
