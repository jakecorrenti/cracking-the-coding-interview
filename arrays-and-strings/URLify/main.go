package main

import (
    "fmt"
)

func main() {
    URLify([]rune{'M', 'r', ' ', 'J', 'o', 'h', 'n', ' ', 'S', 'm', 'i', 't', 'h', ' ', ' ', ' ', ' '}, 13)
}

func URLify(chars []rune, l int) {
    indexes := []int{}
    for i := 0; i < l; i++ {
        if chars[i] == ' ' {
            indexes = append(indexes, i)
        }
    }

    for i := l - 1; i >= indexes[0]; i-- {
        if i > indexes[len(indexes) - 1] {
            chars[i + (2 * len(indexes))] = chars[i]
        }

        if i == indexes[len(indexes) - 1] {
            chars[i + (2 * len(indexes) - 2)] = '%'
            chars[i + (2 * len(indexes) -1)] = '2'
            chars[i + (2 * (len(indexes)))] = '0'
            indexes = indexes[:len(indexes) -1]
            if len(indexes) < 1 {
                break
            }
        }
    }

    for _, c := range chars {
        fmt.Print(string(c))
    }
    fmt.Println()
}
