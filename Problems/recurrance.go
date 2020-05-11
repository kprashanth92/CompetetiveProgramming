package problems

import "fmt"

func FindRepetetiveValues(str string) {
	a := make(map[string]int)
	for _, val := range str {
		a[string(val)] = a[string(val)] + 1
	}
	fmt.Println(a)
}
