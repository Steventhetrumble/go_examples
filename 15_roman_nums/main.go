package main

import (
	"fmt"
	"strings"
)

func value(s string) int {
	fmt.Printf("Hello rune %s \n", s)
	switch s {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	default:
		return -1
	}
}

func main() {
	a := "IV"
	b := 0
	c := strings.Split(a, "")

	for i := 0; i < len(c); i++ {
		temp1 := value(c[i])
		fmt.Println(temp1)
		if i+1 < len(c) {
			temp2 := value(c[i+1])
			fmt.Println(temp2)
			if temp1 >= temp2 {
				fmt.Println("up top")
				b += temp1
			} else {
				fmt.Println("down low")
				b += temp2 - temp1
				i++
			}
		} else {
			b += temp1
		}
	}
	fmt.Printf("Hello worlds %d", b)

}
