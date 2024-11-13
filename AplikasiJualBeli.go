package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const NMAX = 1024

type barang struct {
	Id_Barang       int
	Nama_Barang     string
	Kategori_Barang string
	Harga_Modal     int
	Harga_Jual      int
	Stock_Barang    int
	Jumlah_Terjual  int
}

type pembelian struct {
	Kode_Transaksi        string
	Nama_Pembeli          string
	Tanggal_Pembelian     time.Time
	ID_Barang_Yang_Dibeli int
	Jumlah_Pembelian      int
	Total                 int
	Status                string
}

type array struct {
	inventori        [NMAX]barang
	jumlah_inventory int
	transaksi        [NMAX]pembelian
	jumlah_transaksi int
}

func main() {
	var data array
	addDummyData(&data)
	mainMenu(&data)
}

//Menampilkan menu utama
func mainMenu(data *array) {
	var nData int
	for {
		fmt.Println()
		fmt.Println("===============================")
		fmt.Println()
		fmt.Println("      APLIKASI JUAL BELI")
		fmt.Println()
		fmt.Println("===============================")
		fmt.Println("Pilih menu:")
		fmt.Println("1. Menu Barang")
		fmt.Println("2. Print Data Barang")
		fmt.Println("3. Urutkan Data Barang")
		fmt.Println("4. Menu Transaksi")
		fmt.Println("5. Print Data Transaksi")
		fmt.Println("6. Urutkan Data Transaksi")
		fmt.Println("7. Pencarian Barang")
		fmt.Println("8. Tampilkan Modal dan Pendapatan")
		fmt.Println("9. Tampilkan Barang Terbanyak Terjual")
		fmt.Println("10. Keluar")
		fmt.Println("===============================")
		fmt.Print("Pilih (1/2/3/4/5/6/7/8/9/10): ")
		fmt.Scan(&nData)

		switch nData {
		case 1:
			menuBarang(data)
		case 2:
			cetakDataBarang(data)
		case 3:
			sortDataBarang(data)
		case 4:
			menuTransaksi(data)
		case 5:
			cetakDataTransaksi(data)
		case 6:
			sortDataTransaksi(data)
		case 7:
			var keyword string
			fmt.Print("Masukkan kata kunci pencarian: ")
			fmt.Scan(&keyword)
			cariBarang(data, keyword)
		case 8:
			tampilkanModalPendapatan(data)
		case 9:
			tampilkanBarangTerbanyakTerjual(data)
		case 10:
			fmt.Println("Keluar. Terima kasih telah mencoba aplikasi ini.")
			return
		default:
			fmt.Println("Silahkan input pilihan yang benar.")
		}
	}
}

//Menampilkan menu Kustomisasi Data Barang
func menuBarang(data *array) {
	var pilihKustomisasi int
	for {
		fmt.Println()
		fmt.Println("*")
		fmt.Println("    Kustomisasi Data Barang")
		fmt.Println("*")
		fmt.Println("Pilih menu:")
		fmt.Println("1. Tambah data barang")
		fmt.Println("2. Edit data barang")
		fmt.Println("3. Hapus data barang")
		fmt.Println("4. Kembali ke menu utama")
		fmt.Println("*")
		fmt.Print("Pilih (1/2/3/4): ")
		fmt.Scan(&pilihKustomisasi)

		switch pilihKustomisasi {
		case 1:
			tambahDataBarang(data)
		case 2:
			editDataBarang(data)
		case 3:
			hapusDataBarang(data)
		case 4:
			return
		default:
			fmt.Println("Silahkan input pilihan yang benar.")
		}
	}
}

//Menampilkan menu Kustomisasi Data Transaksi
func menuTransaksi(data *array) {
	var pilihKustomisasi int
	for {
		fmt.Println()
		fmt.Println("*")
		fmt.Println("    Kustomisasi Data Transaksi")
		fmt.Println("*")
		fmt.Println("Pilih menu:")
		fmt.Println("1. Tambah data transaksi")
		fmt.Println("2. Edit data transaksi")
		fmt.Println("3. Update Status Transaksi")
		fmt.Println("4. Hapus data transaksi")
		fmt.Println("5. Kembali ke menu utama")
		fmt.Println("*")
		fmt.Print("Pilih (1/2/3/4/5): ")
		fmt.Scan(&pilihKustomisasi)

		switch pilihKustomisasi {
		case 1:
			tambahDataTransaksi(data)
		case 2:
			editDataTransaksi(data)
		case 3:
			updateStatusTransaksi(data)
		case 4:
			hapusDataTransaksi(data)
		case 5:
			return
		default:
			fmt.Println("Silahkan input pilihan yang benar.")
		}
	}
}

