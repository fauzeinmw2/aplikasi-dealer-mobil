package main

import (
	"fmt"
	"sort"
)

// Tipe bentukan Mobil
type Mobil struct {
	Nama     string
	Tahun    int
	Penjualan int
}

// Tipe bentukan Pabrik
type Pabrik struct {
	Nama       string
	ListMobil  []Mobil
}

// Array utama yang akan digunakan
var daftarPabrik []Pabrik

// Fungsi untuk menambahkan data pabrik
func tambahPabrik(nama string) {
	pabrik := Pabrik{
		Nama:      nama,
		ListMobil: []Mobil{},
	}
	daftarPabrik = append(daftarPabrik, pabrik)
	fmt.Println("Data pabrik berhasil ditambahkan.")
}

// Fungsi untuk mengubah data pabrik
func ubahPabrik(namaLama string, namaBaru string) {
	for i := range daftarPabrik {
		if daftarPabrik[i].Nama == namaLama {
			daftarPabrik[i].Nama = namaBaru
			fmt.Println("Data pabrik berhasil diubah.")
			return
		}
	}
	fmt.Println("Pabrik tidak ditemukan.")
}

// Fungsi untuk menghapus data pabrik
func hapusPabrik(nama string) {
	for i := range daftarPabrik {
		if daftarPabrik[i].Nama == nama {
			copy(daftarPabrik[i:], daftarPabrik[i+1:])
			daftarPabrik = daftarPabrik[:len(daftarPabrik)-1]
			fmt.Println("Data pabrik berhasil dihapus.")
			return
		}
	}
	fmt.Println("Pabrik tidak ditemukan.")
}

// Fungsi untuk menambahkan data mobil ke pabrik
func tambahMobil(namaPabrik string, mobil Mobil) {
	for i := range daftarPabrik {
		if daftarPabrik[i].Nama == namaPabrik {
			daftarPabrik[i].ListMobil = append(daftarPabrik[i].ListMobil, mobil)
			fmt.Println("Data mobil berhasil ditambahkan.")
			return
		}
	}
	fmt.Println("Pabrik tidak ditemukan.")
}

// Fungsi untuk mengubah data mobil di pabrik
func ubahMobil(namaPabrik string, namaMobil string, mobil Mobil) {
	for i := range daftarPabrik {
		if daftarPabrik[i].Nama == namaPabrik {
			for j := range daftarPabrik[i].ListMobil {
				if daftarPabrik[i].ListMobil[j].Nama == namaMobil {
					daftarPabrik[i].ListMobil[j] = mobil
					fmt.Println("Data mobil berhasil diubah.")
					return
				}
			}
			fmt.Println("Mobil tidak ditemukan.")
			return
		}
	}
	fmt.Println("Pabrik tidak ditemukan.")
}

// Fungsi untuk menghapus data mobil di pabrik
func hapusMobil(namaPabrik string, namaMobil string) {
	for i := range daftarPabrik {
		if daftarPabrik[i].Nama == namaPabrik {
			for j := range daftarPabrik[i].ListMobil {
				if daftarPabrik[i].ListMobil[j].Nama == namaMobil {
					copy(daftarPabrik[i].ListMobil[j:], daftarPabrik[i].ListMobil[j+1:])
					daftarPabrik[i].ListMobil = daftarPabrik[i].ListMobil[:len(daftarPabrik[i].ListMobil)-1]
					fmt.Println("Data mobil berhasil dihapus.")
					return
				}
			}
			fmt.Println("Mobil tidak ditemukan.")
			return
		}
	}
	fmt.Println("Pabrik tidak ditemukan.")
}

// Fungsi untuk mencari mobil berdasarkan nama pabrik
func cariMobilByPabrik(namaPabrik string) []Mobil {
	var result []Mobil
	for _, pabrik := range daftarPabrik {
		if pabrik.Nama == namaPabrik {
			result = append(result, pabrik.ListMobil...)
			break
		}
	}
	return result
}

// Fungsi untuk mencari mobil berdasarkan nama mobil
func cariMobilByNama(namaMobil string) []Mobil {
	var result []Mobil
	for _, pabrik := range daftarPabrik {
		for _, mobil := range pabrik.ListMobil {
			if mobil.Nama == namaMobil {
				result = append(result, mobil)
			}
		}
	}
	return result
}

