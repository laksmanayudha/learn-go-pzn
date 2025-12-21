package main

// import "fmt"

// /*
// Pada golang, tidak perlu eksplisit menyebutkan suatu struct implement ke suatu interface
// Jika pada struct sudah memiliki kontrak yang sama (nama method, tipe data param, dan return value sudah sama)
// maka otomatis terdeteksi sudah implement interface tersebut
// */

// type HasName interface {
// 	// GetName(firstName string, lastName string) string
// 	GetName() string
// }

// type Person struct {
// 	Name string
// }

// type Animal struct {
// 	Name string
// }

// func (self Person) GetName() string {
// 	return self.Name
// }

// func (self Animal) GetName() string {
// 	return self.Name
// }

// func sayHello(impl HasName) {
// 	fmt.Println("Hello", impl.GetName())
// }

// func main() {
// 	person := Person{Name: "Yudha"}
// 	animal := Animal{Name: "Popo"}

// 	sayHello(person)
// 	sayHello(animal)
// }