# The Go Programming Language (Golang)

## Penjelasan

Golang adalah bahasa pemograman yang dibuat oleh Google. Dikembangkan dari tahun 2007 dan dirilis ke publik pada tahun 2012. Bahasa pemograman berbasis prosedural atau POP ( Procedural Oriented Programming). Perbedaan dengan python lainnya adalah bahasa ini bersifat compiled, proses evaluasi atau compiling dilakukan sebelum di run dalam proses build sehinga mengahsilkan compiled file, dan file tersebut yang di eksekusi (run). 

## Struktur Data (Golang)

## Kondisional dan Tipe Data

### Kondisional dalam Golang

Kondisional adalah pembuatan / pengambilan keputusan dan mengeksekusi code berdasarkan kondisi tertentu pada suatu program. Jenis kondisional yang umum digunakan di golang meliputi ``if``, ``else``, dan ``else if``.

#### Statement `if`

Kondisional `if` membuat program dapat menentukan dan mengeksekusi suatu code bila kondisi yang dibuat terpenuhi.

```go
package main

import "fmt"

func main() {
	nilai := 70

	if nilai >= 65 {
		fmt.Println("Lulus KKM")
	}
}

// Output
// Lulus KKM
```

#### Statement ``if-else``

Kondisional `if-else` membuat program dapat menentukan dan mengeksekusi suatu code sesuai dengan kondisi yang terpenuhi diantara dua kemungkinan.

```go
package main

import "fmt"

func main() {
	nilai := 60

	if nilai >= 65 {
		fmt.Println("Lulus KKM")
	} else {
		fmt.Println("Tidak Lulus KKM")
	}
}

// Output
// Tidak Lulus KKM
```

### Statement ``if-else if-else``

Kondisional `if-else if-else` digunakan ketika terdapat lebih dari dua kondisi dan program harus dapat menentukan dan mengekesekusi suatu kode berdasarkan kondisi yang terpenuhi.

```go
package main

import "fmt"

func main() {
	nilai := 70

	if nilai >= 90 {
		fmt.Println("Nilai Akhir: A")
	} else if nilai >= 80 {
		fmt.Println("Nilai Akhir: B")
	} else if nilai >= 70 {
		fmt.Println("Nilai Akhir: C")
	} else if nilai >= 60 {
		fmt.Println("Nilai Akhir: D")
	} else {
		fmt.Println("Tidak Lulus KKM")
	}
}

// Output
// Nilai Akhir: C
```

### Tipe Data dalam Golang

Golang memiliki tiga tipe data dasar (basic). Tipe data tersebut meliputi numberic, string, dan boolean.

### 1. Tipe Data Numeric

Tipe data numeric mewakili bilangan bulat (`integer`), bilangan pecahan (`float`), dan bilangan kompleks (`complex`) yang memungkinkan operasi aritmatika.

Contoh:

```go
package main
import ("fmt")

func main() {
  var a int = 5     // Integer
  var b int = 3     // Integer
  
  fmt.Println("Integer:   ", a)
  fmt.Println("Integer:   ", b)
  
  penjumlahan := a + b
  pengurangan := a - b
  perkalian := a * b
  pembagian := a / b
  
  fmt.Println("Penjumlahan: ", penjumlahan)
  fmt.Println("Pengurangan: ", pengurangan)
  fmt.Println("Perkalian:   ", perkalian)
  fmt.Println("Pembagian:   ", pembagian)  
}

// Output
// Integer:    5
// Integer:    3
// Penjumlahan:  8
// Pengurangan:  2
// Perkalian:    15
// Pembagian:    1
```

Pada golang, kita tidak dapat melakukan operasi aritmatika bila terdapat perbedaan jenis tipe data numeric, misal penjumlahan bilangan integer dengan bilangan float, ketika operasi aritmatika dilakukan terhadap dua bilangan tersebut maka akan muncul error ketika program di eksekusi.

Contoh:

```go
package main
import ("fmt")

func main() {
  var a int = 5            // Integer
  var b float64 = 3.14     // Floating point number
  
  fmt.Println("Integer: ", a)
  fmt.Println("Float:   ", b)
  
  penjumlahan := a + b
  pengurangan := a - b
  perkalian := a * b
  pembagian := a / b
  
  fmt.Println("Penjumlahan: ", penjumlahan)
  fmt.Println("Pengurangan: ", pengurangan)
  fmt.Println("Perkalian:   ", perkalian)
  fmt.Println("Pembagian:   ", pembagian)  
}

// Output
// ./prog.go:14:17: invalid operation: a + b (mismatched types int and float64)
// ./prog.go:15:17: invalid operation: a - b (mismatched types int and float64)
// ./prog.go:16:15: invalid operation: a * b (mismatched types int and float64)
// ./prog.go:17:15: invalid operation: a / b (mismatched types int and float64)
```

#### 2. Tipe Data String

Tipe data string memungkinkan penyimpanan data berupa teks dan memiliki berbagai metode untuk manipulasi data tersebut.

