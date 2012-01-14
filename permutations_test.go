package main

import (
    "./permutations"
    "fmt"
)


const dict = "short_list.txt"

func main() {
    words := permutations.MakeAnagrams(dict)
    for word, ana := range words {
        fmt.Print(word, ": ")
        for _, a := range ana {
            fmt.Print("\"", string(a), "\" ")            
        }
        fmt.Print("\n")
    }
}
