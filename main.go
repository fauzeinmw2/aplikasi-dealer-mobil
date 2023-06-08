package main
import ( "fmt"; "os"; "os/exec"; "strconv")

const NMAXPABRIK int = 100
const NMAXALLMOBIL int = 1000
const NMAXMOBIL int = 10

type Pabrik struct {
	merek string
	daftarMobil[NMAXMOBIL] Mobil
	nMobil int
}

type Mobil struct {
	nama string
	tahunKeluar, harga, penjualan, jumlah int
}

type AllMobil struct {
	mobil Mobil
	pabrik string
}

type arrPabrik[NMAXPABRIK] Pabrik
type arrAllMobil[NMAXALLMOBIL] AllMobil
type arrInt[100] int
type arrStr[100] string

var pesan string = ""
var dataPabrik arrPabrik
var dataAllMobil arrAllMobil
var nPabrik, nAllMobil int

// DONE
func main(){
	dataPabrik[0].merek = "Toyota"
	dataPabrik[1].merek = "Daihatsu"
	dataPabrik[2].merek = "Hyundai"
	dataPabrik[3].merek = "Mitsubishi"
	nPabrik+=4

	dataPabrik[0].daftarMobil[0].nama = "Avanza"
	dataPabrik[0].daftarMobil[0].tahunKeluar = 2019
	dataPabrik[0].daftarMobil[0].harga = 100000000
	dataPabrik[0].daftarMobil[0].penjualan = 12
	dataPabrik[0].daftarMobil[0].jumlah = 5
	dataPabrik[0].nMobil++

	dataPabrik[0].daftarMobil[1].nama = "Rush"
	dataPabrik[0].daftarMobil[1].tahunKeluar = 2023
	dataPabrik[0].daftarMobil[1].harga = 150000000
	dataPabrik[0].daftarMobil[1].penjualan = 2
	dataPabrik[0].daftarMobil[1].jumlah = 2
	dataPabrik[0].nMobil++

	dataPabrik[1].daftarMobil[0].nama = "Sirion"
	dataPabrik[1].daftarMobil[0].tahunKeluar = 2023
	dataPabrik[1].daftarMobil[0].harga = 228000000
	dataPabrik[1].daftarMobil[0].penjualan = 1
	dataPabrik[1].daftarMobil[0].jumlah = 3
	dataPabrik[1].nMobil++

	dataPabrik[2].daftarMobil[0].nama = "Stargazer"
	dataPabrik[2].daftarMobil[0].tahunKeluar = 2022
	dataPabrik[2].daftarMobil[0].harga = 307000000
	dataPabrik[2].daftarMobil[0].penjualan = 5
	dataPabrik[2].daftarMobil[0].jumlah = 1
	dataPabrik[2].nMobil++

	dataPabrik[3].daftarMobil[0].nama = "XPander"
	dataPabrik[3].daftarMobil[0].tahunKeluar = 2021
	dataPabrik[3].daftarMobil[0].harga = 250000000
	dataPabrik[3].daftarMobil[0].penjualan = 0
	dataPabrik[3].daftarMobil[0].jumlah = 4
	dataPabrik[3].nMobil++

	mainMenu()
}

