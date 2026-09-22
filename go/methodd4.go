package main

import "fmt"

func main() {
	var y1, y2, jumlah, sebelum, sesudah int

	fmt.Scan(&y1, &y2)

	jumlah = 0
	cariSemuaKabisat(y1, y2, &jumlah)

	sebelum = cariKabisatSebelum(y1)
	sesudah = cariKabisatSesudah(y2)

	fmt.Printf("sebelum: %d, sesudah: %d\n", sebelum, sesudah)
}

func isKabisat(tahun int) bool {
	return (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)
}

func cariSemuaKabisat(y1, y2 int, jumlah *int) {
	var i int
	var ada bool
	ada = false
	*jumlah = 0

	for i = y1; i <= y2; i++ {
		if isKabisat(i) {
			if ada {
				fmt.Printf(" %d", i)
			} else {
				fmt.Printf("%d", i)
				ada = true
			}
			*jumlah++
		}
	}

	if !ada {
		fmt.Println("NONE")
	} else {
		fmt.Println()
	}

	fmt.Println(*jumlah)
}

func cariKabisatSebelum(tahun int) int {
	var i int
	for i = tahun - 1; i >= 1; i-- {
		if isKabisat(i) {
			return i
		}
	}
	return -1
}

func cariKabisatSesudah(tahun int) int {
	var i int
	for i = tahun + 1; i <= 9999; i++ {
		if isKabisat(i) {
			return i
		}
	}
	return -1
}

func hitungKabisat(y1, y2 int) int {
	var i, jumlah int
	jumlah = 0
	for i = y1; i <= y2; i++ {
		if isKabisat(i) {
			jumlah++
		}
	}
	return jumlah
}