package main

import (
	"fmt"
	"math"
)

func main() {
	// FizzBuzz()
	// Reverse("こんにちは")
	// list := []int{-40, 6, 9, 50}
	// fmt.Println(MinNumber(list))
	// fmt.Println(CountDigit(1000))
	// fmt.Println(MergeArray([]int{1, 3, 2, 4}, []int{2, 5, 1, 8}))
	// fmt.Println(Average([]int{1, 2}))
	// fmt.Println(AreaRectangle(3, 7))
	// fmt.Println(CountCaractar("addddeecccsd", 'd'))
	fmt.Println(IsOddArr([]int{1, 3, 6}))
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

func CountDigit(num int) int {
	if num == 0 {
		return 1
	}
	return int(math.Log10(float64(num))) + 1
}

func MergeArray(first, second []int) []int {
	arr := make([]int, 0, len(first)+len(second))
	arr = append(arr, first...)
	arr = append(arr, second...)
	return arr
}

func Average(list []int) (float64, error) {
	if len(list) == 0 {
		return 0, fmt.Errorf("リストに数値が入っていません")
	}
	sum := 0
	for _, num := range list {
		sum += num
	}
	return float64(sum) / float64(len(list)), nil
}

func AreaRectangle(width, height int) int {
	return width * height
}

func CountCaractar(s string, char rune) int {
	count := 0
	for _, str := range s {
		if str == char {
			count++
		}
	}
	return count
}

func IsOddArr(arr []int) string {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	if sum%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}