// DONE
func clear(){
	cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

// DONE
func header(){
	fmt.Println("=================================")
	fmt.Println("   Aplikasi Dealer Mobil Josss   ")
	fmt.Println("=================================")
}

// DONE
func tampilTabel(n, nData int, lebar arrInt, title arrStr, dataP arrPabrik, dataM arrAllMobil, flag string){
	var pemisah string
	var i, penjualan int

	for i = 0; i < n; i++ {
		pemisah += "+"
		for j := 0; j < lebar[i]+2; j++ {
			pemisah += "-"
		}
	}
	pemisah += "+"
	fmt.Println(pemisah)

	for i=0; i<n; i++{
		fmt.Printf("| %-*s ", lebar[i], title[i])
	}
	fmt.Print("|\n")
	i=0
	for i < nData && nData != 0{
		fmt.Println(pemisah)

		if flag == "pabrik" {
			var row = dataP[i]
			penjualan = totalPenjualanPabrik(dataP, i)
			fmt.Printf("| %-*d ", lebar[0], i+1)
			fmt.Printf("| %-*s ", lebar[1], row.merek)
			fmt.Printf("| %-*d ", lebar[2], row.nMobil)
			fmt.Printf("| %-*d |\n", lebar[3], penjualan)
		}else{
			var row = dataAllMobil[i]
			fmt.Printf("| %-*d ", lebar[0], i+1)
			fmt.Printf("| %-*s ", lebar[1], row.mobil.nama)
			fmt.Printf("| %-*s ", lebar[2], row.pabrik)
			fmt.Printf("| %-*d ", lebar[3], row.mobil.tahunKeluar)
			fmt.Printf("| %-*d ", lebar[4], row.mobil.harga)
			fmt.Printf("| %-*d ", lebar[5], row.mobil.penjualan)
			fmt.Printf("| %-*d |\n", lebar[6], row.mobil.jumlah)
		}
		i++
	}

	if n == 0 {
		fmt.Printf("\n%s", pemisah)
		fmt.Println("Data Masih kosong !!!")
	}
	fmt.Println(pemisah)
}

// DONE
func totalPenjualanPabrik(data arrPabrik, idx int) int {
	var i, total int
	for i < data[idx].nMobil {
		total += data[idx].daftarMobil[i].penjualan
		i++
	}
	return total
}

// DONE
func setAllMobil() {
	var i, j int
	nAllMobil = 0
	for i=0; i<nPabrik; i++{
		for j=0; j<dataPabrik[i].nMobil; j++ {
			dataAllMobil[nAllMobil].mobil = dataPabrik[i].daftarMobil[j]
			dataAllMobil[nAllMobil].pabrik = dataPabrik[i].merek
			nAllMobil++
		}
	}
}

// DONE
func mainMenu(){
	var lanjut bool = true
	var pil int

	for lanjut {
		clear(); header()
		fmt.Println(pesan)
		pesan = ""
		fmt.Println("1. Kelola Data Pabrikan")
		fmt.Println("2. Kelola Data Mobil")
		fmt.Println("3. Lihat Data Penjualan")
		fmt.Println("4. Exit")
		fmt.Println("=================================")
		fmt.Print("Masukkan Pilihan [1-4] : ")
		fmt.Scan(&pil)
	
		switch pil {
			case 1: menuPabrik(); lanjut = true
			case 2: menuMobil(); lanjut = true
			case 3: menuPenjualan(dataPabrik, nPabrik); lanjut = true
			case 4: lanjut = false
			default: pesan = "Pilihan tidak tersedia !!!\n"; lanjut = true
		}
	}
	
}

// DONE
func menuPabrik(){
	var lanjut bool = true
	var pil, penjualan, idx int
	var x, confirm string

	for lanjut {
		clear(); header()
		tampilPabrik(dataPabrik, nPabrik)
		fmt.Print("\n", pesan)
		pesan = ""
		subMenu()
		fmt.Print("Masukkan Pilihan [1-6] : ")
		fmt.Scan(&pil)
	
		switch pil {
			case 1: tambahPabrik(); lanjut = true
			case 2: editPabrik(); lanjut = true
			case 3: hapusPabrik(); lanjut = true
			case 4: 
				for lanjut {
					fmt.Print("Masukkan Merek Pabrik : ")
					fmt.Scan(&x)
					idx = cariPabrik(x)
					if idx >= 0 {
						fmt.Println("\nMerek Pabrik :", dataPabrik[idx].merek)
						fmt.Println("Banyak Mobil :", dataPabrik[idx].nMobil)
						penjualan = totalPenjualanPabrik(dataPabrik, idx)
						fmt.Println("Penjualan :", penjualan)
					}else{
						fmt.Println("Data Tidak Tersedia!!!")
					}

					fmt.Print("\nCari Lagi ? [y/n] : ")
					fmt.Scan(&confirm)
					if confirm == "n" || confirm == "N" {
						lanjut = false
					}
				}
				lanjut = true
			case 5: menuUrutPabrik(dataPabrik, nPabrik); lanjut = true
			case 6: lanjut = false
			default: pesan = "Pilihan tidak tersedia !!!\n"; lanjut = true
		}
	}
}

// DONE
func menuMobil(){
	var lanjut bool = true
	var pil, idx int
	var x, confirm string

	for lanjut {
		clear(); header()
		setAllMobil()
		tampilMobil(dataAllMobil, nAllMobil)
		fmt.Print("\n", pesan)
		pesan = ""
		subMenu()
		fmt.Print("Masukkan Pilihan [1-6] : ")
		fmt.Scan(&pil)
	
		switch pil {
			case 1: tambahMobil(); lanjut = true
			case 2: editMobil(); lanjut = true
			case 3: hapusMobil(); lanjut = true
			case 4: 
				for lanjut {
					fmt.Print("Masukkan Nama Mobil : ")
					fmt.Scan(&x)
					idx = cariMobil(x)
					if idx >= 0 {
						fmt.Println("\nNama Mobil :", dataAllMobil[idx].mobil.nama)
						fmt.Println("Merek Pabrik :", dataAllMobil[idx].pabrik)
						fmt.Println("Tahun Keluar :", dataAllMobil[idx].mobil.tahunKeluar)
						fmt.Println("Harga :", dataAllMobil[idx].mobil.harga)
						fmt.Println("Jumlah :", dataAllMobil[idx].mobil.jumlah)
					}else{
						fmt.Println("Data Tidak Tersedia!!!")
					}

					fmt.Print("\nCari Lagi ? [y/n] : ")
					fmt.Scan(&confirm)
					if confirm == "n" || confirm == "N" {
						lanjut = false
					}
				}
				lanjut = true
			case 5: menuUrutMobil(); lanjut = true
			case 6: lanjut = false
			default: pesan = "Pilihan tidak tersedia !!!\n"; lanjut = true
		}
	}
}


func menuPenjualan(dataP arrPabrik, nP int){
	var pabrik arrPabrik
	var mobil arrAllMobil
	var x string
	var lanjut bool = true
	setAllMobil()

	for lanjut {
		clear()
		header()
		fmt.Println("=================================")
		fmt.Println("    3 Data Penjualan Tertinggi   ")
		fmt.Println("=================================")
	
		urutPabrik(&dataP, nP, 2, 2)
		urutMobil(2, 4)
	
		for i:=0; i<3; i++{
			pabrik[i] = dataP[i]
			mobil[i] = dataAllMobil[i]
		}
	
		tampilPabrik(pabrik, 3)
		tampilMobil(mobil, 3)
	
		fmt.Print("\nKembali [Y/N] : ")
		fmt.Scan(&x)

		if x == "Y" || x == "y" {
			lanjut = false
		}
	}
	
}

// DONE
func subMenu(){
	fmt.Println("=================================")
	fmt.Println("1. Tambah Data")
	fmt.Println("2. Edit Data")
	fmt.Println("3. Hapus Data")
	fmt.Println("4. Cari Data")
	fmt.Println("5. Urutkan Data")
	fmt.Println("6. Kembali")
	fmt.Println("=================================")
}

// DONE
func tampilPabrik(dataP arrPabrik, n int){
	var row Pabrik
	var lebar arrInt
	var i, penjualan int
	var title arrStr

	title[0] = "No"; title[1] = "Merek"; title[2] = "Total Jenis Mobil"; title[3] = "Total Penjualan"
	lebar[0] = len(title[0]); lebar[1] = len(title[1]); lebar[2] = len(title[2]); lebar[3] = len(title[3])

	for i = 0; i < n; i++ {
		row = dataP[i]
		if len(strconv.Itoa(i+1)) > lebar[0] {
			lebar[0] = len(strconv.Itoa(i+1))
		}
		if len(row.merek) > lebar[1] {
			lebar[1] = len(row.merek)
		}
		if len(strconv.Itoa(row.nMobil)) > lebar[2] {
			lebar[2] = len(strconv.Itoa(row.nMobil))
		}
		penjualan = totalPenjualanPabrik(dataPabrik, i)
		if len(strconv.Itoa(penjualan)) > lebar[3] {
			lebar[3] = len(strconv.Itoa(penjualan))
		}
	}

	// clear()
	// header()
	fmt.Println("           Data Pabrikan        ")
	tampilTabel(4, n, lebar, title, dataP, dataAllMobil, "pabrik")
}

// DONE
func tambahPabrik(){
	var confirm, merek string
	clear()
	header()

	fmt.Println("       Tambah Data Pabrik        ")
	fmt.Println("=================================")
	fmt.Print("\nMasukkan Merek Pabrik : ")
	fmt.Scan(&merek)
	fmt.Print("\nSimpan Data Ini? [Y/N] : ")
	fmt.Scan(&confirm)

	pesan = "Data Pabrik Gagal Ditambahkan...\n"
	if confirm == "Y" || confirm == "y" {
		dataPabrik[nPabrik].merek = merek
		nPabrik++
		pesan = "Data Pabrik Baru Berhasi Ditambahkan...\n"
	}
}

// DONE
func editPabrik(){
	var confirm, merek, cari string
	var idx int

	clear()
	header()
	fmt.Println("         Edit Data Pabrik        ")
	fmt.Println("=================================")
	fmt.Print("\nMasukkan Merek Pabrik Yang Akan Diedit : ")
	fmt.Scan(&cari)
	idx = cariPabrik(cari)

	pesan = "Data Pabrik Batal Diedit...\n"
	if idx >= 0 {
		fmt.Print("\nMasukkan Merek Pabrik Yang Baru : ")
		fmt.Scan(&merek)
		fmt.Print("\nEdit Data Ini? [Y/N] : ")
		fmt.Scan(&confirm)

		if confirm == "Y" || confirm == "y" {
			dataPabrik[idx].merek = merek
			pesan = "Data Pabrik Berhasil Diedit...\n"
		}
	}else{
		pesan = "Gagal... Data Pabrik Tidak Ditemukan...\n"
	}	
}

// DONE
func hapusPabrik(){
	var idx int
	var cari, confirm string

	fmt.Print("\nMasukkan Merek Pabrik yang Akan Dihapus : ")
	fmt.Scan(&cari)
	idx = cariPabrik(cari)

	pesan = "Data Pabrik Batal Dihapus...\n"
	if idx >= 0 {
		fmt.Print("\nHapus Data Ini? [Y/N] : ")
		fmt.Scan(&confirm)

		if confirm == "Y" || confirm == "y" {
			for idx < nPabrik-1 {
				dataPabrik[idx] = dataPabrik[idx+1]
				idx++
			}
			nPabrik--
			pesan = "Data Pabrik Berhasil Dihapus...\n"
		}
	}else{
		pesan = "Gagal... Data Pabrik Tidak Ditemukan...\n"
	}	
}

// DONE
func cariPabrik(x string) int {
	var i, idx int = 0, -1
	var found bool
	for i < nPabrik && !found {
		if dataPabrik[i].merek == x {
			idx = i
			found = false
		}
		i++
	}
	return idx
}

// DONE
func menuUrutPabrik(data arrPabrik, n int){
	var lanjut bool = true
	var pil int

	for lanjut {
		clear(); header()
		tampilPabrik(dataPabrik, nPabrik)
		fmt.Print("\n", pesan)
		pesan = ""
		fmt.Println("=================================")
		fmt.Println("1. Berdasarkan Jumlah Mobil")
		fmt.Println("2. Berdasarkan Jumlah Penjualan")
		fmt.Println("3. Kembali")
		fmt.Println("=================================")
		fmt.Print("Masukkan Pilihan [1-3] : ")
		fmt.Scan(&pil)
	
		switch pil {
			case 1: 
				for lanjut {
					fmt.Println("1. Ascending")
					fmt.Println("2. Descending")
					fmt.Println("3. Kembali")
					fmt.Print("Pilih : ")
					fmt.Scan(&pil)
			
					if pil == 3 {
						lanjut = false
					}
					clear()
					header()
					urutPabrik(&data, n, 1, pil)
					tampilPabrik(data, n)
				}
				lanjut = true
			case 2: 
				for lanjut {
					fmt.Println("1. Ascending")
					fmt.Println("2. Descending")
					fmt.Println("3. Kembali")
					fmt.Print("Pilih : ")
					fmt.Scan(&pil)
			
					if pil == 3 {
						lanjut = false
					}
					clear()
					header()
					urutPabrik(&data, n, 2, pil)
					tampilPabrik(data, n)
				}
				lanjut = true
			case 3: lanjut = false
			default: pesan = "Pilihan tidak tersedia !!!\n"; lanjut = true
		}
	}
}

func urutPabrik(data *arrPabrik, n, flag, pil int){
	var i, pass, penjualanT int
	var temp Pabrik
	var penjualanP arrInt

	i = 0
	for i < n {
		penjualanP[i] = totalPenjualanPabrik(*data, i)
		i++
	}

	// Insertion Sort
	pass = 1
	for pass <= n-1 && pil != 3 {
		penjualanT = 0
		i = pass
		temp = data[pass]

		var j int = 0
		for j < temp.nMobil {
			penjualanT += temp.daftarMobil[j].penjualan
			j++
		}

		for i > 0 && ( (temp.nMobil < data[i-1].nMobil && pil == 1 && flag == 1) || (temp.nMobil > data[i-1].nMobil && pil == 2 && flag == 1) ||
			(penjualanT < penjualanP[i-1] && pil == 1 && flag == 2) || (penjualanT > penjualanP[i-1] && pil == 2 && flag == 2) ){
			data[i] = data[i-1]
			penjualanP[i] = penjualanP[i-1]
			i--
		}

		data[i] = temp
		penjualanP[i] = penjualanT
		pass++
	}
}

// DONE
func tampilMobil(data arrAllMobil, n int){
	var row AllMobil
	var lebar arrInt
	var i int
	var title arrStr

	title[0] = "No"; title[1] = "Nama"; title[2] = "Pabrik"; title[3] = "Tahun Keluar"; title[4] = "Harga"; title[5] = "Penjualan"; title[6] = "Jumlah"
	lebar[0] = len(title[0]); lebar[1] = len(title[1]); lebar[2] = len(title[2]); lebar[3] = len(title[3]); lebar[4] = len(title[4]); lebar[5] = len(title[5]); lebar[6] = len(title[6])

	for i = 0; i < n; i++ {
		row = data[i]
		if len(strconv.Itoa(i+1)) > lebar[0] {
			lebar[0] = len(strconv.Itoa(i+1))
		}
		if len(row.mobil.nama) > lebar[1] {
			lebar[1] = len(row.mobil.nama)
		}
		if len(row.pabrik) > lebar[2] {
			lebar[2] = len(row.pabrik)
		}
		if len(strconv.Itoa(row.mobil.tahunKeluar)) > lebar[3] {
			lebar[3] = len(strconv.Itoa(row.mobil.tahunKeluar))
		}
		if len(strconv.Itoa(row.mobil.harga)) > lebar[4] {
			lebar[4] = len(strconv.Itoa(row.mobil.harga))
		}
		if len(strconv.Itoa(row.mobil.penjualan)) > lebar[5] {
			lebar[5] = len(strconv.Itoa(row.mobil.penjualan))
		}
		if len(strconv.Itoa(row.mobil.jumlah)) > lebar[6] {
			lebar[6] = len(strconv.Itoa(row.mobil.jumlah))
		}
	}

	// clear()
	// header()
	fmt.Println("            Data Mobil           ")
	fmt.Println("=================================")

	tampilTabel(7, n, lebar, title, dataPabrik, data, "mobil")
}

// DONE
func tambahMobil(){
	var confirm, merek, nama string
	var idx, tahunKeluar, harga, penjualan, jumlah int
	clear()
	header()

	fmt.Println("        Tambah Data Mobil        ")
	fmt.Println("=================================")
	fmt.Print("\nMasukkan Merek Pabrik : ")
	fmt.Scan(&merek)

	idx = cariPabrik(merek)

	pesan = "Data Mobil Batal Ditambahkan...\n"
	if idx >= 0 {
		fmt.Print("\nNama Mobil : ")
		fmt.Scan(&nama)
		fmt.Print("Tahun Keluar : ")
		fmt.Scan(&tahunKeluar)
		fmt.Print("Harga : ")
		fmt.Scan(&harga)
		fmt.Print("Total Penjualan (unit) : ")
		fmt.Scan(&penjualan)
		fmt.Print("Jumlah Mobil (Unit) : ")
		fmt.Scan(&jumlah)

		fmt.Print("\nSimpan Data Ini? [Y/N] : ")
		fmt.Scan(&confirm)

		if confirm == "Y" || confirm == "y" {
			var nMobil = dataPabrik[idx].nMobil
			dataPabrik[idx].daftarMobil[nMobil].nama = nama
			dataPabrik[idx].daftarMobil[nMobil].tahunKeluar = tahunKeluar
			dataPabrik[idx].daftarMobil[nMobil].harga = harga
			dataPabrik[idx].daftarMobil[nMobil].penjualan = penjualan
			dataPabrik[idx].daftarMobil[nMobil].jumlah = jumlah
			dataPabrik[idx].nMobil++
			pesan = "Data Mobil Baru Berhasi Ditambahkan...\n"
		}
	}else{
		pesan = "Gagal... Data Pabrik Tidak Ditemukan...\n"
	}
}

// DONE
func getIndexMobil(idxP int, x string) int {
	var i, idx int = 0, -1
	var found bool = false
	for i < dataPabrik[idxP].nMobil && !found{
		if dataPabrik[idxP].daftarMobil[i].nama == x {
			found = true
			idx = i
		}
		i++
	}
	return idx
}

// DONE
func editMobil(){
	var confirm, merek, nama string
	var idxM, idxP, tahunKeluar, harga, penjualan, jumlah int
	clear()
	header()

	fmt.Println("         Edit Data Mobil         ")
	fmt.Println("=================================")
	fmt.Print("\nMasukkan Nama Mobil : ")
	fmt.Scan(&nama)

	idxM = cariMobil(nama)
	pesan = "Data Mobil Batal Diubah...\n"

	if idxM >= 0 {
		merek = dataAllMobil[idxM].pabrik
		idxP = cariPabrik(merek)
		idxM = getIndexMobil(idxP, nama)

		fmt.Print("\nNama Mobil : ")
		fmt.Scan(&nama)
		fmt.Print("Tahun Keluar : ")
		fmt.Scan(&tahunKeluar)
		fmt.Print("Harga : ")
		fmt.Scan(&harga)
		fmt.Print("Total Penjualan (unit) : ")
		fmt.Scan(&penjualan)
		fmt.Print("Jumlah Mobil (Unit) : ")
		fmt.Scan(&jumlah)
		fmt.Print("\nSimpan Data Ini? [Y/N] : ")
		fmt.Scan(&confirm)

		if confirm == "Y" || confirm == "y" {
			dataPabrik[idxP].daftarMobil[idxM].nama = nama
			dataPabrik[idxP].daftarMobil[idxM].tahunKeluar = tahunKeluar
			dataPabrik[idxP].daftarMobil[idxM].harga = harga
			dataPabrik[idxP].daftarMobil[idxM].penjualan = penjualan
			dataPabrik[idxP].daftarMobil[idxM].jumlah = jumlah
			pesan = "Data Mobil Baru Berhasi Diubah...\n"
		}
	}else{
		pesan = "Gagal... Data Mobil Tidak Ditemukan...\n"
	}
}

// DONE
func hapusMobil(){
	var confirm, cari string
	var idxM, idxP int
	
	fmt.Print("\nMasukkan Nama Mobil yang Akan Dihapus : ")
	fmt.Scan(&cari)

	idxM = cariMobil(cari)
	pesan = "Data Mobil Batal Dihapus...\n"
	if idxM >= 0 {
		idxP = cariPabrik(dataAllMobil[idxM].pabrik)
		idxM = getIndexMobil(idxP, cari)
		fmt.Print("\nHapus Data Ini? [Y/N] : ")
		fmt.Scan(&confirm)

		if confirm == "Y" || confirm == "y" {
			for idxM < dataPabrik[idxP].nMobil-1 {
				dataPabrik[idxP].daftarMobil[idxM] = dataPabrik[idxP].daftarMobil[idxM+1]
				idxM++
			}
			dataPabrik[idxP].nMobil--
			pesan = "Data Mobil Berhasil Dihapus...\n"
		}
	}else{
		pesan = "Gagal... Data Mobil Tidak Ditemukan...\n"
	}
}

// DONE
func cariMobil(x string) int {
	var i, idx int = 0, -1
	var found bool
	for i < nAllMobil && !found {
		if dataAllMobil[i].mobil.nama == x {
			idx = i
			found = false
		}
		i++
	}
	return idx
}

// DONE
func menuUrutMobil(){
	var lanjut bool = true
	var pil int

	for lanjut {
		clear(); header()
		tampilMobil(dataAllMobil, nAllMobil)
		fmt.Print("\n", pesan)
		pesan = ""
		fmt.Println("=================================")
		fmt.Println("1. Berdasarkan Nama Mobil")
		fmt.Println("2. Berdasarkan Merek Pabrik")
		fmt.Println("3. Berdasarkan Tahun Keluaran")
		fmt.Println("4. Kembali")
		fmt.Println("=================================")
		fmt.Print("Masukkan Pilihan [1-4] : ")
		fmt.Scan(&pil)
	
		switch pil {
			case 1:
				for lanjut {
					fmt.Println("1. Ascending")
					fmt.Println("2. Descending")
					fmt.Println("3. Kembali")
					fmt.Print("Pilih : ")
					fmt.Scan(&pil)
			
					if pil == 3 {
						lanjut = false
					}
					clear(); header()
					urutMobil(pil, 1)
					tampilMobil(dataAllMobil, nAllMobil)
				}
				lanjut = true
			case 2:
				for lanjut {
					fmt.Println("1. Ascending")
					fmt.Println("2. Descending")
					fmt.Println("3. Kembali")
					fmt.Print("Pilih : ")
					fmt.Scan(&pil)
			
					if pil == 3 {
						lanjut = false
					}
					clear(); header()
					urutMobil(pil, 2)
					tampilMobil(dataAllMobil, nAllMobil)
				}
				lanjut = true
			case 3:
				for lanjut {
					fmt.Println("1. Ascending")
					fmt.Println("2. Descending")
					fmt.Println("3. Kembali")
					fmt.Print("Pilih : ")
					fmt.Scan(&pil)
			
					if pil == 3 {
						lanjut = false
					}
					clear(); header()
					urutMobil(pil, 3)
					tampilMobil(dataAllMobil, nAllMobil)
				}
				lanjut = true
			case 4: lanjut = false
			default: pesan = "Pilihan tidak tersedia !!!\n"; lanjut = true
		}
	}
}

// DONE
func urutMobil(pil, flag int){
	var pass, i, idx int
	var temp AllMobil

	// Selection Sort
	pass = 1
	for pass < nAllMobil {
		idx = pass-1
		i = pass
		for i < nAllMobil {
			if (dataAllMobil[idx].mobil.nama > dataAllMobil[i].mobil.nama && pil == 1 && flag == 1) || (dataAllMobil[idx].mobil.nama < dataAllMobil[i].mobil.nama && pil == 2 && flag == 1) ||
				(dataAllMobil[idx].pabrik > dataAllMobil[i].pabrik && pil == 1 && flag == 2) || (dataAllMobil[idx].pabrik < dataAllMobil[i].pabrik && pil == 2 && flag == 2) ||
				(dataAllMobil[idx].mobil.tahunKeluar > dataAllMobil[i].mobil.tahunKeluar && pil == 1 && flag == 3) || (dataAllMobil[idx].mobil.tahunKeluar < dataAllMobil[i].mobil.tahunKeluar && pil == 2 && flag == 3) ||
				(dataAllMobil[idx].mobil.penjualan < dataAllMobil[i].mobil.penjualan && flag == 4) {
					idx = i
			}
			i++
		}
		temp = dataAllMobil[idx]
		dataAllMobil[idx] = dataAllMobil[pass-1]
		dataAllMobil[pass-1] = temp
		pass++
	}
}