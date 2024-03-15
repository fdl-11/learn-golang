// Defer function adalah function yg bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai dieksekusi
// Defer function akan selalu dieksekusi walau terjadi error di function yg dieksekusi

package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging()		// taruh defer di atas || misal kita ingin membuat function yg tetap dieksekusi walau terjadi error || defer function tetap dieksekusi walau ada error || kalau ada error, defer function akan dieksekusi di akhir
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runApplication(10)
}