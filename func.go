package main

import "fmt"

func main() {
	// sayHello();
	// sayHelloTo("Dede", "YudhaL")
	// sayHelloTo("Laksmana", "Yudha")

	// greetingMessage := getHello("Yudha")
	// fmt.Println(greetingMessage)

	// firstName, lastName := getFullName()
	// fmt.Println(firstName, lastName)

	// fristNameDua, _ := getFullName()
	// fmt.Println(fristNameDua)
	
	// firstName, lastName, age := getCompleteProfile()
	// fmt.Println(firstName, lastName, age)

	// fmt.Println("Total ", sumAll(1, 2, 3, 4))
	// numbers := []int{1, 2, 3, 4}
	// total := sumAll(numbers...)
	// fmt.Println("Total ", total)

	// goodbye := getGoodBye
	// myOtherGoodbyte := getGoodBye
	// fmt.Println(goodbye("YudhaL"))
	// fmt.Println(myOtherGoodbyte)
	// fmt.Println(myOtherGoodbyte("Dede"))
}

func sayHello() {
	fmt.Println("Hello, Yudha")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func getHello(name string) string {
	return "Good morning, " + name
}

// multiple return values
func getFullName() (string, string) {
	return "Laksmana", "Yudha"
}

// named return values
func getCompleteProfile() (firstName string, lastName string, age int) {
	firstName = "Laksmana"
	lastName = "Yudha"
	age = 26
	return firstName, lastName, age
}

// variadic function
func sumAll(numbers ...int) int {
	// numbers jadi slice []int
	total := 0;
	for _, number := range numbers {
		total += number
	}

	return total
}

// function value
func getGoodBye(name string) string {
	return "Good Bye " + name
}