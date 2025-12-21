package main

// import "fmt"

// /*
// Multi-line comment
// Multi-line comment
// Multi-line comment
// Multi-line comment
// Multi-line comment
// Multi-line comment
// Multi-line comment
// */

// func endApp() {
// 	fmt.Println("End App")

// 	errMessage := recover() // ketika panic, tapi bisa recover, panic sebelumnya akan ditangani (tidak jadi panic / program tidak terhenti)
// 	fmt.Println("Terjadi panic", errMessage)
// }

// func runApp(err bool) {
// 	defer endApp()

// 	if (err) {
// 		panic("Ups Error")
// 	}

// 	// endApp()
// }

// func main() {
// 	runApp(true)
// 	fmt.Println("print on main")
// }