package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(checkPermutation("abc", "cba"))
	fmt.Println(checkPermutation("abc", "cbap"))
}

func checkPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1Bytes := []byte(s1)
	s2Bytes := []byte(s2)

	sort.Slice(s1Bytes, func(a, b int) bool {
		return s1Bytes[a] < s1Bytes[b]
	})
	sort.Slice(s2Bytes, func(a, b int) bool {
		return s2Bytes[a] < s2Bytes[b]
	})

	return string(s1Bytes) == string(s2Bytes)

	// s1Seen := make([]int, 26)
	// s2Seen := make([]int, 26)
	//
	// for i := 0; i < len(s1); i++ {
	//     s1Seen[s1[i] - 'a'] += 1
	//     s2Seen[s1[i] - 'a'] += 1
	// }
	//
	// for i := 0; i < len(s1Seen); i++ {
	//     if s1Seen[i] != s2Seen[i] {
	//         return false
	//     }
	// }
	//
	// return true
}
