package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/go-pdf/fpdf"
)

// Sistem Manajemen Pesanan Restoran
type Pesanan struct {
	ID      string
	Menu    string
	Meja    int
	Jumlah  int
	Tanggal time.Time
}

var ListPesanan []Pesanan

// TODO :
// Pesan dapat dimasukan secara draft dan banyak
func TambahPesanan() {
	inputanUser := bufio.NewReader(os.Stdin)
	mejaPelanggan := 0
	jumlahPesananPelanggan := 0

	fmt.Println("=================================")
	fmt.Println("Tambah Pesanan")
	fmt.Println("=================================")

	draftPesanan := []Pesanan{}

	for {
		fmt.Print("Silakan Masukan Menu : ")
		menuPelanggan, err := inputanUser.ReadString('\n')
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}
		menuPelanggan = strings.Replace(menuPelanggan, "\n", "", 1)
		// _, err := fmt.Scanln(&menuPelanggan)
		// if err != nil {
		// 	fmt.Println("Terjadi Error:", err)
		// 	return
		// }

		fmt.Print("Silakan Masukan Meja : ")
		_, err = fmt.Scanln(&mejaPelanggan)
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}

		fmt.Print("Silakan Masukan Jumlah : ")
		_, err = fmt.Scanln(&jumlahPesananPelanggan)
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}

		draftPesanan = append(draftPesanan, Pesanan{
			ID:      fmt.Sprintf("PSN-%d", time.Now().Unix()),
			Menu:    menuPelanggan,
			Jumlah:  jumlahPesananPelanggan,
			Meja:    mejaPelanggan,
			Tanggal: time.Now(),
		})

		var pilihanMenuPersanan = 0
		fmt.Print("Ketikan 1 untuk tambah pesanan, ketika 0 untuk keluar ")
		_, err = fmt.Scanln(&pilihanMenuPersanan)
		if err != nil {
			fmt.Println("Terjadi Error:", err)
			return
		}
		if pilihanMenuPersanan == 0 {
			break
		}
	}

	fmt.Println("Menambah Pesanan ...")
	os.Mkdir("pesanan", 0777)

	ch := make(chan Pesanan)

	wg := sync.WaitGroup{}

	jumlahPelayan := 5

	// Mendaftarkan receiver/pemroses data
	for i := 0; i < jumlahPelayan; i++ {
		wg.Add(1)
		go simpanPesanan(ch, &wg, i)
	}

	// Mengirimkan data ke channel
	for _, pesanan := range draftPesanan {
		ch <- pesanan
	}

	close(ch)
	wg.Wait()
	fmt.Println("Berhasil Menambah Pesanan!")
}

func simpanPesanan(ch <-chan Pesanan, wg *sync.WaitGroup, noPelayan int) {
	for pesanan := range ch {
		dataJson, err := json.Marshal(pesanan)

		if err != nil {
			fmt.Println("Terjadi Error:", err)
		}

		err = os.WriteFile(fmt.Sprintf("pesanan/%s.json", pesanan.ID), dataJson, 0644)
		if err != nil {
			fmt.Println("Terjadi Error:", err)
		}

		fmt.Printf("Pelayan No %d Memperoses ID : %s!]\n", noPelayan, pesanan.ID)
	}
	wg.Done()
}

func lihatPesanan(ch <-chan string, chPesanan chan Pesanan, wg *sync.WaitGroup) {
	var pesanan Pesanan
	for idPesanan := range ch {
		dataJson, err := os.ReadFile(fmt.Sprintf("pesanan/%s", idPesanan))
		if err != nil {
			fmt.Println("Terjadi Error:", err)
		}

		// untuk mengurai data JSON yang diberikan dataJson ke dalam objek Pesanan
		err = json.Unmarshal(dataJson, &pesanan)
		if err != nil {
			fmt.Println("Terjadi Error:", err)
		}

		// mengirimkan objek Pesanan ke channel
		chPesanan <- pesanan
	}
	wg.Done()
}

