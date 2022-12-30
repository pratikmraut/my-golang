package main

import "fmt"

func main() {
	fmt.Println("Structs in golangs")

	// no inheritance in golang; No super or parent

	pratik := User{"Pratik", "pratikmraut123@gmail.com", true, 21}
	fmt.Println(pratik)
	fmt.Printf("pratik details are: %+v \n", pratik)

	fmt.Printf("Name is %v and email is %v\n", pratik.Name, pratik.Email)

	pratik.GetStatus()
	pratik.NewMail()

	fmt.Printf("Name is %v and email is %v\n", pratik.Name, pratik.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
