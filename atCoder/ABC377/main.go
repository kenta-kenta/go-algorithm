package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var str string
	fmt.Scanf("%s", &str)

	strSlice := strings.Split(str, "")
	sort.Strings(strSlice)

	if strings.Join(strSlice, "") != "ABC" {
		fmt.Print("No")
	} else {
		fmt.Print("Yes")
	}
}
