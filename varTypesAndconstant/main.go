package main

import "fmt"

// fist charcter is caps is a way to define d-tpes so
// Can be used or accesible by any files or folder
const LoginToken string = "mytoken"

func main() {
	var userName string = "maijd"
	fmt.Println(userName)
	fmt.Printf("Vriable is of type %T \n", userName)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Vriable is of type %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Vriable is of type %T \n", smallVal)

	var smallFloat float32 = 256.2355776767557 //will get 5 value after decimal, ise float64 for getting all te values
	fmt.Println(smallFloat)
	fmt.Printf("Vriable is of type %T \n", smallFloat)

	//default values and some alias
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Vriable is of type %T \n", anotherVariable)

	var defaultUserName string
	fmt.Println(defaultUserName)
	fmt.Printf("Vriable is of type %T \n", defaultUserName)

	//imlicit type or other ways to declare variables
	var website = "www.xyz.com"
	fmt.Println(website)

	//no variable
	//only can be use inside methods
	//outside the mehtods use standars ways
	numOfUser := 300000
	fmt.Println(numOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Vriable is of type %T \n", LoginToken)

}
