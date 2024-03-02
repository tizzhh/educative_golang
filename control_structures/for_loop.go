package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// str := "Go is a beautiful language!"
	// fmt.Printf("The length of str is: %d\n", len(str))

	// for ix := 0; ix < len(str); ix++ {
	// 	fmt.Printf("char on pos %d is %c \n", ix, str[ix])
	// }
	// str2 := "Chinese: 日本語"
	// fmt.Printf("The length of str2 is: %d\n", len(str2))
	// for ix := 0; ix < len(str2); ix++ {
	// 	fmt.Printf("char on pos %d is %c \n", ix, str2[ix])
	// }

	i := 0

	for i < 5 {
		fmt.Println("iteration", i)
		i += 1
	}
}
