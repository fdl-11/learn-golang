// Panic function adalah function yg bisa kita gunakan untuk menghentikan program
// Panic function biasanya dipanggil ketika terjadi error pada saat program berjalan
// Saat panic function dipanggil, program akan terhenti, namun defer function akan tetap dieksekusi

package main

import "fmt"

func endApp() {
	fmt.Println("Bye!")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("APLIKASI ERROR")
	}

	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(true)
}