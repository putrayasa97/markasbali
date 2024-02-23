package main

import "fmt"

type Hewan struct {
	Nama  string
	Suara string
}

// method dari struct
func (h Hewan) bunyi() {
	fmt.Println(h.Suara)
}

func (h Hewan) nama() {
	fmt.Println(h.Nama)
}

func main() {
	Hewan := Hewan{
		Nama:  "kucing",
		Suara: "Meoong",
	}
	// cara memanggil method
	Hewan.bunyi()
	Hewan.nama()
}
