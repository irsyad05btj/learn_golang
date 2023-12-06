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
