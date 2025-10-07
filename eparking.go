package main

import (
	"fmt"
	"strings"
)

// Struktur Data
type Petugas struct {
	ID       int
	Nama     string
	Username string
	Password string
}

type Transaksi struct {
	NoPol         string
	Jenis         string
	TanggalMasuk  string // format: "DD-MM-YYYY"
	TanggalKeluar string // format: "DD-MM-YYYY"
	Masuk         string
	Keluar        string
	Biaya         int
}

// Array Statis
const MAX_DATA = 100

var DataPetugas [MAX_DATA]Petugas
var DataTransaksi [MAX_DATA]Transaksi
var JumlahPetugas, JumlahTransaksi int

// Fungsi Utama
func main() {
	fmt.Println("====================================")
	fmt.Println("        Aplikasi E-Parking")
	fmt.Println("====================================")
	fmt.Println("1. Login sebagai Admin")
	fmt.Println("2. Login sebagai Petugas")
	fmt.Println("3. Keluar")
	fmt.Println("====================================")
	fmt.Print("Pilih menu: ")
	var pilihan int
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		LoginAdmin()
	case 2:
		LoginPetugas()
	case 3:
		fmt.Println("Terima kasih!")
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

// Login Admin
func LoginAdmin() {
	fmt.Println("Masukkan Username Admin: ")
	var username, password string
	fmt.Scan(&username)
	fmt.Println("Masukkan Password: ")
	fmt.Scan(&password)

	if username == "admin" && password == "admin" {
		fmt.Println("Login Admin Berhasil!")
		MenuAdmin()
	} else {
		fmt.Println("Username atau Password salah!")
	}
}

// Login Petugas
func LoginPetugas() {
	fmt.Println("Masukkan Username Petugas: ")
	var username, password string
	fmt.Scan(&username)
	fmt.Println("Masukkan Password: ")
	fmt.Scan(&password)

	// Sequential Search
	for i := 0; i < JumlahPetugas; i++ {
		if DataPetugas[i].Username == username && DataPetugas[i].Password == password {
			fmt.Println("Login Petugas Berhasil!")
			MenuPetugas()
			return
		}
	}
	fmt.Println("Username atau Password salah!")
}

// Menu Admin
func MenuAdmin() {
	for {
		fmt.Println("====================================")
		fmt.Println("          Menu Admin")
		fmt.Println("====================================")
		fmt.Println("1. Tambah Data Petugas")
		fmt.Println("2. Ubah Data Petugas")
		fmt.Println("3. Hapus Data Petugas")
		fmt.Println("4. Tampilkan Data Petugas")
		fmt.Println("5. Logout")
		fmt.Println("====================================")
		fmt.Print("Pilih menu: ")
		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			TambahPetugas()
		case 2:
			UbahPetugas()
		case 3:
			HapusPetugas()
		case 4:
			TampilkanPetugas()
		case 5:
			main()
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Menu Petugas
func MenuPetugas() {
	for {
		fmt.Println("====================================")
		fmt.Println("          Menu Petugas")
		fmt.Println("====================================")
		fmt.Println("1. Tambah Transaksi Parkir")
		fmt.Println("2. Ubah Transaksi Parkir")
		fmt.Println("3. Hapus Transaksi Parkir")
		fmt.Println("4. Tampilkan Transaksi Parkir")
		fmt.Println("5. Pencarian Kendaraan")
		fmt.Println("6. Logout")
		fmt.Println("====================================")
		fmt.Print("Pilih menu: ")
		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			TambahTransaksi()
		case 2:
			UbahTransaksi()
		case 3:
			HapusTransaksi()
		case 4:
			TampilkanTransaksi()
		case 5:
			PencarianKendaraan()
		case 6:
			main()
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Fungsi Tambah Petugas
func TambahPetugas() {
	if JumlahPetugas >= MAX_DATA {
		fmt.Println("Data petugas penuh!")
		return
	}
	var p Petugas
	p.ID = JumlahPetugas + 1
	fmt.Print("Nama: ")
	fmt.Scan(&p.Nama)
	fmt.Print("Username: ")
	fmt.Scan(&p.Username)
	fmt.Print("Password: ")
	fmt.Scan(&p.Password)
	DataPetugas[JumlahPetugas] = p
	JumlahPetugas++
	fmt.Println("Petugas berhasil ditambahkan!")
}

// Fungsi Ubah Petugas
func UbahPetugas() {
	fmt.Print("Masukkan ID Petugas yang akan diubah: ")
	var id int
	fmt.Scan(&id)
	if id <= 0 || id > JumlahPetugas {
		fmt.Println("ID tidak ditemukan!")
		return
	}
	fmt.Print("Nama Baru: ")
	fmt.Scan(&DataPetugas[id-1].Nama)
	fmt.Print("Username Baru: ")
	fmt.Scan(&DataPetugas[id-1].Username)
	fmt.Print("Password Baru: ")
	fmt.Scan(&DataPetugas[id-1].Password)
	fmt.Println("Data petugas berhasil diubah!")
}

// Fungsi Hapus Petugas
func HapusPetugas() {
	fmt.Print("Masukkan ID Petugas yang akan dihapus: ")
	var id int
	fmt.Scan(&id)

	for i := 0; i < JumlahPetugas; i++ {
		if DataPetugas[i].ID == id {
			// Menghapus petugas dengan cara menggantinya dengan data petugas terakhir
			DataPetugas[i] = DataPetugas[JumlahPetugas-1]
			JumlahPetugas--
			fmt.Println("Data petugas berhasil dihapus!")
			return
		}
	}
	fmt.Println("Petugas tidak ditemukan!")
}

// Fungsi Tampilkan Petugas
func TampilkanPetugas() {
	if JumlahPetugas == 0 {
		fmt.Println("Data petugas kosong!")
		return
	}

	// Selection Sort berdasarkan Nama
	for i := 0; i < JumlahPetugas-1; i++ {
		minIdx := i
		for j := i + 1; j < JumlahPetugas; j++ {
			if strings.ToLower(DataPetugas[j].Nama) < strings.ToLower(DataPetugas[minIdx].Nama) {
				minIdx = j
			}
		}
		// Tukar posisi
		DataPetugas[i], DataPetugas[minIdx] = DataPetugas[minIdx], DataPetugas[i]
	}

	fmt.Println("Daftar Petugas (diurutkan berdasarkan Nama):")
	for i := 0; i < JumlahPetugas; i++ {
		fmt.Printf("ID: %d, Nama: %s, Username: %s\n", DataPetugas[i].ID, DataPetugas[i].Nama, DataPetugas[i].Username)
	}
}

// Fungsi Pencarian Kendaraan
func PencarianKendaraan() {
	if JumlahTransaksi == 0 {
		fmt.Println("Data transaksi kosong!")
		return
	}

	// Selection Sort berdasarkan NoPol
	for i := 0; i < JumlahTransaksi-1; i++ {
		minIdx := i
		for j := i + 1; j < JumlahTransaksi; j++ {
			if DataTransaksi[j].NoPol < DataTransaksi[minIdx].NoPol {
				minIdx = j
			}
		}
		DataTransaksi[i], DataTransaksi[minIdx] = DataTransaksi[minIdx], DataTransaksi[i]
	}

	fmt.Print("Masukkan nomor polisi kendaraan: ")
	var noPol string
	fmt.Scan(&noPol)

	// Binary Search
	low, high := 0, JumlahTransaksi-1
	for low <= high {
		mid := (low + high) / 2
		if DataTransaksi[mid].NoPol == noPol {
			fmt.Printf("Kendaraan %s ditemukan!\n", noPol)
			fmt.Printf("Jenis: %s, Masuk: %s, Keluar: %s, Biaya: %d\n",
				DataTransaksi[mid].Jenis,
				DataTransaksi[mid].Masuk,
				DataTransaksi[mid].Keluar,
				DataTransaksi[mid].Biaya)
			return
		} else if noPol < DataTransaksi[mid].NoPol {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("Kendaraan tidak ditemukan!")
}

// Fungsi Hapus Transaksi Parkir
func HapusTransaksi() {
	fmt.Print("Masukkan nomor polisi kendaraan yang ingin dihapus: ")
	var noPol string
	fmt.Scan(&noPol)

	for i := 0; i < JumlahTransaksi; i++ {
		if DataTransaksi[i].NoPol == noPol {
			// Menghapus transaksi dengan cara menggantinya dengan data transaksi terakhir
			DataTransaksi[i] = DataTransaksi[JumlahTransaksi-1]
			JumlahTransaksi--
			fmt.Println("Transaksi parkir berhasil dihapus!")
			return
		}
	}
	fmt.Println("Transaksi parkir tidak ditemukan!")
}

// Fungsi Tampilkan Transaksi
func TampilkanTransaksi() {
	if JumlahTransaksi == 0 {
		fmt.Println("Data transaksi kosong!")
		return
	}

	// Insertion Sort berdasarkan NoPol
	for i := 1; i < JumlahTransaksi; i++ {
		key := DataTransaksi[i]
		j := i - 1
		for j >= 0 && DataTransaksi[j].NoPol > key.NoPol {
			DataTransaksi[j+1] = DataTransaksi[j]
			j--
		}
		DataTransaksi[j+1] = key
	}

	fmt.Println("===== Daftar Transaksi Parkir (urut NoPol) =====")
	for i := 0; i < JumlahTransaksi; i++ {
		fmt.Printf("No. Polisi     : %s\n", DataTransaksi[i].NoPol)
		fmt.Printf("Jenis Kendaraan: %s\n", DataTransaksi[i].Jenis)
		fmt.Printf("Tanggal Masuk  : %s\n", DataTransaksi[i].TanggalMasuk)
		fmt.Printf("Jam Masuk      : %s\n", DataTransaksi[i].Masuk)
		fmt.Printf("Tanggal Keluar : %s\n", DataTransaksi[i].TanggalKeluar)
		fmt.Printf("Jam Keluar     : %s\n", DataTransaksi[i].Keluar)
		fmt.Printf("Biaya Parkir   : Rp %d\n", DataTransaksi[i].Biaya)
		fmt.Println("-----------------------------------")
	}
}

// Fungsi Tambah Transaksi
func TambahTransaksi() {
	if JumlahTransaksi >= len(DataTransaksi) {
		fmt.Println("Data transaksi sudah penuh!")
		return
	}

	var t Transaksi
	fmt.Print("Masukkan nomor polisi: ")
	fmt.Scan(&t.NoPol)

	fmt.Print("Masukkan jenis kendaraan (Motor/Mobil): ")
	fmt.Scan(&t.Jenis)

	fmt.Print("Masukkan tanggal masuk (DD-MM-YYYY): ")
	fmt.Scan(&t.TanggalMasuk)

	fmt.Print("Masukkan waktu masuk (HH:MM): ")
	fmt.Scan(&t.Masuk)

	fmt.Print("Masukkan tanggal keluar (DD-MM-YYYY): ")
	fmt.Scan(&t.TanggalKeluar)

	fmt.Print("Masukkan waktu keluar (HH:MM): ")
	fmt.Scan(&t.Keluar)

	// Hitung durasi jam
	jam := HitungTotalJam(t.TanggalMasuk, t.Masuk, t.TanggalKeluar, t.Keluar)

	// Hitung biaya
	switch strings.ToLower(t.Jenis) {
	case "motor":
		t.Biaya = jam * 2000
	case "mobil":
		t.Biaya = jam * 4000
	default:
		t.Biaya = 0
		fmt.Println("Jenis kendaraan tidak dikenali, biaya dianggap 0")
	}

	DataTransaksi[JumlahTransaksi] = t
	JumlahTransaksi++

	fmt.Println("Transaksi berhasil ditambahkan!")
	fmt.Printf("Durasi: %d jam\n", jam)
	fmt.Printf("Biaya parkir: Rp %d\n", t.Biaya)
}

// Fungsi Ubah Transaksi
func UbahTransaksi() {
	fmt.Print("Masukkan nomor polisi kendaraan yang ingin diubah: ")
	var noPol string
	fmt.Scan(&noPol)

	for i := 0; i < JumlahTransaksi; i++ {
		if DataTransaksi[i].NoPol == noPol {
			fmt.Println("Transaksi ditemukan!")
			fmt.Println("1. Ubah Tanggal & Waktu Keluar")
			fmt.Println("2. Ubah Biaya Manual")
			fmt.Print("Pilih yang ingin diubah: ")
			var pilihan int
			fmt.Scan(&pilihan)

			switch pilihan {
			case 1:
				// Ubah waktu keluar
				fmt.Print("Masukkan Tanggal Keluar Baru (format: DD-MM-YYYY): ")
				fmt.Scan(&DataTransaksi[i].TanggalKeluar)

				fmt.Print("Masukkan Jam Keluar Baru (format: HH:MM): ")
				fmt.Scan(&DataTransaksi[i].Keluar)

				// Hitung ulang biaya
				jam := HitungTotalJam(DataTransaksi[i].TanggalMasuk, DataTransaksi[i].Masuk, DataTransaksi[i].TanggalKeluar, DataTransaksi[i].Keluar)

				switch strings.ToLower(DataTransaksi[i].Jenis) {
				case "motor":
					DataTransaksi[i].Biaya = jam * 2000
				case "mobil":
					DataTransaksi[i].Biaya = jam * 4000
				default:
					DataTransaksi[i].Biaya = 0
				}

				fmt.Println("Waktu keluar dan biaya berhasil diperbarui!")
				fmt.Printf("Durasi parkir: %d jam\n", jam)
				fmt.Printf("Biaya baru: Rp %d\n", DataTransaksi[i].Biaya)

			case 2:
				// Ubah biaya manual
				fmt.Print("Masukkan Biaya Baru: ")
				var biaya int
				fmt.Scan(&biaya)
				DataTransaksi[i].Biaya = biaya
				fmt.Println("Biaya berhasil diubah secara manual!")

			default:
				fmt.Println("Pilihan tidak valid!")
			}
			return
		}
	}
	fmt.Println("Transaksi parkir tidak ditemukan!")
}

// Fungsi Hitung Parkir
func HitungTotalJam(tanggalMasuk, waktuMasuk, tanggalKeluar, waktuKeluar string) int {
	var tglM, blnM, thnM int
	var tglK, blnK, thnK int
	var jamM, menitM, jamK, menitK int

	// Ambil tanggal dan waktu
	fmt.Sscanf(tanggalMasuk, "%d-%d-%d", &tglM, &blnM, &thnM)
	fmt.Sscanf(waktuMasuk, "%d:%d", &jamM, &menitM)

	fmt.Sscanf(tanggalKeluar, "%d-%d-%d", &tglK, &blnK, &thnK)
	fmt.Sscanf(waktuKeluar, "%d:%d", &jamK, &menitK)

	// Konversi semua ke menit
	totalMenitMasuk := (thnM*365*24*60 + blnM*30*24*60 + tglM*24*60) + jamM*60 + menitM
	totalMenitKeluar := (thnK*365*24*60 + blnK*30*24*60 + tglK*24*60) + jamK*60 + menitK

	selisih := totalMenitKeluar - totalMenitMasuk
	if selisih < 0 {
		return 0
	}

	jam := selisih / 60
	if selisih%60 != 0 {
		jam++ // bulatkan ke atas kalau ada menit sisa
	}
	return jam
}

// Second Commit
