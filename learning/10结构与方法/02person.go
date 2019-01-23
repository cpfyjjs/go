package main

import (
	"fmt"
	"strings"
)

type Person struct{
	firstName string
	lastName string
}

func upPerson(p *Person){
	p.firstName =strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}
func main() {
	var per1 Person
	per1.firstName = "Chris"
	per1.lastName = "Woodward"

	upPerson(&per1)
	fmt.Printf("The name of the person is %s %s\n",per1.firstName,per1.lastName)

	per2 := &Person{"Chris","Woodward"}
	fmt.Printf("The name of the person is %s %s\n",per2.firstName,per2.lastName)

	var per3 = new(Person)
	per3.firstName = "Chris"
	upPerson(per3)
	fmt.Printf("The name of the person is %s %s\n",per3.firstName,per3.lastName)





}
