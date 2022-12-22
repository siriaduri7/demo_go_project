package main
import ("fmt")

func main() {

	var greeting string = "hello world"
	fmt.Print(greeting) //Now we are going to print this variable using the println method by fmt package.
}

// First, we specify the name of the package to be used, which is fmt.
	//Here the function's name is println, which means print, and then go to new line.
	//The .operator between fmt and println is known as the member access operator.It is used to access members of a package.

//########### Let's see how we can print a variable using the same method by fmt package,###############//

/*

package main
import ("fmt")

func main() {
	var city string ="pune"
	fmt.Print(city)
}

