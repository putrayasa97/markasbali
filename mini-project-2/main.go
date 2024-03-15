package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"text/tabwriter"
	"time"

	"github.com/go-pdf/fpdf"
)

// struct atau data type untuk atribut buku
type Buku struct {
	Kode        string
	Judul       string
	Pengarang   string
	Penerbit    string
	JumlahHal   int
	TahunTerbit int
}

// variable slice dengan type Buku
// untuk menampung list buku yang ditambahkan
var listBuku []Buku

// Method untuk mencari 1 buku berdasarkan kode buku
func (buku Buku) getByCode(code string) (Buku, int) {
	for index, buku := range listBuku {
		if buku.Kode == code {
			return buku, index
		}
	}
	return Buku{}, 0
}

var draftListBuku []Buku

// Method untuk mencari 1 buku berdasarkan kode buku pada list draft
func (buku Buku) getDratByCode(code string) (Buku, int) {
	for index, buku := range draftListBuku {
		if buku.Kode == code {
			return buku, index
		}
	}
	return Buku{}, 0
}

// fungsi Main yang dijalankan pertama kali
func main() {
	optionMenu()
}

// fungsi untuk mendapatkan data buku yg di simpan pada file json
func getJsonBuku() []Buku {
	listBuku := []Buku{}
	listJsonBuku, err := os.ReadDir("books")
	if err != nil {
		fmt.Println("Terjadi Error:", err)
	}
	wg := sync.WaitGroup{}

	ch := make(chan string)
	chBuku := make(chan Buku, len(listJsonBuku))

	jmlThread := 5

	for i := 0; i < jmlThread; i++ {
		wg.Add(1)
		go func(ch <-chan string, chBuku chan Buku, wg *sync.WaitGroup) {
			var buku Buku
			for kodeBuku := range ch {
				dataJson, err := os.ReadFile(fmt.Sprintf("books/%s", kodeBuku))
				if err != nil {
					fmt.Println("Terjadi Error:", err)
				}
				err = json.Unmarshal(dataJson, &buku)
				if err != nil {
					fmt.Println("Terjadi Error:", err)
				}
				chBuku <- buku
			}
			wg.Done()
		}(ch, chBuku, &wg)
	}

	for _, filePesanan := range listJsonBuku {
		ch <- filePesanan.Name()
	}

	close(ch)
	wg.Wait()
	close(chBuku)

	for dataBuku := range chBuku {
		listBuku = append(listBuku, dataBuku)
	}

	sort.Slice(listBuku, func(i, j int) bool {
		return listBuku[i].Kode < listBuku[j].Kode
	})

	return listBuku
}

// fungsi untuk menambah daftar buku
func tambahBuku() {
	var modelBuku Buku
	kodeBuku := ""
	judulBuku := ""
	pengarangBuku := ""
	penerbitBuku := ""
	jumlahHalBuku := 0
	tahunTerbitBuku := 0
	draftListBuku = []Buku{}

	os.Mkdir("books", 0777)
	listBuku = getJsonBuku()

	line()
	fmt.Println("Tambah Buku")
	line()

	for {
		lineInput("Masukan Kode Buku : ", &kodeBuku)
		buku, _ := modelBuku.getByCode(kodeBuku)
		draftBuku, _ := modelBuku.getDratByCode(kodeBuku)

		if (buku.Kode == kodeBuku) || (draftBuku.Kode == kodeBuku) {
			fmt.Println("Kode Buku sudah digunakan !")
		} else {
			lineInput("Masukan Judul Buku : ", &judulBuku)
			lineInput("Masukan Pengarang Buku : ", &pengarangBuku)
			lineInput("Masukan Penerbit Buku : ", &penerbitBuku)
			lineInput("Masukan Jumlah Halaman Buku : ", &jumlahHalBuku)
			lineInput("Masukan Tahun Terbit Buku : ", &tahunTerbitBuku)

			draftListBuku = append(draftListBuku, Buku{
				Kode:        kodeBuku,
				Judul:       judulBuku,
				Pengarang:   pengarangBuku,
				Penerbit:    penerbitBuku,
				JumlahHal:   jumlahHalBuku,
				TahunTerbit: tahunTerbitBuku,
			})
		}

		var pilihanMenuPersanan = 0
		lineInput("Ketikan 1 untuk tambah buku, ketika 0 untuk keluar : ", &pilihanMenuPersanan)
		if pilihanMenuPersanan == 0 {
			break
		}

	}

	prosesSimpanBuku()
	fmt.Println("Berhasil Menambah Buku!")
	optionMenu()
}

