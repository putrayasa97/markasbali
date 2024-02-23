package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Sistem Manajemen Pesanan Restoran
type Pesanan struct {
	Menu   string
	Meja   int
	Jumlah int
}

var ListPesanan []Pesanan

func TambahPesanan() {
	inputanUser := bufio.NewReader(os.Stdin)
	menuPelanggan := ""
	mejaPelanggan := 0
	jumlahPesananPelanggan := 0

	fmt.Println("=================================")
	fmt.Println("Tambah Pesanan")
	fmt.Println("=================================")
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

	itemPesan := Pesanan{
		Menu:   menuPelanggan,
		Jumlah: jumlahPesananPelanggan,
		Meja:   mejaPelanggan,
	}

	ListPesanan = append(ListPesanan, itemPesan)

	fmt.Println("Berhasil Menambah Pesanan!")
}

func LiatPesanan() {
	fmt.Println("=================================")
	fmt.Println("Lihat Pesanan")
	fmt.Println("=================================")
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
	ListPesanan = append(
		ListPesanan[:urutanPesanan-1],
		ListPesanan[urutanPesanan:]...,
	)
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
	fmt.Println("4. Keluar")
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
		os.Exit(0)
	}
	main()
}

// go mod tidy
