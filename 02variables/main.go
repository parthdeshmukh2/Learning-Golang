package main

import "fmt"

func main() {
	var username string = "Parth"
	fmt.Println(username, "username")
	fmt.Printf("Variable Type %T", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn, "isLoggedIn")
	fmt.Printf("Variable Type %T", isLoggedIn)

}