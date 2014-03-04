package main

import "fmt"

func main() {
    fmt.Printf("Example of different data types in GoLang:\n")
    myMap := make(map[string]int)
    fmt.Printf("Map usage!\n")
    myMap["jordan"] = 14
    myMap["corey"] = 87
    myMap["tuesday"] = 2
    fmt.Printf("Length: %v\n", len(myMap))
    fmt.Printf("Key 'jordan' produces %v value\n", myMap["jordan"])

    //Numbers: ints or floating point
    i := 9999
    p := 3.14159
    fmt.Printf("Variable i is of type: %T\n", i)
    fmt.Printf("Variable i is of type: %T\n", p)

}