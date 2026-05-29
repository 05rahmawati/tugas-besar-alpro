package main
import "fmt"

type Aset struct {
	Nama       string
	Satuan     string  
	Modal      float64
	Jumlah     float64 
	HargaPasar float64
}

func tampilkanMenu() {
	fmt.Println("\n=== APLIKASI MANAJEMEN INVESTASI ===")
	fmt.Println("1. Tambah Investasi")
	fmt.Println("2. Tarik / Hapus Investasi")
	fmt.Println("3. Profil & Laporan Aset")
	fmt.Println("4. Urutkan Aset")
	fmt.Println("5. Keluar")
	fmt.Print("Pilih menu (1-5): ")
}

func cariIndeksAset(arr [3]Aset, target string) int {
	var i int
	
	for i = 0; i < len(arr); i++ {
		if arr[i].Nama == target {
			return i
		}
	}
	return -1
}

func prosesTambahInvestasi(portofolio [3]Aset) [3]Aset {
	var namaAset string
	var indeks int
	var hargaBaru float64
	var tambahModal float64
	var jumlahDidapat float64

	fmt.Print("\nMasukkan nama aset (Saham/Emas/Reksadana): ")
	fmt.Scanln(&namaAset)

	indeks = cariIndeksAset(portofolio, namaAset)

	if indeks != -1 {
		fmt.Print("Masukkan harga pasar saat ini: Rp ")
		fmt.Scanln(&hargaBaru)
		fmt.Print("Masukkan nominal investasi: Rp ")
		fmt.Scanln(&tambahModal)

		jumlahDidapat = tambahModal / hargaBaru

		portofolio[indeks].HargaPasar = hargaBaru
		portofolio[indeks].Modal = portofolio[indeks].Modal + tambahModal
		portofolio[indeks].Jumlah = portofolio[indeks].Jumlah + jumlahDidapat

		fmt.Printf("Berhasil! Anda mendapatkan %.4f %s %s.\n", jumlahDidapat, portofolio[indeks].Satuan, portofolio[indeks].Nama)
	} else {
		fmt.Println("Gagal: Aset tidak ditemukan. Pastikan huruf awal kapital.")
	}
	return portofolio
}

func prosesTarikInvestasi(portofolio [3]Aset) [3]Aset {
	var namaAset string
	var indeks int
	var tarikModal float64
	var jumlahDitarik float64

	fmt.Print("\nMasukkan nama aset yang ingin ditarik (Saham/Emas/Reksadana): ")
	fmt.Scanln(&namaAset)

	indeks = cariIndeksAset(portofolio, namaAset)

	if indeks != -1 {
		if portofolio[indeks].Modal > 0 {
			fmt.Printf("Modal %s Anda saat ini: Rp %.2f\n", portofolio[indeks].Nama, portofolio[indeks].Modal)
			fmt.Print("Masukkan nominal uang yang ingin ditarik: Rp ")
			fmt.Scanln(&tarikModal)

			if tarikModal <= portofolio[indeks].Modal {
				jumlahDitarik = tarikModal / portofolio[indeks].HargaPasar
				
				portofolio[indeks].Modal = portofolio[indeks].Modal - tarikModal
				portofolio[indeks].Jumlah = portofolio[indeks].Jumlah - jumlahDitarik

				fmt.Printf("Penarikan berhasil! %.4f %s %s telah dijual atau ditarik.\n", jumlahDitarik, portofolio[indeks].Satuan, portofolio[indeks].Nama)
			} else {
				fmt.Println("Gagal: Nominal penarikan melebihi batas modal yang Anda miliki.")
			}
		} else {
			fmt.Println("Gagal: Anda belum memiliki saldo pada investasi ini.")
		}
	} else {
		fmt.Println("Gagal: Aset tidak ditemukan. Pastikan huruf awal kapital.")
	}
	return portofolio
}

func tampilkanProfil(portofolio [3]Aset) {
	var i int
	var nilaiAset float64
	var profit float64
	var totalModal float64
	var totalNilai float64
	var totalProfit float64

	totalModal = 0
	totalNilai = 0
	totalProfit = 0

	fmt.Println("\n=== PROFIL & LAPORAN ASET ===")
	for i = 0; i < len(portofolio); i++ {
		nilaiAset = portofolio[i].Jumlah * portofolio[i].HargaPasar
		profit = nilaiAset - portofolio[i].Modal

		totalModal = totalModal + portofolio[i].Modal
		totalNilai = totalNilai + nilaiAset
		totalProfit = totalProfit + profit

		fmt.Printf("\n- %s:\n", portofolio[i].Nama)
		fmt.Printf("  Kepemilikan : %.4f %s\n", portofolio[i].Jumlah, portofolio[i].Satuan)
		fmt.Printf("  Modal       : Rp %.2f\n", portofolio[i].Modal)
		fmt.Printf("  Nilai Aset  : Rp %.2f\n", nilaiAset)
		
		if profit > 0 {
			fmt.Printf("  Profit      : +Rp %.2f\n", profit)
		} else if profit < 0 {
			fmt.Printf("  Loss        : -Rp %.2f\n", profit * -1) 
		} else {
			fmt.Printf("  Profit/Loss : Rp 0.00\n")
		}
	}

	fmt.Println("\n--- RINGKASAN KESELURUHAN ---")
	fmt.Printf("Total Modal Investasi : Rp %.2f\n", totalModal)
	fmt.Printf("Total Nilai Aset      : Rp %.2f\n", totalNilai)
	if totalProfit >= 0 {
		fmt.Printf("Total Keuntungan      : +Rp %.2f\n", totalProfit)
	} else {
		fmt.Printf("Total Kerugian        : -Rp %.2f\n", totalProfit * -1)
	}
	fmt.Println("==============================")
}

func urutkanAset(portofolio [3]Aset) [3]Aset {
	var pass int
	var idx int
	var i int
	var temp Aset

	for pass = 1; pass <= len(portofolio)-1; pass++ {
		idx = pass - 1
		for i = pass; i < len(portofolio); i++ {
			// Mencari nilai Harga Pasar paling kecil
			if portofolio[i].HargaPasar < portofolio[idx].HargaPasar {
				idx = i
			}
		}
		temp = portofolio[pass-1]
		portofolio[pass-1] = portofolio[idx]
		portofolio[idx] = temp
	}
	fmt.Println("Berhasil! Daftar aset telah diurutkan berdasarkan harga pasar termurah.")
	return portofolio
}

func main() {
	var portofolio = [3]Aset{
		{"Saham", "Lembar", 0, 0, 0},
		{"Emas", "Gram", 0, 0, 0},
		{"Reksadana", "Unit", 0, 0, 0},
	}

	var pilihan int
	var jalan bool

	jalan = true

	for jalan == true {
		tampilkanMenu()
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			fmt.Println("\n-- Menu Tambah Investasi --")
			portofolio = prosesTambahInvestasi(portofolio)
			
		} else if pilihan == 2 {
			fmt.Println("\n-- Menu Tarik / Hapus Investasi --")
			portofolio = prosesTarikInvestasi(portofolio)
			
		} else if pilihan == 3 {
			tampilkanProfil(portofolio)
			
		} else if pilihan == 4 {
			fmt.Println("\n-- Urutkan Aset (Selection Sort) --")
			portofolio = urutkanAset(portofolio)
			tampilkanProfil(portofolio)
			
		} else if pilihan == 5 {
			fmt.Println("Terima kasih telah menggunakan aplikasi!")
			jalan = false
			
		} else {
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}