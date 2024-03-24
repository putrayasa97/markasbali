package book

import (
	"fmt"
	"os"
	"sekolahbeta/mini-project-3/config"
	"sekolahbeta/mini-project-3/model"
	"sekolahbeta/mini-project-3/utils/helpers"
	"sync"
	"text/tabwriter"
)

func BookGet() {
	helpers.Line()
	fmt.Println("Daftar Buku")
	helpers.Line()
	modelBook := model.Book{}

	w := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "|ID\tISBN\tJudul\tPenulis\tTahun\tStok\t")

	books, err := modelBook.GetAll(config.Mysql.DB)
	if err != nil {
		fmt.Println(err)
	}

	wg := sync.WaitGroup{}
	chGetData := make(chan model.Book)
	chShow := make(chan model.Book, len(books))
	jmlThread := 1

	for i := 0; i < jmlThread; i++ {
		wg.Add(1)
		go proccesGet(chGetData, chShow, w, &wg)
	}

	for _, book := range books {
		chGetData <- book
	}

	close(chGetData)
	wg.Wait()
	close(chShow)

	for book := range chShow {
		fmt.Fprintf(w, "|%d\t%s\t%s\t\n",
			book.ID,
			book.ISBN,
			book.Judul,
		)
	}

	w.Flush()
}

func proccesGet(chGetData <-chan model.Book, chShow chan model.Book, w *tabwriter.Writer, wg *sync.WaitGroup) {
	for book := range chGetData {
		chShow <- book
	}
	wg.Done()
}