//Fungsi Tambah Data Barang
func tambahDataBarang(data *array) {
	var jumlahBarang int
	if data.jumlah_inventory < NMAX {
		fmt.Print("Masukkan jumlah barang yang ingin ditambahkan: ")
		fmt.Scan(&jumlahBarang)
		for jumlahBarang <= 0 {
			fmt.Println("Jumlah barang harus lebih besar dari 0 dan tidak boleh mengandung huruf. Silakan coba lagi.")
			fmt.Print("Masukkan jumlah barang yang ingin ditambahkan: ")
			fmt.Scan(&jumlahBarang)
		}
		for i := 0; i < jumlahBarang; i++ {
			var idBarang, hargaModal, hargaJual, jumlahBarangInput int
			var namaBarang, kategoriBarang string

			fmt.Print("Masukkan ID barang: ")
			fmt.Scan(&idBarang)
			for idBarang <= 0 {
				fmt.Println("ID barang harus lebih besar dari 0 dan tidak boleh mengandung huruf. Silakan coba lagi.")
				fmt.Print("Masukkan ID barang: ")
				fmt.Scan(&idBarang)
			}

			fmt.Print("Masukkan nama barang: ")
			fmt.Scan(&namaBarang)
			for len(namaBarang) == 0 || hasNumber(namaBarang) {
				fmt.Println("Nama barang tidak boleh kosong dan tidak boleh mengandung angka. Silakan coba lagi.")
				fmt.Print("Masukkan nama barang: ")
				fmt.Scan(&namaBarang)
			}

			fmt.Print("Masukkan kategori barang: ")
			fmt.Scan(&kategoriBarang)
			for len(kategoriBarang) == 0 || hasNumber(kategoriBarang) {
				fmt.Println("Kategori barang tidak boleh kosong dan tidak boleh mengandung angka. Silakan coba lagi.")
				fmt.Print("Masukkan kategori barang: ")
				fmt.Scan(&kategoriBarang)
			}

			fmt.Print("Masukkan harga modal: ")
			fmt.Scan(&hargaModal)
			for hargaModal <= 0 {
				fmt.Println("Harga modal harus lebih besar dari 0 dan tidak boleh mengandung huruf. Silakan coba lagi.")
				fmt.Print("Masukkan harga modal: ")
				fmt.Scan(&hargaModal)
			}

			fmt.Print("Masukkan harga jual: ")
			fmt.Scan(&hargaJual)
			for hargaJual <= 0 {
				fmt.Println("Harga jual harus lebih besar dari 0 dan tidak boleh mengandung huruf. Silakan coba lagi.")
				fmt.Print("Masukkan harga jual: ")
				fmt.Scan(&hargaJual)
			}

			fmt.Print("Masukkan jumlah barang: ")
			fmt.Scan(&jumlahBarangInput)
			for jumlahBarangInput <= 0 {
				fmt.Println("Jumlah barang harus lebih besar dari 0 dan tidak boleh mengandung huruf. Silakan coba lagi.")
				fmt.Print("Masukkan jumlah barang: ")
				fmt.Scan(&jumlahBarangInput)
			}

			data.inventori[data.jumlah_inventory] = barang{idBarang, namaBarang, kategoriBarang, hargaModal, hargaJual, jumlahBarangInput, 0}
			data.jumlah_inventory++
		}
	} else {
		fmt.Println("Kapasitas inventori sudah penuh.")
	}
}

//Fungsi Edit Data Barang
func editDataBarang(data *array) {
	var index, cari_ID int

	if data.jumlah_inventory == 0 {
		fmt.Println("Inventori masih kosong.")
		return
	}

	fmt.Print("Masukkan ID barang yang akan dicari: ")
	fmt.Scan(&cari_ID)
	index = findBarang(data, cari_ID)

	if index == -1 {
		fmt.Println("Barang tidak ditemukan.")
		return
	}

	fmt.Println("Data apa yang ingin diubah?")
	fmt.Println("1. ID")
	fmt.Println("2. Nama")
	fmt.Println("3. Kategori")
	fmt.Println("4. Harga Modal")
	fmt.Println("5. Harga Jual")
	fmt.Println("6. Jumlah")
	fmt.Print("Pilih opsi(1/2/3/4/5/6): ")
	var opsi int
	fmt.Scan(&opsi)

	switch opsi {
	case 1:
		fmt.Print("Masukkan ID barang baru: ")
		fmt.Scan(&data.inventori[index].Id_Barang)
	case 2:
		fmt.Print("Masukkan nama barang baru: ")
		fmt.Scan(&data.inventori[index].Nama_Barang)
	case 3:
		fmt.Print("Masukkan kategori barang baru: ")
		fmt.Scan(&data.inventori[index].Kategori_Barang)
	case 4:
		fmt.Print("Masukkan harga modal baru: ")
		fmt.Scan(&data.inventori[index].Harga_Modal)
	case 5:
		fmt.Print("Masukkan harga jual baru: ")
		fmt.Scan(&data.inventori[index].Harga_Jual)
	case 6:
		fmt.Print("Masukkan jumlah barang baru: ")
		fmt.Scan(&data.inventori[index].Stock_Barang)
	default:
		fmt.Println("Opsi tidak valid.")
	}
	fmt.Println("Data berhasil diubah.")
}

//Fungsi Hapus Data Barang
func hapusDataBarang(data *array) {
	var index, cari_ID int

	if data.jumlah_inventory == 0 {
		fmt.Println("Inventori masih kosong.")
		return
	}

	fmt.Print("Masukkan ID barang yang akan dihapus: ")
	fmt.Scan(&cari_ID)
	index = findBarang(data, cari_ID)

	if index == -1 {
		fmt.Println("Barang tidak ditemukan.")
		return
	}

	for i := index; i < data.jumlah_inventory-1; i++ {
		data.inventori[i] = data.inventori[i+1]
	}
	data.jumlah_inventory--
	fmt.Println("Data berhasil dihapus.")
}

