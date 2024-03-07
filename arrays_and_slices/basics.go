package main

import "fmt"

func f(a [3]int)   { a[0] = 100 }
func fp(a *[3]int) { a[0] = 100 }

func Sum(a *[3]float64) (sum float64) {
	for _, v := range a {
		sum += v
	}
	return sum
}

func main() {
	// var arr1 [5]int

	// for i := 0; i < len(arr1); i++ {
	// 	arr1[i] = i * 2
	// }

	// for i := range arr1 {
	// 	arr1[i] = i * 2
	// }

	// fmt.Println(arr1)

	// var ar [3]int
	// f(ar)
	// fp(&ar)

	// var arrLazy = [...]int{1, 2, 3}
	// var arrKeyVal = [5]int{3: 100, 4: 200}

	// array := [3]float64{7.0, 8, 9}
	// x := Sum(&array)
}
