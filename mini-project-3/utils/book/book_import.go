package book

import (
	"encoding/csv"
	"fmt"
	"os"
	"sekolahbeta/mini-project-3/config"
	"sekolahbeta/mini-project-3/model"
	"strconv"
	"sync"
)

func BookImport() {
	file, err := os.Open("sample_books.csv")
	if err != nil {
		fmt.Println("Terjadi Error:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Terjadi Error:", err)
		return
	}

	fmt.Println("Menunggu import data buku ...")

	wg := sync.WaitGroup{}

	ch := make(chan []string)
	chBooks := make(chan model.Book, len(lines))

	jmlThread := 5

	for i := 0; i < jmlThread; i++ {
		wg.Add(1)
		go func(ch <-chan []string, chBooks chan model.Book, wg *sync.WaitGroup) {
			for book := range ch {
				chBooks <- model.Book{
					Model: model.Model{
						ID: uint(convertToInt(book[0])),
					},
					ISBN:    book[1],
					Penulis: book[2],
					Tahun:   uint(convertToInt(book[3])),
					Judul:   book[4],
					Gambar:  book[5],
					Stok:    uint(convertToInt(book[6])),
				}
			}
			wg.Done()
		}(ch, chBooks, &wg)
	}

	for i, book := range lines {
		if i == 0 {
			continue
		}
		ch <- book
	}

	close(ch)
	wg.Wait()
	close(chBooks)

	for book := range chBooks {
		_, err := book.GetByID(config.Mysql.DB)
		if err != nil {
			err := book.Create(config.Mysql.DB)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err := book.UpdateByID(config.Mysql.DB)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	fmt.Println("Berhasil import data buku !")
}

func convertToInt(value string) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println(err)
	}
	return result
}
