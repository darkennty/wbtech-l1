package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Println(checkForUnique(str1)) // true
	fmt.Println(checkForUnique(str2)) // false
	fmt.Println(checkForUnique(str3)) // false
}

func checkForUnique(str string) bool {
	str = strings.ToLower(str)
	uniqueMap := make(map[byte]struct{})

	for i := range str {
		if _, ok := uniqueMap[str[i]]; ok {
			return false
		}
		uniqueMap[str[i]] = struct{}{}
	}
	return true
}
