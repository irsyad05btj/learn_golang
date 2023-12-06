# Skenario Kompleks: Sistem Manajemen Inventaris

Bayangkan Anda ditugaskan untuk membuat sistem manajemen inventaris untuk sebuah toko kecil. Sistem ini harus dapat:

- Menambahkan produk ke inventaris.
- Memperbarui jumlah produk.
- Menampilkan inventaris saat ini.
- Menghitung total nilai inventaris.

## Pembagian menjadi Komponen Non-Kompleks

### Menambahkan Produk ke Inventaris

- Pengguna memberikan detail produk (nama, harga, jumlah).
- Program menambahkan produk ke dalam inventaris.

### Memperbarui Jumlah Produk

- Pengguna menentukan produk yang ingin diperbarui dan jumlah baru.
- Program menemukan produk di dalam inventaris dan memperbarui jumlahnya.

### Menampilkan Inventaris Saat Ini

- Program mencetak daftar semua produk dalam inventaris beserta detailnya (nama, harga, jumlah).

### Menghitung Total Nilai Inventaris

- Program mengulang melalui inventaris, mengalikan jumlah setiap produk dengan harganya, dan menjumlahkan nilai-nilai ini.

## Boilerplate Code

```go
package main

import (
    "fmt"
)

// Inventaris awal kosong
var item map[string]int
var inventaris []item

// Fungsi untuk menambahkan produk ke inventaris
func tambah_produk(nama string, harga int, jumlah int){
    fmt.Println(nama, harga)
    // proses untuk menambahkan produk
}

// Fungsi untuk memperbarui jumlah produk
func perbarui_jumlah(nama_produk string , jumlah_baru int){
    // proses untuk memperbarui jumlah produk
    fmt.Println(nama_produk, jumlah_baru)
}

// Fungsi untuk menampilkan inventaris saat ini
func tampilkan_inventaris(){
    print("Inventaris Saat Ini:")
    //  proses untuk menampilkan inventaris saat ini (menggunakan loops)
}

// Fungsi untuk menghitung total nilai inventaris
func hitung_total_nilai(){
    total_nilai := 0
    // proses untuk menghitung total nilai
    fmt.Println("Total nilai inventaris: ", total_nilai)
}

func main() {
    
}
```