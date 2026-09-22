package main

import "fmt"

const NMAX int = 100

type barang struct {
	nama     string
	harga    int
	qty      int
	subtotal int
}

type tabBarang [100]barang

func inputTransaksi(T *tabBarang, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&T[i].nama, &T[i].harga, &T[i].qty)
		T[i].subtotal = T[i].harga * T[i].qty
	}
}

func cetakStruk(T tabBarang, n int) {
	var i int
	var grandTotal int

	fmt.Println("Struk Belanja")
	fmt.Println("Nama Barang         | Harga Satuan | Qty | Total Harga")
	fmt.Println("--------------------+--------------+-----+------------")

	grandTotal = 0
	for i = 0; i < n; i++ {
		fmt.Printf("%-20s| Rp. %8d | %3d | Rp. %8d\n",
			T[i].nama, T[i].harga, T[i].qty, T[i].subtotal)
		grandTotal = grandTotal + T[i].subtotal
	}

	fmt.Printf("Total Bayar: Rp.%d\n", grandTotal)
}

func main() {
	var keranjang tabBarang
	var nBarang int

	fmt.Scan(&nBarang)
	inputTransaksi(&keranjang, nBarang)
	cetakStruk(keranjang, nBarang)
}