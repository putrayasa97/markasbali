package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

type Buku struct {
	Kode        string
	Judul       string
	Pengarang   string
	Penerbit    string
	JumlahHal   int
	TahunTerbit int
}

var listBuku []Buku

func (buku Buku) GetByCode(code string) (Buku, int) {
	for index, buku := range listBuku {
		if buku.Kode == code {
			return buku, index
		}
	}
	return Buku{}, 0
}

func main() {
	optionMenu()
}

func tambahBuku() {
	var Buku Buku
	kodeBuku := ""
	judulBuku := ""
	pengarangBuku := ""
	penerbitBuku := ""
	jumlahHalBuku := 0
	tahunTerbitBuku := 0

	line()
	fmt.Println("Tambah Buku")
	line()

	lineInput("Masukan Kode Buku : ", &kodeBuku)
	buku, _ := Buku.GetByCode(kodeBuku)
	if buku.Kode == kodeBuku {
		fmt.Println("Kode Buku tidak boleh sama!")
		optionMenu()
	}
	lineInput("Masukan Judul Buku : ", &judulBuku)
	lineInput("Masukan Pengarang Buku : ", &pengarangBuku)
	lineInput("Masukan Penerbit Buku : ", &penerbitBuku)
	lineInput("Masukan Jumlah Halaman Buku : ", &jumlahHalBuku)
	lineInput("Masukan Tahun Terbit Buku : ", &tahunTerbitBuku)

	Buku.Kode = kodeBuku
	Buku.Judul = judulBuku
	Buku.Pengarang = pengarangBuku
	Buku.Penerbit = penerbitBuku
	Buku.JumlahHal = jumlahHalBuku
	Buku.TahunTerbit = tahunTerbitBuku
	itemBuku := Buku

	listBuku = append(listBuku, itemBuku)

	fmt.Println("Berhasil Menambah Buku!")
}

func lihatBuku() {
	line()
	fmt.Println("Daftar Buku")
	line()

	w := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "|Kode\tJudul\tPengarang\tPenerbit\tJumlah Hal\tTahun Terbit\t")
	for _, buku := range listBuku {
		fmt.Fprintf(w, "|%s\t%s\t%s\t%s\t%d\t%d\t\n",
			buku.Kode,
			buku.Judul,
			buku.Pengarang,
			buku.Penerbit,
			buku.JumlahHal,
			buku.TahunTerbit,
		)
	}
	w.Flush()
}

func ubahBuku() {
	var Buku Buku
	kodeBuku := ""
	judulBuku := ""
	pengarangBuku := ""
	penerbitBuku := ""
	jumlahHalBuku := 0
	tahunTerbitBuku := 0

	line()
	fmt.Println("Ubah Buku")
	line()

	lineInput("Masukan Kode Buku yang ingin diubah : ", &kodeBuku)
	buku, index := Buku.GetByCode(kodeBuku)
	if buku.Kode == "" {
		fmt.Println("Kode Buku tidak ditemukan!")
		optionMenu()
	}
	line()
	countChange := 0
	confirm := lineConfirm("Apa anda ingin merubah Judul Buku ? ")
	if confirm {
		lineInput("Judul Buku Sebelumnya '"+buku.Judul+"' : ", &judulBuku)
		listBuku[index].Judul = judulBuku
		countChange += 1
	}
	confirm = lineConfirm("Apa anda ingin merubah Pengarang Buku ? ")
	if confirm {
		lineInput("Pengarang Buku Sebelumnya '"+buku.Pengarang+"' : ", &pengarangBuku)
		listBuku[index].Pengarang = pengarangBuku
		countChange += 1
	}
	confirm = lineConfirm("Apa anda ingin merubah Penerbit Buku? ")
	if confirm {
		lineInput("Penerbit Buku Sebelumnya '"+buku.Penerbit+"' : ", &penerbitBuku)
		listBuku[index].Penerbit = penerbitBuku
		countChange += 1
	}
	confirm = lineConfirm("Apa anda ingin merubah Jumlah Halaman Buku ? ")
	if confirm {
		lineInput("Jumlah Halaman Buku Sebelumnya '"+strconv.Itoa(buku.JumlahHal)+"' : ", &jumlahHalBuku)
		listBuku[index].JumlahHal = jumlahHalBuku
		countChange += 1
	}
	confirm = lineConfirm("Apa anda ingin merubah Tahun Terbit Buku ? ")
	if confirm {
		lineInput("Tahun Terbit Buku Sebelumnya '"+strconv.Itoa(buku.TahunTerbit)+"' : ", &tahunTerbitBuku)
		listBuku[index].TahunTerbit = tahunTerbitBuku
		countChange += 1
	}

	if countChange == 0 {
		fmt.Println("Tidak ada perubahan data buku!")
		return
	}
	fmt.Println("Berhasil Merubah Buku!")
	optionMenu()
}

func hapusBuku() {
	var Buku Buku
	kodeBuku := ""

	line()
	fmt.Println("Hapus Buku")
	line()

	lineInput("Masukan Kode Buku yang ingin diubah : ", &kodeBuku)
	buku, index := Buku.GetByCode(kodeBuku)
	if buku.Kode == "" {
		fmt.Println("Kode Buku tidak ditemukan!")
		optionMenu()
	}

	listBuku = append(
		listBuku[:index],
		listBuku[index+1:]...,
	)

	fmt.Println("Berhasil Menghapus Buku!")
	optionMenu()
}

func optionMenu() {
	var pilihMenu int
	line()
	fmt.Println("Aplikasi Manajemen Daftar Buku Perpustakan")
	line()
	fmt.Println("Silakan Pilih Menu : ")
	fmt.Println("1. Tambah Buku")
	fmt.Println("2. Lihat Buku")
	fmt.Println("3. Ubah Buku")
	fmt.Println("4. Hapus Buku")
	fmt.Println("5. Keluar")
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
		os.Exit(0)
	}
	main()
}

func line() {
	fmt.Println("+-----------------------------------------+")
}

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
