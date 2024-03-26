package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "fadli,darusalam,sragen\n" +
		"bambang,pamungkas,jakarta\n" +
		"darwin,nunez,uruguay"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {				//EOF (End Of File) / Sudah baris terakhir
			break
		}

		fmt.Println(record)
	}
}