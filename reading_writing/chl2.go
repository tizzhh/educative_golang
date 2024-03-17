package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Book struct {
	title    string
	price    float64
	quantity int
}

func main() {
	file, err := os.Open("products2.txt")
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("error", err)
		return
	}
	books := []Book{}
	for _, record := range records {
		price, _ := strconv.ParseFloat(record[1], 64)
		quant, _ := strconv.Atoi(record[2])
		books = append(books, Book{record[0], price, quant})
	}
	fmt.Println(books)
}