func LiatPesanan() {
	fmt.Println("=================================")
	fmt.Println("Lihat Pesanan")
	fmt.Println("=================================")
	fmt.Println("Memuat Data...")

	ListPesanan = []Pesanan{}

	listJsonPesanan, err := os.ReadDir("pesanan")
	if err != nil {
		fmt.Println("Terjadi Error:", err)
	}
	wg := sync.WaitGroup{}

	// channel untuk nama file dari id pesanan
	ch := make(chan string)

	// channel untuk data listJsonPesanan yg sudah diurai ke dalam objek Pesanan
	chPesanan := make(chan Pesanan, len(listJsonPesanan))

	jumlahPelayan := 5

	for i := 0; i < jumlahPelayan; i++ {
		wg.Add(1)
		go lihatPesanan(ch, chPesanan, &wg)
	}

	for _, filePesanan := range listJsonPesanan {
		// mengirimkan nama file ke channel
		ch <- filePesanan.Name()
	}

	close(ch)
	wg.Wait()
	close(chPesanan)

	for dataPesanan := range chPesanan {
		ListPesanan = append(ListPesanan, dataPesanan)
	}

	for urutan, pesanan := range ListPesanan {
		fmt.Printf("%d. Nama Menu : %s, Meja: %d\n",
			urutan+1,
			pesanan.Menu,
			pesanan.Meja)
	}
}

func HapusPesanan() {
	var urutanPesanan int
	fmt.Println("=================================")
	fmt.Println("Hapus Pesanan")
	fmt.Println("=================================")
	LiatPesanan()
	fmt.Println("Masukan Urutan Pesanan : ")
	_, err := fmt.Scanln(&urutanPesanan)
	if err != nil {
		fmt.Println("Terjadi Error:", err)
		return
	}

	if (urutanPesanan-1) < 0 || (urutanPesanan-1) > len(ListPesanan) {
		fmt.Println("Urutan Pesanan Tidak Sesuai")
		HapusPesanan()
		return
	}

	//  [:urutanPesanan-1] batas awal
	//  [urutanPesanan:] batas akhir
	// ListPesanan = append(
	// 	ListPesanan[:urutanPesanan-1],
	// 	ListPesanan[urutanPesanan:]...,
	// )

	err = os.Remove(fmt.Sprintf("pesanan/%s.json", ListPesanan[urutanPesanan-1].ID))
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	fmt.Println("Pesanan Berhasil Dihapus!")
}

func main() {
	var pilihanMenu int
	fmt.Println("=================================")
	fmt.Println("Sistem Manajemen Pesanan Restoran")
	fmt.Println("=================================")
	fmt.Println("Silakan Pilih : ")
	fmt.Println("1. Tambah Pesanan")
	fmt.Println("2. Liat Pesanan")
	fmt.Println("3. Hapus Pesanan")
	fmt.Println("4. Generate Daftar Pesanan")
	fmt.Println("5. Keluar")
	fmt.Println("=================================")
	fmt.Print("Masukan Pilihan : ")

	_, err := fmt.Scanln(&pilihanMenu)
	if err != nil {
		fmt.Println("Terjadi Error:", err)
		return
	}

	switch pilihanMenu {
	case 1:
		TambahPesanan()
	case 2:
		LiatPesanan()
	case 3:
		HapusPesanan()
	case 4:
		GeneratePdfPesanan()
	case 5:
		os.Exit(0)
	}
	main()
}

// go mod tidy
func GeneratePdfPesanan() {
	LiatPesanan()
	fmt.Println("=================================")
	fmt.Println("Membuat Daftar Pesanan ...")
	fmt.Println("=================================")
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "", 12)
	pdf.SetLeftMargin(10)
	pdf.SetRightMargin(10)

	for i, pesanan := range ListPesanan {
		pesananText := fmt.Sprintf(
			"Pesanan #%d:\nID : %s\nMenu : %s\nMeja : %d\nJumlah : %d\nTanggal : %s\n",
			i+1, pesanan.ID, pesanan.Menu,
			pesanan.Meja, pesanan.Jumlah,
			pesanan.Tanggal.Format("2006-01-02 15:04:05"))

		pdf.MultiCell(0, 10, pesananText, "0", "L", false)
		pdf.Ln(5)
	}

	err := pdf.OutputFileAndClose(
		fmt.Sprintf("daftar_pesanan_%s.pdf",
			time.Now().Format("2006-01-02-15-04-05")))

	if err != nil {
		fmt.Println("Terjadi error:", err)
	}
}
