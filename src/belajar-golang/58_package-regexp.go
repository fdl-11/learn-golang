/*
Package regexp adalah utilitas di Golang untuk melakukan pencarian regular expression
Regexp di golang menggunakan library C yg dibuat google bernama RE2 (github.com/google/re2/wiki/Syntax)
*/

package main

import (
	"fmt"
	"regexp"
)

func main() {
	r, _ := regexp.Compile("[0-9]+")
    fmt.Println(r.MatchString("123456789"))
    fmt.Println(r.MatchString("1234567890"))
    fmt.Println(r.MatchString("12345678901"))
    fmt.Println(r.MatchString("123456789012"))
    fmt.Println(r.MatchString("1234567890123"))
    fmt.Println(r.MatchString("12345678901234"))
    fmt.Println(r.MatchString("123456789012345"))

	fmt.Println()

	var regex *regexp.Regexp = regexp.MustCompile("e([a-z]o)")
	fmt.Println(regex.MatchString("hello"))
	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eGo"))

	fmt.Println()

	// search := regex.FindAllString("eko eka edo eki esa ego", 1) // cari satu data yg sesuai regex
	// search := regex.FindAllString("eko eka edo eki esa ego", 2) // cari satu data yg sesuai regex
	search := regex.FindAllString("eko eka edo eki esa ego", -1) // cari semua data yg sesuai regex
	fmt.Println(search)
}