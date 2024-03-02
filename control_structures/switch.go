package main

import "fmt"

func main() {
	var num1 int = 50
	switch num1 {
	case 98, 99:
		fmt.Println("98 or 99")
		fmt.Println("aboba")
	case 100:
		fmt.Println("100")
	default:
		fmt.Println("default")
	}

	switch {
	case num1 < 0:
		println("lower than 100")
	case num1 > 0 && num1 < 10:
		fmt.Println("between 0 and 100")
	default:
		fmt.Println("greater than 10")
	}

	switch num1 := 100; {
	case num1 < 0:
		fmt.Println("negative")
	default:
		fmt.Println(">= 0", num1)
	}

	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
