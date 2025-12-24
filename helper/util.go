package helper

var version = "1.0.0"
var Application = "golang"

func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func GetVersion() string {
	return version
}

// access modifier public = Huruf depan kapital
// access modifier private = Huruf depan kecil
func SayHello(name string) string {
	return "Hello " + name
}