// Fungsi untuk menampilkan daftar mobil terurut berdasarkan tahun keluar secara descending
func tampilkanMobilTerurut() {
	var allMobil []Mobil
	for _, pabrik := range daftarPabrik {
		allMobil = append(allMobil, pabrik.ListMobil...)
	}
	sort.SliceStable(allMobil, func(i, j int) bool {
		return allMobil[i].Tahun > allMobil[j].Tahun
	})

	fmt.Println("Daftar mobil terurut berdasarkan tahun keluar:")
	for _, mobil := range allMobil {
		fmt.Printf("Nama: %s, Tahun: %d, Penjualan: %d\n", mobil.Nama, mobil.Tahun, mobil.Penjualan)
	}
}

// Fungsi untuk menampilkan 3 daftar mobil dan pabrikan dengan jumlah penjualan tertinggi
func tampilkanDaftarPenjualanTertinggi() {
	type Penjualan struct {
		Pabrik string
		Mobil  string
		Jumlah int
	}

	var allPenjualan []Penjualan
	for _, pabrik := range daftarPabrik {
		for _, mobil := range pabrik.ListMobil {
			allPenjualan = append(allPenjualan, Penjualan{
				Pabrik: pabrik.Nama,
				Mobil:  mobil.Nama,
				Jumlah: mobil.Penjualan,
			})
		}
	}

	sort.SliceStable(allPenjualan, func(i, j int) bool {
		return allPenjualan[i].Jumlah > allPenjualan[j].Jumlah
	})

	fmt.Println("Daftar 3 pabrik dengan penjualan tertinggi:")
	for i, penjualan := range allPenjualan {
		if i == 3 {
			break
		}
		fmt.Printf("Pabrik: %s, Mobil: %s, Jumlah Penjualan: %d\n", penjualan.Pabrik, penjualan.Mobil, penjualan.Jumlah)
	}
}

// Fungsi untuk menampilkan seluruh data pabrik
func tampilkanDataPabrik() {
	fmt.Println("Daftar pabrik:")
	for _, pabrik := range daftarPabrik {
		fmt.Println("Nama: ", pabrik.Nama)
	}
}

// Fungsi untuk menampilkan seluruh data mobil
func tampilkanDataMobil() {
	fmt.Println("Daftar mobil:")
	for _, pabrik := range daftarPabrik {
		for _, mobil := range pabrik.ListMobil {
			fmt.Printf("Nama: %s, Tahun: %d, Penjualan: %d\n", mobil.Nama, mobil.Tahun, mobil.Penjualan)
		}
	}
}

func main() {
	tambahPabrik("Pabrik A")
	tambahPabrik("Pabrik B")

	mobilA1 := Mobil{
		Nama:     "Mobil A1",
		Tahun:    2022,
		Penjualan: 100,
	}
	mobilA2 := Mobil{
		Nama:     "Mobil A2",
		Tahun:    2023,
		Penjualan: 200,
	}
	mobilB1 := Mobil{
		Nama:     "Mobil B1",
		Tahun:    2021,
		Penjualan: 150,
	}

	mobilA3 := Mobil{
		Nama:     "Mobil A3",
		Tahun:    2023,
		Penjualan: 112,
	}

	fmt.Println("\n=========")
	tambahMobil("Pabrik A", mobilA1)
	tambahMobil("Pabrik A", mobilA2)
	tambahMobil("Pabrik B", mobilB1)
	tambahMobil("Pabrik A", mobilA3)

	fmt.Println("\n=========")
	tampilkanDataPabrik()

	fmt.Println("\n=========")
	tampilkanDataMobil()

	// Contoh penggunaan fungsi lainnya
	fmt.Println("\n=========")
	tampilkanMobilTerurut()

	fmt.Println("\n=========")
	tampilkanDaftarPenjualanTertinggi()

	// Contoh penggunaan fungsi ubahMobil
	mobilA4 := Mobil{
		Nama:     "Mobil A4",
		Tahun:    2024,
		Penjualan: 300,
	}
	fmt.Println("\n=========")
	ubahMobil("Pabrik A", "Mobil A2", mobilA4)

	fmt.Println("\n=========")
	tampilkanDataMobil()

	// Contoh penggunaan fungsi hapusPabrik
	fmt.Println("\n=========")
	hapusPabrik("Pabrik B")

	fmt.Println("\n=========")
	tampilkanDataPabrik()

	fmt.Println("\n=========")
	tampilkanDataMobil()
}