// fungsi proses untuk menyimpan buku dalam bentuk json
func prosesSimpanBuku() {
	fmt.Println("Menambah Buku ...")

	chSave := make(chan Buku)

	wgSave := sync.WaitGroup{}

	jmlThread := 5

	for i := 0; i < jmlThread; i++ {
		wgSave.Add(1)
		go func(ch <-chan Buku, wg *sync.WaitGroup, noThread int) {
			for buku := range ch {
				dataJson, err := json.Marshal(buku)

				if err != nil {
					fmt.Println("Terjadi Error:", err)
				}

				err = os.WriteFile(fmt.Sprintf("books/book-%s.json", buku.Kode), dataJson, 0644)
				if err != nil {
					fmt.Println("Terjadi Error:", err)
				}
			}
			wgSave.Done()
		}(chSave, &wgSave, i)
	}

	for _, buku := range draftListBuku {
		chSave <- buku
	}

	close(chSave)
	wgSave.Wait()
}

// fungsi untuk melihat daftar buku
func lihatBuku() {
	listBuku = getJsonBuku()
	line()
	fmt.Println("Daftar Buku")
	line()

	w := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "|Kode\tJudul\tPengarang\tPenerbit\tJumlah Hal\tTahun Terbit\t")

	wg := sync.WaitGroup{}
	ch := make(chan Buku)

	jmlThread := 5

	for i := 0; i < jmlThread; i++ {
		wg.Add(1)
		go func(ch <-chan Buku, w *tabwriter.Writer, wg *sync.WaitGroup) {
			for buku := range ch {
				fmt.Fprintf(w, "|%s\t%s\t%s\t%s\t%d\t%d\t\n",
					buku.Kode,
					buku.Judul,
					buku.Pengarang,
					buku.Penerbit,
					buku.JumlahHal,
					buku.TahunTerbit,
				)
			}
			wg.Done()
		}(ch, w, &wg)
	}

	for _, buku := range listBuku {
		ch <- buku
	}

	close(ch)
	wg.Wait()

	w.Flush()
}

// fungsi untuk mengubah data buku
// berdasarkan kode buku
func ubahBuku() {
	var modalBuku Buku
	kodeBuku := ""
	judulBuku := ""
	pengarangBuku := ""
	penerbitBuku := ""
	jumlahHalBuku := 0
	tahunTerbitBuku := 0
	listBuku = []Buku{}

	listBuku = getJsonBuku()
	line()
	fmt.Println("Ubah Buku")
	line()

	lineInput("Masukan Kode Buku yang ingin diubah : ", &kodeBuku)
	buku, _ := modalBuku.getByCode(kodeBuku)
	if buku.Kode == "" {
		fmt.Println("Kode Buku tidak ditemukan!")
		optionMenu()
	}
	modalBuku = buku
	line()
	countChange := 0
	confirm := lineConfirm("Apa anda ingin merubah Judul Buku ?")
	if confirm {
		lineInput("Judul Buku Sebelumnya '"+buku.Judul+"' : ", &judulBuku)
		modalBuku.Judul = judulBuku
		countChange += 1
	}
	confirm = lineConfirm("Apa anda ingin merubah Pengarang Buku ?")
	if confirm {
		lineInput("Pengarang Buku Sebelumnya '"+buku.Pengarang+"' : ", &pengarangBuku)
		modalBuku.Pengarang = pengarangBuku
		countChange += 1
	}
	confirm = lineConfirm("Apa anda ingin merubah Penerbit Buku?")
	if confirm {
		lineInput("Penerbit Buku Sebelumnya '"+buku.Penerbit+"' : ", &penerbitBuku)
		modalBuku.Penerbit = penerbitBuku
		countChange += 1
	}
	confirm = lineConfirm("Apa anda ingin merubah Jumlah Halaman Buku ?")
	if confirm {
		lineInput("Jumlah Halaman Buku Sebelumnya '"+strconv.Itoa(buku.JumlahHal)+"' : ", &jumlahHalBuku)
		modalBuku.JumlahHal = jumlahHalBuku
		countChange += 1
	}
	confirm = lineConfirm("Apa anda ingin merubah Tahun Terbit Buku ?")
	if confirm {
		lineInput("Tahun Terbit Buku Sebelumnya '"+strconv.Itoa(buku.TahunTerbit)+"' : ", &tahunTerbitBuku)
		modalBuku.TahunTerbit = tahunTerbitBuku
		countChange += 1
	}

	if countChange == 0 {
		fmt.Println("Tidak ada perubahan data buku!")
		return
	}

	dataJson, err := json.Marshal(modalBuku)
	if err != nil {
		fmt.Println("Terjadi Error:", err)
	}

	err = os.WriteFile(fmt.Sprintf("books/book-%s.json", buku.Kode), dataJson, 0644)
	if err != nil {
		fmt.Println("Terjadi Error:", err)
	}

	fmt.Println("Berhasil Merubah Buku!")
	optionMenu()
}

