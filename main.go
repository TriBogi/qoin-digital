package main

import (
	"fmt"
	"math/rand"
	"time"
)

func lemparDadu(jumlahDadu int) []int {
	dadu := make([]int, jumlahDadu)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < jumlahDadu; i++ {
		dadu[i] = rand.Intn(6) + 1
	}
	return dadu
}

func evaluasiDadu(pemain *[]int, pemainSebelah *[]int) {
	for i := 0; i < len(*pemain); i++ {
		switch (*pemain)[i] {
		case 6:
			*pemainSebelah = append(*pemainSebelah, 6)
			(*pemain)[i] = 0
		case 1:
			if len(*pemainSebelah) > 0 {
				(*pemainSebelah)[0] = 1
				(*pemain)[i] = 0
			}
		}
	}
	if len(*pemainSebelah) > 0 {
		*pemainSebelah = (*pemainSebelah)[1:]
	}
}

func tampilkanHasil(pemainHasil [][]int, poinPemain []int) {
	for i, hasil := range pemainHasil {
		fmt.Printf("Pemain #%d (%d): ", i+1, poinPemain[i])
		for _, dadu := range hasil {
			if dadu > 0 {
				fmt.Printf("%d,", dadu)
			}
		}
		fmt.Println()
	}
	fmt.Println("==================")
}

func main() {
	var jumlahPemain, jumlahDadu int

	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scan(&jumlahPemain)
	fmt.Print("Masukkan jumlah dadu: ")
	fmt.Scan(&jumlahDadu)

	pemainHasil := make([][]int, jumlahPemain)
	poinPemain := make([]int, jumlahPemain)

	for giliran := 1; ; giliran++ {
		fmt.Printf("Giliran %d lempar dadu:\n", giliran)

		for i := 0; i < jumlahPemain; i++ {
			pemainHasil[i] = lemparDadu(jumlahDadu)
			fmt.Printf("Pemain #%d (%d): %v\n", i+1, poinPemain[i], pemainHasil[i])
		}

		for i := 0; i < jumlahPemain; i++ {
			evaluasiDadu(&pemainHasil[i], &pemainHasil[(i+1)%jumlahPemain])
		}

		tampilkanHasil(pemainHasil, poinPemain)

		selesai := 0
		for i := 0; i < jumlahPemain; i++ {
			if len(pemainHasil[i]) > 0 {
				selesai++
			}
		}

		if selesai == 1 {
			for i := 0; i < jumlahPemain; i++ {
				if len(pemainHasil[i]) > 0 {
					poinPemain[i] += len(pemainHasil[i])
				}
			}
			tampilkanHasil(pemainHasil, poinPemain)
			maxPoin := 0
			pemenang := 0
			for i, poin := range poinPemain {
				if poin > maxPoin {
					maxPoin = poin
					pemenang = i + 1
				}
			}
			fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya.\n", pemenang)
			break
		}
	}
}
