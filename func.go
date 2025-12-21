package main

// import "fmt"

// type Filter func(string) string

// type Blacklist func(string) bool

// func main() {
// 	// sayHello();
// 	// sayHelloTo("Dede", "YudhaL")
// 	// sayHelloTo("Laksmana", "Yudha")

// 	// greetingMessage := getHello("Yudha")
// 	// fmt.Println(greetingMessage)

// 	// firstName, lastName := getFullName()
// 	// fmt.Println(firstName, lastName)

// 	// fristNameDua, _ := getFullName()
// 	// fmt.Println(fristNameDua)

// 	// firstName, lastName, age := getCompleteProfile()
// 	// fmt.Println(firstName, lastName, age)

// 	// fmt.Println("Total ", sumAll(1, 2, 3, 4))
// 	// numbers := []int{1, 2, 3, 4}
// 	// total := sumAll(numbers...)
// 	// fmt.Println("Total ", total)

// 	// goodbye := getGoodBye
// 	// myOtherGoodbyte := getGoodBye
// 	// fmt.Println(goodbye("YudhaL"))
// 	// fmt.Println(myOtherGoodbyte)
// 	// fmt.Println(myOtherGoodbyte("Dede"))

// 	// sayHelloWithFilter("Anjing", spamFilter)

// 	// blacklist := func (name string) bool {
// 	// 	return name == "anjing"
// 	// }

// 	// registerUser("anjing", func (name string) bool  {
// 	// 	return name == "anjing"
// 	// })

// 	nextVal := counter(3)
// 	fmt.Println(nextVal())
// 	fmt.Println(nextVal())
// 	fmt.Println(nextVal())

// 	count := 0
// 	increment := func () {
// 		fmt.Println("Current count", count)
// 		count++
// 	}

// 	increment()
// 	increment()
// 	increment()
// }

// func sayHello() {
// 	fmt.Println("Hello, Yudha")
// }

// func sayHelloTo(firstName string, lastName string) {
// 	fmt.Println("Hello", firstName, lastName)
// }

// func getHello(name string) string {
// 	return "Good morning, " + name
// }

// // multiple return values
// func getFullName() (string, string) {
// 	return "Laksmana", "Yudha"
// }

// // named return values
// func getCompleteProfile() (firstName string, lastName string, age int) {
// 	firstName = "Laksmana"
// 	lastName = "Yudha"
// 	age = 26
// 	return firstName, lastName, age
// }

// // variadic function
// func sumAll(numbers ...int) int {
// 	// numbers jadi slice []int
// 	total := 0;
// 	for _, number := range numbers {
// 		total += number
// 	}

// 	return total
// }

// // function value
// func getGoodBye(name string) string {
// 	return "Good Bye " + name
// }

// // function as parameter
// // func sayHelloWithFilter(name string, filter func(string) string) {
// // 	filteredName := filter(name)
// // 	fmt.Println("Hello, ", filteredName)
// // }

// func sayHelloWithFilter(name string, filter Filter) {
// 	filteredName := filter(name)
// 	fmt.Println("Hello,", filteredName)
// }

// func spamFilter(name string) string {
// 	if name == "Anjing" {
// 		return "..."
// 	} else {
// 		return name
// 	}
// }

// func registerUser(name string, blacklist Blacklist) {
// 	if blacklist(name) {
// 		fmt.Println("You are blocked", name)
// 	} else {
// 		fmt.Println("Welcome", name)
// 	}
// }

// // recursive function
// func factorialRecursive(value int) int {
// 	if value == 1 {
// 		return 1
// 	} else {
// 		return value * factorialRecursive(value - 1)
// 	}
// }

// // closure & higher order function
// func counter(defaultCount int) (func() int) { // higher order function
// 	count := defaultCount

// 	return func () int { // closure
// 		count++
// 		return count
// 	}
// }