// fungsi untuk menghapus data buku
// berdasarkan kode buku
func hapusBuku() {
	var modelBuku Buku
	kodeBuku := ""
	listBuku = []Buku{}
	listBuku = getJsonBuku()
	line()
	fmt.Println("Hapus Buku")
	line()

	lineInput("Masukan Kode Buku yang ingin diubah : ", &kodeBuku)
	buku, index := modelBuku.getByCode(kodeBuku)
	if buku.Kode == "" {
		fmt.Println("Kode Buku tidak ditemukan!")
		optionMenu()
	}

	err := os.Remove(fmt.Sprintf("books/book-%s.json", listBuku[index].Kode))
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	fmt.Println("Berhasil Menghapus Buku!")
	optionMenu()
}

// fungsi untuk cetak data buku
func printBuku() {
	line()
	fmt.Println("Print Buku")
	line()
	os.Mkdir("pdf", 0777)

	optionPrint()

	fmt.Println("Berhasil print buku")
	optionMenu()
}

// fungsi untuk opsi print
func optionPrint() {
	var opsiPrint int
	fmt.Println("Silakan pilih opsi print : ")
	fmt.Println("1. Print Berdasarkan Kode Buku")
	fmt.Println("2. Print Semua Buku")
	fmt.Println("3. Kembali ke menu utama")
	line()
	lineInput("Masukan Pilihan : ", &opsiPrint)

	switch opsiPrint {
	case 1:
		printByCodeBook()
	case 2:
		printAllBook()
	case 3:
		optionMenu()
	case 4:
		os.Exit(0)
	}
}

// fungsi untuk print berdasarkan kode
func printByCodeBook() {
	var modelBuku Buku
	kodeBuku := ""
	listBuku = []Buku{}
	listBuku = getJsonBuku()

	lineInput("Masukan Kode Buku yang ingin diprint : ", &kodeBuku)
	buku, _ := modelBuku.getByCode(kodeBuku)
	if buku.Kode == "" {
		fmt.Println("Kode Buku tidak ditemukan!")
		optionPrint()
	}
	modalBuku := buku

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "", 12)
	pdf.SetLeftMargin(10)
	pdf.SetRightMargin(10)

	bukuText := fmt.Sprintf(
		"Kode Buku : %s\nJudul : %s\nPengarang : %s\nPenerbit : %s\nJumlah Halaman : %d\nTahun Terbit : %d \n",
		modalBuku.Kode, modalBuku.Judul,
		modalBuku.Pengarang, modalBuku.Penerbit,
		modalBuku.JumlahHal, modalBuku.TahunTerbit)

	pdf.MultiCell(0, 10, bukuText, "0", "L", false)
	pdf.Ln(5)

	err := pdf.OutputFileAndClose(
		fmt.Sprintf("pdf/buku_%s_%s.pdf",
			modalBuku.Kode, time.Now().Format("2006-01-02-15-04-05")))

	if err != nil {
		fmt.Println("Terjadi error:", err)
	}
}

