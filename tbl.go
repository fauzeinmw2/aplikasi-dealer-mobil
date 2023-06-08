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
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			col := data[i][j]
			if len(col) > colWidths[j] {
				colWidths[j] = len(col)
			}
		}
	}

	// Membuat baris pemisah
	var separator string
	for k := 0; k < len(colWidths); k++ {
		width := colWidths[k]
		separator += "+" + strings.Repeat("-", width+2)
	}
	separator += "+"

	// Menampilkan tabel
	for l := 0; l < len(data); l++ {
		fmt.Println(separator)
		for m := 0; m < len(data[l]); m++ {
			col := data[l][m]
			fmt.Printf("| %-*s ", colWidths[m], col)
		}
		fmt.Println("|")
	}
	fmt.Println(separator)
}
