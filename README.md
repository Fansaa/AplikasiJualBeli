# README: Aplikasi Jual Beli Barang dan Transaksi

## Deskripsi Aplikasi
Aplikasi ini adalah sistem manajemen inventori dan transaksi berbasis *Go (Golang)* yang digunakan untuk mengelola data barang dan transaksi pembelian secara efisien. Aplikasi ini dapat membantu dalam mengelola inventori, mencatat transaksi penjualan, dan memberikan analisis terkait penjualan serta keuntungan.

## Fitur Aplikasi
1. *Menu Barang*:
   - Menambah data barang.
   - Mengedit data barang.
   - Menghapus data barang.

2. *Menu Transaksi*:
   - Menambah data transaksi.
   - Mengedit data transaksi.
   - Mengupdate status transaksi (Proses atau Sukses).
   - Menghapus data transaksi.

3. *Fitur Lainnya*:
   - Mencetak data barang dan transaksi.
   - Mengurutkan data barang berdasarkan ID, nama, harga, dan jumlah.
   - Mengurutkan data transaksi berdasarkan nama pembeli, tanggal, ID barang, jumlah, dan total.
   - Melakukan pencarian barang berdasarkan nama atau kategori.
   - Menampilkan total modal dan pendapatan.
   - Menampilkan barang yang paling banyak terjual.

## Struktur Data
- *Barang*: Menyimpan informasi terkait produk di inventori.
  - Id_Barang, Nama_Barang, Kategori_Barang, Harga_Modal, Harga_Jual, Stock_Barang, Jumlah_Terjual.

- *Pembelian*: Menyimpan informasi terkait transaksi pembelian.
  - Kode_Transaksi, Nama_Pembeli, Tanggal_Pembelian, ID_Barang_Yang_Dibeli, Jumlah_Pembelian, Total, Status.

- *Array*: Struktur data utama yang menyimpan inventori dan transaksi.
  - inventori, jumlah_inventory, transaksi, jumlah_transaksi.

## Cara Menjalankan Aplikasi
1. *Pastikan Golang sudah terpasang* di komputer Anda. Jika belum, silakan unduh dan pasang dari [Golang](https://golang.org/dl/).
2. Simpan kode aplikasi dalam file bernama main.go.
3. Buka terminal di direktori tempat file main.go disimpan.
4. Jalankan perintah berikut:
   bash
   go run main.go
   
5. Ikuti instruksi yang ditampilkan di layar untuk berinteraksi dengan aplikasi.

## Struktur Menu
Berikut adalah menu utama yang tersedia dalam aplikasi:

### Menu Utama

1. Menu Barang
2. Print Data Barang
3. Urutkan Data Barang
4. Menu Transaksi
5. Print Data Transaksi
6. Urutkan Data Transaksi
7. Pencarian Barang
8. Tampilkan Modal dan Pendapatan
9. Tampilkan Barang Terbanyak Terjual
10. Keluar


### Menu Barang

1. Tambah data barang
2. Edit data barang
3. Hapus data barang
4. Kembali ke menu utama


### Menu Transaksi

1. Tambah data transaksi
2. Edit data transaksi
3. Update Status Transaksi
4. Hapus data transaksi
5. Kembali ke menu utama


## Contoh Dummy Data
Aplikasi sudah dilengkapi dengan *data dummy* sebanyak 20 data barang dan 10 data transaksi, yang dapat digunakan untuk menguji aplikasi.

### Contoh Data Barang
| ID | Nama           | Kategori  | Harga Modal | Harga Jual | Stock | Terjual |
|----|----------------|-----------|-------------|------------|-------|---------|
| 1  | Beras         | Sembako   | 10000       | 12000      | 50    | 5       |
| 2  | Gula          | Sembako   | 12000       | 15000      | 30    | 3       |
| 3  | Minyak Goreng | Sembako   | 14000       | 17000      | 20    | 4       |

### Contoh Data Transaksi
| Kode Transaksi | Nama Pembeli | Tanggal    | ID Barang | Jumlah | Total | Status  |
|----------------|--------------|------------|-----------|--------|-------|---------|
| TRX1           | Ali         | 2024-11-08 | 1         | 5      | 60000 | Proses  |
| TRX2           | Budi        | 2024-11-07 | 2         | 3      | 45000 | Proses  |
| TRX3           | Charlie     | 2024-11-06 | 3         | 4      | 68000 | Proses  |

## Penjelasan Fungsi Utama
- *tambahDataBarang()*: Menambah barang baru ke dalam inventori.
- *editDataBarang()*: Mengedit detail barang yang ada.
- *hapusDataBarang()*: Menghapus barang dari inventori.
- *tambahDataTransaksi()*: Menambah transaksi baru.
- *editDataTransaksi()*: Mengedit detail transaksi.
- *updateStatusTransaksi()*: Mengubah status transaksi menjadi sukses.
- *hapusDataTransaksi()*: Menghapus transaksi dari daftar.
- *sortDataBarang()*: Mengurutkan data barang berdasarkan berbagai kriteria.
- *sortDataTransaksi()*: Mengurutkan data transaksi berdasarkan kriteria tertentu.
- *tampilkanModalPendapatan()*: Menampilkan total modal, pendapatan, dan keuntungan.
- *tampilkanBarangTerbanyakTerjual()*: Menampilkan barang yang paling banyak terjual.

## Teknologi yang Digunakan
- Bahasa Pemrograman: *Go (Golang)*
- Modul: fmt, strconv, strings, time

## Kontribusi
Jika Anda ingin berkontribusi dalam pengembangan aplikasi ini, silakan lakukan pull request pada repositori GitHub.

## Kontak
Jika ada pertanyaan atau saran, silakan hubungi saya di *fathanfsanum@gmail.com*.

Terima kasih telah menggunakan aplikasi ini! 🎉
