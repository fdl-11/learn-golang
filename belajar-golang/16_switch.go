package main

import "fmt"

func main() {
	name := "Fadli"

	switch name {
		case "Fadli":
            fmt.Println("Halo Fadli")
            fmt.Println("Salam kenal...")
        case "Darusalam":
            fmt.Println("Halo Darusalam")
        default:
            fmt.Println("Hi, kenalan yuk")
	}

	// =================================================================

	// switch length := len(name); length > 5 {
	// 	case true:
    //         fmt.Println("Nama terlalu panjang")
    //     case false:
    //         fmt.Println("Nama sudah cukup")
	// }

	// =================================================================

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah cukup")
		
	}
}