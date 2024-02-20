package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "privvVvet"
	fmt.Println(strings.HasPrefix(str, "pr"))
	fmt.Println(strings.HasSuffix(str, "et"))
	fmt.Println(strings.Contains(str, "iv"))
	fmt.Println(strings.Index(str, "r"))
	fmt.Println(strings.LastIndex(str, "v"))
	fmt.Println(strings.LastIndex(str, "x"))
	fmt.Println(strings.Replace(str, "vvvv", "aboba", -1))
	fmt.Println(strings.Count(str, "v"))
	fmt.Println(strings.Repeat(str, 5))
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.ToUpper(str))

	str2 := "    aboba     "
	fmt.Println(str2)
	fmt.Println(strings.TrimSpace(str2))
	fmt.Println(strings.Trim(str, "Vvet"))

	str3 := "hello this, is a sentence"
	split_str := strings.Fields(str3)
	fmt.Println(split_str)
	fmt.Println(strings.Split(str3, ","))
	fmt.Println(strings.Join(split_str, " "))

	fmt.Println(strconv.IntSize)
	fmt.Println(strconv.Itoa(5))
	fmt.Println(strconv.FormatFloat(5.69420, 'e', 10, 64))
	fmt.Println(strconv.Atoi("5"))
	fmt.Println(strconv.ParseFloat("5.69420", 64))
}
