package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
    //     counter++
	// }

	// =================================================================

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	// =================================================================

	names := []string{"Fadli", "Darusalam", "Sragen"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for i, name := range names {
		fmt.Println("Index", i, "=", name)
	}
	
	// Tanda underscore untuk memberitahu bahwa variabel ini tidak diperlukan
	for _, name := range names {
		fmt.Println(name)	// print value
	}
	
	for name := range names {
		fmt.Println(name)	// print index
	}

	// =================================================================

	person := make(map[string]string)
	person["name"] = "Fadli"
	person["address"] = "Sragen"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

}