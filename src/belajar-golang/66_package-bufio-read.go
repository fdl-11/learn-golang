/*
Package bufio (buffered io) digunakan untuk membuat data IO seperti Reader dan Writer
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	input := strings.NewReader("this is long string\nwith new line\n")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()	// baris, prefix, error
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}
}