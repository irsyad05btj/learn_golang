# Kasus: Pembulatan ke atas nilai hasil ujian mahasiswa

## Deskripsi Masalah

Seorang dosen yang mengajar di suatu perguruan tinggi sedang memeriksa ujian mahasiswa dan akan melakukan pembulatan ke atas nilai tersebut berdasarkan aturan berikut:

- Setiap mahasiswa menerima dalam rentang inklusif dari 0 hingga 100
- Nilai kurang dari 40 adalah nilai yang tidak lulus
- Jika selisih nilai dengan kelipatan 5 berikutnya kurang dari 3, bulatkan ke atas ke kelipatan 5 berikutnya.
- Jika nilai nilainya kurang dari 38, tidak dilakukan pembulatan karena hasilnya tetap nilai yang tidak lulus.

### Contoh
- Nilai = 83 -> dibulatkan menjadi 85 (85 - 83 kurang dari 3)
- Nilai = 56 -> jangan dibulatkan (60 - 56 adalah 4 atau lebih dari 3)
- Nilai = 33 -> jangan dibulatkan (33 kurang dari 38 yang dapat dibulatkan menjadi 40 yang merupakan standar kelulusan)

### Objektif

Tulis kode untuk mengotomatiskan proses pembulatan, dengan input awal jumlah mahasiwa diikuti dengan nilai dari masing-masing mahasiswa. 

### Limitasi
1 <= jumlah mahasiswa <= 60
0 <= Nilai <= 100

### Sample Input
```
4
74
34
38
67
```

### Sample Output
```
75
34
40
67
```
