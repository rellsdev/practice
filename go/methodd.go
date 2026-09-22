package main

import "fmt"

const N int = 10

type archer struct {
	nama  string
	hasil [10]int
}

func printPoin(p archer) {
	var i int
	fmt.Print("(")
	for i = 0; i < N; i++ {
		if i < N-1 {
			fmt.Printf("%d, ", p.hasil[i])
		} else {
			fmt.Printf("%d", p.hasil[i])
		}
	}
	fmt.Println(")")
}

func decider(p1 archer, p2 archer, poin1 *int, poin2 *int) {
	var i int
	var streak1 int
	var streak2 int

	*poin1 = 0
	*poin2 = 0
	streak1 = 0
	streak2 = 0

	for i = 0; i < N; i++ {
		*poin1 = *poin1 + p1.hasil[i]
		*poin2 = *poin2 + p2.hasil[i]

		if p1.hasil[i] >= 9 {
			streak1 = streak1 + 1
		} else {
			streak1 = 0
		}
		if p2.hasil[i] >= 9 {
			streak2 = streak2 + 1
		} else {
			streak2 = 0
		}

		if streak1 == 3 {
			*poin1 = *poin1 + 5
			streak1 = 0
		}
		if streak2 == 3 {
			*poin2 = *poin2 + 5
			streak2 = 0
		}

		if p1.hasil[i]-p2.hasil[i] >= 6 {
			*poin1 = *poin1 + 3
		}
		if p2.hasil[i]-p1.hasil[i] >= 6 {
			*poin2 = *poin2 + 3
		}
	}
}

func main() {
	var p1 archer
	var p2 archer
	var poin1 int
	var poin2 int
	var i int

	fmt.Scan(&p1.nama)
	for i = 0; i < N; i++ {
		fmt.Scan(&p1.hasil[i])
	}
	fmt.Scan(&p2.nama)
	for i = 0; i < N; i++ {
		fmt.Scan(&p2.hasil[i])
	}

	decider(p1, p2, &poin1, &poin2)

	if poin1 > poin2 {
		fmt.Println("Pemenang :", p1.nama)
		printPoin(p1)
		fmt.Println("Total poin :", poin1)
	} else if poin2 > poin1 {
		fmt.Println("Pemenang :", p2.nama)
		printPoin(p2)
		fmt.Println("Total poin :", poin2)
	} else {
		fmt.Println("Draw")
		fmt.Println("Poin", p1.nama, ":", poin1)
		fmt.Println("Poin", p2.nama, ":", poin2)
	}
}