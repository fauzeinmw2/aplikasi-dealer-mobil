package main

import (
	"fmt"
	"strings"
)

type Mobil struct {
	Nama   string
	Tahun  int
	Harga  int
	Penjualan int
}

type Pabrik struct {
	Nama        string
	listMobil   [100]Mobil
	JumlahMobil int
}

var listPabrik [100]Pabrik
var jumlahPabrik int

func main() {
	for {
		fmt.Println("=== Dealer Mobil ===")
		fmt.Println("1. Tambah Pabrik")
		fmt.Println("2. Tambah Mobil")
		fmt.Println("3. Ubah Mobil")
		fmt.Println("4. Hapus Mobil")
		fmt.Println("5. Cari Mobil")
		fmt.Println("6. Cari Pabrik")
		fmt.Println("7. Tampilkan Daftar Mobil")
		fmt.Println("8. Tampilkan Daftar Pabrik")
		fmt.Println("9. Tampilkan 3 Daftar Mobil dan Pabrik dengan Jumlah Penjualan Tertinggi")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		var menu int
		fmt.Scanln(&menu)

		switch menu {
		case 1:
			tambahPabrik()
		case 2:
			tambahMobil()
		case 3:
			ubahMobil()
		case 4:
			hapusMobil()
		case 5:
			cariMobil()
		case 6:
			cariPabrik()
		case 7:
			tampilkanDaftarMobil()
		case 8:
			tampilkanDaftarPabrik()
		case 9:
			tampilkanTop3Penjualan()
		case 0:
			fmt.Println("Terima kasih telah menggunakan program Dealer Mobil!")
			return
		default:
			fmt.Println("Menu yang Anda pilih tidak valid.")
		}
	}
}

func tambahPabrik() {
	if jumlahPabrik == 100 {
		fmt.Println("Kapasitas pabrik sudah penuh.")
		return
	}

	var namaPabrik string
	fmt.Print("Masukkan nama pabrik: ")
	fmt.Scanln(&namaPabrik)

	listPabrik[jumlahPabrik] = Pabrik{
		Nama:        namaPabrik,
		JumlahMobil: 0,
	}

	jumlahPabrik++
	fmt.Println("Pabrik berhasil ditambahkan.")
}

func tambahMobil() {
	if jumlahPabrik == 0 {
		fmt.Println("Belum ada pabrik yang tersedia.")
		return
	}

	var namaPabrik string
	fmt.Print("Masukkan nama pabrik: ")
	fmt.Scanln(&namaPabrik)

	pabrikIndex := cariPabrikIndex(namaPabrik)
	if pabrikIndex == -1 {
		fmt.Println("Pabrik tidak ditemukan.")
		return
	}

	if listPabrik[pabrikIndex].JumlahMobil == 100 {
		fmt.Println("Kapasitas mobil pada pabrik ini sudah penuh.")
		return
	}

	var mobil Mobil
	fmt.Print("Masukkan nama mobil: ")
	fmt.Scanln(&mobil.Nama)
	fmt.Print("Masukkan tahun keluar: ")
	fmt.Scanln(&mobil.Tahun)
	fmt.Print("Masukkan harga mobil: ")
	fmt.Scanln(&mobil.Harga)
	fmt.Print("Masukkan jumlah penjualan: ")
	fmt.Scanln(&mobil.Penjualan)

	listPabrik[pabrikIndex].listMobil[listPabrik[pabrikIndex].JumlahMobil] = mobil
	listPabrik[pabrikIndex].JumlahMobil++

	fmt.Println("Mobil berhasil ditambahkan.")
}

