package main

import "fmt"

/*  More GoLang Examples for my reference!
    This focuses on 'classes' / OOP
*/
type Person struct {
    name string
    job string
    ssn int
} //Adds methods to the objects
func (p *Person) SayName() {
    fmt.Printf("My name is %v.\n", p.name)
}
func (p *Person) SayJob() {
    fmt.Printf("My job is %v.\n", p.job)
}
//Embedded type (anonomous fields):
//Android "is a" Person.  person Person would be "has a"
type Android struct {
    Person //No var name!
    modelNum int
    modelName string
}
func (a *Android) SayName() {
    fmt.Printf("...MY NAME IS: %v.\n", a.Person.name)
}

func main() {
    //Base class
    jordan := new(Person)
    jordan.name = "Jordan"
    jordan.job = "Software Engineer"
    jordan.ssn = 333224444
    jordan.SayName()

    //'Inherited' class with Person methods
    beepboop := new(Android)
    beepboop.modelName = "123-AB-v92.23"
    beepboop.modelNum = 12345
    beepboop.name = "ANDY v1.0"
    beepboop.job = "RoboCop" //assigns to Person's job prop

    beepboop.SayName()
    beepboop.SayJob() //use of Person's method on Android object
}