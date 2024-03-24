package book

import (
	"fmt"
	"sekolahbeta/mini-project-3/config"
	"sekolahbeta/mini-project-3/model"
	"sekolahbeta/mini-project-3/utils/helpers"
	"sync"
	"time"

	"github.com/go-pdf/fpdf"
)

func BookPrintByID() {
	var modelBook model.Book
	bookID := 0

	helpers.LineInput("Masukan ID Buku yang ingin diprint : ", &bookID)
	modelBook.Model = model.Model{
		ID: uint(bookID),
	}
	book, err := modelBook.GetByID(config.Mysql.DB)
	if err != nil {
		fmt.Println("Kode Buku tidak ditemukan!")
	}

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "", 12)
	pdf.SetLeftMargin(10)
	pdf.SetRightMargin(10)

	bookText := fmt.Sprintf(
		"ISBN : %s\nJudul : %s\nPenulis : %s\nTahun : %d\nStok : %d\nGambar : %s \n",
		book.ISBN, book.Judul,
		book.Penulis, book.Tahun,
		book.Stok, book.Gambar)

	pdf.MultiCell(0, 10, bookText, "0", "L", false)
	pdf.Ln(5)

	err = pdf.OutputFileAndClose(
		fmt.Sprintf("pdf/buku_%s_%s.pdf",
			book.ISBN, time.Now().Format("2006-01-02-15-04-05")))

	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	fmt.Println("Berhasil Print Buku!")
}

func BookPrintAll() {
	var modelBook model.Book
	books, err := modelBook.GetAll(config.Mysql.DB)
	if err != nil {
		fmt.Println(err)
	}

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "", 12)
	pdf.SetLeftMargin(10)
	pdf.SetRightMargin(10)

	fmt.Println("Menunggu Print Buku ...")

	ch := make(chan model.Book)

	chpdf := make(chan string, len(books))

	wg := sync.WaitGroup{}

	jmlThread := 5
	for i := 0; i < jmlThread; i++ {
		wg.Add(1)
		go func(ch <-chan model.Book, chpdf chan string, wg *sync.WaitGroup) {
			for book := range ch {
				chpdf <- fmt.Sprintf(
					"ISBN : %s\nJudul : %s\nPenulis : %s\nTahun : %d\nStok : %d\nGambar : %s \n",
					book.ISBN, book.Judul,
					book.Penulis, book.Tahun,
					book.Stok, book.Gambar)
			}
			wg.Done()
		}(ch, chpdf, &wg)
	}

	for _, buku := range books {
		ch <- buku
	}

	close(ch)
	wg.Wait()
	close(chpdf)

	for text := range chpdf {
		pdf.MultiCell(0, 10, text, "0", "L", false)
		pdf.Ln(5)
	}

	err = pdf.OutputFileAndClose(
		fmt.Sprintf("pdf/daftar_buku_%s.pdf",
			time.Now().Format("2006-01-02-15-04-05")))

	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	fmt.Println("Berhasil Print Buku!")
}
