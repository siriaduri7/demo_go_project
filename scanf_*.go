//####### syntax : fmt.scanf("%<format specifier> (s)", object_arguments) ##########// ex : fmt.Scanf("%s", &name)
//One of the most used ways of taking user input is through scanner function,which is provided by the FMT package.
// Scanner function takes the format string(%<format specifier> (s)"), along with the list of variable number of arguments(object_arguments)
//This string contains custom specifiers that the scanner function uses to format the final output string. (%<format specifier> 
//The list hood is the list of all arguments where you want to store your data. object_arguments)
//"&" in scanf we use ampersand symbol
//////////////////////////////////////////////////////////////
/* package main
import("fmt")
func main() {
	 //var name string
	 /*fmt.Print("Enter your Name = ")
	 fmt.Scanf("%s", &name)
	 fmt.Println("hey there, ", name)*/

	 ///////////////////////////
	 /*var name string
	 var is_muggle bool

	 fmt.Print("enter your namme & are you a muggle :")
	 fmt.Scanf( "%s, %t" , &name, &is_muggle)
	 fmt.Println(name, is_muggle)
}
*/

//////////////////////////////////
// ########### MULTIPLE INPUT ### can also take multiple inputs through single Scanf statement this way. (check above)

//############ SCANF RETURN VALUE ###########////////////////

//The FMT scanner function returns two values,

/* 1) count:  which is the number of arguments that the function writes to successfully,

2) error, any error that's thrown during the execution of the scanner function. We'll be using this information to see what 
errors can be caused while taking input from the user. */

package main
import ("fmt")
func main() {
	var a string
	var b int
	fmt.Print("enter a string and a number :")
	count, err := fmt.Scanf("%s %d", &a, &b)

	fmt.Println("count :",count)
	fmt.Println("error :", err)
	fmt.Println("a :",a)
	fmt.Println("b:",b)
}

