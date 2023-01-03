package main

import (
	"fmt"
	// "strings"
)

func main() {
    fmt.Println(isUnique("abcd"))
    fmt.Println(isUnique("abca"))
}

func isUnique(s string) bool {
    // seen := map[rune]bool{}
    //
    // for _, c := range s {
    //     if _, ok := seen[c]; ok {
    //         return false
    //     }
    //
    //     seen[c] = true
    // }
    //
    // return true

    // for _, c := range s {
    //     if strings.Count(s, string(c)) > 1 {
    //         return false
    //     }
    // }
    //
    // return true

    seen := make([]bool, 26)

    for _, c := range s {
        if seen[c - 'a'] {
            return false
        }

        seen[c - 'a'] = true
    }

    // NOTE: could also use a bit vector (bit flags) to save space

    return true
}
