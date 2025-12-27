package main

import "fmt"

type validationError struct {
	Message string
}

func (self validationError) Error() string {
	return self.Message
}

type notFoundError struct {
	Message string
}

func (self *notFoundError) Error() string {
	return self.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return validationError{"ID kosong"}
	}

	if id != "yudha" {
		return &notFoundError{"data not found"}
	}

	return nil
}

func main() {
	err := SaveData("dede", nil)

	if err != nil {
		// validationErr, ok := err.(validationError)
		// if ok { fmt.Println("validation error", validationErr.Error()) }

		// notFoundErr, ok := err.(*notFoundError)
		// if ok { fmt.Println("not found error", notFoundErr.Error()) }

		if validationErr, ok := err.(validationError); ok {
			// fmt.Println("validation error:", validationErr.Error())
			fmt.Println("validation error:", validationErr.Message)
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			// fmt.Println("not found error:", notFoundErr.Error())
			fmt.Println("not found error:", notFoundErr.Message)
		} else {
			fmt.Println("unknown error:", err.Error())
		}

		switch myErr := err.(type) {
		case validationError:
			fmt.Println("validation error:", myErr.Message)
		case *notFoundError:
			fmt.Println("not found error:", myErr.Message)
		default:
			fmt.Println("unknown error:", myErr.Error())
		}
	}
}