package main

import (
	"fmt"
	"os"
	"os/exec"
	// "strconv"
)

const (
	NMAX      int = 100
	NMAXMOBIL int = 1000
)

type Mobil struct {
	nama         string
	tahunKeluar  int
	harga        int
	penjualan    int
	jumlah       int
}

type Pabrik struct {
	merek       string
	daftarMobil [10]Mobil
	nMobil      int
}

type AllMobil struct {
	mobil  Mobil
	pabrik string
}

type arrPabrik [NMAX]Pabrik
type arrMobil [NMAXMOBIL]AllMobil

var pesan string = ""

var allMobil arrMobil
var nMobil int

func main() {
	var data arrPabrik
	var pil, n int

	menu(&pil)
	for pil != 4 {

		switch pil {
		case 1:
			tampilPabrik(data, n)
			menuPabrik(&pil)

			for pil != 4 {
				switch pil {
				case 1:
					tambahPabrik(&data, &n)
				case 2:
					editPabrik(&data, n)
				case 3:
					hapusPabrik(&data, &n)
				}

				tampilPabrik(data, n)
				menuPabrik(&pil)
			}

		case 2:
			tampilMobil(data, n)
			menuMobil(&pil)

			for pil != 3 {
				switch pil {
				case 1:
					tambahMobil(&data, &n)
				case 2:
					editMobil(&data, allMobil, n)
				case 3:
					// hapusMobil(&data, &n)
				}

				tampilMobil(data, n)
				menuMobil(&pil)
			}
		}

		menu(&pil)
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func header() {
	fmt.Println("=================================")
	fmt.Println("   Aplikasi Dealer Mobil Josss   ")
	fmt.Println("=================================")
}

func menu(pil *int) {
	clear()
	header()
	fmt.Println("1. Kelola Data Pabrikan")
	fmt.Println("2. Kelola Data Mobil")
	fmt.Println("3. Lihat Data Penjualan")
	fmt.Println("4. Exit")
	fmt.Println("=================================")

	fmt.Print("Masukkan Pilihan [1-4] : ")
	fmt.Scan(pil)

	for *pil > 4 || *pil == 0 {
		fmt.Println("Pilihan Tidak Tersedia!!!")
		fmt.Print("Masukkan Pilihan [1-4] : ")
		fmt.Scan(pil)
	}
}

func menuPabrik(pil *int) {
	fmt.Println("=================================")
	fmt.Println("1. Tambah Data (V)")
	fmt.Println("2. Edit Data")
	fmt.Println("3. Hapus Data")
	fmt.Println("4. Kembali")
	fmt.Println("=================================")

	fmt.Print("Masukkan Pilihan [1-4] : ")
	fmt.Scan(pil)

	for *pil > 4 {
		fmt.Println("Pilihan Tidak Tersedia!!!")
		fmt.Print("Masukkan Pilihan [1-4] : ")
		fmt.Scan(pil)
	}
}

func tampilPabrik(data arrPabrik, n int) {
	clear()
	header()
	fmt.Println("=================================")
	fmt.Println("          Data Pabrikan          ")
	fmt.Println("=================================")
	fmt.Println("No.\tMerek\t\tJumlah Mobil")
	fmt.Println("=================================")

	for i := 0; i < n; i++ {
		fmt.Printf("%d.\t%s\t\t%d\n", i+1, data[i].merek, data[i].nMobil)
	}

	fmt.Println("=================================")
}

func tambahPabrik(data *arrPabrik, n *int) {
	var merk string

	clear()
	header()
	fmt.Println("=================================")
	fmt.Println("         Tambah Data Pabrikan     ")
	fmt.Println("=================================")

	fmt.Print("Merek Pabrik : ")
	fmt.Scan(&merk)

	data[*n].merek = merk
	data[*n].nMobil = 0
	*n++

	pesan = "Data berhasil ditambahkan"
}

func editPabrik(data *arrPabrik, n int) {
	var idx, nMobil int

	clear()
	header()
	fmt.Println("=================================")
	fmt.Println("          Edit Data Pabrikan      ")
	fmt.Println("=================================")

	tampilPabrik(*data, n)

	fmt.Print("Masukkan Nomor Pabrik : ")
	fmt.Scan(&idx)

	for idx <= 0 || idx > n {
		fmt.Println("Nomor Pabrik Tidak Tersedia!!!")
		fmt.Print("Masukkan Nomor Pabrik : ")
		fmt.Scan(&idx)
	}

	fmt.Print("Jumlah Mobil : ")
	fmt.Scan(&nMobil)

	data[idx-1].nMobil = nMobil

	pesan = "Data berhasil diubah"
}

func hapusPabrik(data *arrPabrik, n *int) {
	var idx int

	clear()
	header()
	fmt.Println("=================================")
	fmt.Println("          Hapus Data Pabrikan     ")
	fmt.Println("=================================")

	tampilPabrik(*data, *n)

	fmt.Print("Masukkan Nomor Pabrik : ")
	fmt.Scan(&idx)

	for idx <= 0 || idx > *n {
		fmt.Println("Nomor Pabrik Tidak Tersedia!!!")
		fmt.Print("Masukkan Nomor Pabrik : ")
		fmt.Scan(&idx)
	}

	idx--
	for i := idx; i < *n-1; i++ {
		data[i] = data[i+1]
	}

	*n--

	pesan = "Data berhasil dihapus"
}

func menuMobil(pil *int) {
	fmt.Println("=================================")
	fmt.Println("       Kelola Data Mobil          ")
	fmt.Println("=================================")
	fmt.Println("1. Tambah Data Mobil")
	fmt.Println("2. Edit Data Mobil")
	fmt.Println("3. Kembali")
	fmt.Println("=================================")

	fmt.Print("Masukkan Pilihan [1-3] : ")
	fmt.Scan(pil)

	for *pil > 3 {
		fmt.Println("Pilihan Tidak Tersedia!!!")
		fmt.Print("Masukkan Pilihan [1-3] : ")
		fmt.Scan(pil)
	}
}

func tampilMobil(data arrPabrik, n int) {
	clear()
	header()
	fmt.Println("=================================")
	fmt.Println("        Data Mobil per Pabrik     ")
	fmt.Println("=================================")
	fmt.Println("No.\tMerek\t\tModel\t\tTahun\tHarga\tJumlah\tTerjual")
	fmt.Println("=================================")

	for i := 0; i < n; i++ {
		for j := 0; j < data[i].nMobil; j++ {
			mobil := data[i].daftarMobil[j]
			fmt.Printf("%d.\t%s\t\t%s\t\t%d\t%d\t%d\t%d\n", i+1, data[i].merek, mobil.nama, mobil.tahunKeluar, mobil.harga, mobil.jumlah, mobil.penjualan)
		}
	}

	fmt.Println("=================================")
}

func tambahMobil(data *arrPabrik, n *int) {
	var idx, harga, jumlah int
	var nama string
	var tahunKeluar int

	clear()
	header()
	fmt.Println("=================================")
	fmt.Println("         Tambah Data Mobil        ")
	fmt.Println("=================================")

	tampilPabrik(*data, *n)

	fmt.Print("Masukkan Nomor Pabrik : ")
	fmt.Scan(&idx)

	for idx <= 0 || idx > *n {
		fmt.Println("Nomor Pabrik Tidak Tersedia!!!")
		fmt.Print("Masukkan Nomor Pabrik : ")
		fmt.Scan(&idx)
	}

	fmt.Print("Nama Mobil : ")
	fmt.Scan(&nama)

	fmt.Print("Tahun Keluar : ")
	fmt.Scan(&tahunKeluar)

	fmt.Print("Harga Mobil : ")
	fmt.Scan(&harga)

	fmt.Print("Jumlah Mobil : ")
	fmt.Scan(&jumlah)

	mobil := Mobil{
		nama:        nama,
		tahunKeluar: tahunKeluar,
		harga:       harga,
		jumlah:      jumlah,
		penjualan:   0,
	}

	data[idx-1].daftarMobil[data[idx-1].nMobil] = mobil
	data[idx-1].nMobil++

	pesan = "Data berhasil ditambahkan"
}

func editMobil(data *arrPabrik, allMobil arrMobil, n int) {
	var idx, harga, jumlah int
	var nama string
	var tahunKeluar int

	clear()
	header()
	fmt.Println("=================================")
	fmt.Println("         Edit Data Mobil        ")
	fmt.Println("=================================")

	tampilMobil(*data, n)

	fmt.Print("Masukkan Nomor Mobil : ")
	fmt.Scan(&idx)

	for idx <= 0 || idx > n {
		fmt.Println("Nomor Mobil Tidak Tersedia!!!")
		fmt.Print("Masukkan Nomor Mobil : ")
		fmt.Scan(&idx)
	}

	fmt.Print("Nama Mobil : ")
	fmt.Scan(&nama)

	fmt.Print("Tahun Keluar : ")
	fmt.Scan(&tahunKeluar)

	fmt.Print("Harga Mobil : ")
	fmt.Scan(&harga)

	fmt.Print("Jumlah Mobil : ")
	fmt.Scan(&jumlah)

	allMobil[idx-1].mobil.nama = nama
	allMobil[idx-1].mobil.tahunKeluar = tahunKeluar
	allMobil[idx-1].mobil.harga = harga
	allMobil[idx-1].mobil.jumlah = jumlah

	pesan = "Data berhasil diubah"
}
