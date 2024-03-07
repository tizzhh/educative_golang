package main

import "fmt"

func reverse(s string) string {
	tmp := []rune(s)
	for i := 0; i < len(tmp)/2; i++ {
		tmp[i], tmp[len(tmp)-i-1] = tmp[len(tmp)-i-1], tmp[i]
	}

	return string(tmp)
}

func main() {
	fmt.Println(reverse("Google"))
}