//Fungsi Cari Barang
func findBarang(data *array, id int) int {
	selectionSort(data, 1, "ID")

	left, right := 0, data.jumlah_inventory-1
	for left <= right {
		mid := (left + right) / 2
		if data.inventori[mid].Id_Barang < id {
			left = mid + 1
		} else if data.inventori[mid].Id_Barang > id {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
//Fungsi Cetak Data Barang
func cetakDataBarang(data *array) {
	if data.jumlah_inventory == 0 {
		fmt.Println("Inventori masih kosong.")
		return
	}

	fmt.Println("Data Barang")
	fmt.Println("=============================================================")
	fmt.Printf("%-5s %-20s %-15s %-10s %-10s %-10s %-10s\n", "ID", "Nama", "Kategori", "Harga Modal", "Harga Jual", "Jumlah", "Terjual")
	fmt.Println("=============================================================")

	for i := 0; i < data.jumlah_inventory; i++ {
		fmt.Printf("%-5d %-20s %-15s %-10d %-10d %-10d %-10d\n", data.inventori[i].Id_Barang, data.inventori[i].Nama_Barang, data.inventori[i].Kategori_Barang, data.inventori[i].Harga_Modal, data.inventori[i].Harga_Jual, data.inventori[i].Stock_Barang, data.inventori[i].Jumlah_Terjual)
	}
	fmt.Println("=============================================================")
	fmt.Println("Total barang:", data.jumlah_inventory)
	fmt.Println("=============================================================")
}

func sortDataBarang(data *array) {
	var opsi, subOpsi int
	if data.jumlah_inventory == 0 {
		fmt.Println("Inventori masih kosong.")
		return
	}

	fmt.Println("Urutkan Data Inventori")
	fmt.Println("1. Urutkan berdasarkan ID (Selection Sort)")
	fmt.Println("2. Urutkan berdasarkan Nama (Insertion Sort)")
	fmt.Println("3. Urutkan berdasarkan Harga Modal (Insertion Sort)")
	fmt.Println("4. Urutkan berdasarkan Harga Jual (Insertion Sort)")
	fmt.Println("5. Urutkan berdasarkan Jumlah (Selection Sort)")
	fmt.Print("Pilih opsi(1/2/3/4/5): ")
	fmt.Scan(&opsi)

	fmt.Println("1. Urutkan Secara Ascending")
	fmt.Println("2. Urutkan Secara Descending")
	fmt.Print("Pilih opsi(1/2): ")
	fmt.Scan(&subOpsi)

	switch opsi {
	case 1:
		selectionSort(data, subOpsi, "ID")
	case 2:
		insertionSort(data, subOpsi, "Nama")
	case 3:
		insertionSort(data, subOpsi, "Harga Modal")
	case 4:
		insertionSort(data, subOpsi, "Harga Jual")
	case 5:
		selectionSort(data, subOpsi, "Jumlah")
	default:
		fmt.Println("Tipe pencarian tidak valid.")
	}
}

func selectionSort(data *array, order int, key string) {
	for i := 1; i <= data.jumlah_inventory-1; i++ {
		idx_min := i - 1
		for j := i; j < data.jumlah_inventory; j++ {
			var condition bool
			switch key {
			case "ID":
				condition = (order == 1 && data.inventori[j].Id_Barang < data.inventori[idx_min].Id_Barang) || (order == 2 && data.inventori[j].Id_Barang > data.inventori[idx_min].Id_Barang)
			case "Nama":
				condition = (order == 1 && strings.ToLower(data.inventori[j].Nama_Barang) < strings.ToLower(data.inventori[idx_min].Nama_Barang)) || (order == 2 && strings.ToLower(data.inventori[j].Nama_Barang) > strings.ToLower(data.inventori[idx_min].Nama_Barang))
			case "Harga Modal":
				condition = (order == 1 && data.inventori[j].Harga_Modal < data.inventori[idx_min].Harga_Modal) || (order == 2 && data.inventori[j].Harga_Modal > data.inventori[idx_min].Harga_Modal)
			case "Harga Jual":
				condition = (order == 1 && data.inventori[j].Harga_Jual < data.inventori[idx_min].Harga_Jual) || (order == 2 && data.inventori[j].Harga_Jual > data.inventori[idx_min].Harga_Jual)
			case "Jumlah":
				condition = (order == 1 && data.inventori[j].Stock_Barang < data.inventori[idx_min].Stock_Barang) || (order == 2 && data.inventori[j].Stock_Barang > data.inventori[idx_min].Stock_Barang)
			}
			if condition {
				idx_min = j
			}
		}
		t := data.inventori[idx_min]
		data.inventori[idx_min] = data.inventori[i-1]
		data.inventori[i-1] = t
	}
}

func insertionSort(data *array, order int, key string) {
	for i := 1; i <= data.jumlah_inventory-1; i++ {
		j := i
		temp := data.inventori[j]
		switch key {
		case "ID":
			for ; j > 0 && ((order == 1 && temp.Id_Barang < data.inventori[j-1].Id_Barang) || (order == 2 && temp.Id_Barang > data.inventori[j-1].Id_Barang)); j-- {
				data.inventori[j] = data.inventori[j-1]
			}
		case "Nama":
			for ; j > 0 && ((order == 1 && strings.ToLower(temp.Nama_Barang) < strings.ToLower(data.inventori[j-1].Nama_Barang)) || (order == 2 && strings.ToLower(temp.Nama_Barang) > strings.ToLower(data.inventori[j-1].Nama_Barang))); j-- {
				data.inventori[j] = data.inventori[j-1]
			}
		case "Harga Modal":
			for ; j > 0 && ((order == 1 && temp.Harga_Modal < data.inventori[j-1].Harga_Modal) || (order == 2 && temp.Harga_Modal > data.inventori[j-1].Harga_Modal)); j-- {
				data.inventori[j] = data.inventori[j-1]
			}
		case "Harga Jual":
			for ; j > 0 && ((order == 1 && temp.Harga_Jual < data.inventori[j-1].Harga_Jual) || (order == 2 && temp.Harga_Jual > data.inventori[j-1].Harga_Jual)); j-- {
				data.inventori[j] = data.inventori[j-1]
			}
		case "Jumlah":
			for ; j > 0 && ((order == 1 && temp.Stock_Barang < data.inventori[j-1].Stock_Barang) || (order == 2 && temp.Stock_Barang > data.inventori[j-1].Stock_Barang)); j-- {
				data.inventori[j] = data.inventori[j-1]
			}
		}
		data.inventori[j] = temp
	}
}

func isValidBarangID(data *array, id int) bool {
	for i := 0; i < data.jumlah_inventory; i++ {
		if data.inventori[i].Id_Barang == id {
			return true
		}
	}
	return false
}

func findBarangIndex(data *array, id int) int {
	for i := 0; i < data.jumlah_inventory; i++ {
		if data.inventori[i].Id_Barang == id {
			return i
		}
	}
	return -1
}

func tambahDataTransaksi(data *array) {
	var jumlahTransaksi int
	if data.jumlah_transaksi < NMAX {
		fmt.Print("Masukkan jumlah transaksi yang ingin ditambahkan: ")
		fmt.Scan(&jumlahTransaksi)
		for jumlahTransaksi <= 0 {
			fmt.Println("Jumlah transaksi harus lebih besar dari 0 dan tidak boleh mengandung huruf. Silakan coba lagi.")
			fmt.Print("Masukkan jumlah transaksi yang ingin ditambahkan: ")
			fmt.Scan(&jumlahTransaksi)
		}
		for i := 0; i < jumlahTransaksi; i++ {
			var namaPembeli, inputTanggal string
			var idBarangYangDibeli, jumlahPembelian, total int
			var tanggalPembelian time.Time
			kodeTransaksi := "TRX" + strconv.Itoa(data.jumlah_transaksi+1)

			fmt.Print("Masukkan nama pembeli: ")
			fmt.Scan(&namaPembeli)
			for len(namaPembeli) == 0 || hasNumber(namaPembeli) {
				fmt.Println("Nama pembeli tidak boleh kosong dan tidak boleh mengandung angka. Silakan coba lagi.")
				fmt.Print("Masukkan nama pembeli: ")
				fmt.Scan(&namaPembeli)
			}

			fmt.Println("Format tanggal: yyyy-mm-dd")
			fmt.Print("Masukkan tanggal pembelian: ")
			fmt.Scan(&inputTanggal)
			tanggalPembelian, err := time.Parse("2006-01-02", inputTanggal)
			for err != nil {
				fmt.Println("Tanggal pembelian tidak valid. Silakan coba lagi.")
				fmt.Print("Masukkan tanggal pembelian (yyyy-mm-dd): ")
				fmt.Scan(&inputTanggal)
				tanggalPembelian, err = time.Parse("2006-01-02", inputTanggal)
			}

			fmt.Print("Masukkan ID barang yang dibeli: ")
			fmt.Scan(&idBarangYangDibeli)
			for idBarangYangDibeli <= 0 || !isValidBarangID(data, idBarangYangDibeli) {
				fmt.Println("ID barang tidak valid. Silakan coba lagi.")
				fmt.Print("Masukkan ID barang yang dibeli: ")
				fmt.Scan(&idBarangYangDibeli)
			}

			barangIndex := findBarangIndex(data, idBarangYangDibeli)
			if barangIndex == -1 {
				fmt.Println("Barang tidak ditemukan. Silakan coba lagi.")
				continue
			}

			fmt.Print("Masukkan jumlah pembelian: ")
			fmt.Scan(&jumlahPembelian)
			for jumlahPembelian <= 0 || jumlahPembelian > data.inventori[barangIndex].Stock_Barang {
				fmt.Printf("Jumlah pembelian harus lebih besar dari 0 dan tidak boleh melebihi stok barang (%d). Silakan coba lagi.\n", data.inventori[barangIndex].Stock_Barang)
				fmt.Print("Masukkan jumlah pembelian: ")
				fmt.Scan(&jumlahPembelian)
			}

			total = jumlahPembelian * data.inventori[barangIndex].Harga_Jual
			data.transaksi[data.jumlah_transaksi] = pembelian{kodeTransaksi, namaPembeli, tanggalPembelian, idBarangYangDibeli, jumlahPembelian, total, "Proses"}
			data.inventori[barangIndex].Stock_Barang -= jumlahPembelian
			data.jumlah_transaksi++
		}
	} else {
		fmt.Println("Kapasitas transaksi sudah penuh.")
	}
}

func editDataTransaksi(data *array) {
	var index, opsi int
	var cari_Kode string

	if data.jumlah_transaksi == 0 {
		fmt.Println("Data transaksi masih kosong.")
		return
	}

	fmt.Println("Edit Data Transaksi")
	fmt.Print("Masukkan Kode Transaksi yang akan dicari: ")
	fmt.Scan(&cari_Kode)
	index = sequentialSearchTransaksi(data, cari_Kode)

	if index == -1 {
		fmt.Println("Data transaksi tidak ditemukan.")
		return
	}

	if data.transaksi[index].Status == "Sukses" {
		fmt.Println("Transaksi sudah sukses dan tidak bisa diubah.")
		return
	}

	fmt.Println("Data apa yang ingin diubah?")
	fmt.Println("1. Nama Pembeli")
	fmt.Println("2. Tanggal Pembelian")
	fmt.Println("3. ID Barang")
	fmt.Println("4. Jumlah Pembelian")
	fmt.Print("Pilih opsi(1/2/3/4): ")
	fmt.Scan(&opsi)

	switch opsi {
	case 1:
		fmt.Print("Masukkan Nama Pembeli baru: ")
		fmt.Scan(&data.transaksi[index].Nama_Pembeli)
	case 2:
		fmt.Println("Format tanggal: yyyy-mm-dd")
		fmt.Print("Masukkan Tanggal Pembelian baru: ")
		var inputTanggal string
		fmt.Scan(&inputTanggal)
		tanggalPembelian, err := time.Parse("2006-01-02", inputTanggal)
		for err != nil {
			fmt.Println("Tanggal pembelian tidak valid. Silakan coba lagi.")
			fmt.Print("Masukkan tanggal pembelian (yyyy-mm-dd): ")
			fmt.Scan(&inputTanggal)
			tanggalPembelian, err = time.Parse("2006-01-02", inputTanggal)
		}
		data.transaksi[index].Tanggal_Pembelian = tanggalPembelian
	case 3:
		oldBarangID := data.transaksi[index].ID_Barang_Yang_Dibeli
		oldBarangIndex := findBarangIndex(data, oldBarangID)

		fmt.Print("Masukkan ID Barang baru: ")
		var idBarangBaru int
		fmt.Scan(&idBarangBaru)
		for idBarangBaru <= 0 || !isValidBarangID(data, idBarangBaru) {
			fmt.Println("ID barang tidak valid. Silakan coba lagi.")
			fmt.Print("Masukkan ID Barang baru: ")
			fmt.Scan(&idBarangBaru)
		}
		barangIndexBaru := findBarangIndex(data, idBarangBaru)

		// Update Stock Barang Lama
		data.inventori[oldBarangIndex].Stock_Barang += data.transaksi[index].Jumlah_Pembelian
		data.inventori[oldBarangIndex].Jumlah_Terjual -= data.transaksi[index].Jumlah_Pembelian 

		// Update Stock Barang Baru
		if data.inventori[barangIndexBaru].Stock_Barang >= data.transaksi[index].Jumlah_Pembelian {
			data.inventori[barangIndexBaru].Stock_Barang -= data.transaksi[index].Jumlah_Pembelian
			data.inventori[barangIndexBaru].Jumlah_Terjual += data.transaksi[index].Jumlah_Pembelian 
			data.transaksi[index].ID_Barang_Yang_Dibeli = idBarangBaru
		} else {
			fmt.Println("Jumlah stok barang baru tidak mencukupi.")
			return
		}
	case 4:
		fmt.Print("Masukkan Jumlah Pembelian baru: ")
		var jumlahPembelianBaru int
		fmt.Scan(&jumlahPembelianBaru)
		oldJumlahPembelian := data.transaksi[index].Jumlah_Pembelian
		barangIndex := findBarangIndex(data, data.transaksi[index].ID_Barang_Yang_Dibeli)

		// Update Stock
		data.inventori[barangIndex].Stock_Barang += oldJumlahPembelian
		data.inventori[barangIndex].Jumlah_Terjual -= oldJumlahPembelian 

		for jumlahPembelianBaru <= 0 || jumlahPembelianBaru > data.inventori[barangIndex].Stock_Barang {
			fmt.Printf("Jumlah pembelian harus lebih besar dari 0 dan tidak boleh melebihi stok barang (%d). Silakan coba lagi.\n", data.inventori[barangIndex].Stock_Barang)
			fmt.Print("Masukkan Jumlah Pembelian baru: ")
			fmt.Scan(&jumlahPembelianBaru)
		}

		data.transaksi[index].Jumlah_Pembelian = jumlahPembelianBaru
		data.transaksi[index].Total = jumlahPembelianBaru * data.inventori[barangIndex].Harga_Jual

		// Update Stock
		data.inventori[barangIndex].Stock_Barang -= jumlahPembelianBaru
		data.inventori[barangIndex].Jumlah_Terjual += jumlahPembelianBaru 
	default:
		fmt.Println("Opsi tidak valid.")
	}
	fmt.Println("Data berhasil diubah.")
}

func updateStatusTransaksi(data *array) {
	var index, opsi int
	var cari_Kode string

	if data.jumlah_transaksi == 0 {
		fmt.Println("Data transaksi masih kosong.")
		return
	}

	fmt.Println("Update Status Transaksi")
	fmt.Print("Masukkan Kode Transaksi yang akan dicari: ")
	fmt.Scan(&cari_Kode)
	index = sequentialSearchTransaksi(data, cari_Kode)

	if index == -1 {
		fmt.Println("Data transaksi tidak ditemukan.")
		return
	}

	fmt.Println("Status Transaksi saat ini:", data.transaksi[index].Status)
	if data.transaksi[index].Status == "Sukses" {
		fmt.Println("Transaksi sudah sukses dan tidak bisa diubah.")
		return
	}

	fmt.Println("Apakah Anda ingin mengubah status transaksi menjadi Sukses? (1: Ya, 2: Tidak)")
	fmt.Print("Pilih opsi(1/2): ")
	fmt.Scan(&opsi)

	if opsi == 1 {
		data.transaksi[index].Status = "Sukses"
		if data.transaksi[index].Status == "Sukses" {
			oldBarangID := data.transaksi[index].ID_Barang_Yang_Dibeli
			oldBarangIndex := findBarangIndex(data, oldBarangID)
			data.inventori[oldBarangIndex].Jumlah_Terjual += data.transaksi[index].Jumlah_Pembelian
			fmt.Printf("Data Terjual di tambahkan ke produk %s dengan jumlah pembelian = %d", data.inventori[oldBarangIndex].Nama_Barang, data.inventori[oldBarangIndex].Jumlah_Terjual)
		}
		fmt.Println("Status transaksi berhasil diubah menjadi Sukses.")
	} else {
		fmt.Println("Status transaksi tidak diubah.")
	}
}

func hapusDataTransaksi(data *array) {
	var index int
	var cari_Kode string

	if data.jumlah_transaksi == 0 {
		fmt.Println("Data transaksi masih kosong.")
		return
	}

	fmt.Println("Hapus Data Transaksi")
	fmt.Print("Masukkan Kode Transaksi yang akan dihapus: ")
	fmt.Scan(&cari_Kode)
	index = sequentialSearchTransaksi(data, cari_Kode)

	if index == -1 {
		fmt.Println("Data tidak ditemukan.")
		return
	}

	// Update stock and sold quantity before deleting
	barangIndex := findBarangIndex(data, data.transaksi[index].ID_Barang_Yang_Dibeli)
	if barangIndex != -1 {
		data.inventori[barangIndex].Stock_Barang += data.transaksi[index].Jumlah_Pembelian
		data.inventori[barangIndex].Jumlah_Terjual -= data.transaksi[index].Jumlah_Pembelian
	}

	for i := index; i < data.jumlah_transaksi-1; i++ {
		data.transaksi[i] = data.transaksi[i+1]
	}
	data.jumlah_transaksi--

	// Mengatur ulang kode transaksi agar tetap unik
	for i := 0; i < data.jumlah_transaksi; i++ {
		data.transaksi[i].Kode_Transaksi = "TRX" + strconv.Itoa(i+1)
	}

	fmt.Println("Data berhasil dihapus.")
}

func sequentialSearchTransaksi(data *array, kodeTransaksi string) int {
	for i := 0; i < data.jumlah_transaksi; i++ {
		if data.transaksi[i].Kode_Transaksi == kodeTransaksi {
			return i
		}
	}
	return -1
}

func cetakDataTransaksi(data *array) {
	if data.jumlah_transaksi == 0 {
		fmt.Println("Data transaksi masih kosong.")
		return
	}

	fmt.Println("Data Transaksi")
	fmt.Println("=============================================================")
	fmt.Printf("%-15s %-20s %-15s %-10s %-10s %-10s %-10s\n", "Kode Transaksi", "Nama Pembeli", "Tanggal", "ID Barang", "Jumlah", "Total", "Status")
	fmt.Println("=============================================================")

	for i := 0; i < data.jumlah_transaksi; i++ {
		fmt.Printf("%-15s %-20s %-15s %-10d %-10d %-10d %-10s\n", data.transaksi[i].Kode_Transaksi, data.transaksi[i].Nama_Pembeli, data.transaksi[i].Tanggal_Pembelian.Format("2006-01-02"), data.transaksi[i].ID_Barang_Yang_Dibeli, data.transaksi[i].Jumlah_Pembelian, data.transaksi[i].Total, data.transaksi[i].Status)
	}
	fmt.Println("=============================================================")
	fmt.Println("Total transaksi:", data.jumlah_transaksi)
	fmt.Println("=============================================================")
}

func sortDataTransaksi(data *array) {
	var opsi, subOpsi int
	if data.jumlah_transaksi == 0 {
		fmt.Println("Data transaksi masih kosong.")
		return
	}

	fmt.Println("Urutkan Data Transaksi")
	fmt.Println("1. Urutkan berdasarkan Nama Pembeli")
	fmt.Println("2. Urutkan berdasarkan Tanggal Pembelian")
	fmt.Println("3. Urutkan berdasarkan ID Barang yang dibeli")
	fmt.Println("4. Urutkan berdasarkan Jumlah barang yang dibeli")
	fmt.Println("5. Urutkan berdasarkan Total")
	fmt.Print("Pilih opsi(1/2/3/4/5): ")
	fmt.Scan(&opsi)

	fmt.Println("1. Urutkan Secara Ascending")
	fmt.Println("2. Urutkan Secara Descending")
	fmt.Print("Pilih opsi(1/2): ")
	fmt.Scan(&subOpsi)

	switch opsi {
	case 1:
		selectionSortTransaksi(data, subOpsi, "Nama")
	case 2:
		insertionSortTransaksi(data, subOpsi, "Tanggal")
	case 3:
		selectionSortTransaksi(data, subOpsi, "ID")
	case 4:
		insertionSortTransaksi(data, subOpsi, "Jumlah")
	case 5:
		selectionSortTransaksi(data, subOpsi, "Total")
	default:
		fmt.Println("Tipe pencarian tidak valid.")
	}
}

func selectionSortTransaksi(data *array, order int, key string) {
	for i := 1; i <= data.jumlah_transaksi-1; i++ {
		idx_min := i - 1
		for j := i; j < data.jumlah_transaksi; j++ {
			var condition bool
			switch key {
			case "ID":
				condition = (order == 1 && data.transaksi[j].ID_Barang_Yang_Dibeli < data.transaksi[idx_min].ID_Barang_Yang_Dibeli) || (order == 2 && data.transaksi[j].ID_Barang_Yang_Dibeli > data.transaksi[idx_min].ID_Barang_Yang_Dibeli)
			case "Nama":
				condition = (order == 1 && strings.ToLower(data.transaksi[j].Nama_Pembeli) < strings.ToLower(data.transaksi[idx_min].Nama_Pembeli)) || (order == 2 && strings.ToLower(data.transaksi[j].Nama_Pembeli) > strings.ToLower(data.transaksi[idx_min].Nama_Pembeli))
			case "Tanggal":
				condition = (order == 1 && data.transaksi[j].Tanggal_Pembelian.Before(data.transaksi[idx_min].Tanggal_Pembelian)) || (order == 2 && data.transaksi[j].Tanggal_Pembelian.After(data.transaksi[idx_min].Tanggal_Pembelian))
			case "Jumlah":
				condition = (order == 1 && data.transaksi[j].Jumlah_Pembelian < data.transaksi[idx_min].Jumlah_Pembelian) || (order == 2 && data.transaksi[j].Jumlah_Pembelian > data.transaksi[idx_min].Jumlah_Pembelian)
			case "Total":
				condition = (order == 1 && data.transaksi[j].Total < data.transaksi[idx_min].Total) || (order == 2 && data.transaksi[j].Total > data.transaksi[idx_min].Total)
			}
			if condition {
				idx_min = j
			}
		}
		t := data.transaksi[idx_min]
		data.transaksi[idx_min] = data.transaksi[i-1]
		data.transaksi[i-1] = t
	}
}

func insertionSortTransaksi(data *array, order int, key string) {
	for i := 1; i <= data.jumlah_transaksi-1; i++ {
		j := i
		temp := data.transaksi[j]
		switch key {
		case "ID":
			for ; j > 0 && ((order == 1 && temp.ID_Barang_Yang_Dibeli < data.transaksi[j-1].ID_Barang_Yang_Dibeli) || (order == 2 && temp.ID_Barang_Yang_Dibeli > data.transaksi[j-1].ID_Barang_Yang_Dibeli)); j-- {
				data.transaksi[j] = data.transaksi[j-1]
			}
		case "Nama":
			for ; j > 0 && ((order == 1 && strings.ToLower(temp.Nama_Pembeli) < strings.ToLower(data.transaksi[j-1].Nama_Pembeli)) || (order == 2 && strings.ToLower(temp.Nama_Pembeli) > strings.ToLower(data.transaksi[j-1].Nama_Pembeli))); j-- {
				data.transaksi[j] = data.transaksi[j-1]
			}
		case "Tanggal":
			for ; j > 0 && ((order == 1 && temp.Tanggal_Pembelian.Before(data.transaksi[j-1].Tanggal_Pembelian)) || (order == 2 && temp.Tanggal_Pembelian.After(data.transaksi[j-1].Tanggal_Pembelian))); j-- {
				data.transaksi[j] = data.transaksi[j-1]
			}
		case "Jumlah":
			for ; j > 0 && ((order == 1 && temp.Jumlah_Pembelian < data.transaksi[j-1].Jumlah_Pembelian) || (order == 2 && temp.Jumlah_Pembelian > data.transaksi[j-1].Jumlah_Pembelian)); j-- {
				data.transaksi[j] = data.transaksi[j-1]
			}
		case "Total":
			for ; j > 0 && ((order == 1 && temp.Total < data.transaksi[j-1].Total) || (order == 2 && temp.Total > data.transaksi[j-1].Total)); j-- {
				data.transaksi[j] = data.transaksi[j-1]
			}
		}
		data.transaksi[j] = temp
	}
}

func cariBarang(data *array, keyword string) {
	fmt.Println("Hasil Pencarian Barang")
	fmt.Println("=============================================================")
	fmt.Printf("%-5s %-20s %-15s %-10s %-10s %-10s\n", "ID", "Nama", "Kategori", "Harga Modal", "Harga Jual", "Jumlah")
	fmt.Println("=============================================================")
	found := false
	for i := 0; i < data.jumlah_inventory; i++ {
		if contains(data.inventori[i].Nama_Barang, keyword) || contains(data.inventori[i].Kategori_Barang, keyword) {
			fmt.Printf("%-5d %-20s %-15s %-10d %-10d %-10d\n", data.inventori[i].Id_Barang, data.inventori[i].Nama_Barang, data.inventori[i].Kategori_Barang, data.inventori[i].Harga_Modal, data.inventori[i].Harga_Jual, data.inventori[i].Stock_Barang)
			found = true
		}
	}
	if !found {
		fmt.Println("Barang tidak ditemukan.")
	}
	fmt.Println("=============================================================")
}

func contains(source, substring string) bool {
	return strings.Contains(strings.ToLower(source), strings.ToLower(substring))
}

func tampilkanModalPendapatan(data *array) {
	var totalModal, totalPendapatan, totalKeuntungan int
	fmt.Println("Modal dan Pendapatan")
	fmt.Println("=============================================================")
	fmt.Printf("%-5s %-20s %-10s %-10s\n", "ID", "Nama", "Harga Modal", "Pendapatan")
	fmt.Println("=============================================================")
	for i := 0; i < data.jumlah_inventory; i++ {
		pendapatan := data.inventori[i].Harga_Jual * (data.inventori[i].Stock_Barang)
		fmt.Printf("%-5d %-20s %-10d %-10d\n", data.inventori[i].Id_Barang, data.inventori[i].Nama_Barang, data.inventori[i].Harga_Modal, pendapatan)
		totalModal += data.inventori[i].Harga_Modal * data.inventori[i].Stock_Barang
		totalPendapatan += pendapatan
	}
	fmt.Println("=============================================================")
	fmt.Printf("Total Modal: %d\n", totalModal)
	fmt.Printf("Total Pendapatan Keseluruhan Produk: %d\n", totalPendapatan)

	// Hitung total keuntungan dari transaksi
	for i := 0; i < data.jumlah_transaksi; i++ {
		if data.transaksi[i].Status == "Sukses" {
			keuntungan := data.transaksi[i].Total - (data.transaksi[i].Jumlah_Pembelian * data.inventori[findBarangIndex(data, data.transaksi[i].ID_Barang_Yang_Dibeli)].Harga_Modal)
			totalKeuntungan += keuntungan
		}
	}
	fmt.Printf("Total Keuntungan: %d\n", totalKeuntungan)
	fmt.Println("=============================================================")
}

func tampilkanBarangTerbanyakTerjual(data *array) {
	if data.jumlah_inventory == 0 {
		fmt.Println("Inventori masih kosong.")
		return
	}
	for i := 0; i < data.jumlah_inventory-1; i++ {
		for j := 0; j < data.jumlah_inventory-i-1; j++ {
			if data.inventori[j].Jumlah_Terjual < data.inventori[j+1].Jumlah_Terjual {
				data.inventori[j], data.inventori[j+1] = data.inventori[j+1], data.inventori[j]
			}
		}
	}
	fmt.Println("Barang yang paling banyak terjual:")
	fmt.Println("============================================================================================")
	fmt.Printf("%-5s | %-15s | %-20s | %-10s | %-10s | %-10s | %-15s\n", "ID", "Nama Barang", "Kategori", "Harga Modal", "Harga Jual", "Stock", "Terjual")
	max := 0
	for i := 0; i < data.jumlah_inventory; i++ {
		if data.inventori[i].Jumlah_Terjual > 0 && max < 5 {

			fmt.Println("============================================================================================")
			fmt.Printf("%-5d | %-15s | %-20s | %-10d | %-10d | %-10d | %-15d\n",
				data.inventori[i].Id_Barang,
				data.inventori[i].Nama_Barang,
				data.inventori[i].Kategori_Barang,
				data.inventori[i].Harga_Modal,
				data.inventori[i].Harga_Jual,
				data.inventori[i].Stock_Barang,
				data.inventori[i].Jumlah_Terjual)
			fmt.Println("============================================================================================")
			max++
		}
	}

}

func hasNumber(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			return true
		}
	}
	return false
}

func addDummyData(data *array) {
	// 20 data dummy barang
	data.inventori[data.jumlah_inventory] = barang{1, "Beras", "Sembako", 10000, 12000, 50, 5}
	data.jumlah_inventory++
	data.inventori[data.jumlah_inventory] = barang{2, "Gula", "Sembako", 12000, 15000, 30, 3}
	data.jumlah_inventory++
	data.inventori[data.jumlah_inventory] = barang{3, "Minyak Goreng", "Sembako", 14000, 17000, 20, 4}
	data.jumlah_inventory++
	data.inventori[data.jumlah_inventory] = barang{4, "Telur", "Sembako", 2000, 2500, 100, 10}
	data.jumlah_inventory++
	data.inventori[data.jumlah_inventory] = barang{5, "Roti", "Makanan", 5000, 7000, 40, 8}
	data.jumlah_inventory++
	data.inventori[data.jumlah_inventory] = barang{6, "Susu", "Minuman", 15000, 18000, 25, 7}
	data.jumlah_inventory++
	data.inventori[data.jumlah_inventory] = barang{7, "Kopi", "Minuman", 7000, 9000, 60, 6}
	data.jumlah_inventory++
	data.inventori[data.jumlah_inventory] = barang{8, "Teh", "Minuman", 5000, 7000, 50, 5}
	data.jumlah_inventory++
	data.inventori[data.jumlah_inventory] = barang{9, "Biskuit", "Makanan", 8000, 10000, 70, 7}
	data.jumlah_inventory++
	data.inventori[data.jumlah_inventory] = barang{10, "Coklat", "Makanan", 10000, 13000, 40, 4}
	data.jumlah_inventory++
	data.inventori[data.jumlah_inventory] = barang{11, "Air Mineral", "Minuman", 3000, 5000, 100, 10}
	data.jumlah_inventory++
	data.inventori[data.jumlah_inventory] = barang{12, "Jus Buah", "Minuman", 8000, 10000, 30, 3}
	data.jumlah_inventory++
	data.inventori[data.jumlah_inventory] = barang{13, "Mie Instan", "Makanan", 2000, 3000, 200, 20}
	data.jumlah_inventory++
	data.inventori[data.jumlah_inventory] = barang{14, "Kacang", "Makanan", 6000, 8000, 80, 8}
	data.jumlah_inventory++
	data.inventori[data.jumlah_inventory] = barang{15, "Saus", "Sembako", 4000, 6000, 90, 9}
	data.jumlah_inventory++
	data.inventori[data.jumlah_inventory] = barang{16, "Kecap", "Sembako", 5000, 7000, 70, 7}
	data.jumlah_inventory++
	data.inventori[data.jumlah_inventory] = barang{17, "Sabun", "Kebutuhan", 3000, 5000, 50, 5}
	data.jumlah_inventory++
	data.inventori[data.jumlah_inventory] = barang{18, "Shampoo", "Kebutuhan", 6000, 8000, 40, 4}
	data.jumlah_inventory++
	data.inventori[data.jumlah_inventory] = barang{19, "Pasta Gigi", "Kebutuhan", 5000, 7000, 60, 6}
	data.jumlah_inventory++
	data.inventori[data.jumlah_inventory] = barang{20, "Tisu", "Kebutuhan", 2000, 4000, 80, 8}
	data.jumlah_inventory++

	// data dummy pembelian
	data.transaksi[data.jumlah_transaksi] = pembelian{"TRX1", "Ali", time.Now().AddDate(0, 0, -1), 1, 5, 60000, "Proses"}
	data.jumlah_transaksi++
	data.transaksi[data.jumlah_transaksi] = pembelian{"TRX2", "Budi", time.Now().AddDate(0, 0, -2), 2, 3, 45000, "Proses"}
	data.jumlah_transaksi++
	data.transaksi[data.jumlah_transaksi] = pembelian{"TRX3", "Charlie", time.Now().AddDate(0, 0, -3), 3, 4, 68000, "Proses"}
	data.jumlah_transaksi++
	data.transaksi[data.jumlah_transaksi] = pembelian{"TRX4", "Dina", time.Now().AddDate(0, 0, -4), 4, 10, 25000, "Proses"}
	data.jumlah_transaksi++
	data.transaksi[data.jumlah_transaksi] = pembelian{"TRX5", "Eva", time.Now().AddDate(0, 0, -5), 5, 8, 56000, "Proses"}
	data.jumlah_transaksi++
	data.transaksi[data.jumlah_transaksi] = pembelian{"TRX6", "Faisal", time.Now().AddDate(0, 0, -6), 6, 7, 126000, "Proses"}
	data.jumlah_transaksi++
	data.transaksi[data.jumlah_transaksi] = pembelian{"TRX7", "Gina", time.Now().AddDate(0, 0, -7), 7, 6, 54000, "Proses"}
	data.jumlah_transaksi++
	data.transaksi[data.jumlah_transaksi] = pembelian{"TRX8", "Hadi", time.Now().AddDate(0, 0, -8), 8, 5, 35000, "Proses"}
	data.jumlah_transaksi++
	data.transaksi[data.jumlah_transaksi] = pembelian{"TRX9", "Ika", time.Now().AddDate(0, 0, -9), 9, 7, 70000, "Proses"}
	data.jumlah_transaksi++
	data.transaksi[data.jumlah_transaksi] = pembelian{"TRX10", "Joko", time.Now().AddDate(0, 0, -10), 10, 4, 52000, "Proses"}
	data.jumlah_transaksi++
}
