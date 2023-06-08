package pabrikmodul

import "fmt"
import "main"

func tampilData(data arrPabrik, n int){

	var row Pabrik
	var lebar[3] int
	var i int
	var title = [3]string{"Merek", "Total Mobil", "Total Penjualan"}

	lebar[0] = len(title[0])
	lebar[1] = len(title[1])
	lebar[2] = len(title[2])

	for i = 0; i < n; i++ {
		row = data[i]
		if len(row.merek) > lebar[0] {
			lebar[0] = len(row.merek)
		}

		if len(strconv.Itoa(row.nMobil)) > lebar[1] {
			lebar[1] = len(strconv.Itoa(row.nMobil))
		}

		if len(strconv.Itoa(23)) > lebar[2] {
			lebar[2] = len(strconv.Itoa(23))
		}

	}

	var pemisah string
	for i = 0; i < 3; i++ {
		pemisah += "+"
		for j := 0; j < lebar[i]+2; j++ {
			pemisah += "-"
		}
	}
	pemisah += "+"

	fmt.Println(pemisah)
	fmt.Printf("| %-*s ", lebar[0], title[0])
	fmt.Printf("| %-*s ", lebar[1], title[1])
	fmt.Printf("| %-*s |\n", lebar[2], title[2])
	
	i=0
	for i < n && n != 0{
		fmt.Println(pemisah)
		row = data[i]
		fmt.Printf("| %-*s ", lebar[0], row.merek)
		fmt.Printf("| %-*d ", lebar[1], row.nMobil)
		fmt.Printf("| %-*d |\n", lebar[2], 23)
		i++
	}

	if n == 0 {
		var lebarTotal = (lebar[0] + lebar[1] + lebar[2] + 6)
		fmt.Println(pemisah)
		fmt.Printf("| %-*s |\n", lebarTotal, "Data Masih kosong !!!")
	}

	fmt.Println(pemisah)
}

func tambahData(data *arrPabrik, n *int){
	var confirm, merek string
	fmt.Println("      Tambah Data Pabrikan       ")
	fmt.Println("=================================")

	fmt.Print("\nMasukkan Merek Pabrik : ")
	fmt.Scan(&merek)
	fmt.Scan(&data[*n].merek)

	fmt.Print("\nSimpan Data Ini? [Y/N] : ")
	fmt.Scan(&confirm)

	if confirm == "Y" || confirm == "y" {
		data[*n].merek = merek
		*n++
	}
	
	main.menuPabrik()
}