Contoh:

```go
package main
import (
  "fmt";
  "strings"
)

func main() {
  // Manipulasi string menggunakan method string   
  message := "Ini adalah contoh message golang."

  fmt.Println(message)
  
  // Konversi string ke huruf besar dengan method ToUpper()
  fmt.Println(strings.ToUpper(message))
  
  // Memisahkan kata-kata dengan method Split()
  splittedString := strings.Split(message, " ")
  fmt.Println(splittedString[0])
}

// Output
// Ini adalah contoh message golang.
// INI ADALAH CONTOH MESSAGE GOLANG.
// Ini
```

#### 3. Tipe Data Bool

Tipe data boolean (`bool`) memiliki nilai `True` atau `False`.

Contoh:

```go
package main
import ("fmt")
func main() {
  var a bool = true     // Boolean

  fmt.Println("Boolean: ", a)
  
  if a {
    fmt.Println("Ini kondisi ketika nilai benar")
  } else {
    fmt.Println("Ini kondisi ketika nilai salah")
  }
}

// Output
// Boolean:  true
// Ini kondisi ketika nilai benar
```

#### Tipe Data dalam Golang - Lanjutan

#### Tipe Data Arrays

Tipe data Arrays adalah urutan elemen dengan panjang tertentu.

Contoh:

```go
package main

import "fmt"

func main() {

    var array1 [4]int
    // print array1
    fmt.Println("empty_array:", array1)
    
    // memasukan data ke array1[4] 
    array1[3] = 80
    
    // print array1
    fmt.Println("set:", array1)
    // print array1[3]
    fmt.Println("get:", array1[3])
    
    // print len array1
    fmt.Println("len:", len(array1))
    
    array2 := [5]int{1, 2, 3, 4, 5}
    
    // print array2
    fmt.Println("array2:", array2)
}

// Output
// empty_array: [0 0 0 0]
// set: [0 0 0 80]
// get: 80
// len: 4
// array2: [1 2 3 4 5]
```

#### Tipe Data Slice

Tipe data Slice adalah reference elemen array. Slice dapat dibuat, atau bisa juga dihasilkan dari manipulasi sebuah array ataupun slice lainnya. Karena merupakan data reference, menjadikan perubahan data di tiap elemen slice akan berdampak pada slice lain yang memiliki alamat memori yang sama.

Contoh:

```go
package main

import (
    "fmt"
    "slices"
)

func main() {

    var slice1 []string
    fmt.Println("slice1:", slice1, slice1 == nil, len(slice1) == 0)

    slice1 = make([]string, 3)
    // print slice1, len slice1, cap slice1
    fmt.Println("empty_slice:", slice1, "len:", len(slice1), "cap:", cap(slice1))
    
    // append value ke slice1
    slice1[0] = "ayam"
    slice1[1] = "sapi"
    slice1[2] = "domba"
    fmt.Println("set:", slice1)
    fmt.Println("get:", slice1[2])

    fmt.Println("len:", len(slice1))

    slice1 = append(slice1, "kelinci")
    slice1 = append(slice1, "rusa", "kuda")
    fmt.Println("slice_append:", slice1)
    
    // copy slice1 ke copy_slice
    copy_slice := make([]string, len(slice1))
    copy(copy_slice, slice1)
    fmt.Println("copy_slice:", copy_slice)

// Output
// slice1: [] true true
// empty_slice: [  ] len: 3 cap: 3
// set: [ayam sapi domba]
// get: domba
// len: 3
// slice_append: [ayam sapi domba kelinci rusa kuda]
// copy_slice: [ayam sapi domba kelinci rusa kuda]
// }
```

#### 3. Tipe Data Maps

Tipe data Maps atau yang lebih familiar disebut dictionary di bahasa pemrograman lain memungkinkan penyimpanan data dalam bentuk pasangan kunci-nilai.

```go
package main

import (
	"fmt"
	"maps"
)

func main() {

	map1 := make(map[string]int)

	// append key dan value
	map1["a1"] = 10
	map1["a2"] = 15

	// print maps map1
	fmt.Println("map:", map1)

	fmt.Println("len:", len(map1))

	// membuat variable v1 dan v3 dari elemen maps
	v1 := map1["a1"]
	fmt.Println("v1:", v1)

	v3 := map1["a3"]
	fmt.Println("v3:", v3)

	// menghapus elemen a2
	delete(map1, "a2")
	fmt.Println("map:", map1)

	// menghapus isi map1
	clear(map1)
	fmt.Println("map:", map1)

	// membuat new_map
	new_map := map[string]int{"coklat": 1, "permen": 2}
	fmt.Println("map:", new_map)

	// membuat new_map_2
	new_map_2 := map[string]int{"coklat": 1, "permen": 2}

	// membandingkan new_map dengan new_map_2
	if maps.Equal(new_map, new_map_2) {
		fmt.Println("new_map == new_map_2")
	}
}

// Output
//map: map[a1:10 a2:15]
// len: 2
// v1: 10
// v3: 0
// map: map[a1:10]
// map: map[]
// map: map[bar:2 foo:1]
// new_map == new_map_2
```

