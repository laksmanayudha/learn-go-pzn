package main

// import "fmt"

// // nil bisa di assign ke beberapa tipe data
// // misalnya slice, map bisa di assign nilai nil kalau kosong
// // string tidak bisa

// func createNewString(name string) string {
// 	if name == "" {
// 		// return nil
// 		return ""
// 	} else {
// 		return name
// 	}
// }

// func createNewMap(name string) map[string]string {
// 	if name == "" {
// 		return nil
// 	} else {
// 		return map[string]string{
// 			"name": name,
// 		}
// 	}
// }

// func main() {
// 	data := createNewMap("")
// 	// fmt.Println(data["name"])

// 	if (data == nil) {
// 		fmt.Println("Data map masih kosong")
// 	} else {
// 		fmt.Println(data["name"])
// 	}
// }