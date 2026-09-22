package main

import "fmt"

func adaHurufBesar(s string) bool {
    var i int
    var hasil bool
    hasil = false
    for i = 0; i < len(s); i++ {
        if s[i] >= 'A' && s[i] <= 'Z' {
            hasil = true
        }
    }
    return hasil
}

func adaHurufKecil(s string) bool {
    var i int
    var hasil bool
    hasil = false
    for i = 0; i < len(s); i++ {
        if s[i] >= 'a' && s[i] <= 'z' {
            hasil = true
        }
    }
    return hasil
}

func adaAngka(s string) bool {
    var i int
    var hasil bool
    hasil = false
    for i = 0; i < len(s); i++ {
        if s[i] >= '0' && s[i] <= '9' {
            hasil = true
        }
    }
    return hasil
}

func validPassword(s string) bool {
    return len(s) >= 8 &&
        adaHurufBesar(s) &&
        adaHurufKecil(s) &&
        adaAngka(s)
}

func main() {
    var s string
    fmt.Scan(&s)
    if validPassword(s) {
        fmt.Println("VALID")
    } else {
        fmt.Println("TIDAK VALID")
    }
}