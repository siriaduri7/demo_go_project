//we can convert variable, from one data type to another.
//The process of converting one data type to another is known as Type Casting.Data types can be converted to other data types
/*package main // int to float64
############ println will print %v nd value ******* but printf prints only value that replace %v with value #########

import(
	"fmt"
	)
func main() {
	var i int = 90
	var f float64 =float64(i)
	fmt.Println("%.2f\n",f)
}

// float64 to int
 /*
 package main
 import ("fmt")
 func main(){
	var f float64 =45.78
	var i int = int(f)
	fmt.Printf("%v\n", i) //******************* check printf above def *****************
}*/

// ######################## strconv package ############################# there are two types
//(INTERGER TO STRING)
//1) Itoa() : It converts integer to string and it just returns one value, which is the string formed with a given integer.
/*package main
import (
	"fmt"
	"strconv"
)
func main() {
	var i int =45
	var s string = strconv.Itoa(i) // convert int to string
	fmt.Printf("%q", s)
}

//################## STRING TO INT Atoi()######################
which converts string to integer, and it returns two values, the corresponding integer, error(if any)
*/
package main
import (
	"fmt"
	"strconv"
)
func main() {
	//var s string = "45"
	var s string = "45abc"
	i,err := strconv.Atoi(s)
	fmt.Printf("%v, %T\n", i,i)
	fmt.Printf("%v, %T\n", err,err)
}
/*the Atoi() function will return as an error for this conversion.
o/p :
sirishaaduri@aduri go_project % go run Typecasting_*.go
0, int
strconv.Atoi: parsing "45abc": invalid syntax, *strconv.NumError

The error data type is a custom data type in the string conversion package. While if you see the error message, it simply
 says that it's an invalid syntax


