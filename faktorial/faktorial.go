package main

// Ini soal no 1
import (
	"fmt"
	"math"
)

// Menghitung faktorial dari n (n!)
func hitungFaktorial(n int) int {
	hasil := int(1)
	for i := 1; i <= n; i++ {
		hasil = hasil * int(i)
	}
	return hasil
}

// Menghitung 2 pangkat n (2^n)
func hitungPangkatDua(n int) float64 {
	hasil := 1.0
	for i := 0; i < n; i++ {
		hasil = hasil * 2
	}
	return hasil
}

// Menghitung f(n) = pembulatan ke atas dari (n! / 2^n)
func hitungFaktor(n int) int {
	faktorial := hitungFaktorial(n)
	nilaiPangkat := hitungPangkatDua(n)
	pembagian := float64(faktorial) / nilaiPangkat
	hasilAkhir := math.Ceil(pembagian) // pembulatan ke atas
	return int(hasilAkhir)
}

func main() {
	fmt.Println("Selamat datang ! Program ini hanya menghitung nilai matematis f(n) = (n!) ÷ (2n)")
	var n int // faktorial n
	fmt.Print("Masukkan nilai faktorial: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Faktorial harus positif")
		return
	}

	hasil := hitungFaktor(n)
	fmt.Printf("Hasil f(%d) = %d\n", n, hasil)
}
