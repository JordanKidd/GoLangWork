package main

import "fmt"

/*  More GoLang Examples for my reference!
    This focuses on 'classes' / OOP
*/

type Person struct {
    name string
    job string
    ssn int
}
//Adds methods to the objects
func (p *Person) SayName() {
    fmt.Printf("My name is %v\n", p.name)
}

//Embedded type (anonomous fields):
//Android "is a" Person
//person Person would be "has a"
type Android struct {
    Person //No var name!
    modelNum int
    modelName string
}
func (a *Android) SayModelName() {
    fmt.Printf("My MoDeL nAmE iS: %v\n", a.modelName)
}


func main() {

    //Base class
    jordan := new(Person)
    jordan.name = "Jordan"
    jordan.job = "Software engineer"
    jordan.ssn = 333224444
    jordan.SayName()

    fmt.Printf("\n")

    //'Inherited' class with Person methods
    beepboop := new(Android)
    beepboop.modelName = "123-AB-v92.23"
    beepboop.modelNum = 12345
    beepboop.Person.name = "Droidy" //
    beepboop.Person.job = "RoboCop"
    beepboop.name = "blah"
    beepboop.Person.SayName()
    beepboop.SayModelName()

}