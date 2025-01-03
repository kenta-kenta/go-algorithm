package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"unicode"
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
	// fmt.Println(IsOddArr([]int{1, 3, 6}))
	// fmt.Println(GCM(30, 12))
	// fmt.Println(areAnagrams("abrtd", "tdrab"))
	// fmt.Println(Palindrome("a leela"))
	// fmt.Println(euclideanDistance(1, 1, 4, 5))
	fmt.Println(IsPrime(2))
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

func GCM(num1, num2 int) int {
	if num2 == 0 {
		return num1
	}
	return GCM(num2, num1%num2)
}

func areAnagrams(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1Slice := strings.Split(s1, "")
	s2Slice := strings.Split(s2, "")
	sort.Strings(s1Slice)
	sort.Strings(s2Slice)
	return strings.Join(s1Slice, "") == strings.Join(s2Slice, "")
}

func Palindrome(s string) bool {
	var filterd []rune
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			filterd = append(filterd, r)
		}
	}
	for i, j := 0, len(filterd)-1; i < j; i, j = i+1, j-1 {
		if filterd[i] != filterd[j] {
			return false
		}
	}
	return true
}

func euclideanDistance(x1, y1, x2, y2 float64) float64 {
	dx := x1 - x2
	dy := y1 - y2
	return math.Pow(math.Pow(dx, 2)+math.Pow(dy, 2), 0.5)
}

func UpperToLower(str string) string {
	return strings.ToLower(str)
}

func IsPrime(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
