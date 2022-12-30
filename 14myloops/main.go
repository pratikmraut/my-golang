package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednusday", "Friday", "Saturday"}

	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v \n", index, day)
	// }

	roughValue := 1

	for roughValue < 10 {

		if roughValue == 2 {
			goto lco
		}

		if roughValue == 5 {
			roughValue++
			continue
		}

		fmt.Println("Value is: ", roughValue)
		roughValue++
	}

lco:
	fmt.Println("Jumping at GOTO")

}
