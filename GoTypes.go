package main

import "fmt"

type Node struct {
    one int
    two float64
    three string
}

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

    //Arrays
    var a [5]*Node
    fmt.Printf("Length of array: %v\n", len(a))

    x := new(Node)
    x.one, x.two, x.three = 14, 3.14159, "pi"
    a[0] = x

    y := new(Node)
    y.one, y.two, y.three = 42, 1.41421, "sqrt(2)"
    a[3] = y

    //A is the array, i number in range
    for i := range a {
        if a[i] != nil {
            fmt.Printf("three val: %v\n", a[i].three)
        }
    }

    //Slices (similar to array, but NO defined length):
    //Literal
    alphabet := []string{"a", "b", "c", "d", "e",
                         "f", "g", "h", "i", "j"}
    fmt.Printf("%v\n", alphabet[3:7]) //d - g

    var s []int
    s = make([]int, 5, 5)
    fmt.Printf("%v\n", s[:]) //all
    //default types: s == []int{0, 0, 0, 0, 0}

    type Transform [3][3]float64
    //2d array, array of arrays
}