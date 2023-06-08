package main

import (
	"fmt"
	"strings"
)

func main() {
	// Data untuk tabel
	data := [][]string{
		{"No", "Nama", "Umur"},
		{"1", "John sbajh sdjhsdhj sdjhvbhjsd sdcvbsdhjc", "30"},
		{"2", "Jane sabfkjdfjkds dskjnfsdjv", "25"},
		{"3", "Doe", "35"},
	}

	// Mendapatkan lebar kolom maksimum
	colWidths := make([]int, len(data[0]))
	for _, row := range data {
		for i, col := range row {
			if len(col) > colWidths[i] {
				colWidths[i] = len(col)
			}
		}
	}

	// Membuat baris pemisah
	var separator string
	for _, width := range colWidths {
		separator += "+" + strings.Repeat("-", width+2)
	}
	separator += "+"

	// Menampilkan tabel
	for _, row := range data {
		fmt.Println(separator)
		for i, col := range row {
			fmt.Printf("| %-*s ", colWidths[i], col)
		}
		fmt.Println("|")
	}
	fmt.Println(separator)
}
