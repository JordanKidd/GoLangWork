package main

import "fmt"

/*  GoLang Quick Examples:
    Created by: Jordan Kidd */

//Very similar 'object' notation to C:
type Node struct {
    i int
    s string
    next *Node  //Note * placement!
}
//Adding 'methods' to a class as this format:
func (n Node) getInt() int { return n.i }
func (n Node) setInt(passed int) { n.i = passed }
func (n Node) getStr() string { return n.s }
func (n Node) setStr(passed string) { n.s = passed }
func (n Node) getNext() *Node { return n.next }
func (n Node) setNext(passed *Node) { n.next = passed }


func main() {

    fmt.Printf("Hello world!\n")
    println("This is an example to print to stderr!") //BIF


    //The assignment operator(s)
    var i int = getNum()
    w := getNum()
    //:= is just syntatic sugar for 'var varName type = blahblah
    //%v is for "value" and the default for gotypes (int, char, string, etc)
    fmt.Printf("Function returned: %v, %v\n", i, w)


    //Go's grammar uses ;'s, but are placed by lexer for most statements:
    //Certain places the programmer but include them...
    //*Note { must be on the same line as a statement!!
    for i = 0; i < 24; i++ {
        if i % 4 == 0 {
            fmt.Printf("\n")
        }
        fmt.Printf("%v\t", i)
    }
    fmt.Printf("\n")


    //Multiple value assignment:
    var a, b, c, d = getMultipleNums(false)
    fmt.Printf("Returning multiple items: %v, %v, %v, and %v\n", a,b,c,d)

    fmt.Printf("\nNow for pointers...!\n")
    var memEx int = 55
    fmt.Printf("MemEx's memory location is: %v\n", &memEx)

    //Dynamic memory:
    //one is a Node 'object' allocated via the 'new' BIF
    one := new(Node)  //all default values inside one defined by GoLang types
    two := Node{12345, "Constrctor Example!", nil}  // In order assignment
    var three Node = Node{s:"just string"} //both!~ (i and next = defaulted)

    fmt.Printf("Node one used the defaults, while two used the order\n" +
               "of vars to construct the object. Three is a combo of both!\n")
    fmt.Printf("One: %v, %v, %v\n", one.i, one.s, one.next)
    fmt.Printf("Two: %v, %v, %v\n", two.i, two.s, two.next)
    fmt.Printf("Three: %v, %v, %v\n", three.i, three.s, three.next)
}


//return type is AFTER the func name + params:
func getNum() int {
    return 14
}


//arg names come BEFORE type, then return types:
func getMultipleNums(e bool) (int, int, int, int) {
    if e {
        return 2, 14, 1404, 4444
    } else {
        return 1, 11, 33, 5317
    }
}