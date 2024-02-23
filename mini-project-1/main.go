package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func tambahBuku() {
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
	lineInput("Masukan Judul Buku : ", &judulBuku)
	lineInput("Masukan Pengarang Buku : ", &pengarangBuku)
	lineInput("Masukan Penerbit Buku : ", &penerbitBuku)
	lineInput("Masukan Jumlah Halaman Buku : ", &jumlahHalBuku)
	lineInput("Masukan Tahun Terbit Buku : ", &tahunTerbitBuku)

	itemBuku := Buku{
		Kode:        kodeBuku,
		Judul:       judulBuku,
		Pengarang:   pengarangBuku,
		Penerbit:    penerbitBuku,
		JumlahHal:   jumlahHalBuku,
		TahunTerbit: tahunTerbitBuku,
	}

	listBuku = append(listBuku, itemBuku)

	fmt.Println("Berhasil Menambah Buku!")
}

func lihatBuku() {
	line()
	fmt.Println("Lihat Pesanan")
	line()
	fmt.Printf("|Kode\t|Judul\t|Pengarang\t|Penerbit\t|JumlahHal\t|TahunTerbit\t|\n")
	for _, buku := range listBuku {
		fmt.Printf("|%s \t|%s\t|%s\t|%s\t|%d\t|%d\t|\n",
			buku.Kode,
			buku.Judul,
			buku.Pengarang,
			buku.Penerbit,
			buku.JumlahHal,
			buku.TahunTerbit,
		)
	}
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
		ubahBuku()
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
		lineInput("Pengarang Buku Sebelumnya '"+buku.Pengarang+"' : ", &penerbitBuku)
		listBuku[index].Penerbit = penerbitBuku
		countChange += 1
	}
	confirm = lineConfirm("Apa anda ingin merubah Jumlah Halaman Buku ? ")
	if confirm {
		lineInput("Pengarang Buku Sebelumnya '"+strconv.Itoa(buku.JumlahHal)+"' : ", &jumlahHalBuku)
		listBuku[index].JumlahHal = jumlahHalBuku
		countChange += 1
	}
	confirm = lineConfirm("Apa anda ingin merubah Tahun Terbit Buku ? ")
	if confirm {
		lineInput("Pengarang Buku Sebelumnya '"+strconv.Itoa(buku.TahunTerbit)+"' : ", &tahunTerbitBuku)
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

}

func main() {
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
	fmt.Println("3. Edit Buku")
	fmt.Println("4. Hapus Buku")
	fmt.Println("5. Keluar")
	line()
	lineInput("Masukan Pilihan: ", &pilihMenu)

	switch pilihMenu {
	case 1:
		tambahBuku()
	case 2:
		lihatBuku()
	case 3:
		ubahBuku()
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
