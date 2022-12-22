/* #### Syntax#### for local variable and global variable
{
// outer block (GLOBAL VARIABLES)
   
     {
		//inner block (LOCAL VARIABLES)
	 }
}

1) inner blocks (LOCAL VARIABLES)can access variables that are declared within outer blocks, that are declared inside a function or a block,
These are not accessible outside the functional block and these variables can also be declared inside looping or conditional 
statements.

2) but outer blocks cannot access variables that are declared within inner blocks. Global variables are variables that are declared
outside of a functional block. They are available throughout the lifetime of a program. They're declared at the top of the program
outside all functional blocks and well, they can be accessed from any part of the program.
below ex:
*/
package main
import ("fmt")
func main() {
	city := "London"
	{
		country := "UK"
		fmt.Println(country)
		fmt.Println(city)
	}
	fmt.Println(city)
	//fmt.Println(country) //./variable_scope.go:24:14: undefined: country //error is that globacl var cannot access local var

}
/*
/// LOCAL VARIABLE ///

func mani() {
	name := "lolla"
	fmt.println(name)
}





/// GLOBAL VARIABLE EXAMPLE .//////
package main 
import("fmt")

var name string = "Lisa" // global variable

func main() {
	fmt.println(name) // can access in local so printed in local
} 