// fungsi untuk print semua buku
func printAllBook() {
	listBuku = []Buku{}
	listBuku = getJsonBuku()

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "", 12)
	pdf.SetLeftMargin(10)
	pdf.SetRightMargin(10)

	fmt.Println("Menambah Print Buku ...")

	ch := make(chan Buku)

	chpdf := make(chan string, len(listBuku))

	wg := sync.WaitGroup{}

	jmlThread := 5

	for i := 0; i < jmlThread; i++ {
		wg.Add(1)
		go func(ch <-chan Buku, chpdf chan string, wg *sync.WaitGroup) {
			for buku := range ch {
				chpdf <- fmt.Sprintf(
					"Kode Buku : %s\nJudul : %s\nPengarang : %s\nPenerbit : %s\nJumlah Halaman : %d\nTahun Terbit : %d \n",
					buku.Kode, buku.Judul,
					buku.Pengarang, buku.Penerbit,
					buku.JumlahHal, buku.TahunTerbit)
			}
			wg.Done()
		}(ch, chpdf, &wg)
	}

	for _, buku := range listBuku {
		ch <- buku
	}

	close(ch)
	wg.Wait()
	close(chpdf)

	for text := range chpdf {
		pdf.MultiCell(0, 10, text, "0", "L", false)
		pdf.Ln(5)
	}

	err := pdf.OutputFileAndClose(
		fmt.Sprintf("pdf/daftar_buku_%s.pdf",
			time.Now().Format("2006-01-02-15-04-05")))

	if err != nil {
		fmt.Println("Terjadi error:", err)
	}
}

// fungsi untuk Menu utama dari
// Aplikasi Manajemen Daftar Buku Perpustakaan
func optionMenu() {
	var pilihMenu int
	line()
	fmt.Println("Aplikasi Manajemen Daftar Buku Perpustakaan")
	line()
	fmt.Println("Silakan Pilih Menu : ")
	fmt.Println("1. Tambah Buku")
	fmt.Println("2. Lihat Daftar Buku")
	fmt.Println("3. Ubah Buku")
	fmt.Println("4. Hapus Buku")
	fmt.Println("5. Print Buku")
	fmt.Println("6. Keluar")
	line()
	lineInput("Masukan Pilihan : ", &pilihMenu)

	switch pilihMenu {
	case 1:
		tambahBuku()
	case 2:
		lihatBuku()
	case 3:
		ubahBuku()
	case 4:
		hapusBuku()
	case 5:
		printBuku()
	case 6:
		os.Exit(0)
	}
	main()
}

// fungsi untuk membuat baris
func line() {
	fmt.Println("+-----------------------------------------+")
}

// fungsi untuk inputan baris
// berdasarkan type data
func lineInput(title string, variable interface{}) {
	var err error
	fmt.Print(title)
	switch v := variable.(type) {
	case *int:
		var input string
		_, err = fmt.Scanln(&input)
		if err == nil {
			*v, err = strconv.Atoi(input)
		}
	case *string:
		var readString string
		input := bufio.NewReader(os.Stdin)
		readString, err = input.ReadString('\n')
		*v = strings.Replace(readString, "\n", "", 1)
	default:
		fmt.Println("Tipe data tidak didukung")
		return
	}

	if err != nil {
		fmt.Println("Terjadi Error:", err)
		return
	}
}

// fungsi untuk konfirmasi inputan data
func lineConfirm(s string) bool {
	r := bufio.NewReader(os.Stdin)

	fmt.Printf("%s [y/n]: ", s)

	res, err := r.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi Error:", err)
		return false
	}

	return strings.ToLower(strings.TrimSpace(res))[0] == 'y'
}
