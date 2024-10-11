package main

import (
	"fmt"
)

func main() {
	// FizzBuzz()
	// Reverse("こんにちは")
	list := []int{-40, 6, 9, 50}
	fmt.Println(MinNumber(list))
}

func FizzBuzz() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func Reverse(str string) {
	runes := []rune(str)
	end := len(runes)
	n := end / 2
	for i := 0; i < n; i++ {
		runes[i], runes[end-i-1] = runes[end-i-1], runes[i]
	}
	fmt.Println(string(runes))
}

func MinNumber(list []int) (int, error) {
	min := list[0]
	if len(list) == 0 {
		return 0, fmt.Errorf("配列が空です")
	} else {
		for _, num := range list {
			if min > num {
				min = num
			}
		}
	}
	return min, nil
}