func ubahMobil() {
	if jumlahPabrik == 0 {
		fmt.Println("Belum ada pabrik yang tersedia.")
		return
	}

	var namaPabrik string
	fmt.Print("Masukkan nama pabrik: ")
	fmt.Scanln(&namaPabrik)

	pabrikIndex := cariPabrikIndex(namaPabrik)
	if pabrikIndex == -1 {
		fmt.Println("Pabrik tidak ditemukan.")
		return
	}

	if listPabrik[pabrikIndex].JumlahMobil == 0 {
		fmt.Println("Belum ada mobil yang tersedia di pabrik ini.")
		return
	}

	var namaMobil string
	fmt.Print("Masukkan nama mobil yang ingin diubah: ")
	fmt.Scanln(&namaMobil)

	mobilIndex := cariMobilIndex(pabrikIndex, namaMobil)
	if mobilIndex == -1 {
		fmt.Println("Mobil tidak ditemukan.")
		return
	}

	var mobil Mobil
	fmt.Print("Masukkan nama mobil baru: ")
	fmt.Scanln(&mobil.Nama)
	fmt.Print("Masukkan tahun keluar: ")
	fmt.Scanln(&mobil.Tahun)
	fmt.Print("Masukkan harga mobil: ")
	fmt.Scanln(&mobil.Harga)
	fmt.Print("Masukkan jumlah penjualan: ")
	fmt.Scanln(&mobil.Penjualan)

	listPabrik[pabrikIndex].listMobil[mobilIndex] = mobil

	fmt.Println("Mobil berhasil diubah.")
}

func hapusMobil() {
	if jumlahPabrik == 0 {
		fmt.Println("Belum ada pabrik yang tersedia.")
		return
	}

	var namaPabrik string
	fmt.Print("Masukkan nama pabrik: ")
	fmt.Scanln(&namaPabrik)

	pabrikIndex := cariPabrikIndex(namaPabrik)
	if pabrikIndex == -1 {
		fmt.Println("Pabrik tidak ditemukan.")
		return
	}

	if listPabrik[pabrikIndex].JumlahMobil == 0 {
		fmt.Println("Belum ada mobil yang tersedia di pabrik ini.")
		return
	}

	var namaMobil string
	fmt.Print("Masukkan nama mobil yang ingin dihapus: ")
	fmt.Scanln(&namaMobil)

	mobilIndex := cariMobilIndex(pabrikIndex, namaMobil)
	if mobilIndex == -1 {
		fmt.Println("Mobil tidak ditemukan.")
		return
	}

	// Menggeser data mobil setelah mobil yang dihapus
	for i := mobilIndex; i < listPabrik[pabrikIndex].JumlahMobil-1; i++ {
		listPabrik[pabrikIndex].listMobil[i] = listPabrik[pabrikIndex].listMobil[i+1]
	}

	listPabrik[pabrikIndex].JumlahMobil--

	fmt.Println("Mobil berhasil dihapus.")
}

func cariMobil() {
	if jumlahPabrik == 0 {
		fmt.Println("Belum ada pabrik yang tersedia.")
		return
	}

	var namaMobil string
	fmt.Print("Masukkan nama mobil yang ingin dicari: ")
	fmt.Scanln(&namaMobil)

	found := false
	for i := 0; i < jumlahPabrik; i++ {
		for j := 0; j < listPabrik[i].JumlahMobil; j++ {
			if strings.EqualFold(listPabrik[i].listMobil[j].Nama, namaMobil) {
				fmt.Printf("Nama Pabrik: %s\n", listPabrik[i].Nama)
				fmt.Printf("Nama Mobil: %s\n", listPabrik[i].listMobil[j].Nama)
				fmt.Printf("Tahun Keluar: %d\n", listPabrik[i].listMobil[j].Tahun)
				fmt.Printf("Harga: %d\n", listPabrik[i].listMobil[j].Harga)
				fmt.Printf("Jumlah Penjualan: %d\n", listPabrik[i].listMobil[j].Penjualan)
				found = true
				break
			}
		}
	}

	if !found {
		fmt.Println("Mobil tidak ditemukan.")
	}
}

func cariPabrik() {
	if jumlahPabrik == 0 {
		fmt.Println("Belum ada pabrik yang tersedia.")
		return
	}

	var namaPabrik string
	fmt.Print("Masukkan nama pabrik yang ingin dicari: ")
	fmt.Scanln(&namaPabrik)

	pabrikIndex := cariPabrikIndex(namaPabrik)
	if pabrikIndex == -1 {
		fmt.Println("Pabrik tidak ditemukan.")
		return
	}

	fmt.Printf("Nama Pabrik: %s\n", listPabrik[pabrikIndex].Nama)
	fmt.Printf("Jumlah Mobil: %d\n", listPabrik[pabrikIndex].JumlahMobil)
}

