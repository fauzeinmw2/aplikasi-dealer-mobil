package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		// Menampilkan menu
		fmt.Println("=== MENU ===")
		fmt.Println("1. Pilihan 1")
		fmt.Println("2. Pilihan 2")
		fmt.Println("3. Pilihan 3")
		fmt.Println("0. Keluar")

		// Membaca pilihan dari pengguna
		var pilihan int
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		// Memproses pilihan
		switch pilihan {
		case 1:
			fmt.Println("Anda memilih pilihan 1")
			// Tambahkan kode untuk menghandle pilihan 1
		case 2:
			fmt.Println("Anda memilih pilihan 2")
			// Tambahkan kode untuk menghandle pilihan 2
		case 3:
			fmt.Println("Anda memilih pilihan 3")
			// Tambahkan kode untuk menghandle pilihan 3
		case 0:
			fmt.Println("Terima kasih telah menggunakan program ini. Sampai jumpa!")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid. Silakan masukkan pilihan yang benar.")
		}
		fmt.Println()
	}
}
