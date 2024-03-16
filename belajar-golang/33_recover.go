// Recover adalah function yang bisa kita gunakan untuk menangkap data panic
// Dengan recover, proses panic akan terhenti, sehingga program dapat tetap berjalan
// Taruh recover di defer function

package main

import "fmt"

func endApp() {
	message := recover()
	
	// Kalau tidak ada error, nilai defaultnya nil
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}

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
	runApp(false)
}