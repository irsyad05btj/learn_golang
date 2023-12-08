package main

import (
  //  "bufio"
//    "fmt"
    "log"
    "os"
)

func main() {
    // var data_file []string
    
    // open file
    f, err := os.Open("file.txt")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file at the end of the program
    defer f.Close()

    _, err_i := f.WriteString("Data Baru 99 \n")
    if err_i != nil {
        log.Fatal(err_i)
    }
}
