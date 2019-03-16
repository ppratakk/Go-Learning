package main

import "fmt"

func main() {
	/*
		There is 2 main ways for variable declaration.

		1. Manual Type Declaration
		- var {variable's name} {type of variable}
		by after declaration we have to assign value to that variable.

		Note: So, What if we didn't assign any value to our variable?
			=> That varaible will has "Zero Value" aka. "Default Value"
			eg. Type		-		Zero Value

				bool				false
				int					0
				float				0.0
				string				""
				function				nil
	*/

	var text01 string
	text01 = "This is my text01."
	fmt.Println(text01)

	/*
		2. Type Inference
		- In case of we need to declare variable and assign value in the same time,
		  we have to use ":=".
		  eg. text02 := "Here is my text02."
		by Golang will automatically define which type of variable by themselves.
	*/
	text02 := "Here is my text02."
	fmt.Println(text02)

}
