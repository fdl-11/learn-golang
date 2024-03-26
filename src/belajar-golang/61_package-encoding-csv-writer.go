package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"a", "b", "c", "d", "e",})	// _ = ignore error, karena func Write() mengembalikan error
	_ = writer.Write([]string{"1", "2", "3", "4", "5",})
	_ = writer.Write([]string{"Fadli", "Darusalam"})

	writer.Flush()
}