package main
import (
	"fmt"
    "os"
    "os/exec"
	"strconv"
)

const NMAX int = 100
const NMAXMOBIL int = 1000

type Pabrik struct {
	merek string
	daftarMobil[10] Mobil
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

type arrPabrik[NMAX] Pabrik
type arrMobil[NMAXMOBIL] AllMobil

var pesan string = ""

var allMobil arrMobil
var nMobil int

func main(){
	var data arrPabrik
	var pil, n int
	
	menu(&pil)
	for pil != 4 {
		// Pilihan Kelola Pabrik
		if pil == 1 {
			
			tampilPabrik(data, n)
			menuPabrik(&pil)
	
			for pil != 6 {
				// Tambah Pabrik
				if pil == 1 {
					tambahPabrik(&data, &n)
					tampilPabrik(data, n)
				// Edit Pabrik
				}else if pil == 2 {
					editPabrik(&data, n)
					tampilPabrik(data, n)
				// Hapus Pabrik 
				}else if pil == 3 {
					hapusPabrik(&data, &n)
					tampilPabrik(data, n)
				// Urut Pabrik 
				}else if pil == 4 {
					urutPabrik(data, n, "jumlah")
				// Urut Pabrik 
				}else if pil == 5 {
					urutPabrik(data, n, "penjualan")
				}
	
				menuPabrik(&pil)
			}
		
		}else if pil == 2 {
			setAllMobil(&allMobil, &nMobil, data, n)
			tampilMobil(allMobil, nMobil)
			menuMobil(&pil)
	
			for pil != 7 {
				// Tambah Mobil
				if pil == 1 {
					tambahMobil(&data, &n)
				// Edit Mobil
				}else if pil == 2 {
					editMobil(&data, allMobil, n)
				// Hapus Mobil 
				}else if pil == 3 {
					hapusMobil(&data, &n, allMobil)
				// Cari Mobil
				}else if pil == 4 {
					menuCariMobil(&pil)
					for pil != 4 {
						cariMobilQuery(&pil, allMobil, nMobil)
						menuCariMobil(&pil)
					}
				// Urut Mobil
				}else if pil == 5 {
					menuUrutMobil(&pil)
					for pil != 4 {
						// urutMobil(allMobil, nMobil, pil)
						menuUrutMobil(&pil)
					}
				}else if pil == 6 {
					for pil != 4 {
						// urutMobil(allMobil, nMobil, pil)
						menuUrutMobil(&pil)
					}
				}
		
				setAllMobil(&allMobil, &nMobil, data, n)
				tampilMobil(allMobil, nMobil)
				menuMobil(&pil)
			}
		}

		// Menampilkan menu lagi
		menu(&pil)
	}
}

func clear(){
	cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func header(){
	fmt.Println("=================================")
	fmt.Println("   Aplikasi Dealer Mobil Josss   ")
	fmt.Println("=================================")
}

func menu(pil *int){
	clear()
	header()
	fmt.Println("1. Kelola Data Pabrikan")
	fmt.Println("2. Kelola Data Mobil")
	fmt.Println("3. Lihat Data Penjualan")
	fmt.Println("4. Exit")
	fmt.Println("=================================")

	fmt.Print("Masukkan Pilihan [1-4] : ")
	fmt.Scan(&*pil)

	for *pil > 4 || *pil == 0{
		fmt.Println("Pilihan Tidak Tersedia!!!")
		fmt.Print("Masukkan Pilihan [1-4] : ")
		fmt.Scan(&*pil)
	}
}


// ============KELOLA PABRIK ==================
func menuPabrik(pil *int){
	fmt.Println("=================================")
	fmt.Println("1. Tambah Data (V)")
	fmt.Println("2. Edit Data")
	fmt.Println("3. Hapus Data")
	fmt.Println("4. Urutkan Berdasarkan Jumlah Mobil")
	fmt.Println("5. 3 Penjulan Tertinggi")
	fmt.Println("6. Kembali")
	fmt.Println("=================================")

	fmt.Print("Masukkan Pilihan [1-6] : ")
	fmt.Scan(&*pil)

	for *pil > 6 {
		fmt.Println("Pilihan Tidak Tersedia!!!")
		fmt.Print("Masukkan Pilihan [1-6] : ")
		fmt.Scan(&*pil)
	}
}

func tampilPabrik(data arrPabrik, n int){
	var row Pabrik
	var lebar[4] int
	var i, penjualan int
	var title = [4]string{"No","Merek", "Total Mobil", "Total Penjualan"}

	lebar[0] = len(title[0])
	lebar[1] = len(title[1])
	lebar[2] = len(title[2])
	lebar[3] = len(title[3])

	for i = 0; i < n; i++ {
		row = data[i]
		if len(strconv.Itoa(i+1)) > lebar[0] {
			lebar[0] = len(strconv.Itoa(i+1))
		}

		if len(row.merek) > lebar[1] {
			lebar[1] = len(row.merek)
		}

		if len(strconv.Itoa(row.nMobil)) > lebar[2] {
			lebar[2] = len(strconv.Itoa(row.nMobil))
		}

		penjualan = totalPenjualanPabrik(data, i)

		if len(strconv.Itoa(penjualan)) > lebar[3] {
			lebar[3] = len(strconv.Itoa(penjualan))
		}

	}

	var pemisah string
	for i = 0; i < 4; i++ {
		pemisah += "+"
		for j := 0; j < lebar[i]+2; j++ {
			pemisah += "-"
		}
	}
	pemisah += "+"

	clear()
	header()

	fmt.Println("           Data Pabrikan        ")
	fmt.Println("=================================")
	
	if pesan != "" {
		fmt.Println(pesan)
		pesan = ""
	}

	fmt.Println(pemisah)
	fmt.Printf("| %-*s ", lebar[0], title[0])
	fmt.Printf("| %-*s ", lebar[1], title[1])
	fmt.Printf("| %-*s ", lebar[2], title[2])
	fmt.Printf("| %-*s |\n", lebar[3], title[3])
	
	i=0
	for i < n && n != 0{
		fmt.Println(pemisah)
		row = data[i]
		penjualan = totalPenjualanPabrik(data, i)
		fmt.Printf("| %-*d ", lebar[0], i+1)
		fmt.Printf("| %-*s ", lebar[1], row.merek)
		fmt.Printf("| %-*d ", lebar[2], row.nMobil)
		fmt.Printf("| %-*d |\n", lebar[3], penjualan)
		i++
	}

	if n == 0 {
		var lebarTotal = (lebar[0] + lebar[1] + lebar[2] + lebar[3] + 9)
		fmt.Println(pemisah)
		fmt.Printf("| %-*s |\n", lebarTotal, "Data Masih kosong !!!")
	}

	fmt.Println(pemisah)
}

func tambahPabrik(data *arrPabrik, n *int){
	// var confirm, merek string
	var confirm string

	clear()
	header()

	fmt.Println("      Tambah Data Pabrikan       ")
	fmt.Println("=================================")

	// fmt.Print("\nMasukkan Merek Pabrik : ")
	// fmt.Scan(&merek)

	fmt.Print("\nSimpan Data Ini? [Y/N] : ")
	fmt.Scan(&confirm)

	if confirm == "Y" || confirm == "y" {
		// data[*n].merek = merek
		// *n++

		data[0].merek = "Toyota"
		data[1].merek = "Daihatsu"
		data[2].merek = "Hyndai"
		data[3].merek = "Mitsubishi"
		*n+=4
	}
	
	pesan = "Data Pabrik Baru Berhasil Ditambahkan..."
}

func editPabrik(data *arrPabrik, n int,){
	var confirm, merek string
	var no int

	fmt.Print("Pilih No Data yang Akan Diedit : ")
	fmt.Scan(&no)

	for no > n || no == 0 {
		fmt.Println("Data Tidak Tersedia !!!")

		fmt.Print("Pilih No Data yang Akan Diedit : ")
		fmt.Scan(&no)
	}

	clear()
	header()

	fmt.Println("        Edit Data Pabrikan       ")
	fmt.Println("=================================")

	fmt.Print("\nMasukkan Merek Pabrik : ")
	fmt.Scan(&merek)

	fmt.Print("\nEdit Data Ini? [Y/N] : ")
	fmt.Scan(&confirm)

	if confirm == "Y" || confirm == "y" {
		data[no-1].merek = merek
	}
	
	pesan = "Data Pabrik Berhasil Diedit..."
}

func hapusPabrik(data *arrPabrik, n *int){
	var no, i int

	fmt.Print("Pilih No Data yang Akan Dihapus : ")
	fmt.Scan(&no)

	for no > *n || no == 0 {
		fmt.Println("Data Tidak Tersedia !!!")

		fmt.Print("Pilih No Data yang Akan Dihapus : ")
		fmt.Scan(&no)
	}
	
	i = no-1
	for i < *n-1 {
		data[i] = data[i+1]
		i++
	}
	*n--

	pesan = "Data Pabrik Berhasil Dihapus..."
}

func totalPenjualanPabrik(data arrPabrik, idx int) int {
	var i, total int
	for i < data[idx].nMobil {
		total += data[idx].daftarMobil[i].penjualan
		i++
	}

	return total
}

func cariPabrik(data arrPabrik, n int, x string) int {
	var i int = 0
	var found bool = false
	for i < n && !found{
		found = data[i].merek == x
		i++
	}

	return i-1
}

func urutPabrik(T arrPabrik, n int, flag string) {
	var i int
	var pass, urut int
	var temp Pabrik
	var data arrPabrik
	var g string

	flag = "dsd"

	if flag == "jumlah" {
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")
		fmt.Print("Pilih : ")
		fmt.Scan(&urut)
		
		// Insertion Sort
		pass = 1
		for pass <= n-1 {
			i = pass
			temp = T[i]
			
			for i > 0 && ((urut == 1 && temp.nMobil < T[i-1].nMobil) || (urut == 2 && temp.nMobil > T[i-1].nMobil)) {
				T[i] = T[i-1]
				i--
			}
			T[i] = temp
			pass++
		}

		data = T
	}else{	
		// Insertion Sort
		var penjualanT, penjualanTemp int
		pass = 1
		for pass <= n-1 {
			penjualanTemp = 0
			i = pass
			temp = T[i]
			var j int = 0
			for j < temp.nMobil {
				penjualanTemp += temp.daftarMobil[j].penjualan
				j++
			}

			penjualanT = totalPenjualanPabrik(T, i-1)

			fmt.Println(penjualanTemp, penjualanT)
			fmt.Scan(&g)
			for i > 0 && penjualanTemp > penjualanT {
				T[i] = T[i-1]
				i--
			}
			T[i] = temp
			pass++
		}

		for i=0; i<n; i++{
			data[i] = T[i]
		}
	}

	tampilPabrik(data, n)
}
// ========== END PABRIK =============================


// ============KELOLA MOBIL ==================
func menuMobil(pil *int){
	fmt.Println("=================================")
	fmt.Println("1. Tambah Data")
	fmt.Println("2. Edit Data")
	fmt.Println("3. Hapus Data")
	fmt.Println("4. Cari Data Mobil")
	fmt.Println("5. Urutkan Data Mobil")
	fmt.Println("6. 3 Penjualan Tertinggi")
	fmt.Println("7. Kembali")
	fmt.Println("=================================")

	fmt.Print("Masukkan Pilihan [1-7] : ")
	fmt.Scan(&*pil)

	for *pil > 7 {
		fmt.Println("Pilihan Tidak Tersedia!!!")
		fmt.Print("Masukkan Pilihan [1-7] : ")
		fmt.Scan(&*pil)
	}

}

func menuCariMobil(pil *int) {
	fmt.Println("=================================")
	fmt.Println("      Cari Mobil Berdasarkan     ")
	fmt.Println("=================================")
	fmt.Println("1. Berdasarkan Merek Pabrik")
	fmt.Println("2. Berdasarkan Nama Mobil")
	fmt.Println("3. Berdasarkan Tahun Keluaran")
	fmt.Println("4. Kembali")
	fmt.Println("=================================")

	fmt.Print("Masukkan Pilihan [1-4] : ")
	fmt.Scan(&*pil)

	for *pil > 4 {
		fmt.Println("Pilihan Tidak Tersedia!!!")
		fmt.Print("Masukkan Pilihan [1-4] : ")
		fmt.Scan(&*pil)
	}
}

func setAllMobil(mobil *arrMobil, nMobil *int, pabrik arrPabrik, nPabrik int) {
	var i, j int
	*nMobil = 0
	for i=0; i<nPabrik; i++{
		for j=0; j<pabrik[i].nMobil; j++ {
			mobil[*nMobil].mobil = pabrik[i].daftarMobil[j]
			mobil[*nMobil].pabrik = pabrik[i].merek
			*nMobil++
		}
	}
}

func tampilMobil(mobil arrMobil, nMobil int){
	var row AllMobil
	var lebar[7] int
	var i int
	var title = [7]string{"No","Nama", "Pabrik", "Tahun Keluar", "Harga", "Penjualan", "Jumlah"}

	lebar[0] = len(title[0])
	lebar[1] = len(title[1])
	lebar[2] = len(title[2])
	lebar[3] = len(title[3])
	lebar[4] = len(title[4])
	lebar[5] = len(title[5])
	lebar[6] = len(title[6])

	for i = 0; i < nMobil; i++ {
		row = mobil[i]
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

	var pemisah string
	for i = 0; i < 7; i++ {
		pemisah += "+"
		for j := 0; j < lebar[i]+2; j++ {
			pemisah += "-"
		}
	}
	pemisah += "+"

	clear()
	header()

	fmt.Println("           Data Mobil            ")
	fmt.Println("=================================")
	
	if pesan != "" {
		fmt.Println(pesan)
		pesan = ""
	}

	fmt.Println(pemisah)
	fmt.Printf("| %-*s ", lebar[0], title[0])
	fmt.Printf("| %-*s ", lebar[1], title[1])
	fmt.Printf("| %-*s ", lebar[2], title[2])
	fmt.Printf("| %-*s ", lebar[3], title[3])
	fmt.Printf("| %-*s ", lebar[4], title[4])
	fmt.Printf("| %-*s ", lebar[5], title[5])
	fmt.Printf("| %-*s |\n", lebar[6], title[6])
	
	i=0
	for i < nMobil && nMobil != 0{
		fmt.Println(pemisah)
		row = mobil[i]
		fmt.Printf("| %-*d ", lebar[0], i+1)
		fmt.Printf("| %-*s ", lebar[1], row.mobil.nama)
		fmt.Printf("| %-*s ", lebar[2], row.pabrik)
		fmt.Printf("| %-*d ", lebar[3], row.mobil.tahunKeluar)
		fmt.Printf("| %-*d ", lebar[4], row.mobil.harga)
		fmt.Printf("| %-*d ", lebar[5], row.mobil.penjualan)
		fmt.Printf("| %-*d |\n", lebar[6], row.mobil.jumlah)
		i++
	}

	if nMobil == 0 {
		var lebarTotal = (lebar[0] + lebar[1] + lebar[2] + lebar[3] + lebar[4] + lebar[5] + lebar[6] + 14)
		fmt.Println(pemisah)
		fmt.Printf("| %-*s |\n", lebarTotal, "Data Masih kosong !!!")
	}

	fmt.Println(pemisah)
}

func tambahMobil(pabrik *arrPabrik, nPabrik *int){
	// var confirm, nama string
	// var i, no, tahunKeluar, harga, penjualan, jumlah int

	var confirm string

	// fmt.Println("      Pilihan Pabrik Mobil       ")
	// fmt.Println("=================================")
	// for i=0; i < *nPabrik; i++ {
	// 	fmt.Printf("%d. %s\n", i+1, pabrik[i].merek)
	// }

	// fmt.Print("Pilih No Pabrik : ")
	// fmt.Scan(&no)

	clear()
	header()

	fmt.Println("        Tambah Data Mobil        ")
	fmt.Println("=================================")

	// fmt.Print("\nNama Mobil : ")
	// fmt.Scan(&nama)
	// fmt.Print("Tahun Keluar : ")
	// fmt.Scan(&tahunKeluar)
	// fmt.Print("Harga : ")
	// fmt.Scan(&harga)
	// fmt.Print("Total Penjualan (unit) : ")
	// fmt.Scan(&penjualan)
	// fmt.Print("Jumlah Mobil (Unit) : ")
	// fmt.Scan(&jumlah)

	fmt.Print("\nSimpan Data Ini? [Y/N] : ")
	fmt.Scan(&confirm)

	// var nMobil int = pabrik[no-1].nMobil
	if confirm == "Y" || confirm == "y" {
		// pabrik[no-1].daftarMobil[nMobil].nama = nama
		// pabrik[no-1].daftarMobil[nMobil].tahunKeluar = tahunKeluar
		// pabrik[no-1].daftarMobil[nMobil].harga = harga
		// pabrik[no-1].daftarMobil[nMobil].penjualan = penjualan
		// pabrik[no-1].daftarMobil[nMobil].jumlah = jumlah
		// pabrik[no-1].nMobil++

		pabrik[0].daftarMobil[0].nama = "Avanza"
		pabrik[0].daftarMobil[0].tahunKeluar = 2019
		pabrik[0].daftarMobil[0].harga = 100000000
		pabrik[0].daftarMobil[0].penjualan = 12
		pabrik[0].daftarMobil[0].jumlah = 5
		pabrik[0].nMobil++

		pabrik[0].daftarMobil[1].nama = "Rush"
		pabrik[0].daftarMobil[1].tahunKeluar = 2023
		pabrik[0].daftarMobil[1].harga = 150000000
		pabrik[0].daftarMobil[1].penjualan = 2
		pabrik[0].daftarMobil[1].jumlah = 2
		pabrik[0].nMobil++

		pabrik[1].daftarMobil[0].nama = "Sirion"
		pabrik[1].daftarMobil[0].tahunKeluar = 2023
		pabrik[1].daftarMobil[0].harga = 228000000
		pabrik[1].daftarMobil[0].penjualan = 1
		pabrik[1].daftarMobil[0].jumlah = 3
		pabrik[1].nMobil++

		pabrik[2].daftarMobil[0].nama = "Stargazer"
		pabrik[2].daftarMobil[0].tahunKeluar = 2022
		pabrik[2].daftarMobil[0].harga = 307000000
		pabrik[2].daftarMobil[0].penjualan = 5
		pabrik[2].daftarMobil[0].jumlah = 1
		pabrik[2].nMobil++

		pabrik[3].daftarMobil[0].nama = "XPander"
		pabrik[3].daftarMobil[0].tahunKeluar = 2021
		pabrik[3].daftarMobil[0].harga = 250000000
		pabrik[3].daftarMobil[0].penjualan = 0
		pabrik[3].daftarMobil[0].jumlah = 4
		pabrik[3].nMobil++
	}
	
	pesan = "Data Mobil Baru Berhasil Ditambahkan..."
}

func cariMobilPabrik(data arrPabrik, idx int, x string) int {
	var i int = 0
	var found bool = false
	for i < data[idx].nMobil && !found{
		found = data[idx].daftarMobil[i].nama == x
		i++
	}

	return i-1
}

func cariMobilQuery(pil *int, mobil arrMobil, nMobil int) {
	var x string
	var i, n, thn int
	var data arrMobil

	if *pil == 1{
		fmt.Print("Masukkan Merek Pabrik : ")
		fmt.Scan(&x)
		i = 0
		for i < nMobil {
			if (mobil[i].pabrik == x) {
				data[n] = mobil[i]
				n++
			}
			i++
		}

		tampilMobil(data, n)
		
	}else if *pil == 2{
		fmt.Print("Masukkan Nama Mobil : ")
		fmt.Scan(&x)
		i = 0
		for i < nMobil {
			if (mobil[i].mobil.nama == x) {
				data[n] = mobil[i]
				n++
			}
			i++
		}

		tampilMobil(data, n)
	}else if *pil == 3{
		fmt.Print("Masukkan Tahun Keluaran : ")
		fmt.Scan(&thn)
		i = 0
		for i < nMobil {
			if (mobil[i].mobil.tahunKeluar == thn) {
				data[n] = mobil[i]
				n++
			}
			i++
		}

		tampilMobil(data, n)
	}
}

func editMobil(data *arrPabrik, mobil arrMobil, n int){
	var confirm, nama string
	var idxP, idxM, no, tahunKeluar, harga, penjualan, jumlah int

	fmt.Print("Pilih No Data yang Akan Diedit : ")
	fmt.Scan(&no)

	for no > n || no == 0 {
		fmt.Println("Data Tidak Tersedia !!!")

		fmt.Print("Pilih No Data yang Akan Diedit : ")
		fmt.Scan(&no)
	}

	idxP = cariPabrik(*data, n, mobil[no-1].pabrik)
	idxM = cariMobilPabrik(*data, idxP, mobil[no-1].mobil.nama)

	clear()
	header()
	fmt.Println("         Edit Data Mobil        ")
	fmt.Println("=================================")

	fmt.Println(idxP, idxM)

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

	fmt.Print("\nEdit Data Ini? [Y/N] : ")
	fmt.Scan(&confirm)

	if confirm == "Y" || confirm == "y" {
		data[idxP].daftarMobil[idxM].nama = nama
		data[idxP].daftarMobil[idxM].tahunKeluar = tahunKeluar
		data[idxP].daftarMobil[idxM].harga = harga
		data[idxP].daftarMobil[idxM].penjualan = penjualan
		data[idxP].daftarMobil[idxM].jumlah = jumlah
	}
	
	pesan = "Data Mobil Berhasil Diedit..."
}

func hapusMobil(data *arrPabrik, n *int, mobil arrMobil){
	var no, idxP, idxM int

	fmt.Print("Pilih No Data yang Akan Dihapus : ")
	fmt.Scan(&no)

	for no > *n || no == 0 {
		fmt.Println("Data Tidak Tersedia !!!")

		fmt.Print("Pilih No Data yang Akan Dihapus : ")
		fmt.Scan(&no)
	}

	idxP = cariPabrik(*data, *n, mobil[no-1].pabrik)
	idxM = cariMobilPabrik(*data, idxP, mobil[no-1].mobil.nama)
	
	for idxM < (data[idxP].nMobil - 1) {
		data[idxP].daftarMobil[idxM] = data[idxP].daftarMobil[idxM+1] 
		idxM++
	}
	data[idxP].nMobil--

	pesan = "Data Pabrik Berhasil Dihapus..."
}

func menuUrutMobil(pil *int) {
	fmt.Println("=================================")
	fmt.Println("      Urut Mobil Berdasarkan     ")
	fmt.Println("=================================")
	fmt.Println("1. Berdasarkan Merek Pabrik")
	fmt.Println("2. Berdasarkan Nama Mobil")
	fmt.Println("3. Berdasarkan Tahun Keluaran")
	fmt.Println("4. Kembali")
	fmt.Println("=================================")

	fmt.Print("Masukkan Pilihan [1-4] : ")
	fmt.Scan(&*pil)

	for *pil > 4 {
		fmt.Println("Pilihan Tidak Tersedia!!!")
		fmt.Print("Masukkan Pilihan [1-4] : ")
		fmt.Scan(&*pil)
	}
}

func urutMobil(T arrMobil, n, pil int, flag string) {
	var i, idx int
	var pass, urut int
	var temp AllMobil

	if flag == "kategori" {
		fmt.Println("1. Ascending")
		fmt.Println("2. Descending")
		fmt.Print("Pilih : ")
		fmt.Scan(&urut)
		
		if pil == 1{
			// Insertion Sort
			pass = 1
			for pass <= n-1 {
				i = pass
				temp = T[i]
				
				for i > 0 && ((urut == 1 && temp.pabrik < T[i-1].pabrik) || (urut == 2 && temp.pabrik > T[i-1].pabrik)) {
					T[i] = T[i-1]
					i--
				}
				T[i] = temp
				pass++
			}
		}else if pil == 2{
			// Insertion Sort
			pass = 1
			for pass <= n-1 {
				i = pass
				temp = T[i]
				
				for i > 0 && ((urut == 1 && temp.mobil.nama < T[i-1].mobil.nama) || (urut == 2 && temp.mobil.nama > T[i-1].mobil.nama)) {
					T[i] = T[i-1]
					i--
				}
				T[i] = temp
				pass++
			}
		} else if pil == 3 {
			// Selection Short
			pass = 1
			for pass <= n-1 {
				idx = pass-1
				i = pass
				for i < n {
					if (urut == 1 && T[idx].mobil.tahunKeluar > T[i].mobil.tahunKeluar) || (urut == 2 && T[idx].mobil.tahunKeluar < T[i].mobil.tahunKeluar) {
						idx = i
					}
					i++
				}
				temp = T[pass-1]
				T[pass-1] = T[idx]
				T[idx] = temp
				pass++
			}
		}
	
	}else{
		pass = 1
		for pass <= n-1 {
			i = pass
			temp = T[i]
			
			for i > 0 && ((urut == 1 && temp.mobil.nama < T[i-1].mobil.nama) || (urut == 2 && temp.mobil.nama > T[i-1].mobil.nama)) {
				T[i] = T[i-1]
				i--
			}
			T[i] = temp
			pass++
		}
	}

	

	tampilMobil(T, n)
}
// ========== END PABRIK =============================

// func edit(t *arrPabrik, nP int, jenis string) }{
// 	if jenis == "mobil" {
// 		header()
// 		pertanyaanMobil()
// 	}
// }

// func pertanyaanMobil


// text mobil/pabrik:
// 1. tambah
// 2. edit
// 3. hapus

// menu mobi:
// text()


// menu pabrik:
// text()