package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    var data_file []string
    
    
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
        // Simpan data lama / text lama
        data_file = append(data_file, scanner.Text())
    }
    
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    
    fmt.Println(data_file)
    
    data_file = append(data_file, "Tambahan baris 1")
    data_file = append(data_file, "Tambahan baris dua")

    fmt.Println(data_file)
    
    err_r := os.Remove("file.txt")
    if err_r != nil {
        log.Fatal(err)
    }
    
    
    f_c, err_c := os.Create("file.txt")
    if err_c != nil {
        log.Fatal(err)
    }
    
    defer f_c.Close()
    
    for _, line := range data_file {
        _, err_i := f_c.WriteString(line + "\n")
        if err_i != nil {
            log.Fatal(err_i)
        }
    }
    
    
}

// Output
// Hello Word!!
// This is New file!!
