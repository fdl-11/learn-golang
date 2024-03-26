/*
[File Management]
Saat kita membuat atau membaca file menggunakan package os, struct File merupakan implementasi dari io.Reader dan io.Writer
Oleh karena itu, kita bisa melakukan baca dan tulis terhadap file tersebut menggunakan package io & bufio

[Open File]
Untuk membaca/membuat file, kita bisa menggunakan os.OpenFile(name, flag, permission)
name berisikan nama file, bisa absolute atau relative/local
flag merupakan penanda file, apakah untuk membaca, menulis, atau lain-lain
permission merupakan aturan permission yang diperlukan ketika membuat file, bisa disimulasikan di : https://chmod-calculator.com/
*/

package main

import (
	"bufio"
	// "fmt"
	"io"
	"os"
)

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	file.WriteString(message)

	return nil
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	file.WriteString(message)

	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func main() {
	// createNewFile("sample.log", "this is sample log")

	// result, _ := readFile("sample.log")
	// fmt.Println(result)

	addToFile("sample.log", "\nthis is add message")
}