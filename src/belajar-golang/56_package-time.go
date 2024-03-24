/*
Package time adalah package yg berisikan fungsionalitas untuk manajemen waktu di golang
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	fmt.Println()

	utc := time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	fmt.Println()

	// layout := time.RFC822
	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2020-10-01")
	fmt.Println(parse)
}