package main

import "fmt"

func main() {
	fmt.Println("Structs in golangs")

	// no inheritance in golang; No super or parent

	pratik := User{"Pratik", "pratikmraut123@gmail.com", true, 21}
	fmt.Println(pratik)
	fmt.Printf("pratik details are: %+v \n", pratik)

	fmt.Printf("Name is %v and email is %v.", pratik.Name, pratik.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