## Perulangan / Iterasi dalam Golang

Perulangan memungkinkan kita untuk melakukan eksekusi terhadap sekelompok pernyataan secara berulang kali. Berikut adalah beberapa contoh iterasi dengan `for` loop.

### `for` Loop

### Basic `for`

Contoh:

```go
package main
import "fmt"

func main() {
    var num, total = 15, 0
    
    for i := 1 ; i <= num; i++ {
        total += i      
    }

    fmt.Println("total =", total)
}

// Output
// total = 120
```

### `for` untuk slice or array

Contoh:

```go
package main

import (
    "fmt"
    "maps"
)

func main() {
    array1 := []string{"Foo", "Bar"}
    for i, v := range array1 {
	    fmt.Println(i, v)
    }
}

// Output
// 0 Foo
// 1 Bar
```

### `for` untuk Maps

Contoh:

```go
package main

import (
    "fmt"
    "maps"
)

func main() {
    m := map[string]int{
	    "ayam":   1,
	    "sapi":   2,
	    "kambing": 3,
    }
    for k, v := range m {
	    fmt.Println(k, v)
    }
}

// Output
// ayam 1
// sapi 2
// kambing 3
```

## Manipulasi File / Berkas dalam Golang

Manipulasi file / berkas adalah proses untuk membaca, menulis, atau memanipulasi informasi dalam file menggunakan Golang. Golang menyediakan sejumlah fungsi bawaan yang memungkinkan kita untuk melakukan operasi-operasi pada file, seperti membaca isi file, menulis ke file, dan lainnya.

### Membuka File / Berkas

### 1. Membuka File dan Membaca Keseluruhan File

Untuk membuka file dan membaca keseluruhan file dalam Golang, kita dapat menggunakan fungsi `os.ReadFile()`.

Contoh:

```go
package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    content, err := os.ReadFile("file.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(content))
}

// Ouput
// Hello World!!
```

### 2. Membuka File dan Membaca baris per baris

Untuk membuka file dan membaca file baris per baris dalam Golang, kita dapat menggunakan fungsi `os.Open()` dan `bufio.NewScanner()`.

Contoh:

```go
package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    // open file
    f, err := os.Open("file.txt")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file at the end of the program
    defer f.Close()

    // read the file line by line using scanner
    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        // do something with a line
        fmt.Printf("line: %s\n", scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

// Output
// Hello Word!!
// This is New file!!
```

### 3. Membuka File dan Membaca kata per kata

Untuk membuka file dan membaca file kata per kata dalam Golang, kita dapat menggunakan fungsi `os.Open()`, `bufio.NewScanner()`, dan `bufio.ScanWord`.

Contoh:

```go
package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    // open file
    f, err := os.Open("file.txt")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file at the end of the program
    defer f.Close()

    // read the file word by word using scanner
    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanWords)

    for scanner.Scan() {
        // do something with a word
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

// Ouput 
// Hello
// Word!!
// This
// is
// New
// file!!
```

### Menulis ke File / Berkas

### 1. Menulis seluruh konten ke dalam file dalam satu waktu

Untuk menulis seluruh konten ke dalam file dalam satu waktu pada Golang, kita dapat menggunakan fungsi `os.WriteFile()` dengan tiga input paramater, yaitu path file, byte data, permission bits dari file yang akan dibuat.

Contoh:

```go
package main

import (
	"log"
	"os"
)

func main() {
	if err := os.WriteFile("file_write.txt", []byte("Hello World!!"), 0666); err != nil {
		log.Fatal(err)
	}
}
```

### 2. Menulis ke dalam file baris per baris

Untuk menulis ke dalam file baris per baris pada Golang, kita dapat menggunakan fungsi `func (*File) WriteString() `.

Contoh:

```go
package main

import (
    "log"
    "os"
)

var lines = []string{
    "BTJ",
    "adalah",
    "singkatan",
    "dari",
    "Bangunindo",
    "Teknusa",
    "Jaya",
}

func main() {
    // create file
    f, err := os.Create("file_write.txt")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file
    defer f.Close()

    for _, line := range lines {
        _, err := f.WriteString(line + "\n")
        if err != nil {
            log.Fatal(err)
        }
    }
}
```



## Run Simple Script

### Installation

1. Download file golang Di: [download_url](https://go.dev/doc/install)

2. Restart terminal

3. Cek Versi untuk memastikan bahwa golang sudah terinstall

```
 go version
```

### Create Dependencies TO run Program
1. Masuk ke dalam direktori berisi program go

2. Pilih lah 1 file yang berisi function `main()` untuk di run, kemudian set sebagai init, contoh: sample.go

```
go mod init sample.go 
```

3. run program tersebut dengan command:

```
go run .
```

4. Program tidak dapat  berjalan jika terdapat lebih dari satu `main()` pada direktori tersebut. Diperlukan pengaturan package dan module kedepannya.