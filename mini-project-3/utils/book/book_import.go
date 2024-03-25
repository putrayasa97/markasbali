package book

import (
	"encoding/csv"
	"fmt"
	"os"
	"sekolahbeta/mini-project-3/config"
	"sekolahbeta/mini-project-3/model"
	"sekolahbeta/mini-project-3/utils/helpers"
	"strconv"
	"sync"
)

func BookImport() {
	pathFile := ""
	helpers.LineInput("Masukan path file yang ingin diimport : ", &pathFile)

	file, err := os.Open("./imports/sample_books.csv")
	if err != nil {
		fmt.Println("Terjadi Error:", err)
		BookImport()
	}
	defer file.Close()

	fmt.Println("Menunggu import data buku ...")

	csvChan, err := loadFile(file)
	if err != nil {
		fmt.Println("Terjadi Error:", err)
	}

	jmlGoroutine := 10

	var bookChanTemp []<-chan model.Book

	for i := 0; i < jmlGoroutine; i++ {
		bookChanTemp = append(bookChanTemp, convertStruct(csvChan))
	}

	mergedCh := appendBooks(bookChanTemp...)

	for chBook := range mergedCh {
		_, err := chBook.GetByID(config.Mysql.DB)
		if err != nil {
			err := chBook.Create(config.Mysql.DB)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err := chBook.UpdateByID(config.Mysql.DB)
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

func loadFile(file *os.File) (<-chan []string, error) {
	bookChan := make(chan []string)
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return bookChan, err
	}

	go func() {
		for i, book := range records {
			if i == 0 {
				continue
			}
			bookChan <- book
		}
		close(bookChan)
	}()

	return bookChan, nil

}

func convertStruct(csvChan <-chan []string) <-chan model.Book {
	booksChan := make(chan model.Book)

	go func() {
		for book := range csvChan {
			booksChan <- model.Book{
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

		close(booksChan)
	}()

	return booksChan
}

func appendBooks(bookChanMany ...<-chan model.Book) <-chan model.Book {
	wg := sync.WaitGroup{}

	mergedChan := make(chan model.Book)

	wg.Add(len(bookChanMany))
	for _, ch := range bookChanMany {
		go func(ch <-chan model.Book) {
			for books := range ch {
				mergedChan <- books
			}
			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(mergedChan)
	}()

	return mergedChan
}
