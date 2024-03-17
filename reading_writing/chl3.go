package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	file, err := os.OpenFile("output/"+p.Title, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, str := range p.Body {
		writer.WriteByte(str)
	}
	writer.Flush()
	return nil
}

func load(title string) (*Page, error) {
	file, err := os.Open(title)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	page := new(Page)
	page.Title = title
	for {
		str, err := reader.ReadBytes('\n')
		switch {
		case err == io.EOF:
			return page, nil
		case err != nil:
			return nil, err
		}
		page.Body = append(page.Body, str...)
	}
}

func main() {
	page, err := load("products.txt")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	page.save()
	page2 := &Page{"aboba", []byte("hello how are you\n")}
	page2.save()
}
