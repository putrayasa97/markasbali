package model

import (
	"gorm.io/gorm"
)

type Book struct {
	Model
	ISBN    string `json:"isbn"`
	Penulis string `json:"penulis"`
	Tahun   uint   `json:"tahun"`
	Judul   string `json:"judul"`
	Gambar  string `json:"gambar"`
	Stok    uint   `json:"stok"`
}

// metode untuk create data buku
func (typeBook *Book) Create(db *gorm.DB) error {
	err := db.Model(Book{}).Create(&typeBook).Error

	if err != nil {
		return err
	}

	return nil
}

// metode untuk mengambil data buku berdasarkan id
func (typeBook *Book) GetByID(db *gorm.DB) (Book, error) {
	book := Book{}
	err := db.Model(Book{}).Where("id = ?", typeBook.Model.ID).Take(&book).Error

	if err != nil {
		return Book{}, err
	}

	return book, nil
}

// metode untuk mengambil data buku berdasarkan id tanpa scope soft delete
func (typeBook *Book) GetByIDWithDelete(db *gorm.DB) (Book, error) {
	book := Book{}
	err := db.Unscoped().Model(Book{}).Where("id = ?", typeBook.Model.ID).Take(&book).Error

	if err != nil {
		return Book{}, err
	}

	return book, nil
}

// metode untuk mengambil semua data buku
func (typeBook *Book) GetAll(db *gorm.DB) ([]Book, error) {
	books := []Book{}
	err := db.Model(Book{}).Find(&books).Error

	if err != nil {
		return []Book{}, err
	}

	return books, nil
}

// metode untuk ubah data buku berdasarkan id
func (typeBook *Book) UpdateByID(db *gorm.DB) error {
	err := db.Model(Book{}).
		Select(
			"isbn",
			"penulis",
			"tahun",
			"judul",
			"gambar",
			"stok").
		Where("id = ?", typeBook.Model.ID).
		Updates(map[string]interface{}{
			"isbn":    typeBook.ISBN,
			"penulis": typeBook.Penulis,
			"tahun":   typeBook.Tahun,
			"judul":   typeBook.Judul,
			"gambar":  typeBook.Gambar,
			"stok":    typeBook.Stok,
		}).Error

	if err != nil {
		return err
	}

	return nil
}

// metode untuk hapus data buku berdasarkan id
func (typeBook *Book) DeleteByID(db *gorm.DB) error {
	err := db.Model(Book{}).
		Where("id = ?", typeBook.Model.ID).
		Delete(&typeBook).Error

	if err != nil {
		return err
	}

	return nil
}