func tampilkanDaftarMobil() {
	if jumlahPabrik == 0 {
		fmt.Println("Belum ada pabrik yang tersedia.")
		return
	}

	for i := 0; i < jumlahPabrik; i++ {
		fmt.Printf("Nama Pabrik: %s\n", listPabrik[i].Nama)
		fmt.Println("Daftar Mobil:")
		for j := 0; j < listPabrik[i].JumlahMobil; j++ {
			fmt.Printf("- %s\n", listPabrik[i].listMobil[j].Nama)
		}
		fmt.Println()
	}
}

func tampilkanDaftarPabrik() {
	if jumlahPabrik == 0 {
		fmt.Println("Belum ada pabrik yang tersedia.")
		return
	}

	fmt.Println("Daftar Pabrik:")
	for i := 0; i < jumlahPabrik; i++ {
		fmt.Printf("- %s\n", listPabrik[i].Nama)
	}
}

func tampilkanTop3Penjualan() {
	if jumlahPabrik == 0 {
		fmt.Println("Belum ada pabrik yang tersedia.")
		return
	}

	type PabrikPenjualan struct {
		NamaPabrik   string
		JumlahPenjualan int
	}

	type MobilPenjualan struct {
		NamaMobil   string
		JumlahPenjualan int
	}

	pabrikPenjualan := make([]PabrikPenjualan, jumlahPabrik)
	mobilPenjualan := make([]MobilPenjualan, 0)

	for i := 0; i < jumlahPabrik; i++ {
		pabrikPenjualan[i] = PabrikPenjualan{
			NamaPabrik:   listPabrik[i].Nama,
			JumlahPenjualan: 0,
		}

		for j := 0; j < listPabrik[i].JumlahMobil; j++ {
			pabrikPenjualan[i].JumlahPenjualan += listPabrik[i].listMobil[j].Penjualan

			mobilPenjualan = append(mobilPenjualan, MobilPenjualan{
				NamaMobil:   listPabrik[i].listMobil[j].Nama,
				JumlahPenjualan: listPabrik[i].listMobil[j].Penjualan,
			})
		}
	}

	// Mengurutkan pabrik berdasarkan jumlah penjualan
	for i := 0; i < jumlahPabrik-1; i++ {
		minIndex := i
		for j := i + 1; j < jumlahPabrik; j++ {
			if pabrikPenjualan[j].JumlahPenjualan > pabrikPenjualan[minIndex].JumlahPenjualan {
				minIndex = j
			}
		}
		pabrikPenjualan[i], pabrikPenjualan[minIndex] = pabrikPenjualan[minIndex], pabrikPenjualan[i]
	}

	// Mengurutkan mobil berdasarkan jumlah penjualan
	for i := 0; i < len(mobilPenjualan)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(mobilPenjualan); j++ {
			if mobilPenjualan[j].JumlahPenjualan > mobilPenjualan[minIndex].JumlahPenjualan {
				minIndex = j
			}
		}
		mobilPenjualan[i], mobilPenjualan[minIndex] = mobilPenjualan[minIndex], mobilPenjualan[i]
	}

	fmt.Println("Top 3 Pabrik dengan Jumlah Penjualan Tertinggi:")
	for i := 0; i < jumlahPabrik && i < 3; i++ {
		fmt.Printf("%d. %s - Jumlah Penjualan: %d\n", i+1, pabrikPenjualan[i].NamaPabrik, pabrikPenjualan[i].JumlahPenjualan)
	}

	fmt.Println()

	fmt.Println("Top 3 Mobil dengan Jumlah Penjualan Tertinggi:")
	for i := 0; i < len(mobilPenjualan) && i < 3; i++ {
		fmt.Printf("%d. %s - Jumlah Penjualan: %d\n", i+1, mobilPenjualan[i].NamaMobil, mobilPenjualan[i].JumlahPenjualan)
	}
}

func cariPabrikIndex(namaPabrik string) int {
	for i := 0; i < jumlahPabrik; i++ {
		if strings.EqualFold(listPabrik[i].Nama, namaPabrik) {
			return i
		}
	}
	return -1
}

func cariMobilIndex(pabrikIndex int, namaMobil string) int {
	for i := 0; i < listPabrik[pabrikIndex].JumlahMobil; i++ {
		if strings.EqualFold(listPabrik[pabrikIndex].listMobil[i].Nama, namaMobil) {
			return i
		}
	}
	return -1
}
