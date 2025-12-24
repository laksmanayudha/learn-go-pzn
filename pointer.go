package main

// import "fmt"

// type Address struct {
// 	City, Province, Country string
// }

// type Man struct {
// 	Name string
// }

// func ChangeCountryToIndonesia(address *Address) {
// 	address.Country = "Jepang"
// }

// func (self Man) Married() {
// 	self.Name = "Mr. " + self.Name
// }

// func (self *Man) Aged() {
// 	self.Name = "Mr. " + self.Name
// }

// func main() {
// 	// address := Address{"Subang", "Jawa Barat", "Indonesia"}
// 	// address2 := address

// 	// address2.City = "Bandung"
// 	// fmt.Println(address)
// 	// fmt.Println(address2)

// 	// // address3 := &address
// 	// var address3 *Address = &address;
// 	// address3.City = "Denpasar"
// 	// fmt.Println(address)
// 	// fmt.Println(address3)
// 	// // fmt.Println(*address3)

// 	// ChangeCountryToIndonesia(&address)
// 	// fmt.Println(address)

// 	// defaultnya go selalu duplikasi value ketika assign ataupun pass ke parameter
// 	yudha := Man{"Yudha"}
// 	// yudha.Married()
// 	yudha.Aged();

// 	fmt.Println(yudha.Name)
// }